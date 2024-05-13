-- Напишите SQL-запрос, который выведет пять столбцов: place — место спортсмена, last_name — фамилию, first_name — имя,
-- time — время преодоления дистанции в формате ЧЧ:ММ:СС и champion_lag — время отставания от первого места в формате
-- ЧЧ:ММ:СС

-- init
DROP TABLE IF EXISTS runners;
CREATE TABLE runners (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    first_name VARCHAR(50) NULL,
    last_name VARCHAR(50) NULL,
    start_time TIME NULL,
    end_time TIME NULL
);
INSERT INTO runners (id, first_name, last_name, start_time, end_time)
VALUES
    (1, 'Виктор', 'Алтушев', '10:00:17', '12:35:28'),
    (2, 'Светлана', 'Иванова', '10:00:18', '12:45:19'),
    (3, 'Елена', 'Абрамова', '10:03:22', '12:30:13'),
    (4, 'Василиса', 'Кац', '10:05:58', '12:38:41'),
    (5, 'Антон', 'Сорокин', '10:06:04', '12:50:18'),
    (6, 'Алёна', 'Алясева', '10:00:54', '12:55:17'),
    (7, 'Лиана', 'Белая', '10:01:01', '12:39:14'),
    (8, 'Карина', 'Белая', '10:04:43', '12:45:55'),
    (9, 'Анастасия', 'Дейчман', '10:05:14', '12:37:33'),
    (10, 'Юлия', 'Фёдорова', '10:00:06', '12:40:30');

-- solution
SELECT
    *,
    DATE_FORMAT(TIMEDIFF(time, FIRST_VALUE(time) OVER()), "%H:%i:%s") AS champion_lag
FROM (SELECT RANK() OVER (ORDER BY TIMEDIFF(runners.end_time, runners.start_time)) AS place,
             last_name,
             first_name,
             DATE_FORMAT(TIMEDIFF(end_time, start_time), "%H:%i:%s")               AS time
      FROM runners
      ORDER BY time) as temp
ORDER BY place