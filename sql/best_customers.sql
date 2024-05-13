-- Напишите SQL запрос, который выведет 5 колонок: name — название магазина, first_name — имя покупателя, last_name —
-- фамилия покупателя, amount — общая сумма выполненных заказов (status="success") покупателя в текущем магазине,
-- c_level — группа (уровень) покупателя. Всех покупателей магазин делит на четыре равные группы, нумеруя их от 1
-- (потратили больше всего) до 4 (потратили меньше всего). При этом для каждого магазина группы считаются по
-- отдельности. Итоговые данные отсортируйте по названию магазина, а после по группам в возрастающем порядке.

-- init
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    first_name VARCHAR(50) NULL,
    last_name VARCHAR(50) NULL,
    email VARCHAR(50) NULL
);
INSERT INTO users (id, first_name, last_name, email)
VALUES
    (1, 'Виктор', 'Алтушев', 'user1@domain.com'),
    (2, 'Светлана', 'Иванова', 'user2@domain.com'),
    (4, 'Василиса', 'Кац', 'user4@domain.com'),
    (5, 'Антон', 'Сорокин', 'user5@domain.com'),
    (7, 'Лиана', 'Белая', 'user7@domain.com'),
    (9, 'Анастасия', 'Дейчман', 'user9@domain.com'),
    (10, 'Валерий', 'Артемьев', 'user10@domain.com'),
    (11, 'Юлия', 'Фёдорова', 'user11@domain.com'),
    (12, 'Семен', 'Лутц', 'user12@domain.com'),
    (15, 'Диана', 'Слепакова', 'user15@domain.com');
DROP TABLE IF EXISTS shops;
CREATE TABLE shops (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NULL
);
INSERT INTO shops (id, name)
VALUES
    (1, 'Пролетарская'),
    (2, 'Киевская'),
    (3, 'ТЦ Азимут');
DROP TABLE IF EXISTS orders;
CREATE TABLE orders (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_id INTEGER NULL,
    date DATETIME NULL,
    amount INTEGER NULL,
    status VARCHAR(50) NULL,
    shop_id INTEGER NULL
);
INSERT INTO orders (id, user_id, date, amount, status, shop_id)
VALUES
    (1, 7, '2017-01-04 18:23:09', 500, 'success', 1),
    (2, 1, '2017-01-04 18:25:27', 8700, 'cancelled', 1),
    (3, 4, '2017-01-12 09:23:14', 1350, 'success', 2),
    (4, 10, '2017-01-14 17:16:39', 600, 'new', 1),
    (5, 4, '2017-01-23 17:04:04', 4500, 'success', 3),
    (6, 12, '2017-02-01 13:04:47', 980, 'success', 2),
    (7, 1, '2017-02-01 13:32:17', 680, 'success', 1),
    (8, 10, '2017-02-12 08:30:23', 8000, 'success', 1),
    (9, 5, '2017-02-12 12:12:43', 700, 'success', 2),
    (10, 2, '2017-02-14 23:21:25', 1600, 'success', 2),
    (11, 7, '2017-02-16 14:44:05', 1400, 'success', 3),
    (12, 5, '2017-02-28 02:00:47', 4300, 'cancelled', 3),
    (13, 10, '2017-03-02 08:53:25', 1240, 'new', 2),
    (14, 11, '2017-03-04 11:25:27', 8700, 'success', 1),
    (15, 2, '2017-03-11 17:11:32', 1200, 'new', 2),
    (16, 1, '2017-03-17 08:17:33', 6500, 'success', 3),
    (17, 15, '2017-04-01 20:28:01', 3400, 'success', 2),
    (18, 7, '2017-04-14 14:53:07', 1400, 'success', 1),
    (19, 2, '2017-04-22 16:44:47', 3200, 'cancelled', 1),
    (20, 9, '2017-05-05 12:12:17', 870, 'success', 3);

-- solution
SELECT
    month,
    first_name,
    last_name,
    amount
FROM
    (SELECT
        *,
        NTILE(4) OVER(PARTITION BY month ORDER BY amount DESC) AS c_level
    FROM
        (SELECT
            users.first_name,
            users.last_name,
            MONTH(o.date) as month,
            SUM(amount) AS amount
        FROM users
        INNER JOIN orders o on users.id = o.user_id
        INNER JOIN shops s on o.shop_id = s.id
        WHERE status = 'success'
        GROUP BY MONTH(o.date), users.first_name, users.last_name
        ) AS temp
    ORDER BY amount) as temp_2
WHERE c_level = 1
ORDER BY month, amount