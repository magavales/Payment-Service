CREATE TABLE data_balance (
    id bigserial primary key not null,
    user_id varchar(32) not null,
    balance bigint not null
)