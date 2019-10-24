package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/MaxHalford/eaopt"
	"github.com/gorilla/mux"
)

// RequestBody Representação da requisição
type RequestBody struct {
	Slots  []Slot  `json:"slots"`
	Matchs []Match `json:"matchs"`
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
		EarlyStop: earlyStop, NPops: 1, PopSize: 80, HofSize: 1, NGenerations: 500, ParallelEval: true,
		Model: eaopt.ModGenerational{
			Selector: eaopt.SelTournament{
				NContestants: 3,
			},
			MutRate:   0.5,
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

	scheduleData := ScheduleData{Slots: body.Slots, Matchs: body.Matchs}

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

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	addr := ":" + os.Getenv("PORT")
	http.ListenAndServe(addr, r)

}
