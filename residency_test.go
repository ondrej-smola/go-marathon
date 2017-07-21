package marathon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResidency(t *testing.T) {
	app := NewDockerApplication()

	app = app.SetResidency(TaskLostBehaviorTypeWaitForever)

	if assert.NotNil(t, app.Residency) {
		res := app.Residency

		assert.Equal(t, res.TaskLostBehavior, TaskLostBehaviorTypeWaitForever)

		res.SetRelaunchEscalationTimeoutSeconds(60)
		assert.Equal(t, app.Residency.RelaunchEscalationTimeoutSeconds, 60)

		res.SetTaskLostBehavior(TaskLostBehaviorTypeRelaunchAfterTimeout)
		assert.Equal(t, res.TaskLostBehavior, TaskLostBehaviorTypeRelaunchAfterTimeout)
	}
}
