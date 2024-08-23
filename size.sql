-- サイズを確認する

-- テーブル作成
drop table if exists temp;

CREATE TABLE temp (
  value double
);

-- 10万件のdoubleデータを挿入
DELIMITER //
CREATE PROCEDURE insert_temp()
BEGIN
    DECLARE counter INT DEFAULT 0;

    label: LOOP
        INSERT INTO temp (value) VALUES (0.001);
        SET counter = counter + 1;

        IF counter >= 100000 THEN
            LEAVE label;
        END IF;
    END LOOP label;
END //

DELIMITER ;

CALL insert_temp();

-- 10万件のdoubleデータを挿入した後のサイズ
SET SESSION information_schema_stats_expiry= 1;
select
    table_name,
    data_length + index_length,
    floor((data_length+index_length)/1024) AS KB,
    floor((data_length+index_length)/1024/1024) AS MB,
    floor((data_length+index_length)/1024/1024/1024) AS GB
from information_schema.tables
where table_schema = database();

-- TABLE_NAME  data_length + index_length     KB  MB  GB
-- ----------  --------------------------  -----  --  --
-- temp                         3,686,400  3,600   3   0

-- decimalに変更(精度に変更 double 8byte -> decimal 20桁少数10桁)
alter table temp
  modify column value decimal(20, 10);

-- decimalに変更した後のサイズ(alterだけでは、サイズが変わらない)
SET SESSION information_schema_stats_expiry= 1;
select
    table_name,
    data_length + index_length,
    floor((data_length+index_length)/1024) AS KB,
    floor((data_length+index_length)/1024/1024) AS MB,
    floor((data_length+index_length)/1024/1024/1024) AS GB
from information_schema.tables
where table_schema = database();

-- TABLE_NAME  data_length + index_length     KB  MB  GB
-- ----------  --------------------------  -----  --  --
-- temp                         3,686,400  3,600   3   0

-- 一度削除して同じデータを挿入
DELETE FROM temp;
DELIMITER //
CREATE PROCEDURE insert_temp()
BEGIN
    DECLARE counter INT DEFAULT 0;

    label: LOOP
        INSERT INTO temp (value) VALUES (0.001);
        SET counter = counter + 1;

        IF counter >= 100000 THEN
            LEAVE label;
        END IF;
    END LOOP label;
END //

DELIMITER ;

CALL insert_temp();

-- decimal(20, 10)に変更した後のサイズ
SET SESSION information_schema_stats_expiry= 1;
select
    table_name,
    data_length + index_length,
    floor((data_length+index_length)/1024) AS KB,
    floor((data_length+index_length)/1024/1024) AS MB,
    floor((data_length+index_length)/1024/1024/1024) AS GB
from information_schema.tables
where table_schema = database();

-- TABLE_NAME  data_length + index_length     KB  MB  GB
-- ----------  --------------------------  -----  --  --
-- temp                         3,686,400  3,600   3   0


-- 一度データを削除して、精度ギリギリのデータを挿入
DELETE FROM temp;

DROP PROCEDURE insert_temp;

DELIMITER //
CREATE PROCEDURE insert_temp()
BEGIN
    DECLARE counter INT DEFAULT 0;

    label: LOOP
        INSERT INTO temp (value) VALUES (9999999999.9999999999);
        SET counter = counter + 1;

        IF counter >= 100000 THEN
            LEAVE label;
        END IF;
    END LOOP label;
END //

DELIMITER ;

CALL insert_temp();

-- decimal(20, 10)に変更した後のサイズ
SET SESSION information_schema_stats_expiry= 1;
select
    table_name,
    data_length + index_length,
    floor((data_length+index_length)/1024) AS KB,
    floor((data_length+index_length)/1024/1024) AS MB,
    floor((data_length+index_length)/1024/1024/1024) AS GB
from information_schema.tables
where table_schema = database();

-- TABLE_NAME  data_length + index_length     KB  MB  GB
-- ----------  --------------------------  -----  --  --
-- temp                         4,734,976  4,624   4   0

-- doubleに戻す
alter table temp
  modify column value double;

DELETE from temp;

-- doubleの最大値
DROP PROCEDURE insert_temp;

DELIMITER //
CREATE PROCEDURE insert_temp()
BEGIN
    DECLARE counter INT DEFAULT 0;

    label: LOOP
        INSERT INTO temp (value) VALUES (9007199254740992);
        SET counter = counter + 1;

        IF counter >= 100000 THEN
            LEAVE label;
        END IF;
    END LOOP label;
END //

DELIMITER ;

CALL insert_temp();

-- doubleの最大値を挿入した後のサイズ
SET SESSION information_schema_stats_expiry= 1;
select
    table_name,
    data_length + index_length,
    floor((data_length+index_length)/1024) AS KB,
    floor((data_length+index_length)/1024/1024) AS MB,
    floor((data_length+index_length)/1024/1024/1024) AS GB
from information_schema.tables
where table_schema = database();

-- TABLE_NAME  data_length + index_length     KB  MB  GB
-- ----------  --------------------------  -----  --  --
-- temp                         3,686,400  3,600   3   0

-- decimal(16, 0)に変更
alter table temp
  modify column value decimal(16, 0);

-- decimal(16, 0)に変更した後のサイズ
SET SESSION information_schema_stats_expiry= 1;
select
    table_name,
    data_length + index_length,
    floor((data_length+index_length)/1024) AS KB,
    floor((data_length+index_length)/1024/1024) AS MB,
    floor((data_length+index_length)/1024/1024/1024) AS GB
from information_schema.tables
where table_schema = database();

-- TABLE_NAME  data_length + index_length     KB  MB  GB
-- ----------  --------------------------  -----  --  --
-- temp                         3,686,400  3,600   3   0

-- 一度削除
delete from temp;

-- decimal(16, 0)の最大値
DROP PROCEDURE insert_temp;

DELIMITER //
CREATE PROCEDURE insert_temp()
BEGIN
    DECLARE counter INT DEFAULT 0;

    label: LOOP
        INSERT INTO temp (value) VALUES (9999999999999999);
        SET counter = counter + 1;

        IF counter >= 100000 THEN
            LEAVE label;
        END IF;
    END LOOP label;
END //

DELIMITER ;

CALL insert_temp();

-- TABLE_NAME  data_length + index_length     KB  MB  GB
-- ----------  --------------------------  -----  --  --
-- temp                         7,880,704  7,696   7   0

-- 削除
DELETE FROM temp;

-- 同じデータを挿入
DROP PROCEDURE insert_temp;

DELIMITER //
CREATE PROCEDURE insert_temp()
BEGIN
    DECLARE counter INT DEFAULT 0;

    label: LOOP
        INSERT INTO temp (value) VALUES (9007199254740992);
        SET counter = counter + 1;

        IF counter >= 100000 THEN
            LEAVE label;
        END IF;
    END LOOP label;
END //

DELIMITER ;

CALL insert_temp();

SET SESSION information_schema_stats_expiry= 1;
select
    table_name,
    data_length + index_length,
    floor((data_length+index_length)/1024) AS KB,
    floor((data_length+index_length)/1024/1024) AS MB,
    floor((data_length+index_length)/1024/1024/1024) AS GB
from information_schema.tables
where table_schema = database();

-- TABLE_NAME  data_length + index_length     KB  MB  GB
-- ----------  --------------------------  -----  --  --
-- temp                         4,210,688  4,112   4   0
