// domain/deployment/enum/release.go
package enum

type ReleaseChannel string

const (
	ReleaseChannel_STABLE ReleaseChannel = "STABLE"
	ReleaseChannel_BETA   ReleaseChannel = "BETA"
	ReleaseChannel_ALPHA  ReleaseChannel = "ALPHA"
)
