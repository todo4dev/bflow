legend:
- ⛔ not implemented yet
- ✅ doc generated and implemented
- ✔️ end to end generated

1.  system
    1.  ✅ healthcheck
2.  auth
    1.  ✔️ check if email is available
    2.  ✔️ register account using email and password
    3.  ✔️ activate account using email and code
    4.  ✔️ resend activation code for email
    5.  ✔️ send reset password for email
    6.  ✔️ reset password using email, new password and code
    7.  ✔️ login using email and password
    8.  ✔️ refresh authorization token
3.  auth/sso
    1.  ✅ redirect to sso provider passing callback url as state
    2.  ✅ callback from sso provider callbacking to initial url with internal code
    3.  ✅ login using internal code
4.  me/account
    1.  ⛔ get me account
    2.  ⛔ delete me account
    3.  ⛔ logoff
5.  me/profile
    1.  ⛔ update me profile
6.  me/credential
    1.  ⛔ update me credential
7.  me/avatar
    1.  ⛔ get me avatar
    2.  ⛔ upload me avatar
    3.  ⛔ delete me avatar
8.  me/cover
    1.  ⛔ get me cover
    2.  ⛔ upload me cover
    3.  ⛔ delete me cover
9.  me/preference
    1.  ⛔ get preference
    2.  ⛔ update preference
10. my/organization
    1.  ⛔ create organization
    2.  ⛔ update my organization
    3.  ⛔ delete my organization
    4.  ⛔ get my organization
    5.  ⛔ search my organizations
11. organization/membership
    1.  ⛔ get my membership
    2.  ⛔ list members by organization
    3.  ⛔ add member to organization
    4.  ⛔ update member role (viewer/manager/admin)
    5.  ⛔ remove member from organization
    6.  ⛔ leave organization (self)
12. organization/membership/invite
    1.  ⛔ create membership invite (organization, email, role)
    2.  ⛔ resend membership invite
    3.  ⛔ cancel membership invite
    4.  ⛔ accept membership invite (code)
    5.  ⛔ get membership invite by code
    6.  ⛔ list membership invites by organization
13. infra/cluster
    1.  ⛔ list clusters by organization
    2.  ⛔ create cluster
    3.  ⛔ get cluster details
    4.  ⛔ update cluster
    5.  ⛔ delete cluster
14. infra/cluster/enrollment
    1.  ⛔ generate agent enrollment token
    2.  ⛔ list enrollment tokens by cluster
    3.  ⛔ revoke enrollment token
    4.  ⛔ validate enrollment token (agent -> control)
15. infra/agent
    1.  ⛔ register agent using enrollment token (agent -> control)
    2.  ⛔ list agents by cluster
    3.  ⛔ get agent details
    4.  ⛔ update agent metadata
    5.  ⛔ delete agent
    6.  ⛔ poll actions (agent -> control)
    7.  ⛔ submit action result (agent -> control) (status, logs, output_meta)
16. infra/runtime
    1.  ⛔ list runtimes by organization
    2.  ⛔ get runtime (by id)
    3.  ⛔ get runtime status
    4.  ⛔ update runtime config
    5.  ⛔ toggle runtime readonly
    6.  ⛔ delete runtime
    7.  ⛔ get runtime current release
17. organization/billing/plan
    1.  ⛔ list plans
    2.  ⛔ get plan details
18. organization/billing/subscription
    1.  ⛔ get subscription by organization
    2.  ⛔ change subscription plan
    3.  ⛔ cancel subscription
    4.  ⛔ resume subscription
19. organization/billing/invoice
    1.  ⛔ list invoices by organization
    2.  ⛔ get invoice details
    3.  ⛔ download invoice
20. organization/billing/payment
    1.  ⛔ list payments by invoice
    2.  ⛔ get payment details
21. asset/artifact
    1.  ⛔ list artifacts
    2.  ⛔ get artifact details
22. asset/release
    1.  ⛔ list available releases
    2.  ⛔ get release details
    3.  ⛔ mark release recommended/stable
    4.  ⛔ register/upload release 
23. pipeline
    1.  ⛔ request deploy (organization, target release)
    2.  ⛔ request update config (organization, payload)
    3.  ⛔ request redeploy (clone from previous pipeline)
    4.  ⛔ cancel pipeline
    5.  ⛔ get pipeline details
    6.  ⛔ get pipeline current status summary
    7.  ⛔ list pipelines by organization
24. pipeline/action
    1.  ⛔ get pipeline action summary
    2.  ⛔ list pipeline actions
25. pipeline/action/stage
    1.  ⛔ list stages by action
    2.  ⛔ get stage details
    3.  ⛔ get stage output metadata
    4.  ⛔ download stage output
    5.  ⛔ purge logs/output (admin-only)
26. me/notification
    1.  ⛔ list my notifications (with filters)
    2.  ⛔ get notification
    3.  ⛔ mark notification as read
    4.  ⛔ mark all notifications as read
    5.  ⛔ get unread notification count
    6.  ⛔ publish notification (agent -> control)
27. me/activity
    1.  ⛔ list my activities (filters: kind, date range, cluster)
    2.  ⛔ get activity
    3.  ⛔ export activities
    4.  ⛔ append activity (agent -> control)
28. document
    1.  ⛔ create document (organization, upload file, kind, title, metadata)
    2.  ⛔ get document
    3.  ⛔ download document
    4.  ⛔ search documents (filters: kind/status/title/date, organization)
    5.  ⛔ delete document (soft delete)
    6.  ⛔ replace document (reemit document, invalidate previous)
    7.  ⛔ get document signatures
    8.  ⛔ get document validity
    9.  ⛔ create document from agent (agent -> control)
    10. ⛔ delete document from agent (agent -> control)
    11. ⛔ replace document from agent (agent -> control)
29. signature
    1.  ⛔ create signature request (document, account)
    2.  ⛔ list signatures by document
    3.  ⛔ get signature details
    4.  ⛔ verify signature (internal/optional)
    5.  ⛔ list my signatures
    6.  ⛔ create signature request from agent (agent -> control)
30. my/document
    1.  ⛔ list my documents (where I am signer or requester)
    2.  ⛔ get my document (must be related)
    3.  ⛔ download my document (must be related)
    4.  ⛔ get my document signatures
    5.  ⛔ get my document validity