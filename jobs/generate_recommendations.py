import json
import time

import redis
from redis.commands.search.query import Query
import numpy as np
from sentence_transformers import SentenceTransformer


# Redis connection
redis_client = redis.Redis(host='localhost', port=6379, decode_responses=True)

# Load embedding model
# embedding_model = SentenceTransformer("all-MiniLM-L6-v2")
embedding_model = SentenceTransformer('msmarco-distilbert-base-v4')


def get_movie_embedding(movie_id: str):
    movie_embedding = redis_client.json().get(f'movie:{movie_id}', 'embedding')
    return np.array(movie_embedding) if movie_embedding else None


def track_user_search(user_id, movie_id):
    timestamp = time.time()
    redis_client.zadd(f"user:search_history:{user_id}", {movie_id: timestamp})
    # Keep only the latest 20 searches
    redis_client.zremrangebyrank(f"user:search_history:{user_id}", 0, -21)


def compute_user_profile_embedding(user_id):
    # Function to compute user profile embedding with weighted averaging
    # Get latest 20 searches
    movie_ids = redis_client.zrevrange(f'user:search_history:{user_id}', 0, 19)

    embeddings = []
    for mid in movie_ids:
        movie_embedding = get_movie_embedding(mid)
        if movie_embedding.size > 0:
            embeddings.append(movie_embedding)

    if not embeddings:
        return None

    # Give higher weight to recent searches
    weights = np.linspace(1.0, 2.0, num=len(embeddings))
    # Normalize weights
    weights = weights / np.sum(weights)
    
    weighted_embeddings = np.average(embeddings, axis=0, weights=weights)
    user_embedding = weighted_embeddings.tolist()

    redis_client.set(f'user:profile_embedding:{user_id}', json.dumps(user_embedding))
    return user_embedding


def find_similar_movies(user_embedding, top_n=10):
    # Function to find similar movies using Redis Vector Search
    print('user_embedding: ', user_embedding)
    user_embedding = np.array(user_embedding, dtype=np.float32).tobytes()

    INDEX_NAME = 'idx:movies'
    query = (
        Query(f'*=>[KNN {top_n} @embedding $vector AS score]')
        .sort_by('score')
        .return_fields('id', 'title', 'score')
        # .paging(0, 2)
        .dialect(2)
    )

    query_params = {'vector': user_embedding}
    return redis_client.ft(INDEX_NAME).search(query, query_params)


def store_recommendations(user_id, recommended_movies):
    # Function to store recommendations
    recommended_movies_norm = []
    for recommended_movie in recommended_movies.docs:
        recommended_movies_norm.append({
            'id': recommended_movie['id'].split(':')[1],
            'title': recommended_movie['title'],
            'score': recommended_movie['score'],
        })

    redis_client.json().set(f'user:recommendations:{user_id}', '$', recommended_movies_norm)
    print(f'Stored recommendations for user {user_id}')

# Main function to generate recommendations
def generate_recommendations(user_id, movie_ids):
    for movie_id in movie_ids:
        track_user_search(user_id, movie_id)
    
    user_embedding = compute_user_profile_embedding(user_id)
    if user_embedding:
        recommended_movies = find_similar_movies(user_embedding)
        print('recommended_movies:', recommended_movies)
        store_recommendations(user_id, recommended_movies)
    else:
        print('No valid embeddings found for given movies.')


if __name__ == '__main__':
    generate_recommendations('ef511c64-73aa-486b-ab60-939411b55adf', [149])
