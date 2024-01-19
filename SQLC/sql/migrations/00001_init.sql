-- +goose Up
CREATE TABLE categories (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    name TEXT,
    description TEXT
);

INSERT INTO categories (id, name, description) VALUES 
('00000000-0000-0000-0000-000000000001', 'Category1', 'Description1'),
('00000000-0000-0000-0000-000000000002', 'Category2', 'Description2'),
-- ... (outros INSERTs)
('00000000-0000-0000-0000-000000099999', 'Category99999', 'Description99999'),
('00000000-0000-0000-0000-000000100000', 'Category100000', 'Description100000');

CREATE TABLE courses (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    name TEXT,
    category_id VARCHAR(36) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
 
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE courses;
DROP TABLE categories;

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd