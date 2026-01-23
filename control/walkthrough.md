# Walkthrough - Refresh Authorization Token

I have implemented the **Refresh Authorization Token** use case (Case 2.8).

## Changes

### Domain Layer (New Issues)
- **[NEW]** `validation/issue/auth.go`: Added `AuthInvalidToken` and `AuthSessionExpired` error types.

### Application Layer
- **[NEW]** `application/auth/refresh_authorization_token/refresh_authorization_token.go`: Implements the core logic:
    - Decodes Refresh Token to get Session ID (JTI).
    - Checks if Session exists in Cache (Redis) and validates `maxAge`.
    - Verifies if Account is active.
    - Rotates tokens (issues new Access/Refresh tokens) preserving the Session ID.
    - Updates Session TTL in Cache.
- **[MODIFY]** `application/auth/auth.go`: Registered the new handler in the DI container.

### Presentation Layer
- **[NEW]** `presentation/http/rest/auth/refresh_authorization_token/refresh_authorization_token.go`: HTTP Handler and OpenAPI definition.
- **[MODIFY]** `presentation/http/rest/auth/auth.go`: Registered `POST /auth/refresh` route.

## Verification

To verify the implementation:

1.  **Login**: Perform a login (`POST /auth/login/credential`) to obtain a `refresh_token`.
2.  **Refresh**: Use the token in `POST /auth/refresh`.
    ```bash
    curl -X POST http://localhost:30000/auth/refresh \
      -H "Content-Type: application/json" \
      -d '{"refresh_token": "YOUR_REFRESH_TOKEN"}'
    ```
3.  **Expectation**:
    - **200 OK**: With new `access_token` and `refresh_token`.
    - **401 Unauthorized**: If token is invalid, expired, or session is gone.
