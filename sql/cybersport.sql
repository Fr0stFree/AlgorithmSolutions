-- В киберспортивном турнире участвует 10 команд (таблица cyber_teams) и каждой из них нужно сыграть по три игры
-- (таблица cyber_games) против компьютера. Результаты игр фиксируются в таблице cyber_results и для подсчета набранных
-- баллов используется следующая формула: количество убийств (kills) минус количество смертей (deaths) помноженное на
-- три. Например, если команда в одной игре убила 60 противников и при этом 5 раз умерла, то её итоговый результат будет
-- равен 60 - 5x3 = 45 баллов. Напишите SQL-запрос для формирования итоговой таблицы мест после проведения турнира.
-- В таблице должно быть 3 поля: place — место, team — название команды и points — общее количество набранных очков.
-- В случае, если команды набирают одинаковое количество очков, они разделяют место между собой. Итоговую таблицу
-- отсортируйте по местам, а в случае одинакового места отсортируйте команды в алфавитном порядке.

-- init
DROP TABLE IF EXISTS cyber_teams;
CREATE TABLE cyber_teams (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    team VARCHAR(50) NULL
);
INSERT INTO cyber_teams (id, team)
VALUES
    (1, 'Natus Vincere'),
    (2, 'Astralis'),
    (3, 'Vitality'),
    (4, 'Liquid'),
    (5, 'Evil Geniuses'),
    (6, 'G2 Esports'),
    (7, 'Furia'),
    (8, 'Complexity Gaming'),
    (9, '789'),
    (10, 'Magicians');
DROP TABLE IF EXISTS cyber_games;
CREATE TABLE cyber_games (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    game VARCHAR(50) NULL
);
INSERT INTO cyber_games (id, game)
VALUES
    (1, 'Зимняя ночь'),
    (2, 'Жаркая пустыня'),
    (3, 'Выживание в лесу');
DROP TABLE IF EXISTS cyber_results;
CREATE TABLE cyber_results (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    team_id INTEGER NULL,
    game_id INTEGER NULL,
    kills INTEGER NULL,
    deaths INTEGER NULL
);
INSERT INTO cyber_results (id, team_id, game_id, kills, deaths)
VALUES
    (1, 1, 1, 78, 3),
    (2, 2, 1, 22, 2),
    (3, 3, 1, 57, 4),
    (4, 4, 1, 56, 1),
    (5, 5, 1, 77, 0),
    (6, 6, 1, 80, 0),
    (7, 7, 1, 69, 6),
    (8, 8, 1, 71, 3),
    (9, 9, 1, 45, 2),
    (10, 10, 1, 70, 2),
    (11, 1, 2, 81, 3),
    (12, 2, 2, 89, 4),
    (13, 3, 2, 88, 4),
    (14, 4, 2, 71, 4),
    (15, 5, 2, 79, 1),
    (16, 6, 2, 80, 1),
    (17, 7, 2, 66, 0),
    (18, 8, 2, 0, 0),
    (19, 9, 2, 84, 3),
    (20, 10, 2, 77, 7),
    (21, 1, 3, 65, 4),
    (22, 2, 3, 65, 4),
    (23, 3, 3, 70, 1),
    (24, 4, 3, 66, 0),
    (25, 5, 3, 60, 0),
    (26, 6, 3, 64, 1),
    (27, 7, 3, 68, 1),
    (28, 8, 3, 69, 3),
    (29, 9, 3, 61, 5),
    (30, 10, 3, 71, 3);

-- solution
SELECT
    DENSE_RANK() OVER(ORDER BY points DESC) as place,
    temp.team,
    points
FROM
    (SELECT
        cyber_results.team_id,
        cyber_teams.team,
        MAX(SUM(cyber_results.kills - (cyber_results.deaths * 3))) OVER(PARTITION BY team_id) AS points
    FROM
        cyber_results
    INNER JOIN
        cyber_teams ON cyber_results.team_id = cyber_teams.id
    INNER JOIN
        cyber_games ON cyber_results.game_id = cyber_games.id
    GROUP BY cyber_results.team_id
    ORDER BY points DESC) AS temp