# Usando uma imagem base com Go
FROM golang:1.20

# Definir o diretório de trabalho
WORKDIR /app

# Copiar os arquivos de configuração do Go
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod download

# Copiar o código-fonte para dentro do contêiner
COPY . .

# Mudar para o diretório onde o main.go está localizado
WORKDIR /app/cmd/app

# Compilar o binário da aplicação
RUN go build -o main .

# Definir o comando para rodar a aplicação
CMD ["./main"]
