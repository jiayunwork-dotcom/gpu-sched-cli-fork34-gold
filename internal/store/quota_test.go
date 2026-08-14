package store

import (
	"testing"

	"github.com/gpu-sched-cli/internal/model"
)

func TestIsOverQuota_EqualIsNotOver(t *testing.T) {
	cfg := model.DefaultSchedulerConfig()
	cfg.UserQuotas["bob"] = 10
	s := NewStore(model.NewCluster(&model.ClusterConfig{}), cfg)
	s.AddUserUsage("bob", 10)
	if s.IsOverQuota("bob") {
		t.Fatal("usage equal to quota should not be over quota")
	}
}
