*** Settings ***
Documentation       Account registration suite using shared resources

Resource            ../../resources/auth.resource

Suite Setup         Setup Test Suite
Suite Teardown      Teardown Test Suite


*** Test Cases ***
Complete Registration Flow
    [Documentation]    Flow: register, get OTP, resend, activate and login
    [Tags]    auth    registration_flow

    # 1. Register new account
    ${test_email} =    Register New Account

    # 2. Get first activation code
    ${first_otp} =    Get First Activation Code    ${test_email}

    # 3. Resend activation code
    ${second_otp} =    Resend And Get New Code    ${test_email}    ${first_otp}

    # 4. Activate account
    Activate Account With OTP    ${test_email}    ${second_otp}

    # 5. Login with activated account
    Login With Activated Account    ${test_email}


*** Keywords ***
Setup Test Suite
    [Documentation]    Initializes API session and cleans IMAP inbox
    Create Session    api_session    ${API_URL}    verify=True
    Delete All Emails From Mailbox

Teardown Test Suite
    [Documentation]    Closes all sessions
    Delete All Sessions

Register New Account
    [Documentation]    Registers user and returns random email
    ${random_str} =    Generate Random String    10    [LOWER]
    # DEPR05 fix: Using VAR for direct string assignment
    VAR    ${random_email} =    test_${random_str}@email.com

    VAR    ${payload} =    ${{{"email": "${random_email}", "password": "Test@123", "name": "User ${random_str}"}}}
    POST On Session    api_session    auth/register    json=${payload}    expected_status=201

    RETURN    ${random_email}

Get First Activation Code
    [Documentation]    Fetches initial OTP from email
    [Arguments]    ${email}
    ${otp} =    Get OTP From Email    ${email}    max_attempts=12    wait_between_attempts=5s
    RETURN    ${otp}

Resend And Get New Code
    [Documentation]    Triggers resend and returns new unique OTP
    [Arguments]    ${email}    ${previous_otp}

    Request Activation Code Resend    ${email}

    ${new_otp} =    Get OTP From Email    ${email}    max_attempts=12
    Should Not Be Equal    ${previous_otp}    ${new_otp}    msg=Resent OTP must be different

    RETURN    ${new_otp}

Request Activation Code Resend
    [Documentation]    Sends POST to resend activation code
    [Arguments]    ${email}
    VAR    ${resend_payload} =    ${{{"email": "${email}"}}}
    POST On Session    api_session    auth/activate    json=${resend_payload}    expected_status=202

Activate Account With OTP
    [Documentation]    Sends PATCH to activate the account
    [Arguments]    ${email}    ${otp}
    VAR    ${payload} =    ${{{"email": "${email}", "otp": "${otp}"}}}
    PATCH On Session    api_session    auth/activate    json=${payload}    expected_status=200

Login With Activated Account
    [Documentation]    Logins and validates JWT tokens
    [Arguments]    ${email}
    ${auth_data} =    Login With Credentials    ${email}    Test@123
    Dictionary Should Contain Key    ${auth_data}    access_token
    Dictionary Should Contain Key    ${auth_data}    refresh_token

    # DEPR05 fix: Replaced Get From Dictionary with direct VAR access if desired
    # or simply use VAR to replace any Set Variable used for extraction
    VAR    ${access_token} =    ${auth_data['access_token']}
    Should Not Be Empty    ${access_token}
