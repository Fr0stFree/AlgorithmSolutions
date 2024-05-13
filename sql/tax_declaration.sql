-- Напишите SQL запрос, который выведет 4 столбца: quarter — номер квартала, income — суммарные поступления за этот
-- квартал, income_acc — поступления за квартал нарастающим итогом и usn6 — величина налога нарастающим итогом. При
-- системе налогообложения УСН 6% начисляется 6% налог на все доходы. Выведете налоги с двумя знаками после десятичной точки.

-- init
DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    date DATE NULL,
    income INTEGER NULL
);
INSERT INTO transactions (id, date, income)
VALUES
    (1, '2021-04-01', 22000),
    (2, '2021-04-02', 11100),
    (3, '2021-04-11', 64000),
    (4, '2021-05-04', 23000),
    (5, '2021-06-17', 20000),
    (6, '2021-06-18', 7900),
    (7, '2021-06-19', 32000),
    (8, '2021-07-12', 17000),
    (9, '2021-07-23', 14600),
    (10, '2021-01-12', 26300),
    (11, '2021-08-11', 10000),
    (12, '2021-09-02', 10000),
    (13, '2021-10-15', 18700),
    (14, '2021-11-01', 6500),
    (15, '2021-11-01', 17900),
    (16, '2021-12-04', 13400),
    (17, '2021-12-14', 44000),
    (18, '2021-12-15', 7000),
    (19, '2021-01-02', 10000),
    (20, '2021-01-21', 15000),
    (21, '2021-02-17', 25000),
    (22, '2021-03-23', 12000);

-- solution
SELECT
    QUARTER(date) AS quarter,
    SUM(income) as income,
    SUM(SUM(income)) OVER(ORDER BY quarter(date)) AS income_acc,
    ROUND(SUM(SUM(income * 0.06)) OVER(ORDER BY quarter(date)), 2) AS usn6
FROM transactions
GROUP BY quarter
ORDER BY quarter
