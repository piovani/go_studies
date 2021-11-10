CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS seguidores;

CREATE TABLE usuarios (
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(255) not null,
    criado_em timestamp default current_timestamp
) ENGINE=INNODB;

CREATE TABLE seguidores (
	usuario_id int not null,
	FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE,
	
	seguidor_id int not null,	
	FOREIGN KEY (seguidor_id) REFERENCES usuarios(id) ON DELETE CASCADE,
	
	PRIMARY KEY(usuario_id, seguidor_id)
) ENGINE=INNODB;

CREATE TABLE publicacoes (
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,

    autor_id int not null,
    FOREIGN KEY (autor_id) REFERENCES usuarios(id) ON DELETE CASCADE,

    curtidas int default 0,
    criado_em timestamp default current_timestamp
) ENGINE=INNODB

INSERT INTO usuarios (nome, nick, email, senha) VALUES 
("Usuario 1", "usuario_1", "usuario1@email.com", "$2a$10$wPd3TvVbIyidBJJXy65t1em2sixC7pz/u3G7xeaSkFQGcQB79fafK"),
("Usuario 2", "usuario_2", "usuario2@email.com", "$2a$10$wPd3TvVbIyidBJJXy65t1em2sixC7pz/u3G7xeaSkFQGcQB79fafK"),
("Usuario 3", "usuario_3", "usuario3@email.com", "$2a$10$wPd3TvVbIyidBJJXy65t1em2sixC7pz/u3G7xeaSkFQGcQB79fafK");

INSERT INTO seguidores(usuario_id, seguidor_id) VALUES
(1, 2),
(3, 1),
(1, 3);

INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES 
("Titulo 1", "conteudo 1", "1"),
("Titulo 2", "conteudo 2", "2"),
("Titulo 3", "conteudo 3", "3");