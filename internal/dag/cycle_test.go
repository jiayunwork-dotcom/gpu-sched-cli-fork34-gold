package dag

import "testing"

func TestDetectCycle_TwoNodeLoop(t *testing.T) {
	g := NewDependencyGraph()
	g.AddEdge("A", "B")
	g.AddEdge("B", "A")
	_, found := DetectCycle(g)
	if !found {
		t.Fatal("A→B→A should be a cycle")
	}
}
