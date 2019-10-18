
import names
import random

slots = [
    {"id": "1A", "description": "Dia  1 - Quadra 1"},
    {"id": "1B", "description": "Dia  1 - Quadra 2"},
    {"id": "2A", "description": "Dia  2 - Quadra 1"},
    {"id": "2B", "description": "Dia  2 - Quadra 2"},
    # {"id": "3A", "description": "Dia  3 - Quadra 1"},
    # {"id": "3B", "description": "Dia  3 - Quadra 2"},
    # {"id": "4A", "description": "Dia  4 - Quadra 1"},
    # {"id": "4B", "description": "Dia  4 - Quadra 2"},
    # {"id": "5A", "description": "Dia  5 - Quadra 1"},
    # {"id": "5B", "description": "Dia  5 - Quadra 2"},
    # {"id": "6A", "description": "Dia  6 - Quadra 1"},
    # {"id": "6B", "description": "Dia  6 - Quadra 2"},
    # {"id": "7A", "description": "Dia  7 - Quadra 1"},
    # {"id": "7B", "description": "Dia  7 - Quadra 2"},
    # {"id": "8A", "description": "Dia  8 - Quadra 1"},
    # {"id": "8B", "description": "Dia  8 - Quadra 2"},
    # {"id": "9A", "description": "Dia  9 - Quadra 1"},
    # {"id": "9B", "description": "Dia  9 - Quadra 2"},
    # {"id": "10A", "description": "Dia  10 - Quadra 1"},
    # {"id": "10B", "description": "Dia  10 - Quadra 2"},
    # {"id": "11A", "description": "Dia  11 - Quadra 1"},
    # {"id": "11B", "description": "Dia  11 - Quadra 2"},
    # {"id": "12A", "description": "Dia  12 - Quadra 1"},
    # {"id": "12B", "description": "Dia  12 - Quadra 2"},
    # {"id": "13A", "description": "Dia  13 - Quadra 1"},
    # {"id": "13B", "description": "Dia  13 - Quadra 2"},
    # {"id": "14A", "description": "Dia  14 - Quadra 1"},
    # {"id": "14B", "description": "Dia  14 - Quadra 2"},
    # {"id": "15A", "description": "Dia  15 - Quadra 1"},
    # {"id": "15B", "description": "Dia  15 - Quadra 2"},
    # {"id": "16A", "description": "Dia  16 - Quadra 1"},
    # {"id": "16B", "description": "Dia  16 - Quadra 2"},
    # {"id": "C1A", "description": "Dia  1 - Quadra 1"},
    # {"id": "C1B", "description": "Dia  1 - Quadra 2"},
    # {"id": "C2A", "description": "Dia  2 - Quadra 1"},
    # {"id": "C2B", "description": "Dia  2 - Quadra 2"},
    # {"id": "C3A", "description": "Dia  3 - Quadra 1"},
    # {"id": "C3B", "description": "Dia  3 - Quadra 2"},
    # {"id": "C4A", "description": "Dia  4 - Quadra 1"},
    # {"id": "C4B", "description": "Dia  4 - Quadra 2"},
    # {"id": "C5A", "description": "Dia  5 - Quadra 1"},
    # {"id": "C5B", "description": "Dia  5 - Quadra 2"},
    # {"id": "C6A", "description": "Dia  6 - Quadra 1"},
    # {"id": "C6B", "description": "Dia  6 - Quadra 2"},
    # {"id": "C7A", "description": "Dia  7 - Quadra 1"},
    # {"id": "C7B", "description": "Dia  7 - Quadra 2"},
    # {"id": "C8A", "description": "Dia  8 - Quadra 1"},
    # {"id": "C8B", "description": "Dia  8 - Quadra 2"},
    # {"id": "C9A", "description": "Dia  9 - Quadra 1"},
    # {"id": "C9B", "description": "Dia  9 - Quadra 2"},
    # {"id": "C10A", "description": "Dia  10 - Quadra 1"},
    # {"id": "C10B", "description": "Dia  10 - Quadra 2"},
    # {"id": "C11A", "description": "Dia  11 - Quadra 1"},
    # {"id": "C11B", "description": "Dia  11 - Quadra 2"},
    # {"id": "C12A", "description": "Dia  12 - Quadra 1"},
    # {"id": "C12B", "description": "Dia  12 - Quadra 2"},
    # {"id": "C13A", "description": "Dia  13 - Quadra 1"},
    # {"id": "C13B", "description": "Dia  13 - Quadra 2"},
    # {"id": "C14A", "description": "Dia  14 - Quadra 1"},
    # {"id": "C14B", "description": "Dia  14 - Quadra 2"},
    # {"id": "C15A", "description": "Dia  15 - Quadra 1"},
    # {"id": "C15B", "description": "Dia  15 - Quadra 2"},
    # {"id": "C16A", "description": "Dia  16 - Quadra 1"},
    # {"id": "C16B", "description": "Dia  16 - Quadra 2"},
]


def generate_players(n):

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
