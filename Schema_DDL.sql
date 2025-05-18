-- Customers
CREATE TABLE customers (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  customer_id varchar(50) NOT NULL,
  name text DEFAULT NULL,
  email text DEFAULT NULL,
  address text DEFAULT NULL,
  created_date datetime DEFAULT NULL,
  updated_date timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY customer_id (customer_id)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- Products
CREATE TABLE products (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  product_id varchar(50) NOT NULL,
  name text DEFAULT NULL,
  category text DEFAULT NULL,
  created_date datetime DEFAULT NULL,
  updated_date datetime DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY product_id (product_id)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- Orders
CREATE TABLE orders (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  order_id varchar(100) NOT NULL,
  customer_id bigint(20) unsigned DEFAULT NULL,
  region text DEFAULT NULL,
  date_of_sale date DEFAULT NULL,
  payment_method text DEFAULT NULL,
  shipping_cost decimal(10,2) DEFAULT NULL,
  created_date datetime DEFAULT NULL,
  updated_date datetime DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY order_id (order_id),
  KEY customer_id (customer_id),
  CONSTRAINT orders_ibfk_1 FOREIGN KEY (customer_id) REFERENCES customers (id)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- Order Items
CREATE TABLE order_items (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  order_id bigint(20) unsigned DEFAULT NULL,
  product_id bigint(20) unsigned DEFAULT NULL,
  quantity_sold int(11) DEFAULT NULL,
  unit_price decimal(10,2) DEFAULT NULL,
  discount decimal(5,2) DEFAULT NULL,
  created_date datetime DEFAULT NULL,
  updated_date datetime DEFAULT NULL,
  PRIMARY KEY (id),
  KEY order_id (order_id),
  KEY product_id (product_id),
  CONSTRAINT order_items_ibfk_1 FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
  CONSTRAINT order_items_ibfk_2 FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;