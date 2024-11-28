from fastapi import FastAPI
from app.routes import SUBJECT_ROUTER
from app.database import Database
from contextlib import asynccontextmanager
from fastapi.middleware.cors import CORSMiddleware

print(">> Powered by FastAPI")

@asynccontextmanager
async def lifespan(app: FastAPI):
  await Database.connect()
  await Database.sync()
  
  yield
  
  await Database.disconnect()

app = FastAPI(lifespan=lifespan)

app.add_middleware(
  CORSMiddleware,
  allow_origins=["http://localhost:4200"], 
  allow_credentials=True,
  allow_methods=["*"],
  allow_headers=["*"],
)

app.include_router(SUBJECT_ROUTER)
