import random
import generate_data
import time


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
# players = generate_data.generate_players(32)
# matchs = generate_data.generate_matchs(players)

players = [{
    'id': 0,
    'name': 'Tesha',
    'able_slots': ['1A', '1B', '2A', '6A', '8A', '9B', '10A', '10B']
}, {
    'id': 1,
    'name': 'Robert',
    'able_slots': ['4B', '6B', '7A', '7B', '8B', '9A', '9B', '10A']
}, {
    'id': 2,
    'name': 'Sara',
    'able_slots': ['5B', '6B', '7A', '8B', '9B']
}, {
    'id': 3,
    'name': 'Thomas',
    'able_slots': ['2B', '3A', '4A', '4B', '6B', '7B', '8A']
}, {
    'id': 4,
    'name': 'Charlie',
    'able_slots': ['1A', '1B', '3A', '5A', '6B', '7A', '7B', '9A']
}, {
    'id': 5,
    'name': 'Terry',
    'able_slots': ['1A', '2B', '3B', '4B', '5A', '9B', '10A', '10B']
}, {
    'id': 6,
    'name': 'Toby',
    'able_slots': ['1B', '4B', '6A', '6B', '7A', '8A', '8B', '9B', '10B']
}, {
    'id': 7,
    'name': 'Carlos',
    'able_slots': ['3A', '3B', '4A', '5B', '9A']
}, {
    'id': 8,
    'name': 'Kevin',
    'able_slots': ['3B', '4A', '5A', '6B', '7A', '8A', '9A', '10B']
}, {
    'id': 9,
    'name': 'Dan',
    'able_slots': ['1A', '1B', '2A', '3A', '6A', '6B', '7B', '8A', '8B', '10B']
}, {
    'id': 10,
    'name': 'William',
    'able_slots': ['1A', '2A', '3A', '3B', '5A', '6A', '6B', '7A', '7B', '8B', '9A']
}, {
    'id': 11,
    'name': 'Savannah',
    'able_slots': ['1A', '3A', '5A', '6A', '8B', '9A']
}, {
    'id': 12,
    'name': 'Reyes',
    'able_slots': ['1B',
                   '3B', '6B', '9B', '10A'
                   ]
}, {
    'id': 13,
    'name': 'Kelly',
    'able_slots': ['1B', '2A', '3A', '4A', '6B', '7B', '9A', '10B']
}, {
    'id': 14,
    'name': 'Colleen',
    'able_slots': ['1A', '1B', '2B', '3A', '4B', '5A', '6A', '7A', '7B', '9A']
}, {
    'id': 15,
    'name': 'Monica',
    'able_slots': ['1A', '2B', '3A', '3B', '5B', '6A', '6B', '7A',
                   '7B', '8B', '10A'
                   ]
}, {
    'id': 16,
    'name': 'Bridget',
    'able_slots': ['2A', '3A', '3B', '4A', '5A', '6A', '6B', '8A', '9A', '10B']
}, {
    'id': 17,
    'name': 'Ryan',
    'able_slots': ['1A', '1B', '3B', '5A', '6A', '6B', '7A', '8B', '9B', '10B']
}, {
    'id': 18,
    'name': 'Dale',
    'able_slots': ['1A', '1B', '2B', '6B', '7A']
}, {
    'id': 19,
    'name': 'June',
    'able_slots': ['1A', '2A', '2B', '5A', '6A', '7A', '8B', '9B']
}, {
    'id': 20,
    'name': 'Matthew',
    'able_slots': ['1B', '2A', '2B', '4A', '8A', '8B', '9B', '10A']
}, {
    'id': 21,
    'name': 'Louise',
    'able_slots': ['1A', '2A', '2B', '4A', '5A', '5B',
                   '6A', '10A', '10B'
                   ]
}, {
    'id': 22,
    'name': 'Lionel',
    'able_slots': ['2A', '4A', '9B', '10A']
}, {
    'id': 23,
    'name': 'Eric',
    'able_slots': ['1B', '2B', '4A', '5A', '7B', '9B']
}, {
    'id': 24,
    'name': 'Marcell',
    'able_slots': ['1A', '1B', '4A', '5A', '6A', '7B', '8B']
}, {
    'id': 25,
    'name': 'Jon',
    'able_slots': ['1A', '1B', '2A', '3A', '4B', '7A', '7B', '8B', '9B']
}, {
    'id': 26,
    'name': 'Edna',
    'able_slots': ['1A', '1B', '2A', '3A', '3B', '5A', '5B', '7A', '7B', '8A', '9B', '10A']
}, {
    'id': 27,
    'name': 'Mark',
    'able_slots': ['1A', '5A', '7A', '10A']
}, {
    'id': 28,
    'name': 'Kent',
    'able_slots': ['4B', '5B', '8A', '8B', '9A', '10A']
}, {
    'id': 29,
    'name': 'John',
    'able_slots': ['1A', '2B', '5B', '6B', '7A', '9A',
                   '10B'
                   ]
}, {
    'id': 30,
    'name': 'Sharon',
    'able_slots': ['1A', '1B', '2A', '2B', '4A', '4B', '5B', '7A', '7B', '8B', '10A']
}, {
    'id': 31,
    'name': 'Andrew',
    'able_slots': ['3B', '4A', '6B', '7A', '8A', '8B', '9A']
}]
matchs = [{"id": 0, "player1": 24, "player2": 3}, {"id": 1, "player1": 12, "player2": 10}, {"id": 2, "player1": 19, "player2": 6}, {"id": 3, "player1": 29, "player2": 18}, {"id": 4, "player1": 4, "player2": 14}, {"id": 5, "player1": 25, "player2": 22}, {"id": 6, "player1": 8, "player2": 27}, {"id": 7, "player1": 9, "player2": 7}, {
    "id": 8, "player1": 15, "player2": 21}, {"id": 9, "player1": 5, "player2": 26}, {"id": 10, "player1": 23, "player2": 28}, {"id": 11, "player1": 17, "player2": 13}, {"id": 12, "player1": 0, "player2": 2}, {"id": 13, "player1": 11, "player2": 20}, {"id": 14, "player1": 30, "player2": 1}, {"id": 15, "player1": 31, "player2": 16}]


population_size = 50
iterations = 300
mutation_rate = 0.01

print(players)
print('')
print(matchs)

best = {"fitness": 0}

first = time.time()

pop = init_population(population_size, len(slots), matchs)

for it in range(iterations):
    compute_fitness(pop, matchs, players, slots)

    # print('Melhor fitness: ', pop[0]["fitness"])
    if pop[0]["fitness"] > best["fitness"]:
        best = pop[0]

    pop = selection(pop)
    pop = cross_over(pop)
    pop = elitism(pop, matchs, players, slots)
    mutation(pop, mutation_rate)


print('Melhor solução: ', best["fitness"])

print("Tempo: ", time.time()-first)
