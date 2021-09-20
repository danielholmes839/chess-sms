import pandas as pd

df = pd.read_csv('lichess.csv')

df = df[(df['Result'] == '1-0') | (df['Result'] == '0-1')]
df['Elo'] = df['BlackElo'] + df['WhiteElo'] / 2
df = df.sort_values(by='Elo', ascending=False)

df[:5000].to_csv('lichess_5000.csv')
