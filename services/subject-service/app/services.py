from app.functions import create_error_response, create_success_response
from app.dtos import CreateSubjetctDto
from app.database import Database

class SubjectService:
  
  @staticmethod
  async def get_all():
    try:
      response = await Database.execute_query("SELECT * FROM subjects")
      return create_success_response(response)
    except Exception:
      return create_error_response("Erro interno no servidor")
    
  @staticmethod
  async def create(data: CreateSubjetctDto):
    try:
      insert_subject_query = """
        INSERT INTO subjects (code, name, turn)
        VALUES ($1, $2, $3)
        RETURNING id, name
      """
      subject = await Database.execute_query(insert_subject_query, data.code, data.name, data.turn)
      if not subject:
        return create_error_response("Erro ao cadastrar matéria")

      insert_group_query = """
        INSERT INTO "group" (subject_id)
        VALUES ($1)
        RETURNING id
      """
      group = await Database.execute_query(insert_group_query, subject[0]['id'])
      if not group:
        return create_error_response("Erro ao cadastrar turma")
          
      return create_success_response(subject[0])
    except Exception as e:
      print(e)
      return create_error_response("Erro interno no servidor")
    