CREATE TABLE orders
(
    id varchar(36) NOT NULL,
    price float NOT NULL,
    tax float NOT NULL,
    final_price float NOT NULL,
    PRIMARY KEY (id)
);