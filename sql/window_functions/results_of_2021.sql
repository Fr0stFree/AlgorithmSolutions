-- Напишите SQL запрос для анализа 2021 года. Итоговая таблица должна содержать четыре колонки: month — номер месяца,
-- income — доход за этот месяц, diff — разницу с апрелем в абсолютных показателях (текущий месяц - апрель),
-- diff_percent — разницу с апрелем в процентах с округлением до одного знака после десятичной точки, например,
-- для мая это будет -13.6 (минус 13.6)

-- init
DROP TABLE IF EXISTS revenues;
CREATE TABLE revenues (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    year INTEGER NULL,
    month INTEGER NULL,
    income INTEGER NULL,
    outcome INTEGER NULL
);
INSERT INTO revenues (id, year, month, income, outcome)
VALUES
    (1, 2020, 1, 400000, 240000),
    (2, 2020, 2, 380000, 228000),
    (3, 2020, 3, 410000, 246000),
    (4, 2020, 4, 490000, 294000),
    (5, 2020, 5, 460000, 276000),
    (6, 2020, 6, 430000, 258000),
    (7, 2020, 7, 410000, 246000),
    (8, 2020, 8, 470000, 282000),
    (9, 2020, 9, 490000, 294000),
    (10, 2020, 10, 520000, 312000),
    (11, 2020, 11, 580000, 348000),
    (12, 2020, 12, 620000, 372000),
    (13, 2021, 1, 520000, 312000),
    (14, 2021, 2, 470000, 282000),
    (15, 2021, 4, 590000, 354000),
    (16, 2021, 3, 520000, 312000),
    (17, 2021, 5, 510000, 306000),
    (18, 2021, 6, 490000, 294000),
    (19, 2021, 7, 480000, 288000),
    (20, 2021, 9, 610000, 366000),
    (21, 2021, 8, 530000, 318000),
    (22, 2021, 10, 680000, 408000),
    (23, 2021, 12, 700000, 420000),
    (24, 2021, 11, 650000, 390000);

-- solution
(SELECT
    year,
    month,
    income,
    income - (SELECT income FROM revenues WHERE month = 4 AND year = 2020) AS diff,
    ROUND((income - (SELECT income FROM revenues WHERE month = 4 AND year = 2020)) / (SELECT income FROM revenues WHERE month = 4 AND year = 2020) * 100, 1) AS diff_percent
FROM revenues
WHERE year = 2020
) UNION (
SELECT
    year,
    month,
    income,
    income - (SELECT income FROM revenues WHERE month = 4 AND year = 2021) AS diff,
    ROUND((income - (SELECT income FROM revenues WHERE month = 4 AND year = 2021)) / (SELECT income FROM revenues WHERE month = 4 AND year = 2021) * 100, 1) AS diff_percent
FROM revenues
WHERE year = 2021)
ORDER BY year, month

