def create_error_response(message: str):
  return { "success": False, "message": message }
  
def create_success_response(message):
  return { "success": True, "message": message }