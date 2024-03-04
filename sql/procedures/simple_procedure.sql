# Напишите хранимую процедуру order_payment(), которая принимает три параметра: order_id (номер заказа), amount
# (сумма платежа) и payment_date (дата и время платежа). Процедура должна устанавливать статус заказа для
# переданного order_id в "paid", а также добавлять запись об оплате в таблицу transactions.

-- init
SET foreign_key_checks = 0;
DROP PROCEDURE IF EXISTS order_payment;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS transactions;
SET foreign_key_checks = 1;
CREATE TABLE orders (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_id INT UNSIGNED NOT NULL,
    status VARCHAR(30) DEFAULT 'new',
    amount INTEGER NOT NULL
);
CREATE TABLE transactions (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    order_id INT UNSIGNED NOT NULL,
    amount INTEGER NOT NULL,
    date DATETIME
);
INSERT INTO orders (user_id, status, amount)
VALUES
    (14, 'new', 5600),
    (56, 'paid', 7800),
    (56, 'success', 4000),
    (5, 'delivery', 4500),
    (45, 'new', 700),
    (87, 'new', 1740);
INSERT INTO transactions (order_id, date, amount)
VALUES
    (2, '2019-12-12 07:04:42', 7800),
    (3, '2019-12-14 16:24:24', 4000),
    (4, '2019-12-15 18:03:11', 4500);

DELIMITER $$
CREATE PROCEDURE order_payment(payment_order_id INT, payment_amount INT, payment_date DATETIME)
BEGIN
    UPDATE orders SET status = 'paid' WHERE id = payment_order_id;
    INSERT INTO transactions(order_id, amount, date) VALUES (payment_order_id, payment_amount, payment_date);
END $$
DELIMITER ;

CALL order_payment(5, 1000, '2023-01-01');
