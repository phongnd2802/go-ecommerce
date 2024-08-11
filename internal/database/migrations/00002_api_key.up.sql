CREATE TABLE IF NOT EXISTS api_keys (
    id int PRIMARY KEY AUTO_INCREMENT,
    akey text NOT NULL,
    status BOOLEAN DEFAULT true,
    permissions varchar(20) NOT NULL UNIQUE
);
