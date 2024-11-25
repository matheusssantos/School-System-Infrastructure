def createErrorResponse(message: str):
  return { "success": False, "message": message }
  
def createSuccessResponse(message):
  return { "success": True, "message": message }