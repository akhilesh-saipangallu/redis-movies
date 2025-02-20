from datetime import datetime, timedelta
import json
import os
import random
import redis

# Connect to Redis
redis_host = os.getenv('REDIS_HOST', 'localhost')
redis_port = int(os.getenv('REDIS_PORT', 6379))

r = redis.Redis(host=redis_host, port=redis_port, db=0)

def random_date(start, end):
    return start + timedelta(days=random.randint(0, (end - start).days))

def generate_movie_data(id):
    genres_list = ['Action', 'Comedy', 'Drama', 'Horror', 'Sci-Fi', 'Romance']
    languages = ['en', 'fr', 'es', 'de', 'it', 'jp']
    spoken_languages = ['English', 'French', 'Spanish', 'German', 'Italian', 'Japanese']
    release_date = random_date(datetime(1980, 1, 1), datetime(2025, 12, 31))

    return {
        "id": id,
        "budget": random.randint(100000, 200000000),
        "genres": random.sample(genres_list, k=random.randint(1, 3)),
        "original_language": random.choice(languages),
        "original_title": f"Original Title {id}",
        "overview": f"This is an overview of movie {id}.",
        "popularity": round(random.uniform(1.0, 100.0), 2),
        "poster": f"https://example.com/poster{id}.jpg",
        "release_date": int(release_date.timestamp()),
        "revenue": random.randint(100000, 1000000000),
        "runtime": random.randint(80, 180),
        "spoken_languages": random.sample(spoken_languages, k=random.randint(1, 2)),
        "tagline": f"Catchphrase for movie {id}",
        "title": f"Movie Title {id}",
        "vote_average": round(random.uniform(1.0, 10.0), 1),
        "vote_count": random.randint(100, 10000)
    }

def insert_movies(n):
    for i in range(1, n + 1):
        movie_data = generate_movie_data(i)
        key = f"movie:{i}"
        # r.json().set(key, '$', json.dumps(movie_data))
        r.json().set(key, '$', movie_data)
        print(f"Inserted {key}")

if __name__ == "__main__":
    n = int(input("Enter the number of movies to insert: "))
    insert_movies(n)
