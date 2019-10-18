import generate_data
import json
import pyperclip

slots = generate_data.generate_slots(64)
players = generate_data.generate_players(64, slots)
matchs = generate_data.generate_matchs(players)

x = json.dumps(
    {"slots": slots, "players": players, "matchs": matchs}, indent=4)
print(x)
pyperclip.copy(x)
