# Напишите хранимую процедуру active_products(), которая выполняет SELECT-запрос на извлечение данных из таблицы
# products. Нужно получить активные товары, которые есть на складе. Вывести нужно id, название, количество и цену.
# Данные должны быть отсортированы по цене.

-- init
SET foreign_key_checks = 0;
DROP PROCEDURE IF EXISTS active_products;
DROP TABLE IF EXISTS products;
SET foreign_key_checks = 1;
CREATE TABLE products (
    id INT UNSIGNED NOT NULL,
    name VARCHAR(255) NULL,
    count INTEGER NULL,
    active BOOL,
    price INTEGER NULL
);
INSERT INTO products (id, name, count, price, active)
VALUES
    (1, 'Стиральная машина', 5, 10000, true),
    (2, 'Холодильник', 0, 10000, false),
    (3, 'Микроволновка', 3, 4000, false),
    (4, 'Пылесос', 2, 4500, true),
    (5, 'Вентилятор', 0, 700, true),
    (6, 'Телевизор', 7, 31740, true),
    (7, 'Тостер', 2, 2500, false),
    (8, 'Принтер', 4, 3000, true);

-- solution
DELIMITER $$
CREATE PROCEDURE IF NOT EXISTS active_products()
BEGIN
    SELECT id, name, count, price FROM products WHERE active = TRUE AND count > 0 ORDER BY price;
END $$
DELIMITER ;

CALL active_products();
