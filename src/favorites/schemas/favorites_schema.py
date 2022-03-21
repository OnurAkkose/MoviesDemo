def favorite_serializer(favorite) -> dict:
    return {
        "id": str(favorite["_id"]),
        "userId": favorite["userId"],
        "movieId": favorite["movieId"],
    }

def favorites_serializer(favorites) -> list:
    return [favorite_serializer(favorite) for favorite in favorites]