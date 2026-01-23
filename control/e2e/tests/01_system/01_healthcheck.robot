*** Settings ***
Documentation    E2E tests for the Control API system health status.
Library          RequestsLibrary
Library          Collections
Variables        ../../variables.py

*** Variables ***
# This path is appended to the BASE_URL provided by variables.py
${HEALTH_PATH}    _system/health

*** Test Cases ***
Verify API System Health Status
    [Tags]             system    healthcheck
    
    # 1. Initialize the Session
    Create Session    api_ctrl    ${API_URL}    verify=True
    
    # 2. Perform the Request
    ${response}=      GET On Session    api_ctrl    ${HEALTH_PATH}    expected_status=200
    
    # 3. Validate Basic Response Info
    Status Should Be    200    ${response}
    ${json}=            Set Variable    ${response.json()}
    
    # 4. Validate Root Keys (Based on your Go API implementation)
    Dictionary Should Contain Key    ${json}    uptime
    Dictionary Should Contain Key    ${json}    status
    Dictionary Should Contain Key    ${json}    services
    
    # 5. Validate System Status Value
    # Using 'Should Be True' because your image shows "status": true
    Should Be True    ${json}[status]
    
    # 6. Validate Nested Services Object
    ${services}=    Set Variable    ${json}[services]
    Dictionary Should Contain Key    ${services}    broker
    Dictionary Should Contain Key    ${services}    cache
    Dictionary Should Contain Key    ${services}    database
    Dictionary Should Contain Key    ${services}    storage

    # 7. Optional: Validate Uptime is not empty
    Should Not Be Empty    ${json}[uptime]