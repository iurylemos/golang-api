INSERT INTO usuarios(nome, nick, email, senha)
VALUES
("Usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$9V.PCfReNgclXmiNxivK..65jsj6rLRleCq6HUgEdeMvAUqhAnGjO"),
("Usuario 2", "usuario_2", "usuario2@gmail.com", "$2a$10$9V.PCfReNgclXmiNxivK..65jsj6rLRleCq6HUgEdeMvAUqhAnGjO"),
("Usuario 3", "usuario_3", "usuario3@gmail.com", "$2a$10$9V.PCfReNgclXmiNxivK..65jsj6rLRleCq6HUgEdeMvAUqhAnGjO");

INSERT INTO seguidores(usuario_id, seguidor_id)
VALUES
(1, 2),
(2, 1),
(3, 1);

INSERT INTO publicacoes(titulo, conteudo, autor_id)
VALUES
("Publicacao do usuario 1", "Essa e a publicacao do usuario 1. Oba!", 1),
("Publicacao do usuario 2", "Essa e a publicacao do usuario 2. Oba!", 2),
("Publicacao do usuario 3", "Essa e a publicacao do usuario 3. Oba!", 3);