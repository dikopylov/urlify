CREATE TABLE references
(
    url        varchar(2000) NOT NULL PRIMARY KEY,
    hash       varchar(172)  NOT NULL UNIQUE,
    created_at timestamp,
)