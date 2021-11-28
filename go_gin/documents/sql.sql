create table books (
    id integer generated always as identity,
    name VARCHAR (255) not null,
    description VARCHAR (255),
    medium_price VARCHAR(255),
    author VARCHAR(255),
    img_url VARCHAR(255),
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted TIMESTAMP,
    PRIMARY KEY (id)
);

INSERT INTO books (name) VALUES ('TESTE')