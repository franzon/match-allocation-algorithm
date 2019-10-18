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

// -------------------- Request --------------------

// // RequestSlot Representação do Slot na requisição
// type RequestSlot struct {
// 	ID          int    `json:"id"`
// 	Description string `json:"description"`
// }

// // RequestPlayer Representação do Player na requisição
// type RequestPlayer struct {
// 	ID           int      `json:"id"`
// 	Name         string   `json:"name"`
// 	AbleSlotsIDs []string `json:"able_slots"`
// }

// RequestMatch Representação da Match na requisição
type RequestMatch struct {
	ID      int `json:"id"`
	Player1 int `json:"player1"`
	Player2 int `json:"player2"`
}

// RequestBody Representação da requisição
type RequestBody struct {
	Slots   []Slot         `json:"slots"`
	Players []Player       `json:"players"`
	Matchs  []RequestMatch `json:"matchs"`
}

// -------------------- Response --------------------

// InitGeneticAlgorithm Inicializa o algoritmo genético
func RunGeneticAlgorithm(body RequestBody) (ResponseBody, error) {

	// Função de parada do algorítmo genético
	earlyStop := func(ga *eaopt.GA) bool {
		return ga.HallOfFame[0].Fitness == -float64(len(body.Matchs)*10)
	}

	// Configuração do algorítmo genético
	// gaConfig := eaopt.GAConfig{
	// 	EarlyStop: earlyStop, NPops: 20, PopSize: 100, HofSize: 1, NGenerations: 300, ParallelEval: true, Model: eaopt.ModGenerational{
	// 		Selector: eaopt.SelTournament{
	// 			NContestants: 3,
	// 		},
	// 		MutRate:   0.5,
	// 		CrossRate: 0.1,
	// 	}}

	gaConfig := eaopt.GAConfig{
		EarlyStop: earlyStop, NPops: 24, PopSize: 80, HofSize: 1, NGenerations: 300, ParallelEval: true,
		Model: eaopt.ModGenerational{
			Selector: eaopt.SelTournament{
				NContestants: 3,
			},
			MutRate:   0.4,
			CrossRate: 0.01,
		}}

	var ga, err = gaConfig.NewGA()

	// ga.Callback = func(ga *eaopt.GA) {
	// 	fmt.Printf("Best fitness at generation %d: %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
	// }

	if err != nil {
		fmt.Println(err)
		return ResponseBody{}, fmt.Errorf("erro")
	}

	// Dados necessários para calcular o fitness do cromossomo

	// Popula os jogadores nas partidas
	matchs := make([]Match, 0)

	for _, match := range body.Matchs {
		player1, _ := FindPlayerByID(match.Player1, body.Players)
		player2, _ := FindPlayerByID(match.Player2, body.Players)

		matchs = append(matchs, Match{ID: match.ID, Player1: player1, Player2: player2})
	}

	scheduleData := ScheduleData{Slots: body.Slots, Players: body.Players, Matchs: matchs}

	// Inicia o algorítmo genético

	err = ga.Minimize(scheduleData.GenomeFactory)

	if err != nil {
		fmt.Println(err)
		return ResponseBody{}, fmt.Errorf("erro")
	}

	return ga.HallOfFame[0].Genome.(Schedule).BuildResponseFromSchedule(), nil

}

// GenerateAllocation Procura uma boa alocação das partidas
func GenerateAllocation(w http.ResponseWriter, r *http.Request) {

	// timer := time.Now()

	var body RequestBody
	_ = json.NewDecoder(r.Body).Decode(&body)

	response, _ := RunGeneticAlgorithm(body)

	json.NewEncoder(w).Encode(response)
}

func main() {

	rand.Seed(time.Now().Unix())

	r := mux.NewRouter()
	r.HandleFunc("/allocate", GenerateAllocation).Methods("POST")

	http.ListenAndServe(":8000", r)

}
