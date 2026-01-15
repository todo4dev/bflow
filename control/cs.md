## action
### list actions by pipeline
### get action details
### get action logs

## action/stage
### list stages by action
### get stage details
### get stage output metadata
### download stage output
### purge logs/output (admin-only)

## activity
### list activities by cluster (filters)
### list activities by account (filters)
### list activities by resource (via metadata)
### get activity
### export activities (optional)

## agent
### create agent enrollment token (ui/admin)
### register agent (agent -> control)
### poll actions (agent -> control)
### submit action result (agent -> control) (status, logs, output_meta)

## agent/activity
### append activity (agent -> control)

## agent/document
### create document (agent -> control)
### delete document (agent -> control)
### replace document (agent -> control)
### create signature request (agent -> control) (document, account)

## agent/notification
### publish notification (agent -> control)
### list notification types supported (optional)

## artifact
### list artifacts
### get artifact details

## auth
### activate account using email and code
### reset password using email, new password and code
### check if email is available
### login using email and password
### refresh authorization token
### register account using email and password
### send reset password for email
### send resend activation code for email
### logoff

## auth/sso
### redirect to sso provider passing callback url as state
### callback from sso provider callbacking to initial url with internal code
### login using internal code

## document
### create document (cluster, upload file, kind, title, metadata)
### get document
### download document
### search documents (filters: kind/status/title/date, cluster optional)
### delete document (soft delete)
### replace document (reemit document, invalidate previous)
### get document signatures
### get document validity

## me/account
### get me account
### delete me account

## me/avatar
### get me avatar
### upload me avatar
### delete me avatar

## me/certificate
### list certificates
### create certificate
### get certificate
### delete certificate

## me/cover
### get me cover
### upload me cover
### delete me cover

## me/credential
### update me credential

## me/profile
### update me profile

## my/document
### list my documents (where I am signer or requester)
### get my document (must be related)
### download my document (must be related)
### get my document signatures
### get my document validity

## my/organization
### create organization
### update organization
### delete organization
### get organization
### search organizations

## my/organization/billing/invoice
### list invoices by organization
### get invoice details
### download invoice

## my/organization/billing/payment
### list payments by invoice
### get payment details

## my/organization/billing/plan
### list plans
### get plan details

## my/organization/billing/subscription
### get subscription by organization
### change subscription plan
### cancel subscription
### resume subscription

## my/organization/membership
### get my membership
### list members by organization
### add member to organization
### update member role (viewer/manager/admin)
### remove member from organization
### leave organization (self)
### transfer viewer/manager/admin (optional)

## my/organization/membership/invite
### create membership invite (organization, email, role)
### resend membership invite
### cancel membership invite
### accept membership invite (code)
### get membership invite by code
### list membership invites by organization

## my/signature
### sign document (document, certificate)
### list my signatures

## notification
### list notifications (by cluster/me)
### get notification
### mark notification as read
### mark all notifications as read
### delete notification (optional)
### update notification preferences
### get unread notification count

## pipeline
### request deploy (organization, target release)
### request update config (organization, payload)
### request redeploy (previous_pipeline_id)
### retry pipeline
### cancel pipeline
### get pipeline details
### get pipeline current status summary
### list pipelines by organization

## pipeline/action
### get pipeline action summary
### list pipeline actions

## release
### list available releases
### get release details
### mark release recommended/stable (optional)
### register/upload release (internal, optional)

## runtime
### list runtimes by organization
### get runtime (by id)
### get runtime status
### update runtime config
### set runtime readonly
### unset runtime readonly
### delete runtime
### get runtime current release
### set runtime current release (internal, on successful pipeline completion)

## signature
### create signature (document, account) (request signature slot)
### list signatures by document
### get signature details
### verify signature (internal/optional)

## system
### healthcheck
