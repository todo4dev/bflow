// domain/enum/artifact_release.go
package enum

type ArtifactReleaseChannel string

const (
	ArtifactReleaseChannel_STABLE ArtifactReleaseChannel = "STABLE"
	ArtifactReleaseChannel_BETA   ArtifactReleaseChannel = "BETA"
	ArtifactReleaseChannel_ALPHA  ArtifactReleaseChannel = "ALPHA"
)

