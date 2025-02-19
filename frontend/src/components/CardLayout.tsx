import "bulma/css/bulma.min.css";
import Card from "./Card";

type Movie = {
    id: string;
    poster: string;
    title: string;
    release_year: string;
    tagline: string;
    original_language: [];
};

type CardLayoutProps = {
    movies: Movie[];
};

function CardLayout({ movies }: CardLayoutProps) {
    return (
        <div className="columns is-multiline has-background-light p-2">
            {movies && movies.length > 0 ? (
                movies.map((movie) => (
                    <Card
                        key={movie.id}
                        id={movie.id}
                        poster={movie.poster}
                        title={movie.title}
                        release_year={movie.release_year}
                        tagline={movie.tagline}
                        original_language={movie.original_language}
                    />
                ))
            ) : (
                <div className="column is-12 has-text-centered">
                    <p className="title is-5 has-text-grey">
                        No movies available
                    </p>
                </div>
            )}
        </div>
    );
}

export default CardLayout;
