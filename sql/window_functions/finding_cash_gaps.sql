# Все операции проводились в соответствии с датой в колонке date. Если в одну дату было несколько операций,
# то они проводились в порядке следования id. Итоговую таблицу отсортируйте по дате и id.
# Кассовый разрыв возникает в тот момент, когда у компании недостаточно средств на счетах для совершения операций.
# В нашей таблице кассовый разрыв будет для тех операций, после которых баланс станет отрицательным.

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
    SUM(money) OVER(ORDER BY date, id) as balance
FROM transactions
ORDER BY date, id