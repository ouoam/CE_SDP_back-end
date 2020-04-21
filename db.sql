create table if not exists public.member
(
    id           serial      not null
        constraint member_pk
            primary key,
    name         varchar(30) not null,
    surname      varchar(30) not null,
    username     varchar(20) not null,
    password     char(60)    not null,
    id_card      bigint,
    email        varchar(30) not null,
    verification smallint,
    bank_account bigint,
    address      text
);

alter table public.member
    owner to postgres;

create unique index if not exists member_username_uindex
    on public.member (username);

create unique index if not exists member_email_uindex
    on public.member (email);

create table if not exists public.message
(
    from_uid integer                 not null
        constraint message_member_id_fk
            references public.member,
    to_uid   integer                 not null
        constraint message_member_id_fk_2
            references public.member,
    time     timestamp default now() not null,
    message  text                    not null,
    constraint message_pk
        primary key (from_uid, to_uid, time)
);

alter table public.message
    owner to postgres;

create table if not exists public.tour
(
    id          serial  not null
        constraint tour_pk
            primary key,
    owner_id    integer not null
        constraint tour_member_id_fk
            references public.member,
    name        varchar not null,
    description text    not null,
    category    varchar not null,
    max_member  integer not null,
    first_day   date    not null,
    last_day    date    not null,
    price       integer not null,
    status      integer not null
);

alter table public.tour
    owner to postgres;

create table if not exists public.transcript
(
    id        serial                  not null
        constraint transcript_pk
            primary key,
    tour_id   integer                 not null
        constraint transcript_tour_id_fk
            references public.tour,
    user_id   integer                 not null
        constraint transcript_member_id_fk
            references public.member,
    file      varchar                 not null,
    confirm   integer,
    is_cancel integer,
    time      timestamp default now() not null
);

alter table public.transcript
    owner to postgres;

create table if not exists public.favorite
(
    user_id integer not null
        constraint favorite_member_id_fk
            references public.member,
    tour_id integer not null
        constraint favorite_tour_id_fk
            references public.tour,
    constraint favorite_pk
        primary key (user_id, tour_id)
);

alter table public.favorite
    owner to postgres;

create index if not exists favorite_tour_id_index
    on public.favorite (tour_id);

create index if not exists favorite_user_id_index
    on public.favorite (user_id);

create table if not exists public.review
(
    id      serial                  not null
        constraint review_pk
            primary key,
    user_id integer                 not null
        constraint review_member_id_fk
            references public.member,
    tour_id integer                 not null
        constraint review_tour_id_fk
            references public.tour,
    comment text                    not null,
    ratting integer                 not null,
    score   integer                 not null,
    time    timestamp default now() not null
);

alter table public.review
    owner to postgres;

create index if not exists review_tour_id_index
    on public.review (tour_id);

create table if not exists public.place
(
    id   serial  not null
        constraint place_pk
            primary key,
    name varchar not null,
    pic  varchar,
    geo  point   not null
);

alter table public.place
    owner to postgres;

create table if not exists public.tour_place_list
(
    tour_id integer not null
        constraint tour_place_list_tour_id_fk
            references public.tour,
    seq     integer not null,
    place   integer not null
        constraint tour_place_list_place_id_fk
            references public.place,
    constraint tour_place_list_pk
        primary key (tour_id, seq)
);

alter table public.tour_place_list
    owner to postgres;


