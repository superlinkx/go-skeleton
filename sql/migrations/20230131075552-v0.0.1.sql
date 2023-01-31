
-- +migrate Up
CREATE TABLE messages (
  id BIGSERIAL PRIMARY KEY,
  message TEXT NOT NULL
);

-- +migrate Down
DROP TABLE messages;