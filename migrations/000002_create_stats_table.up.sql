CREATE TABLE IF NOT EXISTS click(
    id serial primary key,
    url_id int not null references url(id) on delete cascade,
    ip varchar(45),
    user_agent varchar,
    device     VARCHAR(20),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
)
