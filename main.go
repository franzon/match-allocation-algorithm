package main

type Player struct {
	ID           int
	Name         string
	AbleSlotsIds []int // índices disponíveis
}

type Match struct {
	ID        int
	Player1Id int
	Player2Id int
}

type Slot struct {
	ID          string
	Description string
}

type Individual struct {
	Genome  []int
	Fitness float32
}

// type Slot struct {
// 	Id          int
// 	Description string
// 	Match       Match
// }

// // HasMatch Função temporária para verificar se Slot tem uma Match atribuida
// func (slot *Slot) HasMatch() bool {
// 	return slot.Match.Player1.Name != "" || slot.Match.Player2.Name != ""
// }

// type Player struct {
// 	Name        string
// 	UnableSlots []Slot
// }

// type Match struct {
// 	Id      int
// 	Player1 Player
// 	Player2 Player
// }

// type Individual struct {
// 	Genome  []int
// 	Fitness float32
// }

// // ComputeFitness Calcula quantas vezes as restrições foram quebradas
// func (individual *Individual) ComputeFitness() {
// 	constraintFailedCount := 0

// 	for _, slot := range individual.Genome {

// 		if slot.HasMatch() {

// 			for _, slotConstraint := range slot.Match.Player1.UnableSlots {
// 				if slotConstraint.Id == slot.Id {
// 					constraintFailedCount++
// 				}
// 			}

// 			for _, slotConstraint := range slot.Match.Player2.UnableSlots {
// 				if slotConstraint.Id == slot.Id {
// 					constraintFailedCount++
// 				}
// 			}

// 		}
// 	}

// 	if constraintFailedCount == 0 {
// 		fmt.Println("Solução encontrada!")
// 		individual.Fitness = 1.0
// 	} else {
// 		individual.Fitness = float32(1 / (constraintFailedCount + 1))
// 	}
// }

// type Population struct {
// 	Individuals []Individual
// }

// // BuildPopulation Inicia população com genomas aleatórios
// // TODO: otimizar geração para abrangir mais possibilidades
// func BuildPopulation(size int, slots []Slot, matchs []Match) Population {

// 	rand.Seed(time.Now().UTC().UnixNano())

// 	population := Population{Individuals: make([]Individual, 0)}
// 	for index := 0; index < size; index++ {

// 		tmpMatchs := make([]Match, len(matchs))
// 		copy(tmpMatchs, matchs)

// 		individual := Individual{Genome: slots, Fitness: 0}

// 		for matchIdx := 0; matchIdx < len(matchs); matchIdx++ {

// 			pickIndex := 0
// 			if len(tmpMatchs) > 0 {
// 				pickIndex = rand.Intn(len(tmpMatchs))
// 			}

// 			individual.Genome[matchIdx].Match = tmpMatchs[pickIndex]

// 			if len(tmpMatchs) > 1 {
// 				tmpMatchs[len(tmpMatchs)-1], tmpMatchs[pickIndex] = tmpMatchs[pickIndex], tmpMatchs[len(tmpMatchs)-1]
// 				tmpMatchs = tmpMatchs[:len(tmpMatchs)-1]
// 			} else if len(tmpMatchs) == 0 {
// 				tmpMatchs = tmpMatchs[:0]
// 			}
// 		}

// 		population.Individuals = append(population.Individuals, individual)
// 	}

// 	return population
// }

// // ComputeFitness Calcula o fitness
// func (population *Population) ComputeFitness() Individual {

// 	// var wg sync.WaitGroup

// 	// wg.Add(len(population.Individuals))

// 	for index := 0; index < len(population.Individuals); index++ {

// 		// go func(i int) {

// 		population.Individuals[index].ComputeFitness()
// 		// defer wg.Done()
// 		// }(index)
// 	}

// 	// wg.Wait()

// 	// Ordena do maior fitness para o menor
// 	slice.Sort(population.Individuals[:], func(i, j int) bool {
// 		return population.Individuals[i].Fitness > population.Individuals[j].Fitness
// 	})

// 	return population.Individuals[0]
// }

// // SelectParents Seleciona os pais
// func SelectParents(population Population) Population {
// 	newPopulation := Population{Individuals: make([]Individual, 0)}

// 	for index := 0; index < int(math.Ceil(float64(len(population.Individuals))/2)); index++ {
// 		newPopulation.Individuals = append(newPopulation.Individuals, population.Individuals[index])
// 	}

// 	return newPopulation
// }

// func prettyPrint(i interface{}) string {
// 	s, _ := json.MarshalIndent(i, "", "\t")
// 	return string(s)
// }

// func main() {

// 	slots := []Slot{
// 		Slot{Id: 1, Description: "Dia 1 - Quadra 1"},
// 		Slot{Id: 2, Description: "Dia 1 - Quadra 2"},
// 		Slot{Id: 3, Description: "Dia 2 - Quadra 1"},
// 		Slot{Id: 4, Description: "Dia 2 - Quadra 2"},
// 		Slot{Id: 5, Description: "Dia 3 - Quadra 1"},
// 		Slot{Id: 6, Description: "Dia 3 - Quadra 2"},
// 		Slot{Id: 7, Description: "Dia 4 - Quadra 1"},
// 		Slot{Id: 8, Description: "Dia 4 - Quadra 2"},
// 	}

// 	// Redundante repetir o slot para o mesmo horário... O front deve tratar isso.

// 	players := []Player{
// 		Player{Name: "Jorge", UnableSlots: []Slot{slots[0], slots[1]}},
// 		Player{Name: "Rafael" /*  UnableSlots: []Slot{slots[2], slots[3], slots[4], slots[5]} */},
// 		Player{Name: "João"},
// 		Player{Name: "Maria", UnableSlots: []Slot{slots[1]}},
// 		Player{Name: "Pedro"},
// 		Player{Name: "Lucas"},
// 		Player{Name: "Dennis", UnableSlots: []Slot{slots[2]}},
// 		Player{Name: "Larissa", UnableSlots: []Slot{slots[3]}}}

// 	matchs := []Match{
// 		Match{Player1: players[0], Player2: players[1]},
// 		Match{Player1: players[2], Player2: players[3]},
// 		Match{Player1: players[4], Player2: players[5]},
// 		Match{Player1: players[6], Player2: players[7]}}

// 	pop1 := BuildPopulation(100, slots, matchs)

// 	// fmt.Println(prettyPrint(pop1))

// 	pop1.ComputeFitness()

// 	parents := SelectParents(pop1)
// 	fmt.Print(len(parents.Individuals))

// 	// players := []string{"Jorge", "Rafael", "João", "Maria", "Pedro", "Lucas", "Dennis", "Larissa"}

// 	// matchs := make([]Match, 0)
// 	// matchs = append(matchs, Match{player1: "Jorge", player2: "Rafael"})
// 	// matchs = append(matchs, Match{player1: "João", player2: "Maria"})
// 	// matchs = append(matchs, Match{player1: "Pedro", player2: "Lucas"})
// 	// matchs = append(matchs, Match{player1: "Dennis", player2: "Larissa"})

// 	// slots := []string{"1A", "1B", "2A", "2B", "3A", "3B", "4A", "4B"}

// }
