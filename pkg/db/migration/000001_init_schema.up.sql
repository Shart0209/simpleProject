CREATE TABLE documents
(
    document_id integer NOT NULL,
    doc_name    text    NOT NULL,
    doc_number  text    NOT NULL,
    doc_date    date    NOT NULL,
    doc_method  varchar(100) NOT NULL,
    price       float   NOT NULL,
    start_date  date,
    end_date    date,
    description text,
    distributor varchar(100),
    created_at  timestamp DEFAULT (now())
);

CREATE TABLE files
(
    file_id     integer NOT NULL,
    file_name   varchar(100) NOT NULL,
    file_size   float,
    file_path   varchar(255) NOT NULL,
    document_id integer
);

ALTER TABLE ONLY documents
    ADD CONSTRAINT pk_document_id PRIMARY KEY (document_id);

ALTER TABLE ONLY files
    ADD CONSTRAINT pk_file_id PRIMARY KEY (file_id);

ALTER TABLE ONLY files
    ADD CONSTRAINT fk_file_document FOREIGN KEY (document_id) REFERENCES documents;