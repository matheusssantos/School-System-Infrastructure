from fastapi import APIRouter
from app.dtos import CreateSubjetctDto
from app.services import SubjectService
from app.logger import Logger

SUBJECT_ROUTER = APIRouter(prefix="/subjects")

@SUBJECT_ROUTER.get("")
async def getAllSubjects():
  return await SubjectService.get_all()

@SUBJECT_ROUTER.post("/create")
async def create(body: CreateSubjetctDto):  
  print(body)
  return await SubjectService.create(body)

