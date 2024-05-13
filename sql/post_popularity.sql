-- Напишите SQL-запрос, который пронумерует посты каждого блогера в порядке их популярности на основе лайков.
-- Самые «залайканные» посты должны получить единицу и чем меньше лайков, тем больший номер у них должен быть.
-- Для каждого блогера нужно вести свою отдельную нумерацию. Популярность постов выводите в колонке post_popularity,
-- которая должна стоять первой. После неё должны идти имя блогера, название поста и количество лайков.
-- Итоговые данные отсортируйте по имени блогера, а затем по популярности записей (самая популярная вверху).

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
    ROW_NUMBER() OVER(PARTITION BY blogger ORDER BY likes DESC) as post_popularity,
    bloggers_posts.blogger,
    bloggers_posts.post,
    bloggers_posts.likes
FROM
    bloggers_posts
ORDER BY
    bloggers_posts.blogger,
    post_popularity;

-- fin
DROP TABLE bloggers_posts;