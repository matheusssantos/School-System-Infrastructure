from fastapi import FastAPI
from app.routes import SUBJECT_ROUTER 
from app.database import Database
from contextlib import asynccontextmanager

print(">> Powered by FastAPI")

app = FastAPI()

@asynccontextmanager
async def lifespan(app: FastAPI):
  await Database.connect()
  await Database.sync()
  
  yield
  
  await Database.disconnect()

app = FastAPI(lifespan=lifespan)

app.include_router(SUBJECT_ROUTER)
