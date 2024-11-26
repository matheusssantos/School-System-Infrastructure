from fastapi import APIRouter
from app.dtos import CreateSubjetctDto
from app.services import SubjectService

SUBJECT_ROUTER = APIRouter(prefix="/subjects")

@SUBJECT_ROUTER.get("")
async def getAllSubjects():
  return await SubjectService.get_all()

@SUBJECT_ROUTER.post("/create")
async def create(body: CreateSubjetctDto):
  return await SubjectService.create(body)

@SUBJECT_ROUTER.get("/{code}/students")
async def getStudentsBySubject(code: str):
  return await SubjectService.get_students(code)
