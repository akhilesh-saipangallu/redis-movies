import "bulma/css/bulma.min.css";

type MovieProps = {
    id: string;
    poster: string;
    title: string;
    release_year: string;
    tagline: string;
    original_language: string[];
    popularity?: number | null;
};

function Card({
    id,
    poster,
    title,
    release_year,
    tagline,
    original_language,
    popularity,
}: MovieProps) {
    const maxTaglineLength = 70;
    const normalizedPopularity = popularity
        ? (popularity / 10).toFixed(1)
        : "N/A";

    return (
        <div key={id} className="column is-3">
            <div className="card">
                <div className="card-image">
                    <figure className="image is-3by4">
                        <img src={poster} alt={`${title} Poster`} />
                    </figure>
                </div>
                <div className="card-content">
                    <p className="mb-2">
                        <strong>★</strong> {normalizedPopularity}
                    </p>
                    <p className="title is-5 mb-1">{title}</p>
                    <p className="subtitle is-6  is-italic mb-3">
                        {release_year}{" "}
                        {original_language.length > 0 &&
                            `/ ${original_language.join(", ")}`}
                    </p>
                    <p className="content">
                        {tagline.length > maxTaglineLength
                            ? tagline.slice(0, maxTaglineLength) + "..."
                            : tagline}
                    </p>
                </div>
            </div>
        </div>
    );
}

export default Card;
