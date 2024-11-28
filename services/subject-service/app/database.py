import asyncpg
from app.logger import Logger
import os
from dotenv import load_dotenv

load_dotenv()

DB_USER = os.getenv("DB_USER")
DB_PASSWORD = os.getenv("DB_PASSWORD")
DB_HOST = os.getenv("DB_HOST")
DB_PORT = int(os.getenv("DB_PORT", 5432))
DB_NAME = os.getenv("DB_NAME")

class Database:
  
  connection = None
  
  @staticmethod
  async def connect():
    if Database.connection is None:
      try:
        # connection_url = "postgresql://postgres:postgres@postgres:5432/school"
        # Conectando com a URL de conexão
        # Database.connection = await asyncpg.connect(connection_url)

        Database.connection = await asyncpg.connect(
          user=DB_USER,
          password=DB_PASSWORD,
          host=DB_HOST,
          port=DB_PORT,
          database=DB_NAME
        )
        Logger.database("Conexão bem-sucedida com o banco de dados!")
      except Exception as e:
        Logger.database("Falha ao conectar no banco!", error=True)
        
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
      
      await Database.connection.execute("""
        CREATE TABLE IF NOT EXISTS registration (
          id SERIAL PRIMARY KEY,                                
          user_id VARCHAR(255) NOT NULL,
          group_id INTEGER NOT NULL,
          CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES "group"(id) ON DELETE CASCADE
        );
      """)
      Logger.database("Tabela 'registrations' criada com sucesso.")
    except Exception as e:
      Logger.database("Erro ao criar tabelas", error=True)
      
  @staticmethod
  async def sync():
    if Database.connection:
      try:
        MAIN_TABLE = 'subjects'
        
        table_exists = await Database.connection.fetchval(
          f"SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = '{MAIN_TABLE}')"
        )
        
        if not table_exists:
          await Database.create_tables()
          
      except Exception as e:
        Logger.database(f"Erro ao sincronizar tabelas", error=True)
    else:
      Logger.database("Falha ao conectar no banco!", error=True)
      

  @staticmethod
  async def execute_query(query: str, *args):
    if Database.connection:
      try:
        return await Database.connection.fetch(query, *args)
      except Exception as e:
        Logger.database(f"Erro ao executar query: {e}", error=True)

      