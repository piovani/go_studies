CREATE TABLE books (
    id integer generated always as identity,
    name VARCHAR(255) not null,
    description VARCHAR(255),
    medium_price NUMERIC(10, 2),
    author VARCHAR(255),
    img_url VARCHAR(255),
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted TIMESTAMP,
    PRIMARY KEY (id)
);

INSERT INTO books (
                   name,
                   description,
                   medium_price,
                   author,
                   img_url
                   ) VALUES (
                             'Teste 1',
                             'livro teste 1',
                             10.0,
                             'Fulano',
                             'https://images-na.ssl-images-amazon.com/images/I/51Xby92J9KL._SY344_BO1,204,203,200_QL70_ML2_.jpg'
                             )