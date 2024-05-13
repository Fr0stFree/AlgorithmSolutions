-- Выведите в итоговой таблице 5 колонок: blogger — имя блогера, post — название поста, likes — количество лайков поста,
-- total_likes — общее количество лайков лучших постов, percent — процент лайков текущего поста относительно общего
-- количества лайков. Процент следует вывести с двумя знаками после десятичной точки.
-- Итоговые данные отсортируйте по популярности постов: чем популярней — тем выше.

-- init
DROP TABLE IF EXISTS bloggers_posts;
CREATE TABLE bloggers_posts (
    id INT NOT NULL PRIMARY KEY,
    blogger VARCHAR(50) NULL,
    post VARCHAR(50) NULL,
    likes INTEGER NULL,
    comments INTEGER NULL
);
INSERT INTO bloggers_posts (id, blogger, post, likes, comments)
VALUES
    (1, '@superblogger', 'post_1', 981, 90),
    (2, '@thebestparty', 'post_7', 134, 14),
    (3, '@greatcoder', 'post_19', 245, 24),
    (4, '@superblogger', 'post_6', 111, 9),
    (5, '@thebestparty', 'post_5', 741, 72),
    (6, '@greatcoder', 'post_11', 145, 14),
    (7, '@superblogger', 'post_24', 221, 19),
    (8, '@thebestparty', 'post_45', 451, 138),
    (9, '@greatcoder', 'post_178', 147, 4),
    (10, '@superblogger', 'post_44', 78, 2),
    (11, '@thebestparty', 'post_13', 633, 78),
    (12, '@greatcoder', 'post_23', 713, 88);

-- solution
SELECT
    blogger,
    post,
    likes,
    SUM(likes) OVER() as total_likes,
    TRUNC(likes * 100 / SUM(likes) OVER(), 2) AS percent
FROM
    (SELECT
        ROW_NUMBER() OVER(PARTITION BY blogger ORDER BY likes DESC) as post_popularity,
        blogger,
        post,
        likes
    FROM
        bloggers_posts) AS temp
WHERE post_popularity = 1
ORDER BY percent DESC;

-- fin
DROP TABLE bloggers_posts;