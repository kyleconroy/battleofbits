CREATE EXTENSION "uuid-ossp";

CREATE TABLE games (
  id varchar(100) PRIMARY KEY
);

INSERT INTO games (id) VALUES ('tictactoe');
INSERT INTO games (id) VALUES ('fourup');

CREATE TABLE bots (
  id   uuid          PRIMARY KEY DEFAULT uuid_generate_v4(),
  slug varchar(500)  NOT NULL UNIQUE,
  game varchar(100)  references games,
  url  varchar(2000) NOT NULL,
  name varchar(500)  NOT NULL
); 

-- This constraint should probably be an ENUM
CREATE TABLE matches (
  id         uuid          PRIMARY KEY DEFAULT uuid_generate_v4(),
  player_one uuid          references bots,
  player_two uuid          references bots,
  winner     smallint      DEFAULT 0,
  created    timestamp     DEFAULT NOW(),
  CONSTRAINT valid_winner  CHECK (winner = 0 OR winner = 1 OR winner = 2)
);

CREATE TABLE moves (
  id      uuid      PRIMARY KEY DEFAULT uuid_generate_v4(),
  match   uuid      references matches,
  created timestamp DEFAULT NOW(),
  state   json      NOT NULL
);
