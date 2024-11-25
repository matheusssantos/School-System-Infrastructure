from pydantic import BaseModel

class CreateSubjetctDto(BaseModel):
  code: str
  name: str
  turn: str
  classCode: int