from scipy import optimize


import random
import generate_data
import time


slots = generate_data.generate_slots(16)
players = generate_data.generate_players(32, slots)
matchs = generate_data.generate_matchs(players)

for m in matchs:
    m["player1"] = [p for p in players if p["id"] == m["player1"]][0]
    m["player2"] = [p for p in players if p["id"] == m["player2"]][0]


def evaluate_minimize(s):
    count = 0
    for slotIndex, matchId in enumerate(s):
        slot = slots[slotIndex]
        if matchId != -1:

            print(matchs)

            match = [m for m in matchs if m["id"] == matchId][0]

            if slot["id"] not in match["player1"]["able_slots"]:
                count += 1
            if slot["id"] not in match["player2"]["able_slots"]:
                count += 1

    return count


def evaluate_good(s):
    count = 0
    for slotIndex, matchId in enumerate(s):
        slot = slots[slotIndex]
        if matchId != -1:

            match = [m for m in matchs if m["id"] == matchId][0]

            if slot["id"] in match["player1"]["able_slots"] and slot["id"] in match["player2"]["able_slots"]:
                count += 1

    return count


init_state = [-1 for x in range(len(slots))]
for j in range(len(matchs)):
    init_state[j] = matchs[j]["id"]

random.shuffle(init_state)

print('Random: ', evaluate_good(init_state))

print(optimize.basinhopping(evaluate_minimize, init_state))
