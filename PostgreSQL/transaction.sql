-- 事务
BEGIN;
UPDATE weather SET city = 'DongJing'
WHERE city = 'BeiJing';
-- 事务中可以看到事务的更改
SELECT * FROM weather;
ROLLBACK;

-- 回退后无任何更改痕迹
SELECT * FROM weather;

-- 不回退
BEGIN;
UPDATE weather SET city = 'DongJing'
WHERE city = 'BeiJing';
COMMIT;

SELECT * FROM weather;

-- 保存点 - 事务
BEGIN;
UPDATE weather SET city = 'BeiJing'
WHERE city = 'DongJing';
SAVEPOINT tobeijing; -- 记录并起名保存点
UPDATE weather SET city = 'Tokoy'
WHERE city = 'BeiJing';
ROLLBACK TO tobeijing; -- 回滚到保存点，并在之后进行下面的内容
UPDATE weather SET city = 'PeKing'
WHERE city = 'BeiJing';
COMMIT;
SELECT * FROM weather;
