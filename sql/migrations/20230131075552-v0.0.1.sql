
-- +migrate Up
CREATE TABLE messages (
  id BIGSERIAL PRIMARY KEY,
  message TEXT NOT NULL
);

INSERT INTO messages (message) VALUES
('Hello Database 1!'), ('Hello Database 2!'), ('Hello Database 3!'),
('Hello Database 4!'), ('Hello Database 5!'), ('Hello Database 6!');

-- +migrate Down
DROP TABLE messages;