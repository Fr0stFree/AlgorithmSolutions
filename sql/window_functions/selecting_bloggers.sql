-- Напишите SQL-запрос для подбора блогеров. В первую очередь должны идти блогеры с максимальным количеством
-- подписчиков, далее, если количество подписчиков равно, то нужно брать блогеров с максимальным количеством лайков,
-- а затем с максимальным количеством комментариев.
-- В итоговой таблице выведите следующие поля: num — номер по порядку начиная с 1, аккаунт блогера, количество
-- подписчиков, число лайков, число комментариев. Первыми должны идти блогеры с максимальным числом подписчиков.

-- init
DROP TABLE IF EXISTS bloggers;
CREATE TABLE bloggers (
    id INT NOT NULL PRIMARY KEY,
    blogger VARCHAR(50) NULL,
    subs INTEGER NULL,
    avg_likes INTEGER NULL,
    avg_comments INTEGER NULL
);
INSERT INTO bloggers (id, blogger, subs, avg_likes, avg_comments)
VALUES
    (1, '@superblogger', 7800, 122, 24),
    (2, '@thebestparty', 6700, 201, 25),
    (3, '@dangerousmom', 4500, 178, 15),
    (4, '@greatcoder', 8100, 221, 25),
    (5, '@incomparablecat', 7800, 288, 44),
    (6, '@likescheater', 2400, 872, 2),
    (7, '@subscheater', 9800, 25, 1),
    (8, '@realpepper', 9800, 330, 60),
    (9, '@commentscheater', 7800, 122, 233),
    (10, '@highlander', 8100, 233, 32);

-- solution
SELECT
    ROW_NUMBER() OVER(ORDER BY subs DESC, avg_likes DESC, avg_comments DESC) as num,
    bloggers.blogger,
    bloggers.subs,
    bloggers.avg_likes,
    bloggers.avg_comments
FROM
    bloggers
ORDER BY num;

-- finalize
DROP TABLE bloggers;