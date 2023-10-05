CREATE TABLE accounts (
    user_id bigserial primary key not null,
    balance bigint not null
);

CREATE TABLE history (
    id bigserial primary key not null,
    from_id bigint not null,
    to_id bigint not null,
    transaction varchar(16) not null,
    amount bigint not null
)