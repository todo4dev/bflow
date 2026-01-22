// domain/enum/account_notification.go
package enum

type AccountNotificationKind string

const (
	AccountNotificationKind_SYSTEM   AccountNotificationKind = "SYSTEM"
	AccountNotificationKind_BILLING  AccountNotificationKind = "BILLING"
	AccountNotificationKind_SECURITY AccountNotificationKind = "SECURITY"
	AccountNotificationKind_RELEASE  AccountNotificationKind = "RELEASE"
	AccountNotificationKind_RUNTIME  AccountNotificationKind = "RUNTIME"
	AccountNotificationKind_PIPELINE AccountNotificationKind = "PIPELINE"
	AccountNotificationKind_AGENT    AccountNotificationKind = "AGENT"
	AccountNotificationKind_DOCUMENT AccountNotificationKind = "DOCUMENT"
)

type AccountNotificationLevel string

const (
	AccountNotificationLevel_INFO    AccountNotificationLevel = "INFO"
	AccountNotificationLevel_WARNING AccountNotificationLevel = "WARNING"
	AccountNotificationLevel_ERROR   AccountNotificationLevel = "ERROR"
)

type AccountNotificationStatus string

const (
	AccountNotificationStatus_ACTIVE    AccountNotificationStatus = "ACTIVE"
	AccountNotificationStatus_DISMISSED AccountNotificationStatus = "DISMISSED"
)

