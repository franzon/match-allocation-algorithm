package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/MaxHalford/eaopt"
)

type Slot struct {
	ID          string
	Description string
}

type Player struct {
	ID        int
	Name      string
	AbleSlots []string
}

type Match struct {
	ID      int
	Player1 Player
	Player2 Player
}

// A Genome contains int32s.
type Genome []int

var slots []Slot
var players []Player
var matchs []Match

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// Evaluate a Vector with the Drop-Wave function which takes two variables as
// input and reaches a minimum of -1 in (0, 0). The function is simple so there
// isn't any error handling to do.
func (X Genome) Evaluate() (float64, error) {

	fitness := math.MaxFloat64

	count := 0

	for index := 0; index < len(X); index++ {

		matchId := X[index]
		slot := slots[index]

		if matchId != -1 {

			for _, match := range matchs {
				if match.ID == matchId {

					if Contains(match.Player1.AbleSlots, slot.ID) {
						count++
					}

					if Contains(match.Player2.AbleSlots, slot.ID) {
						count++
					}

					break
				}
			}
		}
	}

	withoutEmpty := make([]int, 0)

	for index := 0; index < len(X); index++ {
		if X[index] != -1 {
			withoutEmpty = append(withoutEmpty, X[index])
		}
	}

	slotMap := make(map[int]bool)
	for index := 0; index < len(withoutEmpty); index++ {
		slotMap[withoutEmpty[index]] = true
	}

	if len(slotMap) != len(withoutEmpty) {
		// fmt.Println("Duplic")

	} else if len(slotMap) != len(matchs) {
		// fmt.Println("faltant")
	} else {

		if count == len(matchs)*2 {
			fmt.Println("Solução ", X)
			os.Exit(0)
		}

		fitness = float64(1.0 / float64(count+1.0))
	}

	return fitness, nil
}

// // Mutate a Vector by resampling each element from a normal distribution with
// // probability 0.8.
func (X Genome) Mutate(rng *rand.Rand) {
	eaopt.MutPermuteInt(X, 1, rng)
}

// Crossover a Vector with another Vector by applying uniform crossover.
func (X Genome) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossGNXInt(X, Y.(Genome), 1, rng)
}

// Clone a Vector to produce a new one that points to a different slice.
func (X Genome) Clone() eaopt.Genome {
	var Y = make(Genome, len(X))
	copy(Y, X)
	return Y
}

func RemoveIndex(s []Player, index int) []Player {
	return append(s[:index], s[index+1:]...)
}

func GenomeFactory(rng *rand.Rand) eaopt.Genome {

	tmpPlayers := make([]Player, len(players))
	copy(tmpPlayers, players)

	// i := 0

	// for len(tmpPlayers) > 0 {
	// 	rndIdx := rand.Intn((len(tmpPlayers)))
	// 	player1 := tmpPlayers[rndIdx]

	// 	tmpPlayers = RemoveIndex(tmpPlayers, rndIdx)

	// 	rndIdx = rand.Intn((len(tmpPlayers)))
	// 	player2 := tmpPlayers[rndIdx]

	// 	tmpPlayers = RemoveIndex(tmpPlayers, rndIdx)

	// 	matchs = append(matchs, Match{ID: i, Player1: player1, Player2: player2})
	// 	i++''
	// }

	rndSlots := make([]int, len(slots))

	for index := 0; index < len(slots); index++ {
		rndSlots[index] = -1
	}

	for index := 0; index < len(matchs); index++ {
		rndSlots[index] = matchs[index].ID
	}

	rand.Shuffle(len(rndSlots), func(i, j int) {
		rndSlots[i], rndSlots[j] = rndSlots[j], rndSlots[i]
	})

	return Genome(rndSlots)
}

func main() {

	rand.Seed(time.Now().Unix())

	slots = []Slot{
		Slot{ID: "1A", Description: "Dia  1 - Quadra 1"},
		Slot{ID: "1B", Description: "Dia  1 - Quadra 2"},
		Slot{ID: "2A", Description: "Dia  2 - Quadra 1"},
		Slot{ID: "2B", Description: "Dia  2 - Quadra 2"},
		Slot{ID: "3A", Description: "Dia  3 - Quadra 1"},
		Slot{ID: "3B", Description: "Dia  3 - Quadra 2"},
		Slot{ID: "4A", Description: "Dia  4 - Quadra 1"},
		Slot{ID: "4B", Description: "Dia  4 - Quadra 2"},
		Slot{ID: "5A", Description: "Dia  5 - Quadra 1"},
		Slot{ID: "5B", Description: "Dia  5 - Quadra 2"},
		Slot{ID: "6A", Description: "Dia  6 - Quadra 1"},
		Slot{ID: "6B", Description: "Dia  6 - Quadra 2"},
		Slot{ID: "7A", Description: "Dia  7 - Quadra 1"},
		Slot{ID: "7B", Description: "Dia  7 - Quadra 2"},
		Slot{ID: "8A", Description: "Dia  8 - Quadra 1"},
		Slot{ID: "8B", Description: "Dia  8 - Quadra 2"},
	}

	players = []Player{
		Player{ID: 0, Name: "Jorge", AbleSlots: []string{"1A", "1B"}},
		Player{ID: 1, Name: "Rafael", AbleSlots: []string{"1A", "1B"}},
		Player{ID: 2, Name: "Larissa", AbleSlots: []string{"2A"}},
		Player{ID: 3, Name: "Dennis", AbleSlots: []string{"2A"}},
		Player{ID: 4, Name: "João", AbleSlots: []string{"3A", "1B"}},
		Player{ID: 5, Name: "Maria", AbleSlots: []string{"3A", "1B"}},
		Player{ID: 6, Name: "Pedro", AbleSlots: []string{"4A", "1B"}},
		Player{ID: 7, Name: "Lucas", AbleSlots: []string{"4A", "1B"}},
	}

	matchs = []Match{
		Match{ID: 0, Player1: players[0], Player2: players[1]},
		Match{ID: 1, Player1: players[2], Player2: players[3]},
		Match{ID: 2, Player1: players[4], Player2: players[5]},
		Match{ID: 3, Player1: players[6], Player2: players[7]},
	}

	// Instantiate a GA with a GAConfig
	var ga, err = eaopt.NewDefaultGAConfig().NewGA()

	if err != nil {
		fmt.Println(err)
		return
	}

	// Set the number of generations to run for
	ga.NGenerations = 1000

	// Add a custom print function to track progress
	ga.Callback = func(ga *eaopt.GA) {
		fmt.Printf("Best fitness at generation %d: %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
	}

	// Find the minimum
	err = ga.Minimize(GenomeFactory)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Melhor solução", ga.HallOfFame[0].Genome)
}
