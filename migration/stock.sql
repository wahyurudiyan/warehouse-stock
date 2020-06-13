CREATE TABLE warehouse.stock (
	id INT UNSIGNED auto_increment NOT NULL,
	name varchar(100) NOT NULL,
	category varchar(100) NOT NULL,
	producer varchar(100) NOT NULL,
	price varchar(100) NOT NULL,
	quantity varchar(100) NOT NULL,
	production_date varchar(100) NOT NULL,
	expired_date varchar(100) NOT NULL,
	create_at TIMESTAMP NOT NULL,
	update_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP NOT NULL,
	CONSTRAINT stock_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COLLATE=utf8_general_ci;
