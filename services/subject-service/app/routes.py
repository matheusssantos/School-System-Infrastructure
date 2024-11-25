from fastapi import APIRouter
from app.functions import createSuccessResponse
from app.dtos import CreateSubjetctDto
from app.services import SubjectService

SUBJECT_ROUTER = APIRouter(prefix="/subjects")

@SUBJECT_ROUTER.get("")
def getAllSubjects():
  return SubjectService.getAll()

@SUBJECT_ROUTER.post("/create")
def create(body: CreateSubjetctDto):
  return SubjectService.create(body)

@SUBJECT_ROUTER.get("/{code}/students")
def getStudentsBySubject(code: str):
  return SubjectService.getStudents(code)
