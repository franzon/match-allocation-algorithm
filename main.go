package main

import (
	"fmt"
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
// isn"t any error handling to do.
func (X Genome) Evaluate() (float64, error) {

	fitness := 1.0

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
		} else if count == 0 {
			fitness = 1.0
		} else {

			fitness = float64(count) / float64(len(matchs)*2)
		}

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
	// 	i++""
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

// [{"id": 0, "name": "Tesha", AbleSlots: []string{"1A", "1B", "2A", "6A", "8A", "9B", "10A", "10B"]}, {"id": 1, "name": "Robert", AbleSlots: []string{"4B", "6B", "7A", "7B", "8B", "9A", "9B", "10A"]}, {"id": 2, "name": "Sara", AbleSlots: []string{"5B", "6B", "7A", "8B", "9B"]}, {"id": 3, "name": "Thomas", AbleSlots: []string{"2B", "3A", "4A", "4B", "6B", "7B", "8A"]}, {"id": 4, "name": "Charlie", AbleSlots: []string{"1A", "1B", "3A", "5A", "6B", "7A", "7B", "9A"]}, {"id": 5, "name": "Terry", AbleSlots: []string{"1A", "2B", "3B", "4B", "5A", "9B", "10A", "10B"]}, {"id": 6, "name": "Toby", AbleSlots: []string{"1B", "4B", "6A", "6B", "7A", "8A", "8B", "9B", "10B"]}, {"id": 7, "name": "Carlos", AbleSlots: []string{"3A", "3B", "4A", "5B", "9A"]}, {"id": 8, "name": "Kevin", AbleSlots: []string{"3B", "4A", "5A", "6B", "7A", "8A", "9A", "10B"]}, {"id": 9, "name": "Dan", AbleSlots: []string{"1A", "1B", "2A", "3A", "6A", "6B", "7B", "8A", "8B", "10B"]}, {"id": 10, "name": "William", AbleSlots: []string{"1A", "2A", "3A", "3B", "5A", "6A", "6B", "7A", "7B", "8B", "9A"]}, {"id": 11, "name": "Savannah", AbleSlots: []string{"1A", "3A", "5A", "6A", "8B", "9A"]}, {"id": 12, "name": "Reyes", AbleSlots: []string{"1B",
// "3B", "6B", "9B", "10A"]}, {"id": 13, "name": "Kelly", AbleSlots: []string{"1B", "2A", "3A", "4A", "6B", "7B", "9A", "10B"]}, {"id":
// 14, "name": "Colleen", AbleSlots: []string{"1A", "1B", "2B", "3A", "4B", "5A", "6A", "7A", "7B", "9A"]}, {"id": 15, "name": "Monica", AbleSlots: []string{"1A", "2B", "3A", "3B", "5B", "6A", "6B", "7A",
// "7B", "8B", "10A"]}, {"id": 16, "name": "Bridget", AbleSlots: []string{"2A", "3A", "3B", "4A", "5A", "6A", "6B", "8A", "9A", "10B"]}, {"id": 17, "name": "Ryan", AbleSlots: []string{"1A", "1B", "3B", "5A", "6A", "6B", "7A", "8B", "9B", "10B"]}, {"id": 18, "name": "Dale", AbleSlots: []string{"1A", "1B", "2B", "6B", "7A"]}, {"id": 19, "name": "June", AbleSlots: []string{"1A", "2A", "2B", "5A", "6A", "7A", "8B", "9B"]}, {"id": 20, "name": "Matthew", AbleSlots: []string{"1B", "2A", "2B", "4A", "8A", "8B", "9B", "10A"]}, {"id": 21, "name": "Louise", AbleSlots: []string{"1A", "2A", "2B", "4A", "5A", "5B",
// "6A", "10A", "10B"]}, {"id": 22, "name": "Lionel", AbleSlots: []string{"2A", "4A", "9B", "10A"]}, {"id": 23, "name": "Eric", AbleSlots: []string{"1B", "2B", "4A", "5A", "7B", "9B"]}, {"id": 24, "name":
// "Marcell", AbleSlots: []string{"1A", "1B", "4A", "5A", "6A", "7B", "8B"]}, {"id": 25, "name": "Jon", AbleSlots: []string{"1A", "1B", "2A", "3A", "4B", "7A", "7B", "8B", "9B"]}, {"id": 26, "name": "Edna", AbleSlots: []string{"1A", "1B", "2A", "3A", "3B", "5A", "5B", "7A", "7B", "8A", "9B", "10A"]}, {"id": 27, "name": "Mark", AbleSlots: []string{"1A", "5A", "7A", "10A"]}, {"id": 28, "name": "Kent", AbleSlots: []string{"4B", "5B", "8A", "8B", "9A", "10A"]}, {"id": 29, "name": "John", AbleSlots: []string{"1A", "2B", "5B", "6B", "7A", "9A",
// "10B"]}, {"id": 30, "name": "Sharon", AbleSlots: []string{"1A", "1B", "2A", "2B", "4A", "4B", "5B", "7A", "7B", "8B", "10A"]}, {"id": 31, "name": "Andrew", AbleSlots: []string{"3B", "4A", "6B", "7A", "8A", "8B", "9A"]}]

// [{"id": 0, "player1": 24, "player2": 3}, {"id": 1, "player1": 12, "player2": 10}, {"id": 2, "player1": 19, "player2": 6}, {"id": 3, "player1": 29, "player2": 18}, {"id": 4, "player1": 4, "player2": 14}, {"id": 5, "player1": 25, "player2": 22}, {"id": 6, "player1": 8, "player2": 27}, {"id": 7, "player1": 9, "player2": 7}, {"id": 8, "player1": 15, "player2": 21}, {"id": 9, "player1": 5, "player2": 26}, {"id": 10, "player1": 23, "player2": 28}, {"id": 11, "player1": 17, "player2": 13}, {"id": 12, "player1": 0, "player2": 2}, {"id": 13, "player1": 11, "player2": 20}, {"id": 14, "player1": 30, "player2": 1}, {"id": 15, "player1": 31, "player2": 16}]

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
		Slot{ID: "9A", Description: "Dia  9 - Quadra 1"},
		Slot{ID: "9B", Description: "Dia  9 - Quadra 2"},
		Slot{ID: "10A", Description: "Dia  10 - Quadra 1"},
		Slot{ID: "10B", Description: "Dia  10 - Quadra 2"},
	}

	// players = []Player{
	// 	Player{ID: 0, Name: "Jorge", AbleSlots: []string{]string{"1A"}},
	// 	Player{ID: 1, Name: "Rafael", AbleSlots: []string{]string{"1A", "1B"}},
	// 	Player{ID: 2, Name: "Larissa", AbleSlots: []string{]string{"2A"}},
	// 	Player{ID: 3, Name: "Dennis", AbleSlots: []string{]string{"2A"}},
	// 	Player{ID: 4, Name: "João", AbleSlots: []string{]string{"3A", "1B"}},
	// 	Player{ID: 5, Name: "Maria", AbleSlots: []string{]string{"3A", "1B"}},
	// 	Player{ID: 6, Name: "Pedro", AbleSlots: []string{]string{"4A", "1B"}},
	// 	Player{ID: 7, Name: "Lucas", AbleSlots: []string{]string{"4A", "1B"}},
	// }

	players = []Player{
		Player{ID: 0, Name: "Tesha", AbleSlots: []string{"1A", "1B", "2A", "6A", "8A", "9B", "10A", "10B"}},

		Player{ID: 1,
			Name:      "Robert",
			AbleSlots: []string{"4B", "6B", "7A", "7B", "8B", "9A", "9B", "10A"}},

		Player{ID: 2,
			Name:      "Sara",
			AbleSlots: []string{"5B", "6B", "7A", "8B", "9B"}},

		Player{ID: 3,
			Name:      "Thomas",
			AbleSlots: []string{"2B", "3A", "4A", "4B", "6B", "7B", "8A"}},

		Player{ID: 4,
			Name:      "Charlie",
			AbleSlots: []string{"1A", "1B", "3A", "5A", "6B", "7A", "7B", "9A"}},

		Player{ID: 5,
			Name:      "Terry",
			AbleSlots: []string{"1A", "2B", "3B", "4B", "5A", "9B", "10A", "10B"}},

		Player{ID: 6,
			Name:      "Toby",
			AbleSlots: []string{"1B", "4B", "6A", "6B", "7A", "8A", "8B", "9B", "10B"}},

		Player{ID: 7,
			Name:      "Carlos",
			AbleSlots: []string{"3A", "3B", "4A", "5B", "9A"}},

		Player{ID: 8,
			Name:      "Kevin",
			AbleSlots: []string{"3B", "4A", "5A", "6B", "7A", "8A", "9A", "10B"}},

		Player{ID: 9,
			Name:      "Dan",
			AbleSlots: []string{"1A", "1B", "2A", "3A", "6A", "6B", "7B", "8A", "8B", "10B"}},

		Player{ID: 10,
			Name:      "William",
			AbleSlots: []string{"1A", "2A", "3A", "3B", "5A", "6A", "6B", "7A", "7B", "8B", "9A"}},

		Player{ID: 11,
			Name:      "Savannah",
			AbleSlots: []string{"1A", "3A", "5A", "6A", "8B", "9A"}},

		Player{ID: 12,
			Name:      "Reyes",
			AbleSlots: []string{"1B", "3B", "6B", "9B", "10A"}},
		Player{ID: 13,
			Name:      "Kelly",
			AbleSlots: []string{"1B", "2A", "3A", "4A", "6B", "7B", "9A", "10B"}},

		Player{ID: 14,
			Name:      "Colleen",
			AbleSlots: []string{"1A", "1B", "2B", "3A", "4B", "5A", "6A", "7A", "7B", "9A"}},

		Player{ID: 15,
			Name:      "Monica",
			AbleSlots: []string{"1A", "2B", "3A", "3B", "5B", "6A", "6B", "7A", "7B", "8B", "10A"}},

		Player{ID: 16,
			Name:      "BrIDget",
			AbleSlots: []string{"2A", "3A", "3B", "4A", "5A", "6A", "6B", "8A", "9A", "10B"}},

		Player{ID: 17,
			Name:      "Ryan",
			AbleSlots: []string{"1A", "1B", "3B", "5A", "6A", "6B", "7A", "8B", "9B", "10B"}},

		Player{ID: 18,
			Name:      "Dale",
			AbleSlots: []string{"1A", "1B", "2B", "6B", "7A"}},

		Player{ID: 19,
			Name:      "June",
			AbleSlots: []string{"1A", "2A", "2B", "5A", "6A", "7A", "8B", "9B"}},

		Player{ID: 20,
			Name:      "Matthew",
			AbleSlots: []string{"1B", "2A", "2B", "4A", "8A", "8B", "9B", "10A"}},

		Player{ID: 21,
			Name:      "Louise",
			AbleSlots: []string{"1A", "2A", "2B", "4A", "5A", "5B", "6A", "10A", "10B"}},

		Player{ID: 22,
			Name:      "Lionel",
			AbleSlots: []string{"2A", "4A", "9B", "10A"}},

		Player{ID: 23,
			Name:      "Eric",
			AbleSlots: []string{"1B", "2B", "4A", "5A", "7B", "9B"}},

		Player{ID: 24,
			Name:      "Marcell",
			AbleSlots: []string{"1A", "1B", "4A", "5A", "6A", "7B", "8B"}},

		Player{ID: 25,
			Name:      "Jon",
			AbleSlots: []string{"1A", "1B", "2A", "3A", "4B", "7A", "7B", "8B", "9B"}},

		Player{ID: 26,
			Name:      "Edna",
			AbleSlots: []string{"1A", "1B", "2A", "3A", "3B", "5A", "5B", "7A", "7B", "8A", "9B", "10A"}},

		Player{ID: 27,
			Name:      "Mark",
			AbleSlots: []string{"1A", "5A", "7A", "10A"}},

		Player{ID: 28,
			Name:      "Kent",
			AbleSlots: []string{"4B", "5B", "8A", "8B", "9A", "10A"}},

		Player{ID: 29,
			Name:      "John",
			AbleSlots: []string{"1A", "2B", "5B", "6B", "7A", "9A", "10B"}},

		Player{ID: 30,
			Name:      "Sharon",
			AbleSlots: []string{"1A", "1B", "2A", "2B", "4A", "4B", "5B", "7A", "7B", "8B", "10A"}},

		Player{ID: 31,
			Name:      "Andrew",
			AbleSlots: []string{"3B", "4A", "6B", "7A", "8A", "8B", "9A"}}}

	// matchs = []Match{
	// 	Match{ID: 0, Player1: players[0], Player2: players[1]},
	// 	Match{ID: 1, Player1: players[2], Player2: players[3]},
	// 	Match{ID: 2, Player1: players[4], Player2: players[5]},
	// 	Match{ID: 3, Player1: players[6], Player2: players[7]},
	// }

	matchs = []Match{
		Match{ID: 0, Player1: players[24], Player2: players[3]},
		Match{ID: 1, Player1: players[12], Player2: players[10]},
		Match{ID: 2, Player1: players[19], Player2: players[6]},
		Match{ID: 3, Player1: players[29], Player2: players[18]},
		Match{ID: 4, Player1: players[4], Player2: players[14]},
		Match{ID: 5, Player1: players[25], Player2: players[22]},
		Match{ID: 6, Player1: players[8], Player2: players[27]},
		Match{ID: 7, Player1: players[9], Player2: players[7]},
		Match{ID: 8, Player1: players[15], Player2: players[21]},
		Match{ID: 9, Player1: players[5], Player2: players[26]},
		Match{ID: 10, Player1: players[23], Player2: players[28]},
		Match{ID: 11, Player1: players[17], Player2: players[13]},
		Match{ID: 12, Player1: players[0], Player2: players[2]},
		Match{ID: 13, Player1: players[11], Player2: players[20]},
		Match{ID: 14, Player1: players[30], Player2: players[1]},
		Match{ID: 15, Player1: players[31], Player2: players[16]}}

	first := time.Now()

	// Instantiate a GA with a GAConfig

	gaConfig := eaopt.GAConfig{NPops: 100, PopSize: 50, HofSize: 1, NGenerations: 300, ParallelEval: true, Model: eaopt.ModGenerational{
		Selector: eaopt.SelTournament{
			NContestants: 3,
		},
		MutRate:   0.2,
		CrossRate: 0.8,
	}}
	var ga, err = gaConfig.NewGA()

	if err != nil {
		fmt.Println(err)
		return
	}

	// Set the number of generations to run for
	// ga.NGenerations = 1000

	// Add a custom print function to track progress
	ga.Callback = func(ga *eaopt.GA) {
		// fmt.Printf("Best fitness at generation %d: %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
	}

	// Find the minimum
	err = ga.Minimize(GenomeFactory)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Melhor solução %2.2f (%d/%d) de acertos", (1.0-ga.HallOfFame[0].Fitness)*100, int((1.0-ga.HallOfFame[0].Fitness)*float64(len(matchs))), len(matchs))
	fmt.Println("Tempo total (segundos): ", time.Now().Sub(first).Seconds())

}
