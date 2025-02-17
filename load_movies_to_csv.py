from datetime import datetime
import json
import pandas as pd


def date_to_epoch(date_str):
    dt = datetime.strptime(date_str, "%d %b %Y")
    return int(dt.timestamp())


def extract_rotten_tomato_rating(ratings):
    for rating in ratings:
        if rating['Source'] == 'Rotten Tomatoes':
            return float(rating['Value'][:-1])
    return 0

def get_revenue(movie):
    revenue = movie.get('BoxOffice', '$ 0')
    if revenue == 'N/A':
        return '$ 0'
    return revenue


def normalize_movie(movies):
    results = []
    for i, movie in enumerate(movies):
        try:
            popularity = extract_rotten_tomato_rating(movie['Ratings'])
            release_date = date_to_epoch(movie['Released'])
        except Exception as e:
            print(e)
            continue

        results.append({
            'id': i,
            'budget': 40218400,
            'genres': movie['genres'],
            'original_language': movie['Language'].split(', '),
            'original_title': movie['Title'],
            'overview': movie['description'],
            'popularity': popularity,
            'poster': movie['Poster'],
            'release_year': movie['release_year'],
            'release_date': release_date,
            'release_date_str': movie['Released'],
            'revenue': get_revenue(movie),
            'runtime': movie['Runtime'],
            'tagline': movie['Plot'],
            'title': movie['Title'],
            'vote_average': movie['imdbRating'],
            'vote_count': movie['imdbVotes']
        })

    return results


if __name__ == "__main__":
    with open('movie_data.json', "r", encoding="utf-8") as file:
        movies = json.load(file)
        normalized_movies = normalize_movie(movies)
        df = pd.DataFrame(normalized_movies)
        df_unique = df.drop_duplicates(subset='title')
        df_unique.to_csv('movie_data.csv')
