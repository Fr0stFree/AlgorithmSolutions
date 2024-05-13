-- Внесите в ваш SQL-запрос изменения, чтобы он считал процент лайков каждого поста относительно всех лайков всех постов
-- из таблицы bloggers_posts. При этом в итоговой таблице также должно остаться три лучших поста, просто у них будет
-- меньший процент, так как учитываются все лайки. Количество всех лайков (total_likes) также будет другим.

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
    total_likes,
    ROUND(likes * 100 / total_likes, 2) as percent
FROM
    (SELECT
        ROW_NUMBER() OVER(PARTITION BY blogger ORDER BY likes DESC) as post_popularity,
        SUM(likes) OVER() as total_likes,
        blogger,
        post,
        likes
    FROM
        bloggers_posts) AS temp
WHERE post_popularity = 1
ORDER BY percent DESC;

-- fin
DROP TABLE bloggers_posts;