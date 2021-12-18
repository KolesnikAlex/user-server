-- +goose Up
CREATE TABLE public."users" (
     id SERIAL NOT NULL,
     name VARCHAR(255) NOT NULL,
     login VARCHAR(255) NOT NULL,
     password VARCHAR(255) NOT NULL,
     constraint users_pk
     primary key (id)
);
INSERT INTO public."users" (id, name, login, password)
VALUES (1, 'Name1', 'Login1', 'Pass1');
-- +goose Down
DROP TABLE users;