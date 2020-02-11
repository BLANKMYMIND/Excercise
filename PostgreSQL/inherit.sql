-- 表继承
-- 任务：区分城市和首都
-- 普通做法 (... 代表相同的东西)
-- 创建两个表
CREATE TABLE capitals (
    ...,
    state char(2)
);
CREATE TABLE non_capitals (
    ...
);
-- 再建立他们的视图
CREATE VIEW cities AS
SELECT name, population, altitude FROM capitals
UNION
SELECT name, population, altitude FROM non_capitals;
-- 这样使用 cities 视图查城市的统一信息，若要查某个城市所属的州则使用 capitals


-- 继承做法
CREATE TABLE cities (
    ...
  );
-- 使用 INHERITS (表) 来继承某个表
CREATE TABLE capitals (
    state      char(2)
) INHERITS (cities);
-- 这时可以直接使用 cities 查城市的基本信息（包括继承者 capitals 表）
SELECT name, altitude
FROM cities
WHERE altitude > 500;
-- 表名前加上前缀 only 将只查询该表（不包括继承者）
SELECT name, altitude
FROM ONLY cities
WHERE altitude > 500;

-- 继承没有集成外键和唯一性约束，这限制了他的用途（According 5.10）
