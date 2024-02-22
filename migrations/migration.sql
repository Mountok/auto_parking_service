create table users
(
    id serial not null
        constraint users_pk
            primary key,
    name text,
    rank text
);

alter table users owner to postgres;

create unique index users_id_uindex
	on users (id);

create table cars
(
    id serial not null
        constraint cars_pk
            primary key,
    user_id integer
        constraint cars_users_id_fk
            references users,
    colour text,
    brand text,
    license_plate text
);

alter table cars owner to postgres;

create unique index cars_id_uindex
	on cars (id);



INSERT INTO public.users (id, name, rank) VALUES (1, 'John K.', 'CEO');
INSERT INTO public.users (id, name, rank) VALUES (2, 'Lana V.', 'VP Business');
INSERT INTO public.users (id, name, rank) VALUES (3, 'Jimmy J.', 'CTO');
INSERT INTO public.users (id, name, rank) VALUES (4, 'Betty P.', 'CFO');
INSERT INTO public.users (id, name, rank) VALUES (5, 'Lark K.', 'Senior Engineer');
INSERT INTO public.users (id, name, rank) VALUES (6, 'Rob P.', 'Senior Engineer');
INSERT INTO public.users (id, name, rank) VALUES (7, 'Stephen K.', 'Senior Engineer');
INSERT INTO public.users (id, name, rank) VALUES (8, 'Vanessa M.P.', 'QA Engineer');
INSERT INTO public.users (id, name, rank) VALUES (9, 'Jimmy C.', 'Sales');
INSERT INTO public.users (id, name, rank) VALUES (10, 'Norah H.', 'Sales');


INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (1, 2, 'blue', 'Nissan', 'LB081G');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (2, 1, 'white', 'Mazda', 'CM511A');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (3, 9, 'red', 'Kia', 'HT237C');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (4, 7, 'white', 'Acura', 'AX947X');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (5, 1, 'blue', 'Ford', 'DA387P');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (6, 7, 'black', 'BMW', 'FB026X');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (7, 5, 'gray', 'Kia', 'FR078W');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (8, 4, 'cyan', 'Mazda', 'DN199O');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (9, 9, 'red', 'Range Rover', 'NV451G');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (10, 7, 'red', 'BMW', 'AP202Z');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (11, 4, 'white', 'Nissan', 'QY447Y');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (12, 4, 'yellow', 'Range Rover', 'WJ953P');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (13, 2, 'gray', 'MINI', 'FR891E');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (14, 7, 'red', 'Mercedes', 'TA503E');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (15, 6, 'green', 'Subaru', 'BZ515J');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (16, 1, 'red', 'Nissan', 'NW632K');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (17, 2, 'green', 'Subaru', 'MF267Z');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (18, 5, 'green', 'Audi', 'KX079A');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (19, 4, 'green', 'Ford', 'ZL052T');
INSERT INTO public.cars (id, user_id, colour, brand, license_plate) VALUES (20, 1, 'gray', 'Kia', 'TC818A');

ALTER SEQUENCE cars_id_seq RESTART WITH 21
ALTER SEQUENCE users_id_seq RESTART WITH 11