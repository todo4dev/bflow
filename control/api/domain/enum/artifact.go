// domain/enum/artifact.go
package enum

type ArtifactKind string

const (
	ArtifactKind_DEPLOYMENT ArtifactKind = "DEPLOYMENT"
	ArtifactKind_SERVICE    ArtifactKind = "SERVICE"
)

