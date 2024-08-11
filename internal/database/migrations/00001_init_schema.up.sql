BEGIN;

CREATE TABLE shops (
    id VARCHAR(50) PRIMARY KEY,
    shop_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT false,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE roles (
    id int PRIMARY KEY AUTO_INCREMENT,
    role_name varchar(100) UNIQUE NOT NULL,
    role_description text,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE shop_roles (
    shop_id VARCHAR(50),
    role_id int,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(shop_id, role_id),
    FOREIGN KEY (shop_id) REFERENCES shops(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

CREATE TABLE tokens (
    id varchar(50) PRIMARY KEY,
    public_key text NOT NULL,
    refresh_token text NOT NULL,
    refresh_token_used text,
    shop_id varchar(50) NOT NULL UNIQUE,
    FOREIGN KEY (shop_id) REFERENCES shops(id) ON DELETE CASCADE
);

COMMIT;