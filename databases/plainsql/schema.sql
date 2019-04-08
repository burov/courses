CREATE DATABASE shop;

CREATE TABLE users (
  id         serial,
  first_name varchar(255) NOT NULL,
  last_name  varchar(255) NOT NULL,
  phone      varchar(255),
  email      varchar(1024),
  Primary key(id)
);

CREATE TABLE products (
    id   serial,
    name varchar(255) NOT NULL,
    price int NOT NULL,
    count int DEFAULT 0,
    Primary key(id)
);

CREATE TABLE orders (
    id serial,
    user_id int NOT NULL,
    Primary key (id)
);


CREATE TABLE orders_products (
    order_id   serial,
    product_id serial,

    FOREIGN KEY (order_id) REFERENCES orders,
    FOREIGN KEY (product_id) REFERENCES products
);

CREATE TABLE comments (
    id      serial,
    user_id serial NOT NULL,
    product_id serial NOT NULL,
    comment text,

    Primary key (id),
    Foreign key(user_id) references users,
    Foreign key(product_id) references products
);
