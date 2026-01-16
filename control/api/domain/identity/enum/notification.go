// domain/identity/enum/notification.go
package enum

type NotificationKind string

const (
	NotificationKind_SYSTEM   NotificationKind = "SYSTEM"
	NotificationKind_BILLING  NotificationKind = "BILLING"
	NotificationKind_SECURITY NotificationKind = "SECURITY"
	NotificationKind_RELEASE  NotificationKind = "RELEASE"
	NotificationKind_RUNTIME  NotificationKind = "RUNTIME"
	NotificationKind_PIPELINE NotificationKind = "PIPELINE"
	NotificationKind_AGENT    NotificationKind = "AGENT"
	NotificationKind_DOCUMENT NotificationKind = "DOCUMENT"
)

type NotificationLevel string

const (
	NotificationLevel_INFO    NotificationLevel = "INFO"
	NotificationLevel_WARNING NotificationLevel = "WARNING"
	NotificationLevel_ERROR   NotificationLevel = "ERROR"
)

type NotificationStatus string

const (
	NotificationStatus_ACTIVE    NotificationStatus = "ACTIVE"
	NotificationStatus_DISMISSED NotificationStatus = "DISMISSED"
)
