CREATE TABLE IF NOT EXISTS tb_tweet (
    id bigint not null,
    message varchar(280) not null,
    idUser bigint not null,
    createAt timestamp not null,
    hashtag varchar(256) not null,
    PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS tweet_hashtag on tb_tweet(hashtag);
CREATE INDEX IF NOT EXISTS tweet_creatAt on tb_tweet(createAt);

CREATE TABLE IF NOT EXISTS tb_user (
    id bigint not null,
    name varchar not null,
    followers int not null,
    locate varchar(256) not null,
    PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS user_locate on tb_user(locate);
ALTER TABLE tb_tweet ADD FOREIGN KEY (idUser) REFERENCES tb_user;