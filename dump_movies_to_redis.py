import ast
import json
import os

import numpy as np
import pandas as pd
import redis
from sentence_transformers import SentenceTransformer


redis_host = os.getenv('REDIS_HOST', 'localhost')
redis_port = int(os.getenv('REDIS_PORT', 6379))


redis_client = redis.Redis(host=redis_host, port=redis_port, db=0)
# embedding_model = SentenceTransformer('all-MiniLM-L6-v2')
embedding_model = SentenceTransformer('msmarco-distilbert-base-v4')


def generate_movie_embedding(title, description, genre):
    text = f'{title} {description} {genre}'
    embedding = embedding_model.encode(text).astype(np.float32).tolist()

    return embedding


def dump_movies_to_redis(movies):
    for movie in movies:
        del movie['Unnamed: 0']
        key = f"movie:{movie['id']}"

        movie['id'] = str(movie['id'])
        movie['genres'] = ast.literal_eval(movie['genres'])
        movie['original_language'] = ast.literal_eval(movie['original_language'])
        movie['embedding'] = generate_movie_embedding(movie['original_title'], movie['overview'], str(movie['genres']))

        redis_client.json().set(key, '$', movie)
        print(f'Inserted {key}')



if __name__ == '__main__':
    movies_df = pd.read_csv('movie_data.csv')
    movies = movies_df.to_dict(orient='records')
    dump_movies_to_redis(movies)
