from fastapi import FastAPI
from app.routes import SUBJECT_ROUTER 

app = FastAPI()

app.include_router(SUBJECT_ROUTER)
