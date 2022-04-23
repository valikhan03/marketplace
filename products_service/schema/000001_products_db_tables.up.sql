CREATE TABLE IF NOT EXISTS category(
    id SERIAL UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    parent_id INT DEFAULT NULL,
    FOREIGN KEY (parent_id) REFERENCES category (id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS products(
    id VARCHAR(100) UNIQUE,
    title TEXT,
    description TEXT,
    category_id INT NOT NULL,
    price INT NOT NULL,
    seller_id VARCHAR(100) NOT NULL,
    FOREIGN KEY (category_id) REFERENCES category (id) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE IF NOT EXISTS product_photos(
    filename TEXT NOT NULL,
    number INT,
    product_id VARCHAR(100),
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);

INSERT INTO category (id, name, parent_id) VALUES (0, "category", 0)