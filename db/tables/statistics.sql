create table creator (
    id SERIAL,
    patreon_id text unique not null,
    created_at timestamp,
    primary key (id)
);

create table statistics (
    id serial,
    creator_id int,
    patrons_count int default 0 not null,
    revenues int not null default 0,
    created_at timestamp,
    primary key (id),
    foreign key (creator_id) references creator (id)
);
