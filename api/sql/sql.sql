-- Cria o banco de dados caso ele não exista
CREATE DATABASE IF NOT EXISTS devbook;

-- Usar os próximos comandos dentro do banco devbook
USE devbook;

-- Se tiver uma tabela usuários dentro do banco devbook, vou dropar ela
DROP TABLE IF EXISTS usuarios;

-- Criando a tabela de informações
CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
)ENGINE=INNODB;
