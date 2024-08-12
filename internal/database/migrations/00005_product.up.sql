-- CREATE TABLE clothes (
--     id varchar(100) PRIMARY KEY,
--     brand VARCHAR(255) NOT NULL,
--     size VARCHAR(10) NOT NULL,
--     material VARCHAR(255) NOT NULL,
--     product_shop varchar(50) NOT NULL,
--     created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
--     updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--     FOREIGN KEY (id) REFERENCES products(id),
--     FOREIGN KEY (product_shop) REFERENCES shops(id)
-- );


-- CREATE TABLE electronics(
--     id varchar(100) PRIMARY KEY,
--     manufacturer VARCHAR(255) NOT NULL,
--     model varchar(255) NOT NULL,
--     color varchar(100) NOT NULL,
--     product_shop varchar(50) NOT NULL,
--     created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
--     updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--     FOREIGN KEY (id) REFERENCES products(id),
--     FOREIGN KEY (product_shop) REFERENCES shops(id)
-- );


-- CREATE TABLE furnitures (
--     id varchar(100) PRIMARY KEY,
--     brand VARCHAR(255) NOT NULL,
--     size VARCHAR(255) NOT NULL,
--     material VARCHAR(255) NOT NULL,
--     product_shop varchar(50) NOT NULL,
--     created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
--     updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--     FOREIGN KEY (id) REFERENCES products(id),
--     FOREIGN KEY (product_shop) REFERENCES shops(id)
-- );