from app.functions import create_error_response, create_success_response
from app.dtos import CreateSubjetctDto
from app.database import Database
from app.logger import Logger

class SubjectService:
  
  async def get_all():
    try:
      response = await Database.execute_query("SELECT * FROM subjects")
      return create_success_response(response)
    except Exception:
      return create_error_response("Erro interno no servidor")
    
  async def create(data: CreateSubjetctDto):
    try:
       # 1. Insere a disciplina na tabela subjects
        print("Criando subject!")
        insert_subject_query = """
          INSERT INTO subjects (code, name, turn)
          VALUES ($1, $2, $3)
          RETURNING id;
        """
        subject_id = await Database.execute_query(insert_subject_query, data.code, data.name, data.turn)
        print(subject_id)

        print("Criando group!")
        # 2. Insere a turma (group) na tabela group associando à disciplina criada
        insert_group_query = """
          INSERT INTO "group" (subject_id)
          VALUES ($1)
          RETURNING id;
        """
        group_id = await Database.execute_query(insert_group_query, subject_id.id)
            
        # 3. Se tudo deu certo, retorna sucesso
        return create_success_response({
          "subject": data,  # Retorna os dados da disciplina criada
          "group_id": group_id.id  # Retorna o ID do grupo criado
        })
    except Exception as e:
      print(e)
      return create_error_response("Erro interno no servidor")
    
    