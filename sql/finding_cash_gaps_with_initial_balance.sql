# Напишите SQL-запрос, который выведет все исходные столбцы таблицы transactions + колонку balance, в которой будут
# отображаться остатки на счете после проведения операции. Нужно учесть, что на 1 января 2022 года на счетах компании
# было 10000 рублей, эти деньги также участвуют в расчете баланса. Начальный баланс нужно также отразить в итоговой
# таблице с id равным 0, item равным "Начальный баланс" и суммой в 10000. Итоговую таблицу отсортируйте по дате и id.

-- init
DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    date DATE NULL,
    item VARCHAR(50) NULL,
    money INTEGER NULL
);
INSERT INTO transactions (id, date, item, money)
VALUES
    (1, '2022-01-03', 'Продажа кофе', 25000),
    (2, '2022-01-05', 'Продажа молока', 17000),
    (3, '2022-01-05', 'Аванс менеджера 1', -15000),
    (4, '2022-01-05', 'Аванс менеджера 2', -15000),
    (5, '2022-01-07', 'Продажа кофе', 35000),
    (6, '2022-01-08', 'Продажа молока', 7000),
    (7, '2022-01-12', 'Коммунальные платежи', -4000),
    (8, '2022-01-11', 'Аренда офиса', -20000),
    (9, '2022-01-13', 'Продажа кофе', 22000),
    (10, '2022-01-15', 'Телефония', -2000),
    (11, '2022-01-16', 'Участие в выставке', -35000),
    (12, '2022-01-21', 'Продажа молока', 15000),
    (13, '2022-01-19', 'Продажа кофе', 5500),
    (14, '2022-01-22', 'Продажа кофе', 11000),
    (15, '2022-01-22', 'Продажа кофе', 5000),
    (16, '2022-01-24', 'Закуп кофе', 40000),
    (17, '2022-01-19', 'Зарплата менеджера 1', -25000),
    (18, '2022-01-19', 'Зарплата менеджера 2', -25000),
    (19, '2022-01-24', 'Закуп сухого молока', 10000),
    (20, '2022-01-27', 'Продажа кофе', 8000),
    (21, '2022-01-30', 'Продажа молока', 5000),
    (22, '2022-01-30', 'Продажа молока', 15000);

-- solution
SELECT
    *,
    SUM(money) OVER (ORDER BY date, id) AS balance
FROM
    (SELECT 0 as id, DATE('2022-01-01') AS date, 'Начальный баланс' AS item, 10000 AS money
    UNION
    SELECT
        *
    FROM transactions
    ORDER BY date, id) as temp
