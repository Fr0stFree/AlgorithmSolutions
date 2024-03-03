-- Complete the solution so that it reverses the string passed into it.

-- 'world'  =>  'dlrow'

CREATE TABLE solution(str VARCHAR(100));
INSERT INTO solution(str) VALUES ('world');


SELECT
    str,
    REVERSE(str) AS res
FROM
    solution;
