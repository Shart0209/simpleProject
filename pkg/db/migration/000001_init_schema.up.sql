CREATE TABLE suppliers
(
    supplier_id  integer GENERATED ALWAYS AS IDENTITY,
    name         varchar(255) NOT NULL,
    contact_name varchar(100),
    city         varchar(50),
    region       varchar(100),

    CONSTRAINT pk_supplier_id PRIMARY KEY (supplier_id)
);

CREATE TABLE c_groups
(
    c_groups_id integer GENERATED ALWAYS AS IDENTITY,
    name        varchar(100) NOT NULL,

    CONSTRAINT pk_c_groups_id PRIMARY KEY (c_groups_id)
);

CREATE TABLE categories
(
    category_id integer GENERATED ALWAYS AS IDENTITY,
    name        varchar(100) NOT NULL,

    CONSTRAINT pk_category_id PRIMARY KEY (category_id)
);

CREATE TYPE role AS ENUM ('admin', 'user');

CREATE TABLE authors
(
    author_id integer GENERATED ALWAYS AS IDENTITY,
    login     varchar(50) UNIQUE NOT NULL,
    name      varchar(50)        NOT NULL,
    role      role               NOT NULL,
    pswd_hash varchar            NOT NULL,
    email     varchar(50)        NOT NULL,

    CONSTRAINT pk_author_id PRIMARY KEY (author_id)
);

CREATE TABLE contracts
(
    contract_id integer GENERATED ALWAYS AS IDENTITY,
    title       varchar(100) NOT NULL,
    numb        varchar(50)  NOT NULL,
    date        date         NOT NULL,
    price       float        NOT NULL,
    start_date  date         NOT NULL,
    end_date    date         NOT NULL,
    description varchar(150),
    files       jsonb,

    CONSTRAINT pk_contract_id PRIMARY KEY (contract_id),
    CONSTRAINT chk_contracts_price CHECK (price > 0),
    CONSTRAINT chk_contracts_end_date CHECK (start_date <= end_date ),
    CONSTRAINT chk_contracts_date CHECK (date <= start_date )
);

CREATE TABLE commons
(
    contract_id integer                NOT NULL,
    supplier_id integer                NOT NULL,
    category_id integer                NOT NULL,
    c_groups_id integer                NOT NULL,
    author_id   integer                NOT NULL,
    status      boolean   DEFAULT true NOT NULL,
    created_at  timestamp DEFAULT (now()),
    updated_at  timestamp DEFAULT (now()),

    CONSTRAINT fk_commons_contracts FOREIGN KEY (contract_id) REFERENCES contracts ON DELETE CASCADE,
    CONSTRAINT fk_commons_suppliers FOREIGN KEY (supplier_id) REFERENCES suppliers,
    CONSTRAINT fk_commons_categories FOREIGN KEY (category_id) REFERENCES categories,
    CONSTRAINT fk_commons_groups FOREIGN KEY (c_groups_id) REFERENCES c_groups,
    CONSTRAINT fk_commons_authors FOREIGN KEY (author_id) REFERENCES authors
);

INSERT INTO c_groups
VALUES (DEFAULT, 'Интернет'),
       (DEFAULT, 'Мобильный интернет'),
       (DEFAULT, 'Спутниковый интернет'),
       (DEFAULT, 'Стационарная телефонная связь'),
       (DEFAULT, 'Мобильная телефонная связь'),
       (DEFAULT, 'Гранд-смета'),
       (DEFAULT, 'Техническое обслуживание оргтехники'),
       (DEFAULT, 'Утилизация оргтехники');

INSERT INTO categories
VALUES (DEFAULT, 'Открытый конкурс'),
       (DEFAULT, 'Открытый конкурс в электронной форме'),
       (DEFAULT, 'Электронный аукцион'),
       (DEFAULT, 'Закрытый аукцион'),
       (DEFAULT, 'Закрытый аукцион в электронной форме'),
       (DEFAULT, 'Закрытый конкурс'),
       (DEFAULT, 'Закрытый конкурс в электронной форме'),
       (DEFAULT, 'Конкурс с ограниченным участием'),
       (DEFAULT, 'Конкурс с ограниченным участием в электронной форме'),
       (DEFAULT, 'Закрытый конкурс с ограниченным участием'),
       (DEFAULT, 'Закрытый конкурс с ограниченным участием в электронной форме'),
       (DEFAULT, 'Двухэтапный конкурс'),
       (DEFAULT, 'Двухэтапный конкурс в электронной форме'),
       (DEFAULT, 'Закрытый двухэтапный конкурс'),
       (DEFAULT, 'Закрытый двухэтапный конкурс в электронной форме'),
       (DEFAULT, 'Запрос предложений'),
       (DEFAULT, 'Запрос предложений в электронной форме'),
       (DEFAULT, 'Запрос котировок'),
       (DEFAULT, 'Запрос котировок в электронной форме'),
       (DEFAULT, 'Единственный поставщик (п. 1 ч. 1 ст. 93)'),
       (DEFAULT, 'Единственный поставщик (п. 2 ч. 1 ст. 93)'),
       (DEFAULT, 'Единственный поставщик (п. 4 ч. 1 ст. 93)'),
       (DEFAULT, 'Единственный поставщик (п. 6 ч. 1 ст. 93)'),
       (DEFAULT, 'Единственный поставщик (п. 8 ч. 1 ст. 93)'),
       (DEFAULT, 'Единственный поставщик (п. 9 ч. 1 ст. 93)'),
       (DEFAULT, 'Единственный поставщик (п. 19 ч. 1 ст. 93)'),
       (DEFAULT, 'Единственный поставщик (п. 22 ч. 1 ст. 93)'),
       (DEFAULT, 'Единственный поставщик (п. 23 ч. 1 ст. 93)'),
       (DEFAULT, 'Единственный поставщик (п. 29 ч. 1 ст. 93)'),
       (DEFAULT, 'Единственный поставщик (п. 32 ч. 1 ст. 93)');