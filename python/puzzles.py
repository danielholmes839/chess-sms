import json
import pandas as pd

MOVES = [f'Move_ply_{i}' for i in range(1, 120)]     # Columns for the first 30 moves
df = pd.read_csv('lichess_5000.csv')


def chunks(lst, n):
    """Yield successive n-sized chunks from lst."""
    for i in range(0, len(lst), n):
        yield lst[i:i + n]


def pgn(moves) -> str:
    moves = [f'{i+1}. {" ".join(pair)}'for i, pair in enumerate(chunks(moves, 2))]
    return ' '.join(moves)

def hint(move) -> str:
    if move[0].islower():
        return 'Pawn'

    if move[0] == 'Q':
        return 'Queen'

    if move[0] == 'K':
        return 'King'

    if move[0] == 'R':
        return 'Rook'

    if move[0] == 'B':
        return 'Bishop'

    if move[0] == 'N':
        return 'Knight'


without_resignation = []


for i, df_row in df.iterrows():
    moves = [move for move in list(df_row[MOVES]) if type(move) == str]
    if moves[-1][-1] == '#':

        color = 'White'
        if len(moves) % 2 == 0:
            color = 'Black'

        without_resignation.append({
            'color': color,
            'hint': hint(moves[-1]),
            'answer': moves[-1],
            'pgn': pgn(moves)
        })


with open('puzzles.json', 'w') as f:
    json.dump(without_resignation[:500], f)
