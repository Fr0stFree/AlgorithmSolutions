-- В таблице orders хранится список заказов компании. Напишите запрос, который в первом столбце (year) — выведет год,
-- во втором (user_id) — идентификатор пользователя, в третьем (amount) — сумму выполненных (success) заказов за текущий
-- год для этого пользователя, а в четвертом (percent) — вклад пользователя в процентах в общую сумму доходов в рамках
-- текущего года. Данные отсортируйте по году и по вкладу пользователя в возрастающем порядке.

-- init
DROP TABLE IF EXISTS orders;
CREATE TABLE orders (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_id INTEGER NULL,
    date DATETIME NULL,
    amount INTEGER NULL,
    status VARCHAR(50) NULL
);
INSERT INTO orders (id, user_id, date, amount, status)
VALUES
    (1, 7, '2017-01-04 18:23:09', 500, 'success'),
    (2, 1, '2017-01-04 18:25:27', 8700, 'cancelled'),
    (3, 4, '2017-01-12 09:23:14', 1350, 'success'),
    (4, 10, '2017-01-14 17:16:39', 600, 'new'),
    (5, 3, '2017-01-23 17:04:04', 4500, 'success'),
    (6, 12, '2017-02-01 13:04:47', 980, 'success'),
    (7, 1, '2017-02-01 13:32:17', 680, 'success'),
    (8, 10, '2017-02-12 08:30:23', 8000, 'success'),
    (9, 5, '2017-02-12 12:12:43', 700, 'success'),
    (10, 2, '2017-02-14 23:21:25', 1600, 'success'),
    (11, 3, '2017-02-16 14:44:05', 1400, 'success'),
    (12, 5, '2017-02-28 02:00:47', 4300, 'cancelled'),
    (13, 10, '2017-03-02 08:53:25', 1240, 'new'),
    (14, 6, '2016-11-02 08:43:22', 4800, 'new'),
    (15, 5, '2016-11-06 18:35:21', 3400, 'success'),
    (16, 2, '2016-11-19 04:11:17', 5800, 'cancelled'),
    (17, 2, '2016-11-25 11:17:48', 1000, 'success'),
    (18, 3, '2016-12-05 13:01:53', 1280, 'new'),
    (19, 12, '2016-12-11 21:05:16', 2400, 'cancelled'),
    (20, 10, '2016-12-22 17:28:00', 700, 'success');

-- solution
SELECT
    YEAR(date) AS year,
    user_id,
    SUM(amount) as amount,
    SUM(amount) * 100 / SUM(SUM(amount)) OVER(PARTITION BY YEAR(date)) AS percent
FROM orders
WHERE status = 'success'
GROUP BY user_id, year
ORDER BY year, percent