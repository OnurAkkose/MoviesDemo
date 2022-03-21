from fastapi import APIRouter

from models.favorite_model import Favorite
from config.database import collection_name

from schemas.favorites_schema import favorites_serializer, favorite_serializer
from bson import ObjectId

favorites_api_router = APIRouter()


@favorites_api_router.get("/{userId}")
async def get_favorites(userId: str):
    return favorites_serializer(collection_name.find({"userId": userId}))


# post
@favorites_api_router.post("/")
async def insert_favorite(favorite: Favorite):
    _id = collection_name.insert_one(dict(favorite))
    return favorite


# delete
@favorites_api_router.delete("/{userId}/{movieId}")
async def remove_favorite(userId: str, movieId: str):
    collection_name.find_one_and_delete({"userId": userId, "movieId": movieId})
    return {"status": "ok"}

