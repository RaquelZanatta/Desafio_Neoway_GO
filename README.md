
# Docker - Golang - PostgreSQL

Este é um serviço de ETL desenvolvido na linguagem Golang que realiza as seguintes etapas:

- Lê um arquivo csv em formato tabular;
- Transforma os dados pelas lib;
- Valida se todos os CPFs e CNPJs estão válidos, armazenando as informações em 3 novas;
- Realiza a conexão com o banco de dados PostgreeSQL;
- Cria a tabela e insere os dados no banco de dados.

Para a execução do serviço, foram criados os arquivos DockFile e docker-compose, responsáveis por criar os conteiners necessários para executar a aplicação. Sendo eles:

- Golang: responsável pela execução da lógica do serviço;
- Postgree: responsável por hospedar o banco de dados que armazena os dados;
- Adminer: responsável por hospedar a aplicação Adminer, utilizada para acessar e visualizar o banco de dados.

Arquivo Dockerfile:

* FROM golang: define a imagem base para o container como a imagem oficial do Go;
* RUN mkdir /app: cria um diretório chamado /app no container;
* WORKDIR /app: define o diretório de trabalho padrão como /app;
* COPY main.go /app/: copia o arquivo main.go do diretório local para o diretório /app no container;
* COPY base_teste[802].txt /app/: copia o arquivo base_teste[802].txt do diretório local para o diretório /app no container.
* COPY . /app/: copia todo o conteúdo do diretório local para o diretório /app no container.
* RUN go mod init go-neoway: inicializa um novo módulo Go chamado go-neoway.
* RUN go get github.com/lib/pq: instala a dependência github.com/lib/pq no módulo go-neoway.
* RUN go mod tidy: remove as dependências não utilizadas e atualiza o arquivo go.mod com as versões mais recentes das dependências utilizadas.

Para executar a aplicação, siga os passos abaixo:

- Abra um terminal na pasta raiz do projeto e execute o comando:

  - docker-compose up -d
- Após o término do processo acima, execute o seguinte comando para rodar o serviço:

  - docker-compose run app
- Abra um navegador e acesse a URL http://localhost:8081/ com os seguintes dados:

  - Sistema: PostgreSQL
  - Servidor: db
  - Usuário: neoway2023
  - Senha: neoway2023
  - Base de dados: clients
