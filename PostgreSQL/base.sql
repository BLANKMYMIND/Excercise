-- hba 位置
show hba_file;

-- 版本
select "version"();

-- 日期
select current_date;

-- real 单精度浮点
create table weather (
	city VARCHAR(80),
	temp_lo int,
	temp_hi int,
	prcp real,
	date date
);

-- point 由 x y 确定的一个点
CREATE TABLE cities (
	name varchar(80),
	location point
);

-- 日期使用 'yyyy-mm-dd'
INSERT INTO weather VALUES ('San Francisco', 46, 50, 0.25, '1994-11-27');

-- point 点使用 '(x, y)'
INSERT INTO cities VALUES ('San Francisco', '(-194.0, 53.0)');

-- 其他示例
INSERT INTO weather (city, temp_lo, temp_hi, prcp, date)
    VALUES ('San Francisco', 43, 57, 0.0, '1994-11-29');

INSERT INTO weather (date, city, temp_hi, temp_lo)
    VALUES ('1994-11-29', 'Hayward', 54, 37);

-- 查询超时时间
select * from pg_settings ps WHERE ps.name like '%timeout%';

-- COPY 执行之前
-- 必须使得所有用户(不知道postgre使用什么用户)有权限
-- 读取： 文件 和 文件夹 的 rx
-- 导出： 文件夹 的 wrx
-- chmod a+rxw 文件/文件夹

-- COPY 命令 从 文件 中读取数据到 表
-- 默认一行一数据，字段间隔为空格（空格被忽略），delimiter 手动决定字段间隔标识
copy weather from '/home/dbuser/text' delimiter ',';

-- COPY 命令 导出 表 到 文件
copy weather to '/home/dbuser/example';

-- 其他示例
select * from weather;
SELECT city, temp_lo, temp_hi, prcp, date FROM weather;

-- 2.0 是浮点除法 2 是整数除法（取小整）
SELECT city, (temp_hi + temp_lo) / 2 as temp_Avg, date from weather;

-- postgres 的运算符可以空格分隔
SELECT * from weather WHERE city = 'San Francisco' AND prcp > 0.0;

-- 按 字符字段 排序时，使用字典序
select * from weather order by city;

-- 取单独值 & 排序
SELECT DISTINCT city from weather;
SELECT DISTINCT city from weather order by city;

-- 链接 （* -> 链接相同的两个字段仍然显示）
select * from weather, cities WHERE cities."name" = weather.city;

-- 链接 （指定显示字段）
SELECT city, temp_lo, temp_hi, prcp, date, location
FROM weather, cities
WHERE city = name;

-- 内连接，等同上： 表 inner join 表 on ( 条件 )
SELECT * FROM weather INNER JOIN cities ON (weather.city = cities.name);

-- 左连接，即使不等同也保留 weather 的内容： 表 left join 表 on ( 条件 )
SELECT * FROM weather LEFT OUTER JOIN cities ON (weather.city = cities.name);

-- 其他示例
SELECT max(temp_lo) FROM weather;

-- Postgres 同样不支持 WHERE 语句中使用聚集函数，错误示例：
-- SELECT city FROM weather WHERE temp_lo = max(temp_lo);
-- 但在括号再 select 一次 聚集函数 还是支持的
SELECT city FROM weather WHERE temp_lo = (SELECT max(temp_lo) FROM weather);

-- 使用 GROUP BY 聚集，聚集函数将应用在每个聚集上，HAVING 专门用于操作聚集内容
SELECT city, max(temp_lo) from weather GROUP BY city HAVING max(temp_lo) < 45;

-- 使用 where 会在聚集前先进行一次过滤
SELECT city, max(temp_lo)
FROM weather
WHERE city LIKE 'S%'
GROUP BY city
HAVING max(temp_lo) < 49;

-- update 示例
UPDATE weather
SET temp_hi = temp_hi - 2,  temp_lo = temp_lo - 2
WHERE date > '1994-11-28';

-- delete 示例
DELETE FROM weather WHERE city = 'NanJing';

-- 视图创建 create view 名称 as
create view myview as
SELECT city, temp_lo, temp_hi, prcp, date, location
FROM weather, cities
WHERE city = name;

-- 以表形式正常查询视图
SELECT * FROM myview;

-- 无法直接插入数据到视图
-- DETAIL:  Views that do not select from a single table or view are not automatically updatable.
-- HINT:  To enable inserting into the view, provide an INSTEAD OF INSERT trigger or an unconditional ON INSERT DO INSTEAD rule.
-- 即，要么是单表视图，要么写好了 Insert trigger
INSERT INTO myview (city, temp_lo, temp_hi, prcp, date, location)
VALUES ('ZhuHai', 48, 54, 0.45, '2018-07-04', '(24, 28)');



