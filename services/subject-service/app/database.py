import psycopg2
from psycopg2 import sql
from typing import List, Any, Optional

class PostgresDB:
  def __init__(self, host: str, database: str, user: str, password: str):
    self.host = host
    self.database = database
    self.user = user
    self.password = password
    self.connection = None
    self.cursor = None

  def connect(self):
    """Estabelece a conexão com o banco de dados."""
    try:
      self.connection = psycopg2.connect(
        host=self.host,
        database=self.database,
        user=self.user,
        password=self.password
      )
      self.cursor = self.connection.cursor()
      print("Conexão bem-sucedida ao PostgreSQL!")
    except Exception as error:
      print(f"Erro ao conectar: {error}")

  def close(self):
    """Fecha a conexão e o cursor."""
    if self.cursor:
        self.cursor.close()
    if self.connection:
        self.connection.close()
    print("Conexão fechada.")

  def execute_query(self, query: str, params: Optional[List[Any]] = None) -> List[tuple]:
    """Executa uma consulta SQL e retorna os resultados."""
    try:
      if params:
        self.cursor.execute(query, params)
      else:
        self.cursor.execute(query)
      return self.cursor.fetchall()
    except Exception as error:
      print(f"Erro ao executar a consulta: {error}")
      return []

  def execute_non_query(self, query: str, params: Optional[List[Any]] = None):
    """Executa comandos SQL que não retornam resultados (ex: INSERT, UPDATE, DELETE)."""
    try:
      if params:
        self.cursor.execute(query, params)
      else:
        self.cursor.execute(query)
      self.connection.commit()
      print("Comando executado com sucesso!")
    except Exception as error:
      print(f"Erro ao executar o comando: {error}")
      self.connection.rollback()

# Exemplo de uso da classe:
if __name__ == "__main__":
    # Parâmetros de conexão com o banco de dados
    db = PostgresDB(
        host="localhost",
        database="meu_banco",
        user="meu_usuario",
        password="minha_senha"
    )

    # Conectar ao banco de dados
    db.connect()

    # Exemplo de consulta SELECT
    query = "SELECT * FROM minha_tabela;"
    results = db.execute_query(query)
    print("Resultados da consulta:", results)

    # Exemplo de comando INSERT
    insert_query = "INSERT INTO minha_tabela (coluna1, coluna2) VALUES (%s, %s);"
    db.execute_non_query(insert_query, ("valor1", "valor2"))

    # Fechar a conexão
    db.close()
