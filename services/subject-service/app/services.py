from app.functions import createErrorResponse, createSuccessResponse
from app.dtos import CreateSubjetctDto

class SubjectService:
  
  async def getAll():
    try:
      return createSuccessResponse("Lista de materias")
    except Exception:
      return createErrorResponse("Erro interno no servidor")
    
  async def create(data: CreateSubjetctDto):
    try:
      # Salva na banco
      
      # Valdia se criou
      
      return createSuccessResponse(data)
    except Exception:
      return createErrorResponse("Erro interno no servidor")
    
  async def getStudents(subjectCode: int):
    try:
      # Busca a materia no banco
      
      # Busca todas as matriculas de tal materia
      
      return createSuccessResponse(subjectCode)
    except Exception:
      return createErrorResponse("Erro interno no servidor")

    