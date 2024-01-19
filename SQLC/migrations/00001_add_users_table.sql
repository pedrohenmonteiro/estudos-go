-- +goose Up
CREATE TABLE users (
    id INT NOT NULL PRIMARY KEY,
    username VARCHAR(255),
    name VARCHAR(255),
    surname VARCHAR(255)
);

INSERT INTO users (id, username, name, surname) VALUES (1, 'gallant_almeida7', 'Gallant', 'Almeida7');
INSERT INTO users (id, username, name, surname) VALUES (2, 'brave_spence8', 'Brave', 'Spence8');
-- ... (outros INSERTs)
INSERT INTO users (id, username, name, surname) VALUES (99999, 'jovial_chaum1', 'Jovial', 'Chaum1');
INSERT INTO users (id, username, name, surname) VALUES (100000, 'goofy_ptolemy0', 'Goofy', 'Ptolemy0');

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE users;

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
