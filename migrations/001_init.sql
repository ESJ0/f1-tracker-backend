CREATE TABLE drivers (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    team        VARCHAR(100) NOT NULL,
    nationality VARCHAR(100) NOT NULL,
    number      INT          NOT NULL UNIQUE,
    image_url   TEXT,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);

CREATE TABLE races (
    id          SERIAL PRIMARY KEY,
    grand_prix  VARCHAR(100) NOT NULL,
    circuit     VARCHAR(100) NOT NULL,
    country     VARCHAR(100) NOT NULL,
    race_date   DATE         NOT NULL,
    image_url   TEXT,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);

CREATE TABLE results (
    id              SERIAL PRIMARY KEY,
    driver_id       INT REFERENCES drivers(id) ON DELETE CASCADE,
    race_id         INT REFERENCES races(id)   ON DELETE CASCADE,
    position        INT  NOT NULL,
    points          INT  NOT NULL DEFAULT 0,
    fastest_lap     BOOLEAN DEFAULT FALSE,
    created_at      TIMESTAMP DEFAULT NOW(),
    UNIQUE(driver_id, race_id)
);