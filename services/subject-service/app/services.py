from app.functions import create_error_response, create_success_response
from app.dtos import CreateSubjetctDto
from app.database import Database

class SubjectService:
  
  async def get_all():
    try:
      response = await Database.execute_query("SELECT * FROM subjects")
      return create_success_response(response)
    except Exception:
      return create_error_response("Erro interno no servidor")
    
  async def create(data: CreateSubjetctDto):
    try:
      # Salva na banco
      
      # Valdia se criou
      
      return create_success_response(data)
    except Exception:
      return create_error_response("Erro interno no servidor")
    
  async def get_students(subjectCode: int):
    try:
      # Busca a materia no banco
      
      # Busca todas as matriculas de tal materia
      
      return create_success_response(subjectCode)
    except Exception:
      return create_error_response("Erro interno no servidor")

    