package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/MaxHalford/eaopt"
	"github.com/gorilla/mux"
)

type AllocateRequestBody struct {
	Slots   []Slot   `json:"slots"`
	Players []Player `json:"players"`
	Matchs  []Match  `json:"matchs"`
}

type AllocateRequestResponse struct {
	Slots   []int   `json:"slots"`
	Fitness float64 `json:"fitness"`
}

type Slot struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type Player struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	AbleSlots []string `json:"able_slots"`
}

type Match struct {
	ID      int `json:"id"`
	Player1 int `json:"player1"`
	Player2 int `json:"player2"`
}

type Schedule struct {
	Genome  []int
	Slots   []Slot
	Players []Player
	Matchs  []Match
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func FindPlayerByID(id int, players []Player) (Player, error) {
	for _, p := range players {
		if p.ID == id {
			return p, nil
		}
	}
	return Player{}, fmt.Errorf("Couldn't found Player")
}

func (X Schedule) Evaluate() (float64, error) {

	fitness := 1.0

	count := 0

	for index := 0; index < len(X.Genome); index++ {

		matchID := X.Genome[index]
		slot := X.Slots[index]

		if matchID != -1 {

			for _, match := range X.Matchs {
				if match.ID == matchID {

					player1, _ := FindPlayerByID(match.Player1, X.Players)
					player2, _ := FindPlayerByID(match.Player2, X.Players)

					if Contains(player1.AbleSlots, slot.ID) && Contains(player2.AbleSlots, slot.ID) {
						count++
					}

					break

					// if Contains(player1.AbleSlots, slot.ID) {
					// 	count++
					// }

					// if Contains(player2.AbleSlots, slot.ID) {
					// 	count++
					// }

				}
			}
		}
	}

	withoutEmpty := make([]int, 0)

	for index := 0; index < len(X.Genome); index++ {
		if X.Genome[index] != -1 {
			withoutEmpty = append(withoutEmpty, X.Genome[index])
		}
	}

	slotMap := make(map[int]bool)
	for index := 0; index < len(withoutEmpty); index++ {
		slotMap[withoutEmpty[index]] = true
	}

	if len(slotMap) != len(withoutEmpty) {
		// fmt.Println("Duplic")

	} else if len(slotMap) != len(X.Matchs) {
		// fmt.Println("faltant")
	} else {

		if count == len(X.Matchs) {
			// if count == len(X.Matchs)*2 {
			// fmt.Println("Solução ", X)
			fitness = 0.0

		} else if count == 0 {
			fitness = 1.0
		} else {
			fitness = -float64(count) / float64(len(X.Matchs))
			// fitness = float64(count) / float64(len(X.Matchs)*2)
		}

	}

	return fitness, nil
}

func (X Schedule) Mutate(rng *rand.Rand) {
	eaopt.MutPermuteInt(X.Genome, 1, rng)
}

func (X Schedule) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossGNXInt(X.Genome, Y.(Schedule).Genome, 1, rng)
}

func (X Schedule) Clone() eaopt.Genome {
	var Y = Schedule{Genome: make([]int, len(X.Genome)), Slots: X.Slots, Players: X.Players, Matchs: X.Matchs}
	copy(Y.Genome, X.Genome)
	return Y
}

func RemoveIndex(s []Player, index int) []Player {
	return append(s[:index], s[index+1:]...)
}

func (X Schedule) GenomeFactory(rng *rand.Rand) eaopt.Genome {

	tmpPlayers := make([]Player, len(X.Players))
	copy(tmpPlayers, X.Players)

	rndSlots := make([]int, len(X.Slots))

	for index := 0; index < len(X.Slots); index++ {
		rndSlots[index] = -1
	}

	for index := 0; index < len(X.Matchs); index++ {
		rndSlots[index] = X.Matchs[index].ID
	}

	rand.Shuffle(len(rndSlots), func(i, j int) {
		rndSlots[i], rndSlots[j] = rndSlots[j], rndSlots[i]
	})

	return Schedule{Genome: rndSlots, Slots: X.Slots, Players: X.Players, Matchs: X.Matchs}
}

func GenerateAllocation(w http.ResponseWriter, r *http.Request) {

	var body AllocateRequestBody
	_ = json.NewDecoder(r.Body).Decode(&body)

	first := time.Now()

	earlyStop := func(ga *eaopt.GA) bool {
		// fmt.Println(ga.HallOfFame[0].Fitness)
		return ga.HallOfFame[0].Fitness == 0.0
	}

	gaConfig := eaopt.GAConfig{EarlyStop: earlyStop, NPops: 1, PopSize: 80, HofSize: 1, NGenerations: 300, ParallelEval: false, Model: eaopt.ModGenerational{
		Selector: eaopt.SelTournament{
			NContestants: 3,
		},
		MutRate:   0.5,
		CrossRate: 0.7,
	}}
	var ga, err = gaConfig.NewGA()

	if err != nil {
		fmt.Println(err)
		return
	}

	// ga.Callback = func(ga *eaopt.GA) {
	// 	// fmt.Printf("Best fitness at generation %d: %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
	// }

	data := Schedule{Slots: body.Slots, Matchs: body.Matchs, Players: body.Players}

	err = ga.Minimize(data.GenomeFactory)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Melhor solução %2.2f (%d/%d) de acertos\n", (1.0-ga.HallOfFame[0].Fitness)*100, int((1.0-ga.HallOfFame[0].Fitness)*float64(len(body.Matchs))), len(body.Matchs))
	fmt.Println("Tempo total (segundos): ", time.Now().Sub(first).Seconds())

	slots := ga.HallOfFame[0].Genome.(Schedule).Genome

	for slotIdx := 0; slotIdx < len(slots); slotIdx++ {
		slot := data.Slots[slotIdx]
		matchID := slots[slotIdx]

		if matchID != -1 {

			fmt.Println(matchID, slot.ID)

			for _, match := range data.Matchs {

				if match.ID == matchID {

					player1, _ := FindPlayerByID(match.Player1, data.Players)
					player2, _ := FindPlayerByID(match.Player2, data.Players)

					if !Contains(player1.AbleSlots, slot.ID) || !Contains(player2.AbleSlots, slot.ID) {
						fmt.Println("Erro aqui")
					}

					break
				}
			}

		}
	}

	// for index := 0; index < len(slots); index++ {

	// 	matchID := slots[index]
	// 	slot := data.Slots[index]

	// 	if matchID != -1 {

	// 		for _, matchId := range slots {

	// 			found := false

	// 			for _, match := range data.Matchs {
	// 				if match.ID == matchId {

	// 					player1, _ := FindPlayerByID(match.Player1, data.Players)
	// 					player2, _ := FindPlayerByID(match.Player2, data.Players)

	// 					if Contains(player1.AbleSlots, slot.ID) && Contains(player2.AbleSlots, slot.ID) {
	// 						found := true
	// 						break
	// 					}

	// 				}
	// 			}

	// 		}

	// 	}

	// }

	// fmt.Println(x)
	json.NewEncoder(w).Encode(AllocateRequestResponse{Fitness: (1.0 - ga.HallOfFame[0].Fitness), Slots: slots})
}

func main() {

	rand.Seed(time.Now().Unix())

	r := mux.NewRouter()
	r.HandleFunc("/allocate", GenerateAllocation).Methods("POST")

	http.ListenAndServe(":8000", r)

}
