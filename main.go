package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Universe struct {
	ID          int
	State       string
	Entropy     float64
	HasWormhole bool
}

type BlackHole struct {
	ID                int
	Mass              float64
	UniversesConsumed []int
}

var states = []string{
	"Expanding", "Collapsing", "Exploding", "Reversing Time", "Frozen",
	"Splitting", "Merging", "Heating", "Cooling", "Entangled", "Chaotic",
}

var multiverse = make(map[int]*Universe)
var wg sync.WaitGroup
var mu sync.Mutex
var blackHoles = make(map[int]*BlackHole)

func NewUniverse(id int) *Universe {
	return &Universe{
		ID:          id,
		State:       "Initialized",
		Entropy:     rand.Float64(),
		HasWormhole: rand.Float32() < 0.2,
	}
}

func NewBlackHole(id int) *BlackHole {
	return &BlackHole{
		ID:                id,
		Mass:              0,
		UniversesConsumed: []int{},
	}
}

func (u *Universe) Run() {
	for {
		// Time dilation based on entropy by the way
		timeMultiplier := 1.0 / (1.0 + u.Entropy)
		time.Sleep(time.Duration(float64(rand.Intn(1000)+500)*timeMultiplier) * time.Millisecond)

		u.State = states[rand.Intn(len(states))]

		// Add time-based randomness to entropy calculation
		randFloat := rand.Float64() * (1 + float64(time.Now().UnixNano()%1000)/1000)
		u.Entropy += randFloat*0.1 - 0.05

		if u.Entropy < 0 {
			u.Entropy = 0
		}
		if u.Entropy > 1 {
			fmt.Printf("💥 Universe %d has reached MAX ENTROPY and COLLAPSED!\n", u.ID)
			mu.Lock()
			// Create black hole when universe collapses
			bhID := rand.Intn(1000)
			blackHoles[bhID] = NewBlackHole(bhID)
			blackHoles[bhID].ConsumeUniverse(u)
			delete(multiverse, u.ID)
			mu.Unlock()
			break
		}
		if u.HasWormhole && rand.Float32() < 0.1 {
			mu.Lock()
			targetID := rand.Intn(1000)
			if target, ok := multiverse[targetID]; ok && target.ID != u.ID {
				fmt.Printf("🌀 Universe %d opened a WORMHOLE to Universe %d!\n", u.ID, target.ID)
				u.State = "Entangled"
				target.State = "Entangled"
				// Create cosmic string between them
				createCosmicString(u, target)
			}
			mu.Unlock()
		}

		// ANSI escape codes for colored output
		colorCode := 31 + rand.Intn(6) // Random red(31) to cyan(36)
		fmt.Printf("\033[%dm🌌 Universe %d is %s (Entropy: %.2f)\033[0m\n",
			colorCode, u.ID, u.State, u.Entropy)

		time.Sleep(time.Duration(rand.Intn(1000)+500) * time.Millisecond)
	}
}

func (bh *BlackHole) ConsumeUniverse(u *Universe) {
	bh.Mass += u.Entropy
	bh.UniversesConsumed = append(bh.UniversesConsumed, u.ID)
	fmt.Printf("🕳️  Black Hole %d consumed Universe %d (Mass: %.2f)\n",
		bh.ID, u.ID, bh.Mass)
}

// Cosmic Strings connect two universes
func createCosmicString(u1, u2 *Universe) {
	fmt.Printf("🌀 Cosmic String formed between Universe %d and Universe %d!\n", u1.ID, u2.ID)
	// Swap entropy and states
	u1.Entropy, u2.Entropy = u2.Entropy, u1.Entropy
	u1.State, u2.State = u2.State, u1.State
}

func quantumFluctuation() {
	mu.Lock()
	defer mu.Unlock()
	if len(multiverse) == 0 {
		return
	}
	id := rand.Intn(1000)
	if _, exists := multiverse[id]; exists {
		fmt.Printf("⚛️  Quantum Fluctuation split Universe %d into two!\n", id)
		createUniverse(rand.Intn(1000))
		createUniverse(rand.Intn(1000))
	}
}

// Cosmic Council maintains order
func cosmicCouncilMeeting() {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("👑 The Cosmic Council is in session...")
	for id, u := range multiverse {
		if u.Entropy > 0.8 {
			fmt.Printf("⚖️  Cosmic Council resets entropy for Universe %d\n", id)
			u.Entropy = 0.1
		}
	}
}

func createUniverse(id int) {
	universe := NewUniverse(id)
	mu.Lock()
	multiverse[id] = universe
	mu.Unlock()
	wg.Add(1)
	go func() {
		defer wg.Done()
		universe.Run()
	}()
}

func bigBang() {
	fmt.Println("🌟 BIG BANG EVENT! Creating a cluster of universes...")
	for i := 0; i < rand.Intn(10)+5; i++ {
		createUniverse(rand.Intn(1000))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Initial multiverse creation
	for i := 0; i < 10; i++ {
		createUniverse(i)
	}

	// Dynamically create and destroy universes
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(3000)+1000) * time.Millisecond)
			if rand.Float32() < 0.2 {
				bigBang()
			} else {
				newID := rand.Intn(1000)
				fmt.Printf("✨ A NEW UNIVERSE %d has been BORN!\n", newID)
				createUniverse(newID)
			}
		}
	}()

	// Add quantum fluctuations
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(5000)+2000) * time.Millisecond)
			quantumFluctuation()
		}
	}()

	// Cosmic Council meetings
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(15000)+10000) * time.Millisecond)
			cosmicCouncilMeeting()
		}
	}()

	// Monitor the multiverse
	go func() {
		for {
			mu.Lock()
			fmt.Printf("🌠 Multiverse contains %d universes\n", len(multiverse))
			fmt.Printf("🌠 Multiverse contains %d universes and %d black holes\n", len(multiverse), len(blackHoles))
			mu.Unlock()
			time.Sleep(2 * time.Second)
		}
	}()

	// Allow the user to influence the multiverse
	go func() {
		for {
			var input string
			fmt.Print("🔥 Type 'destroy' to collapse a universe or 'create' for a Big Bang: ")
			fmt.Scanln(&input)
			if input == "destroy" {
				mu.Lock()
				for id := range multiverse {
					fmt.Printf("💀 Universe %d has been DESTROYED by the user!\n", id)
					delete(multiverse, id)
					break
				}
				mu.Unlock()
			} else if input == "create" {
				bigBang()
			}
		}
	}()

	// Wait forever in the chaos
	wg.Wait()
}
