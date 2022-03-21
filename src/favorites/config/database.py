from pymongo import MongoClient

client = MongoClient("***")

db = client.favorites

collection_name = db["favorites"]