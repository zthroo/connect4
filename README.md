# connect4
A connect 4 REST API

## to run the server:
* with go installed on your computer in your terminal in the main folder of the repo run `go build`
* In your terminal and same folder run `connect4.exe`
* to test the existing functionality, use the Curls and URLs listed below

## game_table schema:
```
CREATE TABLE game_table (
    game_id INTEGER PRIMARY KEY AUTOINCREMENT,
    players text,
    columns INTEGER NOT NULL,
    rows INTEGER NOT NULL,
    board_state text,
    game_state text,
    winner text
);
```

## game_audit_table schema:
```
CREATE TABLE game_audit_table (
    action_id INTEGER PRIMARY KEY AUTOINCREMENT,
    game_id INTEGER NOT NULL,
    player text NOT NULL,
    type text NOT NULL,
    column int,
    move_number int
)
```

Triggers too complicated to get done in 4 hours probably.

## Curls for creating a new game:
### correct:
```
curl -X POST localhost:8080/drop_token -H "Content-Type: application/json" -d "{\"players\":[\"player1\",\"player2\"],\"columns\":4,\"rows\":4}"
```
### malformed:
```
curl -X POST localhost:8080/drop_token -H "Content-Type: application/json" -d "{\"players\":[\"player1\",\"player2\"],\"columns\":4}"
```

## Useful urls:
### get all in progress games:
```
http://localhost:8080/drop_token
```
### get state of game 1 (in progress)
```
drop_token/1
```
### get state of game 2 (done)
```
drop_token/2
```
### get error for nonexistent game
```
drop_token/3
```
### get error for bad request
```
drop_token/blah
```

