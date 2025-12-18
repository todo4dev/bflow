This document describes the technical specifications of the **Control Plane** artifact, detailing the functionality of the **Authentication** module and the specific use case for verifying the availability of an email address in the system.


# Artifact - Control Plane

## Module - Authentication

### UseCase - Check if email already in use

This use case aims to verify the availability of an email address before proceeding with registration or profile updates.

1. **Input**: The system receives an email address.
2. **Validation**: The email format is syntactically validated.
3. **Query**: A search is performed in the database for existing accounts with the provided email.
4. **Output**: Returns a boolean value or status indicating whether the email is already linked to an active or pending account.

---

### UseCase - Register account using email

This use case aims to create a new account using an email address and password.

1. **Input**: The system receives an email address and a password.
2. **Validation**: The email format is syntactically validated and the password must have one lower, one upper, one digit and minimal 8 characters.
3. **Command**: A search is performed in the database for existing accounts with the provided email. When doesn't exist, a new account is created and an activation email is sent.
4. **Output**: Returns a boolean value or status indicating whether the email is already linked to an active or pending account.

---

### UseCase - Activate account

This use case aims to activate an account using an email address and activation code.

1. **Input**: The system receives an email address and an activation code.
2. **Validation**: The email format is syntactically validated and the activation code must be valid.
3. **Command**: A search is performed in the database for existing accounts with the provided email.
4. **Output**: Returns a boolean value or status indicating whether the email is already linked to an active or pending account.

---

### UseCase - Resend email to activate account

This use case aims to resend an activation email to an account using an email address and activation code.

1. **Input**: The system receives an email address and an activation code.
2. **Validation**: The email format is syntactically validated and the activation code must be valid.
3. **Command**: A search is performed in the database for existing accounts with the provided email.
4. **Output**: Returns a boolean value or status indicating whether the email is already linked to an active or pending account.

---

### UseCase - Login using email and password

This use case aims to login an account using an email address and password.

1. **Input**: The system receives an email address and a password.
2. **Validation**: The email and password are required.
3. **Command**: A search is performed in the database for existing accounts with the provided email. When exists and account is active, the password is verified.
4. **Output**: Returns an authentication token with access_token and refresh_token.

---

### UseCase - Login using email otp

This use case aims to login an account using an email address and otp.

1. **Input**: The system receives an email address and an otp.
2. **Validation**: The email and otp are required.
3. **Command**: A search is performed in the database for existing accounts with the provided email. When exists and account is active, the otp is verified.
4. **Output**: Returns an authentication token with access_token and refresh_token.

---

### UseCase - Redirect to SSO provider

This use case aims to redirect to SSO provider.

1. **Input**: The system receives the sso provider and callback url.
2. **Validation**: The sso provider and callback url are required.
3. **Query**: The sso provider url are mount using environment variables and the sso provider, setting callback url as state.
4. **Output**: Returns an URL to redirect to SSO provider.

---

### UseCase - Callback from SSO provider

This use case aims to callback from SSO provider.

1. **Input**: The system receives the sso provider, state and code.
2. **Validation**: The sso provider, state and code are required.
3. **Command**: A search is performed in the database for existing accounts with the provided email. When exists, account is forced to be active and profile + picture info is updated. When doesn't exists, a new account is created (already activated) with profile and picture info.
4. **Output**: Returns the callback uri contained in state with internal token used to login using sso token.

---

### UseCase - Login using SSO token

This use case aims to login an account using an internal token.

1. **Input**: The system receives an internal token.
2. **Validation**: The internal token is required.
3. **Command**: Token is validated. When account related with token exists and is active, the token is used to login.
4. **Output**: Returns an authentication token with access_token and refresh_token.

---

### UseCase - Resend email to reset password

This use case aims to resend an email to reset password.

1. **Input**: The system receives an email address.
2. **Validation**: The email address is required.
3. **Command**: A search is performed in the database for existing accounts with the provided email. When exists, an email is sent to the provided email address.
4. **Output**: Returns a boolean value or status indicating whether the email is already linked to an active or pending account.

---

### UseCase - Reset password

This use case aims to reset password.

1. **Input**: The system receives an email address and a password.
2. **Validation**: The email address and password are required.
3. **Command**: A search is performed in the database for existing accounts with the provided email. When exists, the password is updated.
4. **Output**: Returns a boolean value or status indicating whether the email is already linked to an active or pending account.

---

## Module - Session

### UseCase - logout

This use case aims to logout an account.

1. **Input**: The system receives an authentication token.
2. **Validation**: The authentication token is required.
3. **Command**: The authentication token is validated. When valid, the session is closed.
4. **Output**: Returns a boolean value or status indicating whether the authentication token is valid.

---

### UseCase - Refresh authentication token

This use case aims to refresh an authentication token.

1. **Input**: The system receives an authentication token.
2. **Validation**: The authentication token is required.
3. **Command**: The authentication token is validated. When valid, a new authentication token is generated.
4. **Output**: Returns a boolean value or status indicating whether the authentication token is valid.

---

### UseCase - Update profile

This use case aims to update a profile.

1. **Input**: The system receives an authentication token and a profile.
2. **Validation**: The authentication token and profile are required.
3. **Command**: The authentication token is validated. When valid, the profile is updated.
4. **Output**: Returns a boolean value or status indicating whether the authentication token is valid.

---

### UseCase - Get picture

This use case aims to get a picture.

1. **Input**: The system receives an authentication token and a picture id.
2. **Validation**: The authentication token and picture id are required.
3. **Command**: The authentication token is validated. When valid, the picture is retrieved.
4. **Output**: Returns the picture.

---

### UseCase - Upload picture

This use case aims to upload a picture.

1. **Input**: The system receives an authentication token and a picture.
2. **Validation**: The authentication token and picture are required.
3. **Command**: The authentication token is validated. When valid, the picture is uploaded.
4. **Output**: Returns true or false indicating whether the picture was uploaded successfully.

---

### UseCase - Delete picture

This use case aims to delete a picture.

1. **Input**: The system receives an authentication token and a picture id.
2. **Validation**: The authentication token and picture id are required.
3. **Command**: The authentication token is validated. When valid, the picture is deleted.
4. **Output**: Returns true or false indicating whether the picture was deleted successfully.

---

### UseCase - Delete account

This use case aims to delete an account.

1. **Input**: The system receives an authentication token.
2. **Validation**: The authentication token is required.
3. **Command**: The authentication token is validated. When valid, the account is deleted and all sessions related are revoked.
4. **Output**: Returns true or false indicating whether the account was deleted successfully.

---

### UseCase - Search signed document

This use case aims to search signed documents.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the signed documents are searched.
4. **Output**: Returns the signed documents in paginated format.

---

### UseCase - Search activity

This use case aims to search activities.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the activities are searched.
4. **Output**: Returns the activities in paginated format.

---

### UseCase - Search notification

This use case aims to search notifications.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the notifications are searched.
4. **Output**: Returns the notifications in paginated format.

---

## Module - System

### UseCase - Healthcheck

This use case aims to check the health of the system.

1. **Input**: The system receives an authentication token.
2. **Validation**: The authentication token is required.
3. **Command**: The authentication token is validated. When valid, the system health is checked.
4. **Output**: Returns true or false indicating whether the system is healthy.

---

## Module - Tenant

### UseCase - Check if subdomain already in use

This use case aims to check if a subdomain is already in use.

1. **Input**: The system receives a subdomain.
2. **Validation**: The subdomain is required.
3. **Command**: The subdomain is checked.
4. **Output**: Returns true or false indicating whether the subdomain is already in use.

---

### UseCase - Update tenant

This use case aims to update a tenant.

1. **Input**: The system receives an authentication token and a tenant.
2. **Validation**: The authentication token and tenant are required.
3. **Command**: The authentication token is validated. When valid, the tenant is updated.
4. **Output**: Returns true or false indicating whether the tenant was updated successfully.

---

### UseCase - Delete tenant

This use case aims to delete a tenant.

1. **Input**: The system receives an authentication token and a tenant id.
2. **Validation**: The authentication token and tenant id are required.
3. **Command**: The authentication token is validated. When valid, the tenant is deleted.
4. **Output**: Returns true or false indicating whether the tenant was deleted successfully.

---

### UseCase - Search tenant

This use case aims to search tenants.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the tenants are searched.
4. **Output**: Returns the tenants in paginated format.

---

### UseCase - Get tenant picture

This use case aims to get a tenant picture.

1. **Input**: The system receives an authentication token and a tenant id.
2. **Validation**: The authentication token and tenant id are required.
3. **Command**: The authentication token is validated. When valid, the tenant picture is retrieved.
4. **Output**: Returns the tenant picture.

---

### UseCase - Upload tenant picture

This use case aims to upload a tenant picture.

1. **Input**: The system receives an authentication token and a tenant picture.
2. **Validation**: The authentication token and tenant picture are required.
3. **Command**: The authentication token is validated. When valid, the tenant picture is uploaded.
4. **Output**: Returns true or false indicating whether the tenant picture was uploaded successfully.

---

### UseCase - Delete tenant picture

This use case aims to delete a tenant picture.

1. **Input**: The system receives an authentication token and a tenant id.
2. **Validation**: The authentication token and tenant id are required.
3. **Command**: The authentication token is validated. When valid, the tenant picture is deleted.
4. **Output**: Returns true or false indicating whether the tenant picture was deleted successfully.

---

### UseCase - Get tenant configuration

This use case aims to get a tenant configuration.

1. **Input**: The system receives an authentication token and a tenant id.
2. **Validation**: The authentication token and tenant id are required.
3. **Command**: The authentication token is validated. When valid, the tenant configuration is retrieved.
4. **Output**: Returns the tenant configuration.

---

### UseCase - Update tenant configuration

This use case aims to update a tenant configuration.

1. **Input**: The system receives an authentication token and a tenant configuration.
2. **Validation**: The authentication token and tenant configuration are required.
3. **Command**: The authentication token is validated. When valid, the tenant configuration is updated.
4. **Output**: Returns true or false indicating whether the tenant configuration was updated successfully.

---

### UseCase - List tenant memberships

This use case aims to list tenant memberships.

1. **Input**: The system receives an authentication token and a tenant id.
2. **Validation**: The authentication token and tenant id are required.
3. **Command**: The authentication token is validated. When valid, the tenant memberships are listed.
4. **Output**: Returns the tenant memberships in paginated format.

---

### UseCase - Add membership to tenant

This use case aims to add a membership to a tenant.

1. **Input**: The system receives an authentication token and a membership.
2. **Validation**: The authentication token and membership are required.
3. **Command**: The authentication token is validated. When valid, the membership is added to the tenant.
4. **Output**: Returns true or false indicating whether the membership was added successfully.

---

### UseCase - Update membership role in tenant

This use case aims to update a membership role in a tenant.

1. **Input**: The system receives an authentication token and a membership.
2. **Validation**: The authentication token and membership are required.
3. **Command**: The authentication token is validated. When valid, the membership role is updated in the tenant.
4. **Output**: Returns true or false indicating whether the membership role was updated successfully.

---

### UseCase - Remove membership from tenant

This use case aims to remove a membership from a tenant.

1. **Input**: The system receives an authentication token and a membership.
2. **Validation**: The authentication token and membership are required.
3. **Command**: The authentication token is validated. When valid, the membership is removed from the tenant.
4. **Output**: Returns true or false indicating whether the membership was removed successfully.

---

## Module - Billing

### UseCase - Manage subscription

This use case aims to manage a subscription.

1. **Input**: The system receives an authentication token and a subscription.
2. **Validation**: The authentication token and subscription are required.
3. **Command**: The authentication token is validated. When valid, the subscription is managed.
4. **Output**: Returns true or false indicating whether the subscription was managed successfully.

---

### UseCase - Search invoices

This use case aims to search invoices.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the invoices are searched.
4. **Output**: Returns the invoices in paginated format.

---

### UseCase - Search payments

This use case aims to search payments.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the payments are searched.
4. **Output**: Returns the payments in paginated format.

---

# Artifact - Data Plane

## Module - Access Control

### UseCase - Resolve tenant context from request

This use case aims to resolve a tenant context from a request.

1. **Input**: The system receives a request.
2. **Validation**: The request is required.
3. **Command**: The request is resolved.
4. **Output**: Returns the tenant context.

---

### UseCase - Validate access token

This use case aims to validate an access token.

1. **Input**: The system receives an access token.
2. **Validation**: The access token is required.
3. **Command**: The access token is validated.
4. **Output**: Returns true or false indicating whether the access token is valid.

---

### UseCase - Get current account context

This use case aims to get the current account context.

1. **Input**: The system receives an authentication token.
2. **Validation**: The authentication token is required.
3. **Command**: The authentication token is validated. When valid, the account context is retrieved.
4. **Output**: Returns the account context.

---

### UseCase - Get current membership context

This use case aims to get the current membership context.

1. **Input**: The system receives an authentication token.
2. **Validation**: The authentication token is required.
3. **Command**: The authentication token is validated. When valid, the membership context is retrieved.
4. **Output**: Returns the membership context.

---

### UseCase - Resolve project context from request

This use case aims to resolve a project context from a request.

1. **Input**: The system receives a request.
2. **Validation**: The request is required.
3. **Command**: The request is resolved.
4. **Output**: Returns the project context.

---

### UseCase - Build request security context

This use case aims to build a request security context.

1. **Input**: The system receives a request.
2. **Validation**: The request is required.
3. **Command**: The request is resolved.
4. **Output**: Returns the request security context.

---

## Module - RBAC

### UseCase - List permission catalog

This use case aims to list the permission catalog.

1. **Input**: The system receives an authentication token.
2. **Validation**: The authentication token is required.
3. **Command**: The authentication token is validated. When valid, the permission catalog is listed.
4. **Output**: Returns the permission catalog.

---

### UseCase - Create role

This use case aims to create a role.

1. **Input**: The system receives an authentication token and a role.
2. **Validation**: The authentication token and role are required.
3. **Command**: The authentication token is validated. When valid, the role is created.
4. **Output**: Returns true or false indicating whether the role was created successfully.

---

### UseCase - Update role

This use case aims to update a role.

1. **Input**: The system receives an authentication token and a role.
2. **Validation**: The authentication token and role are required.
3. **Command**: The authentication token is validated. When valid, the role is updated.
4. **Output**: Returns true or false indicating whether the role was updated successfully.

---

### UseCase - Delete role

This use case aims to delete a role.

1. **Input**: The system receives an authentication token and a role.
2. **Validation**: The authentication token and role are required.
3. **Command**: The authentication token is validated. When valid, the role is deleted.
4. **Output**: Returns true or false indicating whether the role was deleted successfully.

---

### UseCase - Get role

This use case aims to get a role.

1. **Input**: The system receives an authentication token and a role id.
2. **Validation**: The authentication token and role id are required.
3. **Command**: The authentication token is validated. When valid, the role is retrieved.
4. **Output**: Returns the role.

---

### UseCase - Search role

This use case aims to search roles.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the roles are searched.
4. **Output**: Returns the roles in paginated format.

---

### UseCase - Create group

This use case aims to create a group.

1. **Input**: The system receives an authentication token and a group.
2. **Validation**: The authentication token and group are required.
3. **Command**: The authentication token is validated. When valid, the group is created.
4. **Output**: Returns true or false indicating whether the group was created successfully.

---

### UseCase - Update group

This use case aims to update a group.

1. **Input**: The system receives an authentication token and a group.
2. **Validation**: The authentication token and group are required.
3. **Command**: The authentication token is validated. When valid, the group is updated.
4. **Output**: Returns true or false indicating whether the group was updated successfully.

---

### UseCase - Delete group

This use case aims to delete a group.

1. **Input**: The system receives an authentication token and a group.
2. **Validation**: The authentication token and group are required.
3. **Command**: The authentication token is validated. When valid, the group is deleted.
4. **Output**: Returns true or false indicating whether the group was deleted successfully.

---

### UseCase - Get group

This use case aims to get a group.

1. **Input**: The system receives an authentication token and a group id.
2. **Validation**: The authentication token and group id are required.
3. **Command**: The authentication token is validated. When valid, the group is retrieved.
4. **Output**: Returns the group.

---

### UseCase - Search group

This use case aims to search groups.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the groups are searched.
4. **Output**: Returns the groups in paginated format.

---

### UseCase - Add membership to group

This use case aims to add a membership to a group.

1. **Input**: The system receives an authentication token and a membership.
2. **Validation**: The authentication token and membership are required.
3. **Command**: The authentication token is validated. When valid, the membership is added to the group.
4. **Output**: Returns true or false indicating whether the membership was added successfully.

---

### UseCase - Remove membership from group

This use case aims to remove a membership from a group.

1. **Input**: The system receives an authentication token and a membership.
2. **Validation**: The authentication token and membership are required.
3. **Command**: The authentication token is validated. When valid, the membership is removed from the group.
4. **Output**: Returns true or false indicating whether the membership was removed successfully.

---

### UseCase - List membership groups

This use case aims to list the membership groups.

1. **Input**: The system receives an authentication token and a membership id.
2. **Validation**: The authentication token and membership id are required.
3. **Command**: The authentication token is validated. When valid, the membership groups are listed.
4. **Output**: Returns the membership groups.

---

### UseCase - Evaluate effective permissions for membership

This use case aims to evaluate the effective permissions for a membership.

1. **Input**: The system receives an authentication token and a membership id.
2. **Validation**: The authentication token and membership id are required.
3. **Command**: The authentication token is validated. When valid, the effective permissions for the membership are evaluated.
4. **Output**: Returns the effective permissions for the membership.

---

### UseCase - Evaluate effective permissions for membership in project

This use case aims to evaluate the effective permissions for a membership in a project.

1. **Input**: The system receives an authentication token and a membership id and a project id.
2. **Validation**: The authentication token and membership id and project id are required.
3. **Command**: The authentication token is validated. When valid, the effective permissions for the membership in the project are evaluated.
4. **Output**: Returns the effective permissions for the membership in the project.

---

## Module - System

### UseCase - Healthcheck

This use case aims to perform a health check on the system.

1. **Input**: The system receives an authentication token.
2. **Validation**: The authentication token is required.
3. **Command**: The authentication token is validated. When valid, the health check is performed.
4. **Output**: Returns true or false indicating whether the health check was successful.

---

## Module - Cost Center

### UseCase - Create cost center

This use case aims to create a cost center.

1. **Input**: The system receives an authentication token and a cost center.
2. **Validation**: The authentication token and cost center are required.
3. **Command**: The authentication token is validated. When valid, the cost center is created.
4. **Output**: Returns true or false indicating whether the cost center was created successfully.

---

### UseCase - Update cost center

This use case aims to update a cost center.

1. **Input**: The system receives an authentication token and a cost center.
2. **Validation**: The authentication token and cost center are required.
3. **Command**: The authentication token is validated. When valid, the cost center is updated.
4. **Output**: Returns true or false indicating whether the cost center was updated successfully.

---

### UseCase - Delete cost center

This use case aims to delete a cost center.

1. **Input**: The system receives an authentication token and a cost center.
2. **Validation**: The authentication token and cost center are required.
3. **Command**: The authentication token is validated. When valid, the cost center is deleted.
4. **Output**: Returns true or false indicating whether the cost center was deleted successfully.

---

### UseCase - Get cost center

This use case aims to get a cost center.

1. **Input**: The system receives an authentication token and a cost center id.
2. **Validation**: The authentication token and cost center id are required.
3. **Command**: The authentication token is validated. When valid, the cost center is retrieved.
4. **Output**: Returns the cost center.

---

### UseCase - Search cost center

This use case aims to search cost centers.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the cost centers are searched.
4. **Output**: Returns the cost centers in paginated format.

---

## Module - Site

### UseCase - Create site

This use case aims to create a site.

1. **Input**: The system receives an authentication token and a site.
2. **Validation**: The authentication token and site are required.
3. **Command**: The authentication token is validated. When valid, the site is created.
4. **Output**: Returns true or false indicating whether the site was created successfully.

---

### UseCase - Update site

This use case aims to update a site.

1. **Input**: The system receives an authentication token and a site.
2. **Validation**: The authentication token and site are required.
3. **Command**: The authentication token is validated. When valid, the site is updated.
4. **Output**: Returns true or false indicating whether the site was updated successfully.

---

### UseCase - Delete site

This use case aims to delete a site.

1. **Input**: The system receives an authentication token and a site.
2. **Validation**: The authentication token and site are required.
3. **Command**: The authentication token is validated. When valid, the site is deleted.
4. **Output**: Returns true or false indicating whether the site was deleted successfully.

---

### UseCase - Get site

This use case aims to get a site.

1. **Input**: The system receives an authentication token and a site id.
2. **Validation**: The authentication token and site id are required.
3. **Command**: The authentication token is validated. When valid, the site is retrieved.
4. **Output**: Returns the site.

---

### UseCase - Search site

This use case aims to search sites.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the sites are searched.
4. **Output**: Returns the sites in paginated format.

---

## Module - Department

### UseCase - Create department

This use case aims to create a department.

1. **Input**: The system receives an authentication token and a department.
2. **Validation**: The authentication token and department are required.
3. **Command**: The authentication token is validated. When valid, the department is created.
4. **Output**: Returns true or false indicating whether the department was created successfully.

---

### UseCase - Update department

This use case aims to update a department.

1. **Input**: The system receives an authentication token and a department.
2. **Validation**: The authentication token and department are required.
3. **Command**: The authentication token is validated. When valid, the department is updated.
4. **Output**: Returns true or false indicating whether the department was updated successfully.

---

### UseCase - Delete department

This use case aims to delete a department.

1. **Input**: The system receives an authentication token and a department.
2. **Validation**: The authentication token and department are required.
3. **Command**: The authentication token is validated. When valid, the department is deleted.
4. **Output**: Returns true or false indicating whether the department was deleted successfully.

---

### UseCase - Get department

This use case aims to get a department.

1. **Input**: The system receives an authentication token and a department id.
2. **Validation**: The authentication token and department id are required.
3. **Command**: The authentication token is validated. When valid, the department is retrieved.
4. **Output**: Returns the department.

---

### UseCase - Search department

This use case aims to search departments.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the departments are searched.
4. **Output**: Returns the departments in paginated format.

---

## Module - Customer

### UseCase - Create customer

This use case aims to create a customer.

1. **Input**: The system receives an authentication token and a customer.
2. **Validation**: The authentication token and customer are required.
3. **Command**: The authentication token is validated. When valid, the customer is created.
4. **Output**: Returns true or false indicating whether the customer was created successfully.

---

### UseCase - Update customer

This use case aims to update a customer.

1. **Input**: The system receives an authentication token and a customer.
2. **Validation**: The authentication token and customer are required.
3. **Command**: The authentication token is validated. When valid, the customer is updated.
4. **Output**: Returns true or false indicating whether the customer was updated successfully.

---

### UseCase - Delete customer

This use case aims to delete a customer.

1. **Input**: The system receives an authentication token and a customer.
2. **Validation**: The authentication token and customer are required.
3. **Command**: The authentication token is validated. When valid, the customer is deleted.
4. **Output**: Returns true or false indicating whether the customer was deleted successfully.

---

### UseCase - Get customer

This use case aims to get a customer.

1. **Input**: The system receives an authentication token and a customer id.
2. **Validation**: The authentication token and customer id are required.
3. **Command**: The authentication token is validated. When valid, the customer is retrieved.
4. **Output**: Returns the customer.

---

### UseCase - Search customer

This use case aims to search customers.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the customers are searched.
4. **Output**: Returns the customers in paginated format.

---

## Module - Supplier

### UseCase - Create supplier

This use case aims to create a supplier.

1. **Input**: The system receives an authentication token and a supplier.
2. **Validation**: The authentication token and supplier are required.
3. **Command**: The authentication token is validated. When valid, the supplier is created.
4. **Output**: Returns true or false indicating whether the supplier was created successfully.

---

### UseCase - Update supplier

This use case aims to update a supplier.

1. **Input**: The system receives an authentication token and a supplier.
2. **Validation**: The authentication token and supplier are required.
3. **Command**: The authentication token is validated. When valid, the supplier is updated.
4. **Output**: Returns true or false indicating whether the supplier was updated successfully.

### UseCase - Delete supplier

This use case aims to delete a supplier.

1. **Input**: The system receives an authentication token and a supplier.
2. **Validation**: The authentication token and supplier are required.
3. **Command**: The authentication token is validated. When valid, the supplier is deleted.
4. **Output**: Returns true or false indicating whether the supplier was deleted successfully.

---

### UseCase - Get supplier

This use case aims to get a supplier.

1. **Input**: The system receives an authentication token and a supplier id.
2. **Validation**: The authentication token and supplier id are required.
3. **Command**: The authentication token is validated. When valid, the supplier is retrieved.
4. **Output**: Returns the supplier.

---

### UseCase - Search supplier

This use case aims to search suppliers.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the suppliers are searched.
4. **Output**: Returns the suppliers in paginated format.

---

## Module - Project

### UseCase - Create project

This use case aims to create a project.

1. **Input**: The system receives an authentication token and a project.
2. **Validation**: The authentication token and project are required.
3. **Command**: The authentication token is validated. When valid, the project is created.
4. **Output**: Returns true or false indicating whether the project was created successfully.

---

### UseCase - Update project

This use case aims to update a project.

1. **Input**: The system receives an authentication token and a project.
2. **Validation**: The authentication token and project are required.
3. **Command**: The authentication token is validated. When valid, the project is updated.
4. **Output**: Returns true or false indicating whether the project was updated successfully.

---

### UseCase - Delete project

This use case aims to delete a project.

1. **Input**: The system receives an authentication token and a project.
2. **Validation**: The authentication token and project are required.
3. **Command**: The authentication token is validated. When valid, the project is deleted.
4. **Output**: Returns true or false indicating whether the project was deleted successfully.

---

### UseCase - Get project

This use case aims to get a project.

1. **Input**: The system receives an authentication token and a project id.
2. **Validation**: The authentication token and project id are required.
3. **Command**: The authentication token is validated. When valid, the project is retrieved.
4. **Output**: Returns the project.

---

### UseCase - Search project

This use case aims to search projects.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the projects are searched.
4. **Output**: Returns the projects in paginated format.

---

## Module - Contract

### UseCase - Create contract

This use case aims to create a contract.

1. **Input**: The system receives an authentication token and a contract.
2. **Validation**: The authentication token and contract are required.
3. **Command**: The authentication token is validated. When valid, the contract is created.
4. **Output**: Returns true or false indicating whether the contract was created successfully.

---

### UseCase - Update contract

This use case aims to update a contract.

1. **Input**: The system receives an authentication token and a contract.
2. **Validation**: The authentication token and contract are required.
3. **Command**: The authentication token is validated. When valid, the contract is updated.
4. **Output**: Returns true or false indicating whether the contract was updated successfully.

---

### UseCase - Delete contract

This use case aims to delete a contract.

1. **Input**: The system receives an authentication token and a contract.
2. **Validation**: The authentication token and contract are required.
3. **Command**: The authentication token is validated. When valid, the contract is deleted.
4. **Output**: Returns true or false indicating whether the contract was deleted successfully.

---

### UseCase - Get contract

This use case aims to get a contract.

1. **Input**: The system receives an authentication token and a contract id.
2. **Validation**: The authentication token and contract id are required.
3. **Command**: The authentication token is validated. When valid, the contract is retrieved.
4. **Output**: Returns the contract.

---

### UseCase - Search contract

This use case aims to search contracts.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the contracts are searched.
4. **Output**: Returns the contracts in paginated format.

---

## Module - Resource

### UseCase - Create resource

This use case aims to create a resource.

1. **Input**: The system receives an authentication token and a resource.
2. **Validation**: The authentication token and resource are required.
3. **Command**: The authentication token is validated. When valid, the resource is created.
4. **Output**: Returns true or false indicating whether the resource was created successfully.

---

### UseCase - Update resource

This use case aims to update a resource.

1. **Input**: The system receives an authentication token and a resource.
2. **Validation**: The authentication token and resource are required.
3. **Command**: The authentication token is validated. When valid, the resource is updated.
4. **Output**: Returns true or false indicating whether the resource was updated successfully.

---

### UseCase - Delete resource

This use case aims to delete a resource.

1. **Input**: The system receives an authentication token and a resource.
2. **Validation**: The authentication token and resource are required.
3. **Command**: The authentication token is validated. When valid, the resource is deleted.
4. **Output**: Returns true or false indicating whether the resource was deleted successfully.

---

### UseCase - Get resource

This use case aims to get a resource.

1. **Input**: The system receives an authentication token and a resource id.
2. **Validation**: The authentication token and resource id are required.
3. **Command**: The authentication token is validated. When valid, the resource is retrieved.
4. **Output**: Returns the resource.

---

### UseCase - Search resource

This use case aims to search resources.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the resources are searched.
4. **Output**: Returns the resources in paginated format.

---

## Module - Execution Order

### UseCase - Create execution order

This use case aims to create an execution order.

1. **Input**: The system receives an authentication token and an execution order.
2. **Validation**: The authentication token and execution order are required.
3. **Command**: The authentication token is validated. When valid, the execution order is created.
4. **Output**: Returns true or false indicating whether the execution order was created successfully.

---

### UseCase - Update execution order details

This use case aims to update an execution order.

1. **Input**: The system receives an authentication token and an execution order.
2. **Validation**: The authentication token and execution order are required.
3. **Command**: The authentication token is validated. When valid, the execution order is updated.
4. **Output**: Returns true or false indicating whether the execution order was updated successfully.

---

### UseCase - Initialize execution order

This use case aims to initialize an execution order.

1. **Input**: The system receives an authentication token and an execution order.
2. **Validation**: The authentication token and execution order are required.
3. **Command**: The authentication token is validated. When valid, the execution order is initialized.
4. **Output**: Returns true or false indicating whether the execution order was initialized successfully.

---

### UseCase - Cancel execution order

This use case aims to cancel an execution order.

1. **Input**: The system receives an authentication token and an execution order.
2. **Validation**: The authentication token and execution order are required.
3. **Command**: The authentication token is validated. When valid, the execution order is canceled.
4. **Output**: Returns true or false indicating whether the execution order was canceled successfully.

---

### UseCase - Add execution record into execution order

This use case aims to add an execution record into an execution order.

1. **Input**: The system receives an authentication token and an execution record.
2. **Validation**: The authentication token and execution record are required.
3. **Command**: The authentication token is validated. When valid, the execution record is added into the execution order.
4. **Output**: Returns true or false indicating whether the execution record was added successfully.

---

### UseCase - Update execution record on execution order

This use case aims to update an execution record on an execution order.

1. **Input**: The system receives an authentication token and an execution record.
2. **Validation**: The authentication token and execution record are required.
3. **Command**: The authentication token is validated. When valid, the execution record is updated on the execution order.
4. **Output**: Returns true or false indicating whether the execution record was updated successfully.

---

### UseCase - Remove execution record from execution order

This use case aims to remove an execution record from an execution order.

1. **Input**: The system receives an authentication token and an execution record.
2. **Validation**: The authentication token and execution record are required.
3. **Command**: The authentication token is validated. When valid, the execution record is removed from the execution order.
4. **Output**: Returns true or false indicating whether the execution record was removed successfully.

---

### UseCase - Upload media into execution record

This use case aims to upload media into an execution record.

1. **Input**: The system receives an authentication token and a media.
2. **Validation**: The authentication token and media are required.
3. **Command**: The authentication token is validated. When valid, the media is uploaded into the execution record.
4. **Output**: Returns true or false indicating whether the media was uploaded successfully.

---

### UseCase - Remove media from execution record

This use case aims to remove media from an execution record.

1. **Input**: The system receives an authentication token and a media.
2. **Validation**: The authentication token and media are required.
3. **Command**: The authentication token is validated. When valid, the media is removed from the execution record.
4. **Output**: Returns true or false indicating whether the media was removed successfully.

---

### UseCase - Submit execution order for analysis

This use case aims to submit an execution order for analysis.

1. **Input**: The system receives an authentication token and an execution order.
2. **Validation**: The authentication token and execution order are required.
3. **Command**: The authentication token is validated. When valid, the execution order is submitted for analysis.
4. **Output**: Returns true or false indicating whether the execution order was submitted for analysis successfully.

---

### UseCase - Finish execution order

This use case aims to finish an execution order.

1. **Input**: The system receives an authentication token and an execution order.
2. **Validation**: The authentication token and execution order are required.
3. **Command**: The authentication token is validated. When valid, the execution order is finished.
4. **Output**: Returns true or false indicating whether the execution order was finished successfully.

---

### UseCase - Search execution order

This use case aims to search execution orders.

1. **Input**: The system receives an authentication token and a search filter.
2. **Validation**: The authentication token and search filter are required.
3. **Command**: The authentication token is validated. When valid, the execution orders are searched.
4. **Output**: Returns the execution orders in paginated format.

---

