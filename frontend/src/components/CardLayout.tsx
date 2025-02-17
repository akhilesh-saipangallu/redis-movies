import 'bulma/css/bulma.min.css';
import Card from './Card';

type Movie = {
    id: string;
    poster: string;
    title: string;
    release_year: string;
    tagline: string;
};

type CardLayoutProps = {
    movies: Movie[];
};

function CardLayout({ movies }: CardLayoutProps) {
    // const movies = ["m1", "m2", "m3", "m4", "m5"];
    return (
        <div className="columns is-multiline has-background-grey-lighter m-2 p-2">
                {movies.map((movie, index) => (
                    <Card
                        id={movie.id}
                        poster={movie.poster}
                        title={movie.title}
                        release_year={movie.release_year}
                        tagline={movie.tagline}
                    />
                ))}
            </div>
    )
}

export default CardLayout;
