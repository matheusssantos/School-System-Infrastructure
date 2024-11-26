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
          turn CHAR(1) NOT NULL,
          class_code INTEGER NOT NULL
        );
      """)
      Logger.database("Tabela 'subjects' criada com sucesso.")
      
      await Database.connection.execute("""
        CREATE TABLE students (
          id SERIAL PRIMARY KEY,
          uuid VARCHAR NOT NULL
        );
      """)
      Logger.database("Tabela 'students' criada com sucesso.")
      
      await Database.connection.execute("""
        CREATE TABLE registrations (
          id SERIAL PRIMARY KEY,
          id_student INTEGER NOT NULL,
          id_subject INTEGER NOT NULL,
          FOREIGN KEY (id_student) REFERENCES students(id),
          FOREIGN KEY (id_subject) REFERENCES subjects(id)
        );
      """)
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
          await Database.create_tables()
      except Exception as e:
        Logger.database(f"Erro ao sincronizar tabelas", error=True)
    else:
      Logger.database("Falha ao conectar no banco!", error=True)
      
  @staticmethod
  async def execute_query(query: str):
    if Database.connection:
      try:
        return await Database.connection.fetch(query)
      except Exception as e:
        Logger.database(f"Erro ao executar query", error=True)
      