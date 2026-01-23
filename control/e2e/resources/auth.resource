*** Settings ***
Documentation       Auth resource file with shared keywords

Library             Collections
Library             String
Library             RequestsLibrary
Library             ImapLibrary2
Variables           ../variables.py


*** Keywords ***
Get OTP From Email
    [Documentation]    Retrieves the OTP code from the email with retry logic
    [Arguments]    ${recipient_email}    ${max_attempts}=15    ${wait_between_attempts}=5s

    Log    Starting email retrieval for: ${recipient_email}    level=INFO

    # Convert arguments to integers for calculation
    VAR    ${max_attempts_int} =    ${{int($max_attempts)}}
    VAR    ${range_end} =    ${{${max_attempts_int} + 1}}

    Sleep    3s

    FOR    ${attempt}    IN RANGE    1    ${range_end}
        Log    Attempt ${attempt}/${max_attempts} to retrieve email    level=INFO

        TRY
            # Standard assignment to capture keyword return value
            ${otp_code} =    Try Get OTP From Mailbox    ${recipient_email}
            RETURN    ${otp_code}
        EXCEPT    AS    ${error}
            Log    Attempt ${attempt} failed: ${error}    level=WARN
            Handle Failed OTP Attempt    ${attempt}    ${max_attempts_int}    ${wait_between_attempts}    ${error}
        END
    END

Open Mailbox And Wait For Email
    [Documentation]    Opens the mailbox and waits for a specific recipient email
    [Arguments]    ${recipient_email}

    Open Mailbox
    ...    host=${API_MAILING_IMAP_HOST}
    ...    user=${API_MAILING_IMAP_USERNAME}
    ...    password=${API_MAILING_IMAP_PASSWORD}
    ...    port=${API_MAILING_IMAP_PORT}

    ${emails} =    Wait For Email
    ...    sender=noreply@bflow.dev
    ...    recipient=${recipient_email}
    ...    timeout=10

    RETURN    ${emails}

Try Get OTP From Mailbox
    [Documentation]    Attempts to retrieve OTP from mailbox in a single execution
    [Arguments]    ${recipient_email}

    TRY
        ${emails} =    Open Mailbox And Wait For Email    ${recipient_email}

        # Check if email exists
        IF    $emails == $None or $emails == ""    Fail    No email found yet

        ${email_body} =    Get Email Body    ${emails}

        # Alphanumeric regex for codes like 72082f
        VAR    ${otp_pattern} =    id="otp"[^>]*>\\s*([a-zA-Z0-9]{6})\\s*</span>
        ${matches} =    Get Regexp Matches    ${email_body}    ${otp_pattern}    1
        ${match_count} =    Get Length    ${matches}

        Should Be True    ${match_count} > 0    msg=OTP pattern not found in email body

        VAR    ${otp_code} =    ${matches[0]}

        Delete Email    ${emails}
        Close Mailbox

        RETURN    ${otp_code}
    EXCEPT    AS    ${error}
        # Ensure mailbox is closed even on failure
        Run Keyword And Ignore Error    Close Mailbox
        Fail    ${error}
    END

Handle Failed OTP Attempt
    [Documentation]    Logic to sleep between retries or fail on exhaustion
    [Arguments]    ${attempt}    ${max_attempts}    ${wait_time}    ${error}

    IF    $attempt == $max_attempts
        Fail    Failed to retrieve OTP after ${max_attempts} attempts. Last error: ${error}
    END

    Log    Waiting ${wait_time} before next attempt...    level=INFO
    Sleep    ${wait_time}

Login With Credentials
    [Arguments]    ${email}    ${password}
    VAR    ${payload} =    ${{{"email": "${email}", "password": "${password}"}}}

    ${response} =    POST On Session    api_session    auth/login/credential    json=${payload}    expected_status=200

    VAR    ${auth_data} =    ${response.json()}
    Dictionary Should Contain Key    ${auth_data}    access_token
    RETURN    ${auth_data}

Delete All Emails From Mailbox
    [Documentation]    Cleans the inbox to avoid state interference
    TRY
        Open Mailbox
        ...    host=${API_MAILING_IMAP_HOST}
        ...    user=${API_MAILING_IMAP_USERNAME}
        ...    password=${API_MAILING_IMAP_PASSWORD}
        ...    port=${API_MAILING_IMAP_PORT}

        ${status} =    Run Keyword And Return Status    Delete All Emails
        Close Mailbox
    EXCEPT    AS    ${error}
        Log    Failed to clean mailbox: ${error}    level=WARN
        Run Keyword And Ignore Error    Close Mailbox
    END
