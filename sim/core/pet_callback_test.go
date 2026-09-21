package core

import "testing"

func TestPetEnableCallbacks(t *testing.T) {
	pet := &Pet{}
	calls := 0
	pet.ApplyOnPetEnable(func(*Simulation) { calls++ })
	if pet.OnPetEnable == nil {
		t.Fatal("first callback was discarded")
	}
	pet.ApplyOnPetEnable(func(*Simulation) { calls += 2 })
	pet.OnPetEnable(nil)
	if calls != 3 {
		t.Fatalf("callback total = %d, want 3", calls)
	}
}
