INSERT INTO author (name) VALUES ('Keisuke');
INSERT INTO author (name) VALUES ('Taro');
INSERT INTO author (name) VALUES ('Jiro');
INSERT INTO author (name) VALUES ('Saburo');
INSERT INTO author (name) VALUES ('Shiro');
INSERT INTO author (name) VALUES ('Goro');
INSERT INTO author (name) VALUES ('Rokuro');
INSERT INTO author (name) VALUES ('Shichiro');
INSERT INTO author (name) VALUES ('Hachiro');
INSERT INTO author (name) VALUES ('Kuro');
INSERT INTO author (name) VALUES ('Juro');

INSERT INTO post (id, content, author_id) VALUES (1, 'post1', 1);
INSERT INTO post (id, content, author_id) VALUES (2, 'post2', 2);
INSERT INTO post (id, content, author_id) VALUES (3, 'post3', 3);
INSERT INTO post (id, content, author_id) VALUES (4, 'post4', 4);
INSERT INTO post (id, content, author_id) VALUES (5, 'post5', 5);
INSERT INTO post (id, content, author_id) VALUES (6, 'post6', 6);
INSERT INTO post (id, content, author_id) VALUES (7, 'post7', 7);
INSERT INTO post (id, content, author_id) VALUES (8, 'post8', 8);

INSERT INTO comment (post_id, author_id, content) VALUES (1, 1, 'comment1');    -- post_id : 1
INSERT INTO comment (post_id, author_id, content) VALUES (1, 1, 'comment2');    -- post_id : 1
INSERT INTO comment (post_id, author_id, content) VALUES (1, 1, 'comment3');    -- post_id : 1
INSERT INTO comment (post_id, author_id, content) VALUES (1, 1, 'comment4');    -- post_id : 1
INSERT INTO comment (post_id, author_id, content) VALUES (1, 1, 'comment5');    -- post_id : 1
INSERT INTO comment (post_id, author_id, content) VALUES (1, 2, 'comment6');    -- post_id : 1
INSERT INTO comment (post_id, author_id, content) VALUES (1, 3, 'comment7');    -- post_id : 1
INSERT INTO comment (post_id, author_id, content) VALUES (2, 4, 'comment8');    -- post_id : 2
INSERT INTO comment (post_id, author_id, content) VALUES (2, 5, 'comment9');    -- post_id : 2
INSERT INTO comment (post_id, author_id, content) VALUES (2, 6, 'comment10');   -- post_id : 2
INSERT INTO comment (post_id, author_id, content) VALUES (1, 1, 'comment11');   -- post_id : 1
INSERT INTO comment (post_id, author_id, content) VALUES (1, 1, 'comment12');   -- post_id : 1

COMMIT;