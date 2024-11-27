import asyncpg
from app.logger import Logger
import os
from dotenv import load_dotenv

load_dotenv()

DB_USER = "postgres"  # O valor de POSTGRES_USER
DB_PASSWORD = "postgres"  # O valor de POSTGRES_PASSWORD
DB_HOST = "postgres"  # O nome do host no Docker Desktop
DB_PORT = "5432"  # A porta do banco de dados
DB_NAME = "school"  # O valor de POSTGRES_DB

class Database:
  
  connection = None
  
  @staticmethod
  async def connect():
    if Database.connection is None:
      try:
        connection_url = "postgresql://postgres:postgres@postgres:5432/school"
                
        # Conectando com a URL de conexão
        Database.connection = await asyncpg.connect(connection_url)

        # Database.connection = await asyncpg.connect(
        #   user=DB_USER,
        #   password=DB_PASSWORD,
        #   host=DB_HOST,
        #   port=DB_PORT,
        #   database=DB_NAME
        # )
        Logger.database("Conexão bem-sucedida com o banco de dados!")
      except Exception as e:
        Logger.database("Falha ao conectar no banco!", error=True)
        print(e)
        
  @staticmethod
  async def disconnect():
    if Database.connection:
      await Database.connection.close()
      Database.connection = None
      Logger.database("Conexão encerrada com o banco de dados!")
      
  async def create_tables():
    try:
      await Database.connection.execute("""
        CREATE TABLE subjects (
          id SERIAL PRIMARY KEY,
          code VARCHAR NOT NULL,
          name VARCHAR NOT NULL,
          turn CHAR(1) NOT NULL
        );
      """)
      Logger.database("Tabela 'subjects' criada com sucesso.")
      
      await Database.connection.execute("""
        CREATE TABLE "group" (
            id SERIAL PRIMARY KEY,
              subject_id INTEGER NOT NULL,
              FOREIGN KEY (subject_id) REFERENCES subjects(id) ON DELETE CASCADE
          );
      """)
      
      Logger.database("Tabela 'group' criada com sucesso.")
      print("tentando criar registration")
      await Database.connection.execute("""
        CREATE TABLE IF NOT EXISTS registration (
          id SERIAL PRIMARY KEY,                                
          user_id VARCHAR(255) NOT NULL,
          group_id INTEGER NOT NULL,
          CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES "group"(id) ON DELETE CASCADE
        );
      """)
      print("registratiobn criada")
      Logger.database("Tabela 'registrations' criada com sucesso.")
    except Exception as e:
      Logger.database("Erro ao criar tabelas", error=True)
      print(e)
      
  @staticmethod
  async def sync():
    if Database.connection:
      try:
        MAIN_TABLE = 'subjects'
        
        table_exists = await Database.connection.fetchval(
          f"SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = '{MAIN_TABLE}')"
        )
        
        if not table_exists:
          print("criando tabelas")
          await Database.create_tables()
      except Exception as e:
        Logger.database(f"Erro ao sincronizar tabelas", error=True)
    else:
      Logger.database("Falha ao conectar no banco!", error=True)
      

  @staticmethod
  async def execute_query(query: str, *args):
    if Database.connection:
      try:
        # Passa a query e os argumentos para o método do asyncpg
        print("teste")
        return await Database.connection.fetch(query, *args)
      except Exception as e:
        Logger.database(f"Erro ao executar query: {e}", error=True)
        # print(e)

      