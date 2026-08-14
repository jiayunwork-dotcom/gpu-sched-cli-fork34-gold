package queue

import (
	"testing"

	"github.com/gpu-sched-cli/internal/model"
	"github.com/gpu-sched-cli/internal/store"
)

func TestEnqueue_PriorityEightDequeuedFirst(t *testing.T) {
	st := store.NewStore(model.NewCluster(&model.ClusterConfig{}), model.DefaultSchedulerConfig())
	q := NewPriorityQueue(st)
	med := &model.Task{Spec: model.TaskSpec{Name: "med", Priority: 5}}
	hi := &model.Task{Spec: model.TaskSpec{Name: "hi", Priority: 8}}
	q.Enqueue(med)
	q.Enqueue(hi)
	first := q.Dequeue()
	if first == nil || first.Spec.Name != "hi" {
		name := ""
		if first != nil {
			name = first.Spec.Name
		}
		t.Fatalf("first dequeued %q, want hi", name)
	}
}
