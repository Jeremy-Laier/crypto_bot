CREATE TABLE IF NOT EXISTS coin (
    id              INTEGER NOT NULL,
    symbol          TEXT NOT NULL,
    name            TEXT NOT NULL,
    market_cap      DECIMAL NOT NULL,
    price           DECIMAL NOT NULL,
    day_volume      DECIMAL NOT NULL,
    PRIMARY KEY (id, symbol)
);