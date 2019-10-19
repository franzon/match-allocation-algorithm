
import names
import random


def generate_slots(n):
    return [{"id": str(i), "description": "Slot " + str(i)} for i in range(n)]


def generate_players(n, slots):

    players = []

    for i in range(n):
        p = {"id": i, "name": names.get_first_name(), "able_slots": []}

        for s in slots:
            if random.random() < 0.5:
                p["able_slots"].append(s["id"])

        if len(p["able_slots"]) == 0:
            p["able_slots"].append(random.choice(slots)["id"])

        players.append(p)

    return players


def generate_matchs(players):

    matchs = []

    tmp = players.copy()
    i = 0

    while len(tmp) > 0:
        player1 = random.choice(tmp)
        tmp.remove(player1)

        player2 = random.choice(tmp)
        tmp.remove(player2)

        matchs.append(
            {"id": i, "player1": player1["id"], "player2": player2["id"]})

        i += 1

    return matchs
