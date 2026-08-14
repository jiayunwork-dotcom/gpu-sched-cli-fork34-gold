package queue

import (
	"testing"

	"github.com/gpu-sched-cli/internal/model"
	"github.com/gpu-sched-cli/internal/store"
)

func TestGetEffectivePriority_OverQuotaMinusTwo(t *testing.T) {
	cfg := model.DefaultSchedulerConfig()
	cfg.UserQuotas["alice"] = 10
	st := store.NewStore(model.NewCluster(&model.ClusterConfig{}), cfg)
	st.AddUserUsage("alice", 11)
	f := NewFairShareManager(st)
	task := &model.Task{Spec: model.TaskSpec{Priority: 5, User: "alice"}}
	got := f.GetEffectivePriority(task)
	if got != 3 {
		t.Fatalf("GetEffectivePriority=%d, want 3", got)
	}
}
