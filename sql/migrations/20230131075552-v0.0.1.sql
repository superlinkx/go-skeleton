
-- +migrate Up
CREATE TABLE messages (
  id BIGSERIAL PRIMARY KEY,
  message TEXT NOT NULL
);

INSERT INTO messages (message) VALUES ('Hello Database!');

-- +migrate Down
DROP TABLE messages;