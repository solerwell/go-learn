
CREATE TABLE IF NOT EXISTS main.product
(
    id BIGINT PRIMARY KEY ,
    name VARCHAR(255),
    brand VARCHAR(255),
    barPrice float ,
    -- 每包的价格
    packPrice float,
    -- 评论数
    commentNum INT ,
    -- 评分数
    pingNum INT ,
    -- 综合得分
    comSore FLOAT ,
    -- 口味得分
    scoreWei FLOAT ,
    -- 外观得分
    scoreBao FLOAT ,
    -- 性价比得分
    scoreJia FLOAT
);