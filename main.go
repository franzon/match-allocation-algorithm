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
	gaConfig := eaopt.GAConfig{
		EarlyStop: earlyStop, NPops: 10, PopSize: 200, HofSize: 1, NGenerations: 500, ParallelEval: true, Model: eaopt.ModGenerational{
			Selector: eaopt.SelTournament{
				NContestants: 3,
			},
			MutRate:   0.5,
			CrossRate: 0.1,
		}}

	var ga, err = gaConfig.NewGA()

	ga.Callback = func(ga *eaopt.GA) {
		fmt.Printf("Best fitness at generation %d: %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
	}

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

	// var body AllocateRequestBody
	// _ = json.NewDecoder(r.Body).Decode(&body)

	// data := Schedule{Slots: body.Slots, Matchs: body.Matchs, Players: body.Players}

	// fmt.Printf("Melhor solução: %f\n", ga.HallOfFame[0].Fitness)
	// // fmt.Printf("Melhor solução %2.2f (%d/%d) de acertos\n", (1.0-ga.HallOfFame[0].Fitness)*100, int((1.0-ga.HallOfFame[0].Fitness)*float64(len(body.Matchs))), len(body.Matchs))
	// fmt.Println("Tempo total (segundos): ", time.Now().Sub(first).Seconds())

	// slots := ga.HallOfFame[0].Genome.(Schedule).Genome

	// for slotIdx := 0; slotIdx < len(slots); slotIdx++ {
	// 	slot := data.Slots[slotIdx]
	// 	matchID := slots[slotIdx]

	// 	if matchID != -1 {

	// 		fmt.Println(matchID, slot.ID)

	// 		for _, match := range data.Matchs {

	// 			if match.ID == matchID {

	// 				player1, _ := FindPlayerByID(match.Player1, data.Players)
	// 				player2, _ := FindPlayerByID(match.Player2, data.Players)

	// 				if !Contains(player1.AbleSlots, slot.ID) || !Contains(player2.AbleSlots, slot.ID) {
	// 					fmt.Println("Erro aqui")
	// 				}

	// 				break
	// 			}
	// 		}

	// 	}
	// }

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

	// tmpSlots := make([]Slot, len(data.Slots))
	// copy(tmpSlots, data.Slots)

	// response := AllocateRequestResponse{Slots: tmpSlots}

	// for index := 0; index < len(data.Slots); index++ {
	// 	match, _ := FindMatchByID(slots[index], data.Matchs)

	// 	response.Slots[index].Match = match
	// }

	// json.NewEncoder(w).Encode(response)
}

func main() {

	rand.Seed(time.Now().Unix())

	r := mux.NewRouter()
	r.HandleFunc("/allocate", GenerateAllocation).Methods("POST")

	http.ListenAndServe(":8000", r)

}
