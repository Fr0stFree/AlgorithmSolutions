-- Сгруппируйте данные в таблице users по полу и получите количество мужчин и женщин, а также их процентное соотношение
-- к общему числу людей.В первой колонке (sex) выведите пол, во второй (members) — количество , а в третьей (percent)
-- — процент. Данные отсортируйте по последнему столбцу.

-- init
SET foreign_key_checks = 0;
DROP TABLE IF EXISTS users;
SET foreign_key_checks = 1;
CREATE TABLE users (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    first_name VARCHAR(50) NULL,
    last_name VARCHAR(50) NULL,
    age INTEGER NULL,
    sex VARCHAR(1) NULL
);
INSERT INTO users (id, first_name, last_name, age, sex)
VALUES
    (1, 'Дмитрий', 'Федьков', 18, 'm'),
    (2, 'Светлана', 'Иванова', 30, 'w'),
    (3, 'Никита', 'Трутнев', 17, 'm'),
    (4, 'Денис', 'Кац', 22, 'm'),
    (5, 'Алена', 'Сорокина', 16, 'w'),
    (6, 'Василий', 'Федьков', 44, 'm'),
    (7, 'Ярослав', 'Сорокин', 17, 'm'),
    (8, 'Денис', 'Белый', 23, 'm'),
    (9, 'Анатолий', 'Дейчман', 21, 'm'),
    (10, 'Юлия', 'Фёдорова', 29, 'w'),
    (11, 'Андрей', 'Ройзман', 30, 'm'),
    (12, 'София', 'Турыгина', 15, 'w'),
    (13, 'Анастасия', 'Татаренкова', 21, 'w'),
    (14, 'Валерий', 'Редько', 29, 'm');

-- solution
SELECT
    sex,
    COUNT(sex) as members,
    COUNT(sex) * 100 / SUM(COUNT(sex)) OVER() as percent
FROM users
GROUP BY sex
ORDER BY percent