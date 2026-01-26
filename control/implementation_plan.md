# Implementation Plan - SSO Provider Redirect

## User Review Required

> [!IMPORTANT]
> - Ensure `API_OIDC_*` environment variables are set.
> - Callback URL is passed as a query parameter.
> - Route is `GET /auth/:provider`.
> - **Behavior Change**: Redirects (302) to the provider URL instead of returning it in body.

## Proposed Changes

### Documentation

#### [NEW] [sso-provider-redirect.mdx](file:///c:/dev/todo4dev/bflow/control/doc/src/content/use-cases-api/auth/sso-provider-redirect.mdx)
- Create documentation for the `SSO Provider Redirect` use case.
- Defines `GET /auth/:provider?callback_url=...` with 302 Redirect response.

### Application Layer (Auth Domain)

#### [NEW] [sso_provider_redirect.go](file:///c:/dev/todo4dev/bflow/control/api/application/auth/sso/sso_provider_redirect/sso_provider_redirect.go)
- **Data (Input)**:
    - `Provider` (string, required, oneof=google microsoft) - from Path
    - `CallbackURL` (string, required, url) - from Query
- **Result (Output)**:
    - `RedirectURL` (string)
- **Handler**:
    - Validate input using `gox/validate`.
    - Retrieve `oidc.Adapter` for the given provider.
    - Call `adapter.GetAuthURL(state=Data.CallbackURL, scopes=[openid, profile, email])`.
    - Return `Result{RedirectURL: url}`.

#### [MODIFY] [auth.go](file:///c:/dev/todo4dev/bflow/control/api/application/auth/auth.go)
- Register the new handler in `func Provide()`.

### Presentation Layer (HTTP)

#### [NEW] [sso_provider_redirect.go](file:///c:/dev/todo4dev/bflow/control/api/presentation/http/route/auth/sso_provider_redirect/sso_provider_redirect.go)
- Create the REST handler.
- Map HTTP Path param `:provider` to `Data.Provider`.
- Map HTTP Query param `callback_url` to `Data.CallbackURL`.
- Invoke Application Handler.
- **Return 302 Found** assuming success, setting `Location` header to `Result.RedirectURL`.

#### [MODIFY] [router.go](file:///c:/dev/todo4dev/bflow/control/api/presentation/http/router/router.go)
- Register `GET /auth/:provider` route.

## Verification Plan

### Automated Tests
- Unit test for `Handler` mocking `oidc.Provider`.

### Manual Verification
- Open in Browser: `http://localhost:30000/auth/google?callback_url=http://localhost:3000`.
- Verify browser redirects to Google Login.
