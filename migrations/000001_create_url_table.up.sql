CREATE TABLE IF NOT EXISTS url(
    id serial primary key ,
    alias varchar not null unique,
    url varchar not null
);

CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);