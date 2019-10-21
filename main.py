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

                if slot["id"] in player1["able_slots"] and slot["id"] in player2["able_slots"]:
                    fitness += 1

        individual["fitness"] = fitness

        without_empty_slots = [x for x in individual["genome"] if x != -1]

        if len(set(without_empty_slots)) != len(without_empty_slots):
            individual["fitness"] = 0

        if len(set(without_empty_slots)) != len(matchs):
            individual["fitness"] = 0

        if individual["fitness"] == len(matchs):
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


slots = generate_data.generate_slots(16)
players = generate_data.generate_players(32, slots)
matchs = generate_data.generate_matchs(players)

population_size = 1000
iterations = 500
mutation_rate = 0.01


best = {"fitness": 0}

first = time.time()

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


print('Melhor solução: ', best["fitness"])

print("Tempo: ", time.time()-first)
