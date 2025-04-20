create table if not exists tasks (
    id serial primary key,
    title text not null,
    description text,
    status text check (status in ('new', 'in_progress', 'done')) default 'new',
    created_at timestamp default now(),
    updated_at timestamp default now());

--insert into tasks (id, title, description) values (1, '1', 'smth');

-- select * from tasks;

--drop table tasks;