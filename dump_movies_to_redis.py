import ast
import pandas as pd
import redis


r = redis.Redis(host='localhost', port=6379, db=0)


def dump_movies_to_redis(movies):
    for movie in movies:
        del movie['Unnamed: 0']
        key = f"movie:{movie['id']}"

        movie['genres'] = ast.literal_eval(movie['genres'])
        movie['original_language'] = ast.literal_eval(movie['original_language'])

        r.json().set(key, '$', movie)
        print(f"Inserted {key}")


if __name__ == '__main__':
    movies_df = pd.read_csv('movie_data.csv')
    movies = movies_df.to_dict(orient='records')
    dump_movies_to_redis(movies)
