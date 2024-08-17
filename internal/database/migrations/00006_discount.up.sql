CREATE TABLE discounts (
    id varchar(50) not null,
    discount_name varchar(100) not null,
    discount_description text not null,
    discount_type varchar(100) not null default 'fixed_amount',
    discount_value int not null,
    discount_code varchar(100) not null,
    discount_start_date DATETIME not null,
    discount_end_date DATETIME not null,
    discount_max_uses int not null,
    discount_uses_count int not null,
    discount_max_uses_per_user int not null,
    discount_min_order_value int not null,
    discount_shop_id varchar(100) not null,
    discount_is_active BOOLEAN not null default 1,
    discount_applies_to ENUM('all', 'specific') not null,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (discount_shop_id) REFERENCES shops(id),
    PRIMARY KEY (id)    
);

CREATE TABLE discount_users_used (
    discount_id varchar(50) not null,
    user_id varchar(100) not null,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(discount_id, user_id),
    FOREIGN KEY (discount_id) REFERENCES discounts(id)
);

CREATE TABLE discount_product_ids (
    discount_id varchar(50) not null,
    product_id varchar(100) not null,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(discount_id, product_id),
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (discount_id) REFERENCES discounts(id)
);