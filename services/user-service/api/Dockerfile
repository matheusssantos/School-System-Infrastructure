# Etapa 1: Use uma imagem base do Maven
FROM maven:3.8.6 AS build

WORKDIR /app

# Copie o código fonte para dentro do contêiner
COPY . /app

# Compile o projeto (isso irá gerar o .jar dentro da pasta /app/target)
RUN mvn clean package -DskipTests

# Etapa 2: Use uma imagem base do OpenJDK 17 para rodar a aplicação
FROM openjdk:17-jdk-slim

WORKDIR /app

# Copie o .jar gerado da etapa de build
COPY --from=build /app/target/*.jar app.jar

# Defina o comando para rodar o .jar
CMD ["java", "-jar", "app.jar"]

# Exponha a porta 8080
EXPOSE 8080
