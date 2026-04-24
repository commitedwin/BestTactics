CREATE TABLE IF NOT EXISTS players(
    puuid TEXT PRIMARY KEY,
    processed BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS boards(
    id BIGSERIAL PRIMARY KEY ,
    gameId TEXT NOT NULL,
    puuid TEXT NOT NULL,
    placement INT NOT NULL,
    tft_set_number  INT NOT NULL,
    level INT NOT NULL,
    units JSONB NOT NULL,
    traits JSONB NOT NULL,
    UNIQUE(gameId, puuid)
);