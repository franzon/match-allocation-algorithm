
import names
import random


def generate_slots(n):
    return [{"id": str(i), "description": "Slot " + str(i)} for i in range(n)]


def generate_players(n, slots):

    players = []

    for i in range(n):
        p = {"id": i, "name": names.get_first_name(), "unable_slots": []}

        for s in slots:
            if random.random() < 0.2:
                p["unable_slots"].append(s["id"])

        if len(p["unable_slots"]) == 0:
            p["unable_slots"].appepnd(random.choice(slots)["id"])

        players.append(p)

    return players


def generate_matchs(players):

    matchs = []

    for i in range(len(players) // 2):
        matchs.append(
            {"id": i, "player1": players[i*2]["id"], "player2": players[i*2+1]["id"]})

    return matchs
