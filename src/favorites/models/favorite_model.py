from pydantic import BaseModel

class Favorite(BaseModel):
    movieId: str
    userId: str