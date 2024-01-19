-- +goose Up
CREATE TABLE categories (
    id INT NOT NULL PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255)
);

INSERT INTO categories (id, name, description) VALUES (1, 'Category1', 'Description1');
INSERT INTO categories (id, name, description) VALUES (2, 'Category2', 'Description2');
INSERT INTO categories (id, name, description) VALUES (99999, 'Category99999', 'Description99999');
INSERT INTO categories (id, name, description) VALUES (100000, 'Category100000', 'Description100000');

CREATE TABLE courses (
    id INT NOT NULL PRIMARY KEY,
    name VARCHAR(255),
    category_id INT,
    description VARCHAR(255),
    price DECIMAL(10, 2),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

INSERT INTO courses (id, name, category_id, description, price) VALUES (1, 'Gallant', 1, 'DescriptionGallant', 19.99);
INSERT INTO courses (id, name, category_id, description, price) VALUES (2, 'Brave', 2, 'DescriptionBrave', 29.99);
INSERT INTO courses (id, name, category_id, description, price) VALUES (99999, 'Jovial', 99999, 'DescriptionJovial', 39.99);
INSERT INTO courses (id, name, category_id, description, price) VALUES (100000, 'Goofy', 100000, 'DescriptionGoofy', 49.99);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE users;

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
