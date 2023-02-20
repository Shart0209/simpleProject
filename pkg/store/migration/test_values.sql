INSERT INTO categories
VALUES (DEFAULT, 'Аукцион'),
       (DEFAULT, 'Единственный поставщик');

INSERT INTO distributors (company_name, contact_name, city, region)
VALUES ('ПАО Ростелеком', 'Sponge Bob', 'Чита', 'Забайкальский край'),
       ('ПАО МТС', 'Sponge Bob', 'Чита', 'Забайкальский край'),
       ('ПАО Мегафон', 'Sponge Bob', 'Чита', 'Забайкальский край');

INSERT INTO authors (first_name, last_name)
VALUES ('Лёня', 'Пупкин');

INSERT INTO contracts
VALUES (DEFAULT,
        'Оказание услуг по ТО и поддержке ПО «ГРАНД-Смета» для нужд Читинского филиала ФГКУ Росгранстрой',
        '100297672122100033 ',
        '2022-12-28',
        1,
        168000.00,
        '2023-01-01',
        '2023-12-31',
        '',
        1,
        1,
        '2022-12-28 13:45:00'),

       (DEFAULT,
        'Оказание услуг по ТО',
        'ЧИТ09-67 ',
        '2022-12-28',
        2,
        168000.00,
        '2023-01-01',
        '2023-12-31',
        '',
        2,
        1,
        '2022-12-28 13:45:00'),

       (DEFAULT,
        'Оказание услуг по предоставлению сети Интернет',
        '45736 ',
        '2022-12-28',
        2,
        168000.00,
        '2023-01-01',
        '2023-12-31',
        '',
        3,
        1,
        '2022-12-28 13:45:00');

INSERT INTO files (file_size, file_path, contract_id)
VALUES (123.2, '.\upload', 1),
       (133.2, '.\upload', 2),
       (3.2, '.\upload', 2);