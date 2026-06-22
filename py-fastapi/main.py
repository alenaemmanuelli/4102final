from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import Optional
import uuid # generates unique ids

app = FastAPI() # app object

# blueprint for each movie
class Movie(BaseModel):
    title: str
    director: str
    year: int
    genre: str
    rating: Optional[float] = None # optional

# mock database using a dict
movies_db = {}

@app.get("/movies") # when GET comes in, call this function
def get_movies():
    return list(movies_db.values()) # converts to JSON

@app.get("/movies/{movie_id}") # get movie
def get_movie(movie_id: str):
    if movie_id not in movies_db: # check if ID exists
        raise HTTPException(status_code=404, detail="movie not found")
    return movies_db[movie_id]

@app.post("/movies", status_code=201) # create movie
def create_movie(movie: Movie):
    movie_id = str(uuid.uuid4()) # creates unique ID
    movies_db[movie_id] = {"id": movie_id, **movie.dict()} # attach ID to movie
    return movies_db[movie_id] # adds to database

@app.put("/movies/{movie_id}") # replace movie
def update_movie(movie_id: str, movie: Movie):
    if movie_id not in movies_db:
        raise HTTPException(status_code=404, detail="movie not found")
    movies_db[movie_id] = {"id": movie_id, **movie.dict()} 
    return movies_db[movie_id]

@app.delete("/movies/{movie_id}", status_code=204) # delete movie
def delete_movie(movie_id: str):
    if movie_id not in movies_db:
        raise HTTPException(status_code=404, detail="movie not found")
    del movies_db[movie_id]