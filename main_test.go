package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewUniverseInitialization(t *testing.T) {
	rand.Seed(1) // Fixed seed for deterministic results
	u := NewUniverse(42)
	if u.ID != 42 {
		t.Errorf("Expected ID 42, got %d", u.ID)
	}
	if u.State != "Initialized" {
		t.Error("Universe not initialized properly")
	}
	if u.Entropy < 0 || u.Entropy > 1 {
		t.Error("Entropy out of bounds [0,1]")
	}
}

func TestNewBlackHoleInitialization(t *testing.T) {
	bh := NewBlackHole(99)
	if bh.ID != 99 || bh.Mass != 0 {
		t.Error("Black hole initialization failed")
	}
	if len(bh.UniversesConsumed) != 0 {
		t.Error("New black hole should have empty consumption list")
	}
}

func TestCreateCosmicString(t *testing.T) {
	u1 := &Universe{ID: 1, Entropy: 0.3, State: "Expanding"}
	u2 := &Universe{ID: 2, Entropy: 0.7, State: "Collapsing"}
	createCosmicString(u1, u2)
	if u1.Entropy != 0.7 || u2.Entropy != 0.3 {
		t.Error("Entropy swap failed")
	}
	if u1.State != "Collapsing" || u2.State != "Expanding" {
		t.Error("State swap failed")
	}
}

func TestCosmicCouncilMeeting(t *testing.T) {
	// Reset state
	mu.Lock()
	multiverse = make(map[int]*Universe) // Reset multiverse, clear previous state
	mu.Unlock()
	// Setup test universes
	mu.Lock()
	multiverse[1] = &Universe{Entropy: 0.9}
	multiverse[2] = &Universe{Entropy: 0.5}
	multiverse[3] = &Universe{Entropy: 0.85}
	mu.Unlock()
	cosmicCouncilMeeting()
	// Verify results
	mu.Lock()
	defer mu.Unlock()
	if multiverse[1].Entropy != 0.1 {
		t.Error("Universe 1 entropy not reset")
	}
	if multiverse[2].Entropy != 0.5 {
		t.Error("Universe 2 entropy changed unexpectedly")
	}
	if multiverse[3].Entropy != 0.1 {
		t.Error("Universe 3 entropy not reset")
	}
	// Cleanup
	multiverse = make(map[int]*Universe)
}

func TestQuantumFluctuation(t *testing.T) {
	fmt.Println("🔍 Starting TestQuantumFluctuation")
	testingMode = true
	defer func() { testingMode = false }()
	// Reset state
	mu.Lock()
	multiverse = make(map[int]*Universe)
	mu.Unlock()
	// Seed universe
	mu.Lock()
	multiverse[777] = NewUniverse(777)
	mu.Unlock()
	originalCount := len(multiverse)
	rand.Seed(42) // Fixed seed for deterministic test
	fmt.Println("⚡ Calling quantumFluctuation()...")
	quantumFluctuation()
	fmt.Println("✅ quantumFluctuation() executed.")
	// Lock again to safely read the state
	mu.Lock()
	finalCount := len(multiverse)
	mu.Unlock()
	if finalCount != originalCount+2 {
		t.Errorf("Expected %d universes, got %d", originalCount+2, finalCount)
	}
	fmt.Println("✅ TestQuantumFluctuation PASSED")
}

func TestBlackHoleConsumption(t *testing.T) {
	bh := NewBlackHole(1)
	u := &Universe{ID: 99, Entropy: 0.8}
	bh.ConsumeUniverse(u)
	if bh.Mass != 0.8 {
		t.Error("Black hole mass incorrect")
	}
	if len(bh.UniversesConsumed) != 1 || bh.UniversesConsumed[0] != 99 {
		t.Error("Consumption tracking failed")
	}
}

func TestUniverseLifecycle(t *testing.T) {
	fmt.Println("🔍 Starting TestUniverseLifecycle")
	testingMode = true
	defer func() { testingMode = false }()
	// Reset global state
	mu.Lock()
	multiverse = make(map[int]*Universe)
	blackHoles = make(map[int]*BlackHole)
	mu.Unlock()
	// Create a universe that will collapse
	createUniverse(999)
	// Manually set entropy to force collapse
	mu.Lock()
	if u, exists := multiverse[999]; exists {
		fmt.Println("⚠️ Manually setting entropy of Universe 999 to 1.1 (force collapse)")
		u.Entropy = 1.1
	}
	mu.Unlock()
	// Detect if mutex is locked before acquiring it in the test
	fmt.Println("🕵️ Checking if mutex is locked before acquiring lock in test...")
	if !mu.TryLock() {
		fmt.Println("❌ Deadlock detected! Another function is holding the mutex.")
		t.Fatal("Deadlock detected: Mutex is already locked.")
	} else {
		fmt.Println("✅ Mutex is free, proceeding...")
	}
	mu.Unlock()
	// Trigger only one iteration of the universe lifecycle (simulate collapse)
	if u, exists := multiverse[999]; exists {
		fmt.Println("🔍 Running single iteration of Universe 999 lifecycle...")
		// Call simulateOneIterationOfRun without holding mu
		simulateOneIterationOfRun(u)
	} else {
		fmt.Println("⚠️ Universe 999 does not exist when trying to run iteration!")
	}
	// Give it time to process
	time.Sleep(10 * time.Millisecond)
	// Lock again to safely check final state
	mu.Lock()
	_, universeExists := multiverse[999]
	bhCount := len(blackHoles)
	mu.Unlock()
	// Debug output regarding final state
	if universeExists {
		fmt.Printf("❌ Universe 999 still exists in multiverse after supposed collapse!\n")
		t.Errorf("Universe not removed after collapse")
	} else {
		fmt.Printf("✅ Universe 999 successfully removed from multiverse!\n")
	}
	if bhCount == 0 {
		fmt.Printf("❌ No black hole created in response to Universe 999 collapsing!\n")
		t.Errorf("No black hole created on collapse")
	} else {
		fmt.Printf("✅ A black hole was successfully created from Universe 999!\n")
	}
	fmt.Println("✅ TestUniverseLifecycle PASSED")
}
