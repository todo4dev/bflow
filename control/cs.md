1.  system
    1.  ✅ healthcheck
2.  identity
    1.  auth
        1.  ✅ check if email is available
        2.  ⛔ register account using email and password
        3.  ⛔ activate account using email and code
        4.  ⛔ resend activation code for email
        5.  ⛔ send reset password for email
        6.  ⛔ reset password using email, new password and code
        7.  ⛔ login using email and password
        8.  ⛔ refresh authorization token
    2.  auth/sso
        1.  ⛔ redirect to sso provider passing callback url as state
        2.  ⛔ callback from sso provider callbacking to initial url with internal code
        3.  ⛔ login using internal code
    3.  me/account
        1.  ⛔ get me account
        2.  ⛔ delete me account
        3.  ⛔ logoff
    4.  me/profile
        1.  ⛔ update me profile
    5.  me/credential
        1.  ⛔ update me credential
    6.  me/avatar
        1.  ⛔ get me avatar
        2.  ⛔ upload me avatar
        3.  ⛔ delete me avatar
    7.  me/cover
        1.  ⛔ get me cover
        2.  ⛔ upload me cover
        3.  ⛔ delete me cover
    8.  me/notification
        1.  ⛔ list notifications (by cluster/me)
        2.  ⛔ get notification
        3.  ⛔ mark notification as read
        4.  ⛔ mark all notifications as read
        5.  ⛔ update notification preferences
        6.  ⛔ get unread notification count
        7.  ⛔ publish notification (agent -> control)
    9.  me/activity
        1.  ⛔ list activities by account (filters)
        2.  ⛔ list activities by cluster (filters)
        3.  ⛔ list activities by resource (via metadata)
        4.  ⛔ get activity
        5.  ⛔ export activities
        6.  ⛔ append activity (agent -> control)
3.  tenant
    1.  my/organization
        1.  ⛔ create organization
        2.  ⛔ update organization
        3.  ⛔ delete organization
        4.  ⛔ get organization
        5.  ⛔ search organizations
    2.  my/organization/membership
        1.  ⛔ get my membership
        2.  ⛔ list members by organization
        3.  ⛔ add member to organization
        4.  ⛔ update member role (viewer/manager/admin)
        5.  ⛔ remove member from organization
        6.  ⛔ leave organization (self)
        7.  ⛔ transfer viewer/manager/admin 
    3.  my/organization/membership/invite
        1.  ⛔ create membership invite (organization, email, role)
        2.  ⛔ resend membership invite
        3.  ⛔ cancel membership invite
        4.  ⛔ accept membership invite (code)
        5.  ⛔ get membership invite by code
        6.  ⛔ list membership invites by organization
    4.  infra/cluster [NOVO]
        1.  ⛔ list clusters by organization [NOVO]
        2.  ⛔ create cluster [NOVO]
        3.  ⛔ get cluster details [NOVO]
        4.  ⛔ update cluster [NOVO]
        5.  ⛔ delete cluster [NOVO]
        6.  ⛔ generate agent enrollment token [NOVO]
    5.  infra/agent
        1.  ⛔ register agent (agent -> control)
        2.  ⛔ poll actions (agent -> control)
        3.  ⛔ submit action result (agent -> control) (status, logs, output_meta)
    6.  infra/runtime
        1.  ⛔ list runtimes by organization
        2.  ⛔ get runtime (by id)
        3.  ⛔ get runtime status
        4.  ⛔ update runtime config
        5.  ⛔ set runtime readonly
        6.  ⛔ unset runtime readonly
        7.  ⛔ delete runtime
        8.  ⛔ get runtime current release
        9.  ⛔ set runtime current release (internal, on successful pipeline completion)
4.  billing
    1.  my/organization/billing/plan
        1.  ⛔ list plans
        2.  ⛔ get plan details
    2.  my/organization/billing/subscription
        1.  ⛔ get subscription by organization
        2.  ⛔ change subscription plan
        3.  ⛔ cancel subscription
        4.  ⛔ resume subscription
    3.  my/organization/billing/invoice
        1.  ⛔ list invoices by organization
        2.  ⛔ get invoice details
        3.  ⛔ download invoice
    4.  my/organization/billing/payment
        1.  ⛔ list payments by invoice
        2.  ⛔ get payment details
5.  deployment
    1.  asset/artifact
        1.  ⛔ list artifacts
        2.  ⛔ get artifact details
    2.  asset/release
        1.  ⛔ list available releases
        2.  ⛔ get release details
        3.  ⛔ mark release recommended/stable
        4.  ⛔ register/upload release 
    3.  pipeline
        1.  ⛔ request deploy (organization, target release)
        2.  ⛔ request update config (organization, payload)
        3.  ⛔ request redeploy (previous_pipeline_id)
        4.  ⛔ retry pipeline
        5.  ⛔ cancel pipeline
        6.  ⛔ get pipeline details
        7.  ⛔ get pipeline current status summary
        8.  ⛔ list pipelines by organization
    4.  pipeline/action
        1.  ⛔ get pipeline action summary
        2.  ⛔ list pipeline actions
    5.  pipeline/action/stage
        1.  ⛔ list stages by action
        2.  ⛔ get stage details
        3.  ⛔ get stage output metadata
        4.  ⛔ download stage output
        5.  ⛔ purge logs/output (admin-only)
6.  signing
    1.  doc/certificate
        1.  ⛔ list certificates
        2.  ⛔ upload certificate
        3.  ⛔ get certificate
        4.  ⛔ delete certificate
    2.  doc/document
        1.  ⛔ create document (cluster, upload file, kind, title, metadata)
        2.  ⛔ get document
        3.  ⛔ download document
        4.  ⛔ search documents (filters: kind/status/title/date, cluster optional)
        5.  ⛔ delete document (soft delete)
        6.  ⛔ replace document (reemit document, invalidate previous)
        7.  ⛔ get document signatures
        8.  ⛔ get document validity
        9.  ⛔ create document from agent (agent -> control) [NOVO]
        10. ⛔ delete document from agent (agent -> control) [NOVO]
        11. ⛔ replace document from agent (agent -> control) [NOVO]
    3.  doc/my-document
        1.  ⛔ list my documents (where I am signer or requester)
        2.  ⛔ get my document (must be related)
        3.  ⛔ download my document (must be related)
        4.  ⛔ get my document signatures
        5.  ⛔ get my document validity
    4.  doc/signature
        1.  ⛔ create signature request (document, account)
        2.  ⛔ list signatures by document
        3.  ⛔ get signature details
        4.  ⛔ verify signature (internal/optional)
        5.  ⛔ list my signatures
        6.  ⛔ create signature request from agent (agent -> control) [NOVO]
