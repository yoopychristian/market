create database enigmart;
CREATE TABLE product (
  id INT NOT NULL AUTO_INCREMENT,
  product_code VARCHAR(150) NULL DEFAULT NULL,
  product_name VARCHAR(150) NULL DEFAULT NULL,
  price INT NULL DEFAULT '0',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id));
  
  CREATE TABLE purchase_order (
  id INT NOT NULL AUTO_INCREMENT,
  order_date DATE NULL DEFAULT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id));
  
  CREATE TABLE purchase_order_item (
  qty INT NULL DEFAULT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  order_id INT NOT NULL,
  product_id INT NOT NULL,
  PRIMARY KEY (order_id, product_id));
  
insert into product values 
(1, 'LXV', 'Liveboy', 5000, '2020-06-15', '2020-06-15'),
(2, 'DVA', 'Dave', 10000, '2020-06-16', '2020-06-16'),
(3, 'DCA', 'Clean', 25000, '2020-07-17', '2020-07-17');

insert into purchase_order values 
(1, '2020-06-15', '2020-06-15', '2020-06-15'),
(2, '2020-06-16', '2020-06-16', '2020-06-16'),
(3, '2020-06-16', '2020-06-16', '2020-06-16'),
(4, '2020-06-18', '2020-06-18', '2020-06-18');

insert into purchase_order_item values 
(3, '2020-06-15', '2020-06-15', 1, 1),
(2, '2020-06-16', '2020-06-16', 2, 1),
(1, '2020-06-16', '2020-06-16', 3, 2),
(3, '2020-06-18', '2020-06-18', 4, 2);

select * from purchase_order_item;
select * from product;
select * from purchase_order;
select * from login;

CREATE TABLE login (
  username VARCHAR(150) NULL DEFAULT NULL,
  password VARCHAR(150) NULL DEFAULT NULL);



  
