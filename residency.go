package marathon

type TaskLostBehaviorType string

const (
	TaskLostBehaviorTypeWaitForever          TaskLostBehaviorType = "WAIT_FOREVER"
	TaskLostBehaviorTypeRelaunchAfterTimeout TaskLostBehaviorType = "RELAUNCH_AFTER_TIMEOUT"
)

type Residency struct {
	TaskLostBehavior                 TaskLostBehaviorType `json:"taskLostBehavior,omitempty"`
	RelaunchEscalationTimeoutSeconds int                  `json:"relaunchEscalationTimeoutSeconds,omitempty"`
}

func (r *Residency) SetTaskLostBehavior(behavior TaskLostBehaviorType) *Residency {
	r.TaskLostBehavior = behavior
	return r
}

func (r *Residency) SetRelaunchEscalationTimeoutSeconds(seconds int) *Residency {
	r.RelaunchEscalationTimeoutSeconds = seconds
	return r
}
