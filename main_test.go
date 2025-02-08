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
	multiverse = make(map[int]*Universe) // Clear previous state
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
	fmt.Println("🔍 Starting TestQuantumFluctuation") // Debugging

	testingMode = true
	defer func() { testingMode = false }() // Reset after test

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

	// Call quantumFluctuation()
	fmt.Println("⚡ Calling quantumFluctuation()...")
	quantumFluctuation()
	fmt.Println("✅ quantumFluctuation() executed.") // Debugging

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
	// var wg sync.WaitGroup
	mu.Lock()
	multiverse = make(map[int]*Universe)
	mu.Unlock()

	// Create and run a universe that will immediately collapse
	u := NewUniverse(999)
	u.Entropy = 1.1 // Force collapse
	createUniverse(u.ID)

	// Give it time to process
	time.Sleep(100 * time.Millisecond)

	mu.Lock()
	defer mu.Unlock()

	// Verify universe removal
	if _, exists := multiverse[999]; exists {
		t.Error("Universe not removed after collapse")
	}

	// Verify black hole creation
	if len(blackHoles) == 0 {
		t.Error("No black hole created on collapse")
	}
}
