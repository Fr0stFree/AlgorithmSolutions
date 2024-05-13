-- init
DROP TABLE IF EXISTS flats;
CREATE TABLE flats (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    street VARCHAR(50) NULL,
    house VARCHAR(10) NULL,
    price INTEGER NULL,
    rooms INTEGER NULL
);
INSERT INTO flats (id, street, house, price, rooms)
VALUES
    (1, 'Ефремова', '2', 35000, 3),
    (2, 'Горького', '17', 25000, 2),
    (3, 'Горького', '207', 22000, 2),
    (4, 'Чайковского', '36', 30000, 3),
    (5, 'Гагарина', '7', 32000, 3),
    (6, 'Колоскова', '8', 18000, 1),
    (7, 'Мукомольная', '2', 22000, 2),
    (8, 'Озерная', '35', 20000, 1),
    (9, 'Озерова', '7', 24000, 2),
    (10, 'Земельная', '14', 25000, 2),
    (11, 'Горького', '202', 24000, 3),
    (12, 'Советский', '71', 18000, 1),
    (13, 'Банковская', '6', 35000, 3),
    (14, 'Красносельская', '31', 40000, 3),
    (15, 'Суворова', '40', 25000, 2);

-- solution
SELECT
    street,
    house,
    price,
    rooms
FROM
    (SELECT
        *,
        DENSE_RANK() OVER(ORDER BY price) AS rating
    FROM flats
    WHERE rooms > 1) AS temp
WHERE rating <= 3
ORDER BY price, rooms DESC