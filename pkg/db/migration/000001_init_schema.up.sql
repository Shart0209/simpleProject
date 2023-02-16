CREATE TABLE distributors
(
    distributor_id serial,
    company_name   varchar NOT NULL,
    contact_name   varchar,
    city           varchar,
    region         varchar,

    CONSTRAINT pk_distributor_id PRIMARY KEY (distributor_id)
);

CREATE TABLE categories
(
    category_id   serial,
    category_name character varying(100) NOT NULL,

    CONSTRAINT pk_category_id PRIMARY KEY (category_id)
);

CREATE TABLE contracts
(
    contract_id    serial,
    title          varchar NOT NULL,
    contr_number         varchar NOT NULL,
    doc_date       date    NOT NULL,
    category_id    integer NOT NULL,
    price          float   NOT NULL,
    start_date     date,
    end_date       date,
    description    varchar,
    distributor_id integer NOT NULL,
    primary_author varchar(100),
    created_at     timestamp DEFAULT (now()),

    CONSTRAINT pk_contract_id PRIMARY KEY (contract_id),
    CONSTRAINT chk_contracts_price CHECK (price > 0),
    CONSTRAINT chk_contracts_end_date CHECK (end_date > start_date),
    CONSTRAINT fk_contracts_distributors FOREIGN KEY (distributor_id) REFERENCES distributors,
    CONSTRAINT fk_contracts_categories FOREIGN KEY (category_id) REFERENCES categories
);

CREATE TABLE files
(
    file_id     serial,
    file_name   varchar(100) NOT NULL,
    file_size   float,
    file_path   varchar(255) NOT NULL,
    contract_id integer,

    CONSTRAINT pk_file_id PRIMARY KEY (file_id),
    CONSTRAINT fk_files_contracts FOREIGN KEY (contract_id) REFERENCES contracts
);