CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS publicacoes;

CREATE TABLE IF NOT EXISTS usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    nick VARCHAR(100) NOT NULL unique,
    email VARCHAR(100) NOT NULL unique,
    senha VARCHAR(100) NOT NULL,
    criadoEm timestamp DEFAULT current_timestamp()
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS seguidores(
    usuario_id INT NOT NULL,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
    ON DELETE CASCADE,
    seguidor_id  INT NOT NULL,
    FOREIGN KEY (seguidor_id) REFERENCES usuarios(id)
    ON DELETE CASCADE,
    
    PRIMARY KEY (usuario_id, seguidor_id)
) ENGINE=INNODB;

CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,

    autor_id int not null,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    curtidas int default 0,
    criadaEm timestamp default current_timestamp
) ENGINE=INNODB;