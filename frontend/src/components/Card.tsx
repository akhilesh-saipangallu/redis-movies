import 'bulma/css/bulma.min.css';

type MovieProps = {
    id: string;
    poster: string;
    title: string;
    release_year: string;
    tagline: string;
};

function Card({ id, poster, title, release_year, tagline }: MovieProps) {
    return (
        <div key={id} className="column is-3">
            <div className="card">
                <div className="card-image">
                    <figure className="image is-3by4">
                        <img src={poster} alt="Placeholder image" />
                    </figure>
                </div>
                <div className="card-content">
                    <div className="media">
                        <div className="media-content">
                            <p className="title is-4">{title} | {release_year}</p>
                        </div>
                    </div>
                    <div className="content">{tagline}<br /></div>
                </div>
            </div>
        </div>
    )
}

export default Card;
