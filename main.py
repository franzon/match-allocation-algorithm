import random
import generate_data


def init_population(population_size, num_slots, matchs):

    population = []
    for i in range(population_size):
        slots = [-1 for x in range(num_slots)]
        for j in range(len(matchs)):
            slots[j] = matchs[j]["id"]

        random.shuffle(slots)
        population.append({"genome": slots, "fitness": 0})

    return population


def compute_fitness(population, matchs, players, slots):
    for individual in population:

        fitness = 0

        for slot_idx, match_id in enumerate(individual["genome"]):

            slot = slots[slot_idx]
            match = [m for m in matchs if m["id"] == match_id]

            if len(match):
                player1 = [p for p in players if p["id"]
                           == match[0]["player1"]][0]
                player2 = [p for p in players if p["id"]
                           == match[0]["player2"]][0]

                if slot["id"] in player1["able_slots"]:
                    fitness += 1

                if slot["id"] in player2["able_slots"]:
                    fitness += 1

        individual["fitness"] = fitness

        without_empty_slots = [x for x in individual["genome"] if x != -1]

        if len(set(without_empty_slots)) != len(without_empty_slots):
            individual["fitness"] = 0

        if len(set(without_empty_slots)) != len(matchs):
            individual["fitness"] = 0

        if individual["fitness"] == len(matchs) * 2:
            print('Solução encontrada')
            print(individual["genome"])
            exit(0)

    population.sort(
        key=lambda individual: individual["fitness"], reverse=True)


def selection(population):
    sum_fitness = sum([individual["fitness"] for individual in population])

    if sum_fitness == 0:
        exit(-1)

    roulette = []
    acc = 0.0

    for individual in population:
        acc += individual["fitness"] / sum_fitness
        roulette.append(acc)

    selected_individuals = []

    for i in range(len(population)):
        rnd = random.random()

        for j in range(len(roulette)):
            if rnd <= roulette[j]:
                selected_individuals.append(population[j])
                break

    return selected_individuals


def cross_over(population):

    new_population = population.copy()

    for i in range(len(population) // 2):

        father_1 = random.choice(population)
        father_2 = random.choice(population)

        cross_point = random.randrange(len(father_1["genome"]))

        child_1 = {"genome": father_1["genome"][:cross_point] +
                   father_2["genome"][cross_point:], "fitness": 0}
        child_2 = {"genome": father_2["genome"][:cross_point] +
                   father_1["genome"][cross_point:], "fitness": 0}

        new_population.append(child_1)
        new_population.append(child_2)

    return new_population


def elitism(population, matchs, players, slots):
    compute_fitness(population, matchs, players, slots)
    return population[:len(population) // 2]


def mutation(population, mutation_rate):
    for individual in population:
        for i in range(len(individual["genome"])):
            if random.random() < mutation_rate:

                free_slots = [i for i in range(
                    len(individual["genome"])) if individual["genome"][i] == -1]

                if len(free_slots) > 0:
                    slot_idx = random.choice(free_slots)

                    individual["genome"][slot_idx] = individual["genome"][i]
                    individual["genome"][i] = -1


# slots = [
#     {"id": "1A", "description": "Dia  1 - Quadra 1"},
#     {"id": "1B", "description": "Dia  1 - Quadra 2"},
#     {"id": "2A", "description": "Dia  2 - Quadra 1"},
#     {"id": "2B", "description": "Dia  2 - Quadra 2"},
#     {"id": "3A", "description": "Dia  3 - Quadra 1"},
#     {"id": "3B", "description": "Dia  3 - Quadra 2"},
#     {"id": "4A", "description": "Dia  4 - Quadra 1"},
#     {"id": "4B", "description": "Dia  4 - Quadra 2"},
# ]

# players = [
#     {
#         "id": 1,
#         "name": "Jorge",
#         "able_slots": ['1A', '3B']
#     },
#     {
#         "id": 2,
#         "name": "Rafael",
#         "able_slots": ['3B']
#     },
#     {
#         "id": 3,
#         "name": "João",
#         "able_slots": ['1A', '1B', '3A', '3B', '4A']
#     },
#     {
#         "id": 4,
#         "name": "Larissa",
#         "able_slots": ['3B', '4A']
#     },
#     {
#         "id": 5,
#         "name": "Maria",
#         "able_slots": ['3B', '4A', '4B']
#     },
#     {
#         "id": 6,
#         "name": "Pedro",
#         "able_slots": ['4B']
#     },
#     {
#         "id": 7,
#         "name": "Lucas",
#         "able_slots": ['2B', '3A', '3B']
#     },
#     {
#         "id": 8,
#         "name": "Dennis",
#         "able_slots": ['3A', '3B']
#     }
# ]


# matchs = [
#     {"id": 1, "player1": 1, "player2": 2},
#     {"id": 2, "player1": 3, "player2": 4},
#     {"id": 3, "player1": 5, "player2": 6},
#     {"id": 4, "player1": 7, "player2": 8},
# ]

slots = generate_data.slots
players = generate_data.generate_players(32)
matchs = generate_data.generate_matchs(players)


# não encontrou

# players = [{'id': 0, 'name': 'Deborah', 'able_slots': ['1A', '2A', '3A', '4B', '6A', '6B', '7A', '8B', '9A', '9B', '10A']}, {'id': 1, 'name': 'Steve', 'able_slots': ['2B', '3A', '5A', '6A', '8A', '9A', '9B', '10B']}, {'id': 2, 'name': 'Madeline', 'able_slots': ['1A', '2B', '4A', '5B', '6B', '7B', '10A', '10B']}, {'id': 3, 'name': 'Erin', 'able_slots': ['1A', '2A', '3A', '4A', '5A', '5B', '6A', '7B', '9B']}, {'id': 4, 'name': 'Lorna', 'able_slots': ['2B', '3A', '4B', '5B', '6A', '6B', '7A', '7B', '8A', '8B', '9B', '10A']}, {'id': 5, 'name': 'Ernest', 'able_slots': ['1A', '2B', '4A', '5A', '6A', '8A', '9B', '10A', '10B']}, {'id': 6, 'name': 'Irving', 'able_slots': ['1B', '3B', '4A',
#                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 '5A', '6B', '9A']}, {'id': 7, 'name': 'Robert', 'able_slots': ['1A', '1B', '3A', '6A', '7B', '8B']}, {'id': 8, 'name': 'Daniel', 'able_slots': ['1A', '1B', '3B', '7A', '7B', '8B', '9A', '10B']}, {'id': 9, 'name': 'Anna', 'able_slots': ['1B', '2B', '4B', '5B', '7A', '10A']}, {'id': 10, 'name': 'Lillian', 'able_slots': ['2B', '4B', '6A', '6B', '7A', '8B', '10A', '10B']}, {'id': 11, 'name': 'Edward', 'able_slots': ['2A', '2B', '5A', '6A', '6B', '9B', '10A', '10B']}, {'id': 12, 'name': 'Doris', 'able_slots':
#                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      ['3A', '4B', '5A', '7A']}, {'id': 13, 'name': 'Todd', 'able_slots': ['2B', '3B', '4A', '8A', '8B', '9A']}, {'id': 14, 'name': 'Duane', 'able_slots': ['2B', '3B', '4B', '5A', '5B', '6A', '8A', '8B', '9A', '10A', '10B']}, {'id': 15, 'name': 'Ronda', 'able_slots': ['3B', '4B', '6A', '8A', '8B', '10B']}, {'id': 16, 'name': 'Maria', 'able_slots': ['2A', '3B', '5A', '6B', '7B', '8A', '10A', '10B']}, {'id': 17, 'name': 'James',
#                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    'able_slots': ['1A', '3B', '4A', '7A', '7B', '8A', '8B', '9A', '9B', '10A']}, {'id':
#                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   18, 'name': 'Janet', 'able_slots': ['1A', '1B', '2A', '5B', '6A', '6B', '7A', '8A', '8B', '9A', '10A']}, {'id': 19, 'name': 'Martin', 'able_slots': ['3A', '6A', '6B', '7A', '8A', '8B', '9A', '10A']}, {'id': 20, 'name': 'Ruth', 'able_slots': ['3B', '5A', '5B', '8A']}, {'id': 21, 'name': 'Laura', 'able_slots': ['1B', '2A', '3B', '4A', '4B', '6A', '7B', '8B', '9B', '10B']}, {'id': 22, 'name': 'Terresa', 'able_slots': ['3B',
#                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      '7A', '9A']}, {'id': 23, 'name': 'Carlos', 'able_slots': ['2A', '2B', '3B', '4A', '4B', '5A', '6A', '6B', '7B', '10A', '10B']}, {'id': 24, 'name': 'John', 'able_slots': ['1B', '3B', '6B', '8B', '9B', '10A', '10B']}, {'id': 25, 'name': 'Deborah', 'able_slots': ['1A', '1B', '3A', '7B']}, {'id': 26, 'name': 'Russell', 'able_slots': ['3B', '4A', '5A', '6A', '7B', '8A', '10A', '10B']}, {'id': 27, 'name': 'Tisha', 'able_slots': ['1A', '2B', '3A', '3B', '7B', '8A', '10B']}, {'id': 28, 'name': 'Jane', 'able_slots': ['2A', '4B', '5A', '6A', '7B', '8A', '8B', '9A', '9B']}, {'id': 29, 'name': 'James', 'able_slots': ['1B', '2B', '3B', '4A', '5A', '7B', '10A', '10B']}, {'id': 30, 'name': 'Clarence', 'able_slots': ['3B', '4B', '6A', '6B', '7A', '7B', '10B']}, {'id': 31, 'name': 'Michael', 'able_slots': ['4A', '4B', '5A', '7B', '9B']}]
# matchs = [{'id': 0, 'player1': 2, 'player2': 29}, {'id': 1, 'player1': 10, 'player2': 23}, {'id': 2, 'player1': 1, 'player2': 31}, {'id': 3, 'player1': 5, 'player2': 22}, {'id': 4, 'player1': 15, 'player2': 30}, {'id': 5, 'player1': 24, 'player2': 18}, {'id': 6, 'player1': 17, 'player2': 4}, {'id': 7, 'player1': 13, 'player2': 8}, {'id': 8, 'player1': 26, 'player2': 3}, {'id': 9, 'player1': 21, 'player2': 7}, {'id': 10, 'player1': 16, 'player2': 25}, {'id': 11, 'player1': 28, 'player2': 12}, {'id': 12, 'player1':
#                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   27, 'player2': 6}, {'id': 13, 'player1': 11, 'player2': 0}, {'id': 14, 'player1': 20, 'player2': 9}, {'id': 15, 'player1': 19, 'player2': 14}]

population_size = 50
iterations = 1000
mutation_rate = 0.01

print(players)
print('')
print(matchs)

best = {"fitness": 0}

pop = init_population(population_size, len(slots), matchs)

for it in range(iterations):
    compute_fitness(pop, matchs, players, slots)

    print('Melhor fitness: ', pop[0]["fitness"])
    if pop[0]["fitness"] > best["fitness"]:
        best = pop[0]

    pop = selection(pop)
    pop = cross_over(pop)
    pop = elitism(pop, matchs, players, slots)
    mutation(pop, mutation_rate)


print('Melhor solução: ', best["genome"])
