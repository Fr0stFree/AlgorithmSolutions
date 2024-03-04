# Напишите SQL запрос для формирования списка рассылки. Всего планируется три тестовых письма, которые следует
# пронумеровать от 1 до 3-х с равномерным распределением чисел между пользователями таблицы users.
# Распределять письма будем в порядке следования идентификаторов, то есть первый блок пользователей получит единицу,
# следующий двойку и последний тройку. Номер варианта следует вывести в первой колонке с названием mail_variant,
# далее выведите id пользователя, его email и имя. Итоговые данные отсортируйте по id.

-- init
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INT UNSIGNED NOT NULL PRIMARY KEY,
    first_name VARCHAR(50) NULL,
    last_name VARCHAR(50) NULL,
    email VARCHAR(50) NULL
);
INSERT INTO users (id, first_name, last_name, email)
VALUES
    (1, 'Виктор', 'Алтушев', 'user1@domain.com'),
    (2, 'Светлана', 'Иванова', 'user2@domain.com'),
    (4, 'Василиса', 'Кац', 'user4@domain.com'),
    (5, 'Антон', 'Сорокин', 'user5@domain.com'),
    (7, 'Лиана', 'Белая', 'user7@domain.com'),
    (9, 'Анастасия', 'Дейчман', 'user9@domain.com'),
    (10, 'Валерий', 'Артемьев', 'user10@domain.com'),
    (11, 'Юлия', 'Фёдорова', 'user11@domain.com'),
    (12, 'Семен', 'Лутц', 'user12@domain.com'),
    (15, 'Диана', 'Слепакова', 'user15@domain.com');

-- solution
SELECT
    NTILE(4) OVER(ORDER BY MD5(email)) AS mail_variant,
    id,
    email,
    first_name
FROM users
ORDER BY id