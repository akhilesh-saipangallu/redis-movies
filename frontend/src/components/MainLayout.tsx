import { useState, useEffect } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import axiosInstance from "./axiosInstance";
import "bulma/css/bulma.min.css";
import SideBar from "./SideBar";
import CardLayout from "./CardLayout";

function MainLayout() {
    const [selectedCategory, setSelectedCategory] = useState<string>("");
    const [movies, setMovies] = useState<
        {
            id: string;
            poster: string;
            title: string;
            release_year: string;
            tagline: string;
            original_language: [];
        }[]
    >([]);
    const [offset, setOffset] = useState(0);
    const limit = 10;
    const [loading, setLoading] = useState(false);
    const [hasMore, setHasMore] = useState(true);
    const navigate = useNavigate();
    const location = useLocation();
    const searchParams = new URLSearchParams(location.search);
    const searchText = searchParams.get("search_text") || "";

    const categoryFilters: Record<string, string> = {
        Home: "",
        Popular: "/popular",
        Recommended: "/recommendations",
        Action: "?genre=action",
        Comedy: "?genre=comedy",
        Romance: "?genre=romance",
        Drama: "?genre=drama",
        English: "?original_language=English",
        Italian: "?original_language=Italian",
        Latin: "?original_language=Latin",
        German: "?original_language=German",
        Hindi: "?original_language=Hindi",
    };

    useEffect(() => {
        if (selectedCategory || searchText) {
            setMovies([]);
            setOffset(0);
            setHasMore(true);
            setTimeout(() => fetchMovies(selectedCategory, searchText, 0), 0);
        }
    }, [selectedCategory, searchText]);

    useEffect(() => {
        const handleScroll = () => {
            if (
                window.innerHeight + window.scrollY >=
                    document.body.offsetHeight - 100 &&
                !loading &&
                hasMore
            ) {
                fetchMovies(selectedCategory, searchText, offset);
            }
        };

        window.addEventListener("scroll", handleScroll);
        return () => window.removeEventListener("scroll", handleScroll);
    }, [offset, loading, hasMore, selectedCategory, searchText]);

    const fetchMovies = async (
        category: string,
        searchText: string,
        newOffset: number
    ) => {
        if (loading || !hasMore) return;

        setLoading(true);
        const filter = categoryFilters[category] || "";
        const isPaginated = !["/popular", "/recommendations"].includes(
            categoryFilters[category]
        );

        let url = isPaginated
            ? `/movies${
                  filter ? `${filter}&` : "?"
              }limit=${limit}&offset=${newOffset}`
            : `/movies${filter}`;

        if (searchText && isPaginated) {
            url += `&search_text=${encodeURIComponent(searchText)}`;
        }

        try {
            const response = await axiosInstance.get(url);
            const newMovies = response.data;

            if (isPaginated) {
                setMovies((prevMovies) => [...prevMovies, ...newMovies]);
                setOffset(newOffset + limit);
                if (newMovies.length < limit) {
                    setHasMore(false);
                }
            } else {
                setMovies(newMovies);
                setHasMore(false);
            }
        } catch (error: any) {
            if (error.response?.status === 401) {
                localStorage.removeItem("authToken");
                navigate("/signin");
            }
        } finally {
            setLoading(false);
        }
    };

    return (
        <div id="main-layout" className="section m-0 p-0 main-layout-height">
            <div className="columns">
                <div className="column is-2 has-background-light ml-3">
                    <SideBar onCategorySelect={setSelectedCategory} />
                </div>
                <div className="column is-10 has-background-light m-3">
                    <CardLayout movies={movies} />
                    {loading && (
                        <p className="has-text-centered">
                            Loading more movies...
                        </p>
                    )}
                    {!hasMore && (
                        <p className="has-text-centered">
                            No more movies to show.
                        </p>
                    )}
                </div>
            </div>
        </div>
    );
}

export default MainLayout;
