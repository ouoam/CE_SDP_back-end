create table public.member
(
    id           serial                not null
        constraint member_pk
            primary key,
    name         varchar(30)           not null,
    surname      varchar(30)           not null,
    username     varchar(20)           not null,
    password     char(60)              not null,
    id_card      bigint,
    email        varchar(30)           not null,
    bank_account bigint,
    address      text,
    verify       boolean default false not null
);

alter table public.member
    owner to postgres;

create unique index member_username_uindex
    on public.member (username);

create unique index member_email_uindex
    on public.member (email);

create table public.message
(
    "from"  integer                 not null
        constraint message_member_id_fk
            references public.member,
    "to"    integer                 not null
        constraint message_member_id_fk_2
            references public.member,
    time    timestamp default now() not null,
    message text                    not null,
    constraint message_pk
        primary key ("from", "to", time)
);

alter table public.message
    owner to postgres;

create index message_from_index
    on public.message ("from");

create index message_to_index
    on public.message ("to");

create table public.tour
(
    id          serial   not null
        constraint tour_pk
            primary key,
    owner       integer  not null
        constraint tour_member_id_fk
            references public.member,
    name        varchar  not null,
    description text     not null,
    category    varchar  not null,
    max_member  integer  not null,
    first_day   date     not null,
    last_day    date     not null,
    price       integer  not null,
    status      smallint not null
);

alter table public.tour
    owner to postgres;

create index tour_category_index
    on public.tour (category);

create index tour_description_index
    on public.tour (description);

create index tour_first_day_index
    on public.tour (first_day);

create index tour_name_index
    on public.tour (name);

create table public.place
(
    id   serial           not null
        constraint place_pk
            primary key,
    name varchar          not null,
    pic  varchar,
    lat  double precision not null,
    lon  double precision not null
);

alter table public.place
    owner to postgres;

create index place_name_index
    on public.place (name);

create table public.list
(
    tour  integer not null
        constraint tour_place_list_tour_id_fk
            references public.tour,
    seq   integer not null,
    place integer not null
        constraint tour_place_list_place_id_fk
            references public.place,
    constraint tour_place_list_pk
        primary key (tour, seq)
);

alter table public.list
    owner to postgres;

create table public.transcript
(
    tour    integer                 not null
        constraint transcript_tour_id_fk
            references public.tour,
    "user"  integer                 not null
        constraint transcript_member_id_fk
            references public.member,
    file    varchar,
    confirm boolean   default false not null,
    time    timestamp default now() not null,
    constraint transcript_pk
        primary key (tour, "user")
);

alter table public.transcript
    owner to postgres;

create table public.review
(
    tour    integer                 not null
        constraint review_tour_id_fk
            references public.tour,
    "user"  integer                 not null
        constraint review_member_id_fk
            references public.member,
    comment text                    not null,
    ratting smallint                not null,
    time    timestamp default now() not null,
    constraint review_pk
        primary key ("user", tour)
);

alter table public.review
    owner to postgres;

create index review_tour_index
    on public.review (tour);

create table public.favorite
(
    tour   integer not null
        constraint favorite_tour_id_fk
            references public.tour,
    "user" integer not null
        constraint favorite_member_id_fk
            references public.member,
    constraint favorite_pk
        primary key (tour, "user")
);

alter table public.favorite
    owner to postgres;

create view public.tourdetail
            (id, owner, name, description, category, max_member, first_day, last_day, price, status, member, confirm,
             ratting, favorite)
as
SELECT tu.id,
       tu.owner,
       tu.name,
       tu.description,
       tu.category,
       tu.max_member,
       tu.first_day,
       tu.last_day,
       tu.price,
       tu.status,
       ts.member,
       ts.confirm,
       r.ratting,
       f.favorite
FROM tour tu,
     (SELECT transcript.tour,
             count(transcript."user")                 AS member,
             count(NULLIF(false, transcript.confirm)) AS confirm
      FROM transcript
      GROUP BY transcript.tour) ts,
     (SELECT favorite.tour,
             count(favorite."user") AS favorite
      FROM favorite
      GROUP BY favorite.tour) f,
     (SELECT review.tour,
             avg(review.ratting) AS ratting
      FROM review
      GROUP BY review.tour) r
WHERE tu.id = f.tour
  AND tu.id = ts.tour
  AND tu.id = r.tour;

alter table public.tourdetail
    owner to postgres;


