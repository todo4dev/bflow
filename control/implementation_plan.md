# Implementation Plan - Refresh Authorization Token

## User Review Required

> [!IMPORTANT]
> Verify if `src/port/jwt` exposes a method to parse/validate the refresh token and extract claims (jti) without full verification if we rely on cache, or if we should verify signature first. The plan assumes `jwtProvider` has a method to parse/verify.

## Proposed Changes

### Documentation

#### [NEW] [refresh-authorization-token.mdx](file:///c:/dev/todo4dev/bflow/control/doc/src/content/use-cases-api/auth/refresh-authorization-token.mdx)

### Application Layer (Auth Domain)

#### [NEW] [refresh_authorization_token.go](file:///c:/dev/todo4dev/bflow/control/api/application/auth/refresh_authorization_token/refresh_authorization_token.go)
- Implement `Handler` struct with dependencies:
    - `repository.Account` (to check if active)
    - `jwt.Provider` (to parse and create tokens)
    - `cache.Client` (to validate and update session)
- Implement `Data` struct:
    - `RefreshToken` (string, required)
- Implement `Result` struct:
    - `TokenType`
    - `AccessToken`
    - `RefreshToken`
    - `ExpiresIn`
- Implement `Handle` method:
    1.  Validate input.
    2.  Parse `RefreshToken` to get `jti` (session ID). Handle invalid token error.
    3.  Check cache for `key := fmt.Sprintf("account:%s:session:%s", accountId, jti)`. If missing -> 401.
    4.  Parse cache value to get `maxAge`. If `maxAge < now` -> delete cache, return 401.
    5.  Check account status (active). If not -> 401.
    6.  Generate new JWT pair (Access + Refresh) using same `jti`.
    7.  Update session in cache with new expiration (Access Token TTL).
    8.  Return new tokens.

#### [MODIFY] [usecase.go](file:///c:/dev/todo4dev/bflow/control/api/application/auth/usecase.go)
- Register `refresh_authorization_token` handler.

### Presentation Layer (HTTP)

#### [MODIFY] [router.go](file:///c:/dev/todo4dev/bflow/control/api/presentation/http/router/router.go) (or similar)
- Register `POST /auth/refresh` route mapping to the new handler.

## Verification Plan

### Automated Tests
- Since I cannot run full integration tests easily without setup, I will rely on manual verification or unit tests if possible.
- I will attempt to run the server and hit the endpoint with curl.

### Manual Verification
1.  **Login**: Call `POST /auth/login` to get `access_token` and `refresh_token`.
2.  **Refresh Success**: Call `POST /auth/refresh` with the valid `refresh_token`. Expect 200 and new tokens.
3.  **Refresh Invalid**: Call with invalid string. Expect 401 or 400.
4.  **Refresh Expired**: (Hard to simulate without waiting or mocking time/cache).
5.  **Account Disabled**: (Requires modifying DB, might skip for quick check).
