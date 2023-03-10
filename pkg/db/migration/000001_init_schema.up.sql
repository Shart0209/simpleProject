CREATE TABLE contracts
(
    contract_id    integer GENERATED ALWAYS AS IDENTITY,
    title          varchar      NOT NULL,
    contr_number   varchar(100) NOT NULL,
    contr_date     date         NOT NULL,
    category_id    integer      NOT NULL,
    price          float        NOT NULL,
    start_date     date         NOT NULL,
    end_date       date         NOT NULL,
    distributor_id integer      NOT NULL,
    description    varchar,
    files          jsonb,
    author_id      integer,
    created_at     timestamp DEFAULT (now()),
    updated_at     timestamp DEFAULT (now()),

    CONSTRAINT pk_contract_id PRIMARY KEY (contract_id),
    CONSTRAINT chk_contracts_price CHECK (price > 0),
    CONSTRAINT chk_contracts_end_date CHECK (start_date <= end_date ),
    CONSTRAINT chk_contracts_contr_date CHECK (contr_date <= start_date ),
    CONSTRAINT fk_contracts_distributors FOREIGN KEY (distributor_id) REFERENCES distributors,
    CONSTRAINT fk_contracts_categories FOREIGN KEY (category_id) REFERENCES categories,
    CONSTRAINT fk_contracts_authors FOREIGN KEY (author_id) REFERENCES authors
);

CREATE TABLE distributors
(
    distributor_id integer GENERATED ALWAYS AS IDENTITY,
    company_name   varchar(255) NOT NULL,
    contact_name   varchar(100),
    company_city   varchar(50),
    region         varchar(100),

    CONSTRAINT pk_distributor_id PRIMARY KEY (distributor_id)
);

CREATE TABLE categories
(
    category_id   integer GENERATED ALWAYS AS IDENTITY,
    category_name varchar(100) NOT NULL,

    CONSTRAINT pk_category_id PRIMARY KEY (category_id)
);


CREATE TABLE authors
(
    author_id   integer GENERATED ALWAYS AS IDENTITY,
    first_name  varchar(100) NOT NULL,
    last_name   varchar(100) NOT NULL,
    middle_name varchar(100),

    CONSTRAINT pk_author_id PRIMARY KEY (author_id)
);

