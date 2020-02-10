-- 窗口函数 Window Function 对标于 聚类函数 Aggregate Function

-- Example-1
SELECT depname, empno, salary,
       avg(salary) OVER (PARTITION BY depname)
FROM empsalary;

-- 这样可以每个员工的工资与其部门的平均工资作对比
-- avg(salary) OVER (PARTITION BY depname) 是一个整体，整体是一个窗口函数。
-- PARTITION 是分区的意思，即将表中条目按照部分划分计算 avg
-- 如果不使用窗口，而使用聚集，那么需要用 GROUP BY 语句，共查询两次

-- Example-2
SELECT depname, empno, salary,
       rank() OVER (PARTITION BY depname ORDER BY salary DESC)
FROM empsalary;

-- rank() OVER () 也是一个窗口函数，将按区排名(即可能有多个1名次)，并生成一个排名字段，靠前的行名次低， OVER 子句中可加上 ORDER BY 进行排序

-- Example-3
SELECT salary,
       sum(salary) OVER ()
FROM empsalary;

-- 省略 OVER 括号内容时，分区为整个表，同时也可使用 ORDER BY 等排序语句

-- Example-4
SELECT depname, empno, salary, enroll_date
FROM (
         SELECT depname, empno, salary, enroll_date,
                rank() OVER (PARTITION BY depname ORDER BY salary DESC, empno) AS pos
         FROM empsalary
     ) AS ss
WHERE pos < 3;

-- 使用子句 & 重命名 可以快速处理 查取前n名的 情况

-- Example-5
SELECT
            sum(salary) OVER w,
            avg(salary) OVER w
FROM empsalary
    WINDOW w AS (PARTITION BY depname ORDER BY salary DESC);

-- OVER () 的括号子句可以用符号代替并重用，在语句后面使用
-- WINDOW 符号 AS (内容) 来填写内容
