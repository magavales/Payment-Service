CREATE TABLE accounts (
    user_id bigserial primary key not null,
    balance bigint not null
);

CREATE TABLE history (
    from_id bigint primary key references accounts(user_id) not null,
    to_id bigint not null,
    operation varchar(16) not null,
    amount bigint not null
)