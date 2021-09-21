# Chess Puzzles Over SMS - Hack The North 2021

[Watch the Demo!](https://youtu.be/5lVBIMsYSuU)

I created an application that allows users to solve chess puzzles by texting **613-704-4683** the following commands:

```bash
commands
puzzle
<move>                 # Examples: "e2e4", "e4", "Qxd8", "Qxd8#" (not case sensitive)
hint
answer
```
**General Conversation Flow**

1. Start by sending ```puzzle``` this will randomly select 1 of 500 puzzles for you to solve
2. Send your solution. Your move can be in PGN form, or if you're unfamiliar with PGN form you can write the name of the start and end squares ```e2e4``` which also works.
3. If you got the answer wrong try asking for a hint by sending ```hint```. This will tell you the type of piece you need to move.
4. If you still can't figure out the answer send ```answer```
5. Only after getting the correct answer or sending ```answer``` will you you be allowed to access a new puzzle.
6. Ask for another puzzle by sending ```puzzle```


## Technical Details

The application is a REST API written in Go that provides a callback endpoint for Twilio to send text messages. 

The puzzles were created using a [dataset from Lichess](https://web.chessdigits.com/data) containing 200,000 games played online. I wanted puzzles that contained a checkmate in one move to keep things simple. So I wrote a Python script to filter out 500 games from high rated players that ended in checkmate. Then I saved the PGN, the last move, who played the last move, and the type of piece used on the last move (this was used to give hints). The puzzles can found in the **[puzzles.json](https://github.com/danielholmes839/htn-2021/blob/master/python/puzzles.json) file**.

## Next Steps

I would consider the application finished. But there are other features that would be interesting to implement:
- Puzzles that take multiple moves to solve
- Playing a full chess game with an opponent or engine.

