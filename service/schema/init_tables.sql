CREATE TABLE data_balance (
    id bigserial primary key not null,
    user_id varchar(64) not null,
    balance bigint not null
)

CREATE TABLE history_table (
    user_id varchar(64) not null,
    to_id varchar(64) not null,
    operation varchar(16) not null,
    amount bigint not null
)