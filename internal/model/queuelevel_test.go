package model

import "testing"

func TestQueueName_PriorityEightIsHigh(t *testing.T) {
	task := &Task{Spec: TaskSpec{Priority: 8}}
	if got := task.QueueName(); got != "high" {
		t.Fatalf("QueueName(priority=8)=%q, want high", got)
	}
}
