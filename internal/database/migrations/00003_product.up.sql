BEGIN;

CREATE TABLE products (
    id varchar(100) PRIMARY KEY,
    product_name varchar(255) NOT NULL,
    product_thumb varchar(255) NOT NULL,
    product_description text,
    product_price DECIMAL (10, 2) NOT NULL,
    product_quantity INT NOT NULL,
    product_type ENUM('Electronics', 'Clothing', 'Furniture') NOT NULL,
    product_shop varchar(50) NOT NULL,
    product_attributes JSON NOT NULL,
    product_ratingAverage DECIMAL(3, 1) DEFAULT 4.5 CHECK (product_ratingAverage >= 1.0 AND product_ratingAverage <= 5.0),
    product_variations JSON DEFAULT '[]',
    isDraft BOOLEAN DEFAULT TRUE,
    isPublished BOOLEAN DEFAULT FALSE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (product_shop) REFERENCES shops(id)
);

CREATE TABLE clothes (
    id varchar(100) PRIMARY KEY,
    brand VARCHAR(255) NOT NULL,
    size VARCHAR(10),
    material VARCHAR(255),
    product_shop varchar(50) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id) REFERENCES products(id),
    FOREIGN KEY (product_shop) REFERENCES shops(id)
);


CREATE TABLE electronics (
    id varchar(100) PRIMARY KEY,
    manufacturer VARCHAR(255) NOT NULL,
    model varchar(255),
    color varchar(100),
    product_shop varchar(50) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id) REFERENCES products(id),
    FOREIGN KEY (product_shop) REFERENCES shops(id)
);


CREATE TABLE furnitures (
    id varchar(100) PRIMARY KEY,
    brand VARCHAR(255) NOT NULL,
    size VARCHAR(255),
    material VARCHAR(255),
    product_shop varchar(50) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id) REFERENCES products(id),
    FOREIGN KEY (product_shop) REFERENCES shops(id)
);

COMMIT;