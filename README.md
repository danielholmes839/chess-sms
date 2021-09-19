# Hack The North 2021 - "ChesSMS"

[Watch the Demo!](https://youtu.be/5lVBIMsYSuU)

Allows users to solve chess puzzles by texting **613-704-4683** the following commands:

```bash
commands
puzzle
<move>                 # Examples: "e2e4", "e4", "Qxd8", "Qxd8#" (not case sensitive)
hint
answer
```
**Conversational Flow**

1. Start by sending ```puzzle``` this will randomly select 1 of 500 puzzles for you to solve
2. Send your solution. Your move can be in PGN form, or if you're unfamiliar with PGN form you can write the name of the start and end squares ```e2e4``` which also works.
3. If you got the answer wrong try asking for a hint by sending ```hint```. This will tell you the type of piece you need to move.
4. If you still can't figure out the answer send ```answer```
5. Only after getting the correct answer or sending ```answer``` will you you be allowed to access a new puzzle.
6. Ask for another puzzle by sending ```puzzle```

## How I Built It

1. Setup SMS using Twilio

2. Setup a REST API that Twilio can use as a callback. I used a Linux VPS on Vultr for deploying the REST API, and re-used an old domain name for setting up HTTPS.

3. The REST API and entire backend was written in Go. I generally enjoy using Go, and I found a nice [chess library](https://github.com/notnil/chess) that allowed me to do everything I needed including creating SVGs of chess positions. Those SVGs could then be converted into PNGs so that they could be sent using Twilio. 

4. Creating the chess positions. I found a [dataset](https://web.chessdigits.com/data) from Lichess of 200,000  games that were played online. I filtered this down using Python to the top 5000 highest elo games that did not end in a draw. Once I had those 5000 games I filtered out games that did not end in checkmate meaning a player decided to resign. This was actually around 75% of those games (high elo players know when to resign :D). I then picked the top 500 games, and converted them to JSONs containing: the [PGN notation](https://en.wikipedia.org/wiki/Portable_Game_Notation), the last move, the last piece moved (for showing users hints), and who's move it is. [puzzles.json](https://github.com/danielholmes839/htn-2021/blob/master/python/puzzles.json)
