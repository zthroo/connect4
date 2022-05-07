# connect4
A connect 4 REST API

## game_table schema:
```
CREATE TABLE game_table (
    game_id INTEGER PRIMARY KEY AUTOINCREMENT,
    players text
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
    column int
)
```

Triggers too complicated to get done in 4 hours probably.
