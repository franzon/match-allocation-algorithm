package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/MaxHalford/eaopt"
)

// ResponseBody Representação da resposta
type ResponseBody struct {
	NGoodSlots    int    `json:"n_full_slots"`
	NAverageSlots int    `json:"n_average_slots"`
	NBadSlots     int    `json:"n_bad_slots"`
	Slots         []Slot `json:"slots"`
}

type Slot struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Match       Match  `json:"match"`
	Status      string `json:"status"`
}

type Player struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	AbleSlots []string `json:"able_slots"`
	Status    bool     `json:"status"`
}

type Match struct {
	ID      int
	Player1 Player `json:"player1"`
	Player2 Player `json:"player2"`
}

// ScheduleData Dados necessários para calcular o fitness
type ScheduleData struct {
	Slots   []Slot
	Players []Player
	Matchs  []Match
}

// Schedule Representação do cromossomo
type Schedule struct {
	Genome []int
	Data   ScheduleData
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

func FindMatchByID(id int, matchs []Match) (Match, error) {
	for _, p := range matchs {
		if p.ID == id {
			return p, nil
		}
	}
	return Match{}, fmt.Errorf("Couldn't found Match")
}

func (X Schedule) BuildResponseFromSchedule() ResponseBody {

	body := ResponseBody{
		Slots:         make([]Slot, len(X.Data.Slots)),
		NGoodSlots:    0,
		NAverageSlots: 0,
	}

	copy(body.Slots, X.Data.Slots)

	// Número de condições satisfeitas
	count := 0

	for slotIndex, matchID := range X.Genome {

		slot := X.Data.Slots[slotIndex]

		// Pula os slots vazios
		if matchID != -1 {
			match, _ := FindMatchByID(matchID, X.Data.Matchs)

			body.Slots[slotIndex].Match = match

			tmpCount := count

			if Contains(match.Player1.AbleSlots, slot.ID) && Contains(match.Player2.AbleSlots, slot.ID) {

				body.Slots[slotIndex].Status = "good"
				body.NGoodSlots++
				count += 2

				body.Slots[slotIndex].Match.Player1.Status = true

				body.Slots[slotIndex].Match.Player2.Status = true

			} else {

				if Contains(match.Player1.AbleSlots, slot.ID) {
					body.Slots[slotIndex].Status = "average"
					body.NAverageSlots++
					count++

					body.Slots[slotIndex].Match.Player1.Status = true
				}

				if Contains(match.Player2.AbleSlots, slot.ID) {
					body.Slots[slotIndex].Status = "average"
					body.NAverageSlots++
					count++

					body.Slots[slotIndex].Match.Player2.Status = true
				}
			}

			if count == tmpCount {
				body.Slots[slotIndex].Status = "bad"
				body.NBadSlots++
			}
		} else {
			body.Slots[slotIndex].Status = "empty"

		}
	}

	return body
}

// Evaluate Calcula o fitness de um cromossomo
// Quanto menor o valor, melhor o cromossomo resolve o problema
func (X Schedule) Evaluate() (float64, error) {

	// Descartando soluções inválidas

	withoutEmpty := make([]int, 0)

	for index := 0; index < len(X.Genome); index++ {
		if X.Genome[index] != -1 {
			withoutEmpty = append(withoutEmpty, X.Genome[index])
		}
	}

	// Adiciona em um map para remover partidas que aparecem mais de uma vez
	slotMap := make(map[int]bool)
	for index := 0; index < len(withoutEmpty); index++ {
		slotMap[withoutEmpty[index]] = true
	}

	// Retorna um fitness grande para descartar cromossomos onde a partida aparece mais de uma vez
	// ou não aparecem todas as partidas
	if (len(slotMap) != len(withoutEmpty)) || (len(slotMap) != len(X.Data.Matchs)) {
		return math.MaxFloat64, nil
	}

	response := X.BuildResponseFromSchedule()

	return -float64(response.NGoodSlots*10 + response.NAverageSlots), nil
}

// Mutate Aplica a mutação
func (X Schedule) Mutate(rng *rand.Rand) {
	eaopt.MutPermuteInt(X.Genome, 1, rng)
}

// Crossover Aplica o cruzamento
func (X Schedule) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossGNXInt(X.Genome, Y.(Schedule).Genome, 1, rng)
}

// Clone Clona um genoma
func (X Schedule) Clone() eaopt.Genome {
	var Y = Schedule{Genome: make([]int, len(X.Genome)), Data: X.Data}
	copy(Y.Genome, X.Genome)
	return Y
}

// func RemoveIndex(s []Player, index int) []Player {
// 	return append(s[:index], s[index+1:]...)
// }

// GenomeFactory Cria soluções aleatórias
func (data ScheduleData) GenomeFactory(rng *rand.Rand) eaopt.Genome {

	// tmpPlayers := make([]Player, len(data.Players))
	// copy(tmpPlayers, X.Players)

	rndGenome := make([]int, len(data.Slots))

	// Preenche todos os slots como vazios
	for index := 0; index < len(data.Slots); index++ {
		rndGenome[index] = -1
	}

	// Adiciona as partidas nos slots
	for index := 0; index < len(data.Matchs); index++ {
		rndGenome[index] = data.Matchs[index].ID
	}

	// Embaralha o genoma
	rand.Shuffle(len(rndGenome), func(i, j int) {
		rndGenome[i], rndGenome[j] = rndGenome[j], rndGenome[i]
	})

	return Schedule{Genome: rndGenome, Data: data}
}
