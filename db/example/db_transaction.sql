--Tx1: transfer $10 from account 1 to 2
BEGIN;

UPDATE accounts SET balance = balance - 10 WHERE id = 1 RETURNING *;
UPDATE accounts SET balance = balance + 10 WHERE id = 2 RETURNING *;

ROLLBACK;

--TX2: transfer $10 from account 2 to account 1;
BEGIN;

UPDATE accounts SET balance = balance - 10 WHERE id = 2 RETURNING *;
UPDATE accounts SET balance = balance + 10 WHERE id = 1 RETURNING *

ROLLBACK;