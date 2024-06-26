// Code generated by "stringer -type=JobStateType --trimprefix=JobStateType --output job_string.go"; DO NOT EDIT.

package models

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[JobStateTypeUndefined-0]
	_ = x[JobStateTypePending-1]
	_ = x[JobStateTypeQueued-2]
	_ = x[JobStateTypeRunning-3]
	_ = x[JobStateTypeCompleted-4]
	_ = x[JobStateTypeFailed-5]
	_ = x[JobStateTypeStopped-6]
}

const _JobStateType_name = "UndefinedPendingQueuedRunningCompletedFailedStopped"

var _JobStateType_index = [...]uint8{0, 9, 16, 22, 29, 38, 44, 51}

func (i JobStateType) String() string {
	if i < 0 || i >= JobStateType(len(_JobStateType_index)-1) {
		return "JobStateType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _JobStateType_name[_JobStateType_index[i]:_JobStateType_index[i+1]]
}
