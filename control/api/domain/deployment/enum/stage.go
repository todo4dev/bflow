// domain/deployment/enum/stage.go
package enum

type StageStatus string

const (
	StageStatus_PENDING StageStatus = "PENDING"
	StageStatus_RUNNING StageStatus = "RUNNING"
	StageStatus_SUCCESS StageStatus = "SUCCESS"
	StageStatus_FAILED  StageStatus = "FAILED"
)
