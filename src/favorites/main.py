from fastapi import FastAPI
from routes.favorites_route import favorites_api_router
from fastapi.middleware.cors import CORSMiddleware

origins = [ 
    "http://localhost:3000",
]
app = FastAPI()
app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)
app.include_router(favorites_api_router)