import { useState, useEffect } from "react";
import 'bulma/css/bulma.min.css';
import SideBar from './SideBar';
import CardLayout from "./CardLayout";

function MainLayout() {
    const [selectedCategory, setSelectedCategory] = useState<string>("");
    const [movies, setMovies] = useState<{
        id: string;
        poster: string;
        title: string;
        release_year: string;
        tagline: string;
        original_language: [];
    }[]>([]);

    const categoryFilters: Record<string, string> = {
        "Popular": "/popular",
        // TODO
        "Recommended": "sort=recommendation_score",

        "Action": "?genre=action",
        "Comedy": "?genre=comedy",
        "Romance": "?genre=romance",
        "Drama": "?genre=drama",
        
        "English": "?original_language=English",
        "Italian": "?original_language=Italian",
        "Latin": "?original_language=Latin",
        "German": "?original_language=German",
        "Hindi": "?original_language=Hindi",
    };

    useEffect(() => {
    if (selectedCategory) {
        fetchMovies(selectedCategory);
    }
    }, [selectedCategory]);

    const fetchMovies = async (category: string) => {
        setMovies([]);

        const filter = categoryFilters[category] || "";

        try {
            const response = await fetch(`http://127.0.0.1:8080/movies${filter}`);
            const data = await response.json();
            console.log("movies count:", data)
            setMovies(data);
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
                    <CardLayout movies={movies} />
                </div>
          </div>
      </div>
    )
}

export default MainLayout;
