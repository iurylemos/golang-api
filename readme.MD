## Conteúdos
* [Informações gerais](#informacoes-gerais)
* [Tecnologias](#tecnologias)
* [Setup](#setup)

## Informações gerais
Esse projeto foi criado para simular uma pequena rede social que faz publicações, curtidas,
seguir outros usuários, deixar de seguir
	
## Tecnologias
* Docker
* MySQL
* Json Web Token
* Hash (Bcrypt)
* Arquitetura MVC
	
## Setup
Para rodar esse projeto, basta instalar o Go e ter o Docker instalado em sua máquina:

```
$ cd api-nos-golang
$ docker-compose up -d
$ go build
$ ./api-nos-golang.exe
```