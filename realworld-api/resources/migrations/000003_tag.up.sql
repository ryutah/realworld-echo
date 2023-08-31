create table
    tag (
        name short_text not null,
        created_at timestamp with time zone not null default current_timestamp,
        updated_at timestamp with time zone not null default current_timestamp,
        constraint tag_pkey primary key (name)
    );


comment on table "tag" is 'タグ';


comment on column tag.name is 'タグ名';
