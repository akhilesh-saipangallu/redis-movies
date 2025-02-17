import { useState, useEffect } from "react";
import 'bulma/css/bulma.min.css';
import SideBar from './SideBar';
import CardLayout from "./CardLayout";

function MainLayout() {
    const [selectedCategory, setSelectedCategory] = useState<string>("");
    const [movies, setMovies] = useState<{ id: string; poster: string; title: string; release_year: string; tagline: string; }[]>([]);

    const categoryFilters: Record<string, string> = {
        "Popular": "/popular",
        // TODO
        "Recommended": "sort=recommendation_score",

        "Action": "?genre=action",
        "Comedy": "?genre=comedy",
        "Romance": "?genre=romance",
        "Drama": "?genre=drama",
    };

    useEffect(() => {
    if (selectedCategory) {
        fetchMovies(selectedCategory);
    }
    }, [selectedCategory]);

    const fetchMovies = async (category: string) => {
        setMovies([]); // Clear current movies while loading new ones

        // ✅ Get the filter for the selected category
        const filter = categoryFilters[category] || "";

        try {
            // ✅ Append the filter dynamically in the API call
            const response = await fetch(`http://127.0.0.1:8080/movies${filter}`);
            const data = await response.json();
            console.log("movies count:", data)
            setMovies(data); // Update movie list
        } catch (error) {
            console.error("Error fetching movies:", error);
        }
    };

    return (
        <div id="main-layout" className="section m-0 p-0 main-layout-height">
            <div className="columns">
                <div className="column is-2 has-background-light ml-3">
                    <SideBar onCategorySelect={setSelectedCategory}/>
                </div>
                <div className="column is-10 has-background-light m-3">
                <div className="columns is-multiline has-background-grey-lighter m-2 p-2">
                    <CardLayout movies={movies} />
                </div>
                </div>
          </div>
      </div>
    )
}

export default MainLayout;
