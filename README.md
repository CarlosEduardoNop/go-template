# Go Template

Este é um projeto template para iniciar um novo repo em Go.

### Iniciar projeto

Primeiramente, utilize clone o repositório:
`git clone https://github.com/CarlosEduardoNop/go-template.git`

Após clonar o repositório, é necessário criar o arquivo .env:
`cp .env.example env`

### Comandos iniciais

Para subir o projeto basta rodar o comando `docker compose up -d` e após rodar `go mod tidy`

No início do projeto terá dois comandos disponíveis. O primeiro é o `go run ./cmd/artisan migration -name=`, ele será o comando para criar as migrations. O segundo é o comando `go run ./cmd/artisan migrate`, responsável por rodar as migrations criadas.
Ao rodar o comando **migrate**, já será criado o banco e a tabela de migrações.

### Pacotes iniciais

O projeto já irá iniciar com dois pacotes, sendo eles o **connection** e o **rabbitmq**. Sendo o connection o pacote para conectar com o mysql e o rabbitmq para conectar com a fila.

Também contará com um inicial de como se fazer models, podendo ser encontrado na pasta `internal/models`. Nela consta a organização da pasta e também utilizando funções do pacote internal **repository** para facilitar operações.