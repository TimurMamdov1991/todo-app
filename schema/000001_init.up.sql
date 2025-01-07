CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name varchar(255) NOT NULL,
                       username varchar(255) NOT NULL UNIQUE,
                       password_hash varchar(255) NOT NULL
);

CREATE TABLE todo_lists (
                            id SERIAL PRIMARY KEY,
                            title varchar(255) NOT NULL,
                            description varchar(255)
);

CREATE TABLE users_lists (
                            id      serial                                           not null unique,
                            user_id int references users (id) on delete cascade      not null,
                            list_id int references todo_lists (id) on delete cascade not null
);

CREATE TABLE todo_items (
                            id SERIAL PRIMARY KEY,
                            title varchar(255) NOT NULL,
                            description varchar(255),
                            done BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE lists_items (
                             id SERIAL PRIMARY KEY,
                             item_id INT REFERENCES todo_items(id) ON DELETE CASCADE NOT NULL,
                             list_id INT REFERENCES todo_lists(id) ON DELETE CASCADE NOT NULL
);