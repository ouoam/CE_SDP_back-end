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
    verify       boolean default false not null,
    pic          char(43),
    bank_name    varchar,
    id_card_pic  char(43)
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
    id          serial    not null
        constraint tour_pk
            primary key,
    owner       integer   not null
        constraint tour_member_id_fk
            references public.member,
    name        varchar   not null,
    description text      not null,
    category    varchar   not null,
    max_member  integer   not null,
    first_day   timestamp not null,
    last_day    timestamp not null,
    price       integer   not null,
    status      smallint  not null,
    pic         char(43)
);

alter table public.tour
    owner to postgres;

create index tour_category_index
    on public.tour (category);

create index tour_description_index
    on public.tour (description);

create index tour_name_index
    on public.tour (name);

create index tour_first_day_index
    on public.tour (first_day);

create table public.place
(
    id   serial           not null
        constraint place_pk
            primary key,
    name varchar          not null,
    pic  char(43),
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
    tour    integer               not null
        constraint transcript_tour_id_fk
            references public.tour,
    "user"  integer               not null
        constraint transcript_member_id_fk
            references public.member,
    file    char(43),
    confirm boolean default false not null,
    time    timestamp,
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
            (id, owner, name, description, category, max_member, first_day, last_day, price, status, pic, member,
             confirm, ratting, favorite, g_name, g_surname, bank_account, bank_name, list)
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
       tu.pic,
       COALESCE(ts.member, 0::bigint)  AS member,
       COALESCE(ts.confirm, 0::bigint) AS confirm,
       COALESCE(r.ratting, 0::numeric) AS ratting,
       COALESCE(f.favorite, 0::bigint) AS favorite,
       m.name                          AS g_name,
       m.surname                       AS g_surname,
       m.bank_account,
       m.bank_name,
       l.list
FROM tour tu
         LEFT JOIN (SELECT transcript.tour,
                           count(transcript."user")                 AS member,
                           count(NULLIF(false, transcript.confirm)) AS confirm
                    FROM transcript
                    GROUP BY transcript.tour) ts ON tu.id = ts.tour
         LEFT JOIN (SELECT favorite.tour,
                           count(favorite."user") AS favorite
                    FROM favorite
                    GROUP BY favorite.tour) f ON tu.id = f.tour
         LEFT JOIN (SELECT review.tour,
                           avg(review.ratting) AS ratting
                    FROM review
                    GROUP BY review.tour) r ON tu.id = r.tour
         LEFT JOIN member m ON tu.owner = m.id
         LEFT JOIN (SELECT l_1.tour,
                           array_agg(p.name) AS list
                    FROM (SELECT list.tour,
                                 list.seq,
                                 list.place
                          FROM list
                          ORDER BY list.tour, list.seq) l_1
                             LEFT JOIN place p ON l_1.place = p.id
                    GROUP BY l_1.tour) l ON l.tour = tu.id;

alter table public.tourdetail
    owner to postgres;

create function public.messagewithme(me integer, contact integer)
    returns TABLE
            (
                me      boolean,
                message text,
                "time"  timestamp without time zone
            )
    language sql
as
$$
SELECT case when "from" = $1 then true else false end as me,
       message.message,
       time
FROM message
WHERE ("from" = $1 AND "to" = $2)
   OR ("from" = $2 AND "to" = $1)
ORDER BY time DESC
$$;

alter function public.messagewithme(integer, integer) owner to postgres;

create function public.listupdate(tour integer, list integer[]) returns integer
    language plpgsql
as
$$
DECLARE
    i int := 0;
    x int;
BEGIN
    FOREACH x IN ARRAY $2
        LOOP
            EXECUTE
                'INSERT INTO list(tour, seq, place) VALUES($1, $2, $3)
                ON CONFLICT (tour, seq)
                DO UPDATE SET place = excluded.place' USING $1, i, x;
            i := i + 1;
        END LOOP;
    EXECUTE 'DELETE FROM list WHERE tour = $1 AND seq >= $2' USING $1, i;
    RETURN i;
END;
$$;

alter function public.listupdate(integer, integer[]) owner to postgres;

create function public.reviewwithuser("user" integer)
    returns TABLE
            (
                tour    integer,
                comment text,
                ratting smallint,
                "time"  timestamp without time zone,
                name    text
            )
    language sql
as
$$
SELECT r.tour, r.comment, r.ratting, r.time, t.name
FROM review r
         LEFT JOIN tour t on r.tour = t.id
WHERE r."user" = $1
ORDER BY r.time DESC ;
$$;

alter function public.reviewwithuser(integer) owner to postgres;

create function public.reviewwithtour(tour integer)
    returns TABLE
            (
                "user"  integer,
                comment text,
                ratting smallint,
                "time"  timestamp without time zone,
                name    text,
                surname text
            )
    language sql
as
$$
SELECT r."user", r.comment, r.ratting, r.time, m.name, m.surname
FROM review r
         LEFT JOIN member m on r."user" = m.id
WHERE r.tour = $1
ORDER BY r.time DESC ;
$$;

alter function public.reviewwithtour(integer) owner to postgres;

create function public.favoritewithuser("user" integer)
    returns TABLE
            (
                tour integer,
                name text
            )
    language sql
as
$$
SELECT f.tour, t.name
FROM favorite f
         LEFT JOIN tour t on f.tour = t.id
WHERE f."user" = $1;
$$;

alter function public.favoritewithuser(integer) owner to postgres;

create function public.placesearch(keyword text)
    returns TABLE
            (
                id   integer,
                name text,
                pic  character,
                lat  double precision,
                lon  double precision
            )
    language sql
as
$$
SELECT *
FROM place
WHERE name LIKE ('%' || $1 || '%');
$$;

alter function public.placesearch(text) owner to postgres;

create function public.messagelistme(me integer)
    returns TABLE
            (
                contact integer,
                me      boolean,
                message text,
                "time"  timestamp without time zone,
                name    text,
                surname text,
                pic     character
            )
    language sql
as
$$
SELECT DISTINCT ON (a.contact) a.*, m.name, m.surname, m.pic
FROM (SELECT case when "from" = $1 then "to" else "from" end as contact,
             case when "from" = $1 then true else false end  as me,
             message.message,
             time
      FROM message
      WHERE "from" = $1
         OR "to" = $1
      ORDER BY time DESC) as a
         LEFT JOIN member as m ON m.id = a.contact;
$$;

alter function public.messagelistme(integer) owner to postgres;

create function public.listwithtour(tour integer)
    returns TABLE
            (
                id   integer,
                name text,
                pic  character,
                lat  double precision,
                lon  double precision
            )
    language sql
as
$$
SELECT p.*
FROM list l
         LEFT JOIN place p on l.place = p.id
WHERE l.tour = $1
ORDER BY l.seq;
$$;

alter function public.listwithtour(integer) owner to postgres;

create function public.tourdetailsearch(keyword text)
    returns TABLE
            (
                id           integer,
                owner        integer,
                name         text,
                description  text,
                category     text,
                max_member   integer,
                first_day    timestamp without time zone,
                last_day     timestamp without time zone,
                price        integer,
                status       smallint,
                pic          character,
                member       bigint,
                confirm      bigint,
                ratting      numeric,
                favorite     bigint,
                g_name       text,
                g_surname    text,
                bank_account bigint,
                bank_name    text,
                list         character varying[]
            )
    language sql
as
$$
SELECT *
FROM tourdetail
WHERE description LIKE ('%' || $1 || '%')
   OR name LIKE ('%' || $1 || '%')
   OR array_to_string(list, ',') LIKE ('%' || $1 || '%');
$$;

alter function public.tourdetailsearch(text) owner to postgres;

create function public.transcriptwithuser("user" integer)
    returns TABLE
            (
                tour    integer,
                file    character,
                confirm boolean,
                "time"  timestamp without time zone,
                name    text
            )
    language sql
as
$$
SELECT t.tour, t.file, t.confirm, t.time, t2.name
FROM transcript t
         LEFT JOIN tour t2 on t.tour = t2.id
WHERE t."user" = $1
ORDER BY t.time DESC ;
$$;

alter function public.transcriptwithuser(integer) owner to postgres;

create function public.transcriptwithtour(tour integer)
    returns TABLE
            (
                "user"  integer,
                file    character,
                confirm boolean,
                "time"  timestamp without time zone,
                name    text,
                surname text
            )
    language sql
as
$$
SELECT t."user", t.file, t.confirm, t.time, m.name, m.surname
FROM transcript t
         LEFT JOIN member m on t."user" = m.id
WHERE t.tour = $1
ORDER BY t.time DESC ;
$$;

alter function public.transcriptwithtour(integer) owner to postgres;


