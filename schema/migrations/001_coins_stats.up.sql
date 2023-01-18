create table if not exists coins_stats
(
    id             bigserial not null
        constraint clients_pk primary key,
    symbol         varchar   not null,
    count          decimal   not null,
    purchase_price decimal   not null
);

insert into coins_stats(symbol, count, purchase_price)
values ('TON', 12.58, 2.511);
insert into coins_stats(symbol, count, purchase_price)
values ('ATOM', 6.544, 12.877);
insert into coins_stats(symbol, count, purchase_price)
values ('DOT', 1, 6.17);