CREATE TABLE "user"(
    uid SERIAL PRIMARY KEY ,
    name TEXT UNIQUE NOT NULL ,
    password TEXT NOT NULL ,
    modified INT NOT NULL DEFAULT 0,
    tags TEXT[]
);

CREATE TABLE canvans(
    bid SERIAL PRIMARY KEY ,
    uid INT UNIQUE NOT NULL ,
    pid TEXT UNIQUE NOT NULL ,
    tags TEXT[]   
);