import { useState, useEffect, useRef } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import axiosInstance from "./axiosInstance";
import "bulma/css/bulma.min.css";
import SideBar from "./SideBar";
import CardLayout from "./CardLayout";
import TopNavBar from "./TopNavBar";

import "./MainLayout.css";

function MainLayout() {
    const [selectedCategory, setSelectedCategory] = useState("Home");
    const [movies, setMovies] = useState<
        {
            id: string;
            poster: string;
            title: string;
            release_year: string;
            tagline: string;
            original_language: [];
            popularity: number;
        }[]
    >([]);
    const [offset, setOffset] = useState(0);
    const limit = 10;
    let hasMore: boolean = true,
        loading: boolean = false;
    const [hasMoreButton, setHasMoreButton] = useState(false);
    const navigate = useNavigate();
    const location = useLocation();
    const searchParams = new URLSearchParams(location.search);
    let searchText = searchParams.get("search_text") || "";
    const isFirstRender = useRef(true);

    // Remove query params on refresh
    useEffect(() => {
        if (location.search) {
            navigate(location.pathname, { replace: true });
            searchText = "";
        }
    }, []); // Runs only once on mount

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

    // Fetch movies when category or search text changes
    useEffect(() => {
        if (isFirstRender.current) {
            isFirstRender.current = false;
            return; // Skip the first call
        }
        setMovies([]);
        setOffset(0);
        hasMore = true;
        setHasMoreButton(true);
        loading = false;
        setTimeout(() => fetchMovies(0), 0);
    }, [selectedCategory, searchText]);

    const fetchMovies = async (newOffset: number) => {
        if (loading || !hasMore) return;

        loading = true;
        const filter = categoryFilters[selectedCategory] || "";
        const isPaginated = !["/popular", "/recommendations"].includes(filter);

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
                    hasMore = false;
                    setHasMoreButton(false);
                }
            } else {
                setMovies(newMovies); // Fetch all at once for Popular & Recommended
                hasMore = false;
                setHasMoreButton(false);
            }
        } catch (error: any) {
            if (error.response?.status === 401) {
                localStorage.removeItem("authToken");
                navigate("/signin");
            }
        } finally {
            loading = false;
        }
    };

    return (
        <div id="main-layout" className="section m-0 p-0 main-layout-height">
            <TopNavBar selectedCategory={selectedCategory} />
            <div className="columns">
                <div className="column is-2 has-background-light ml-3">
                    <SideBar
                        onCategorySelect={setSelectedCategory}
                        selectedCategory={selectedCategory}
                    />
                </div>
                <div className="column is-10 has-background-light m-3">
                    <CardLayout movies={movies} />

                    {/* Show Load More Button only for paginated categories */}
                    {hasMore &&
                        hasMoreButton &&
                        !loading &&
                        !["Popular", "Recommended"].includes(
                            selectedCategory
                        ) && (
                            <div className="has-text-centered mt-4">
                                <button
                                    className="button is-primary is-rounded"
                                    onClick={() => fetchMovies(offset)}
                                >
                                    Load More
                                </button>
                            </div>
                        )}

                    {/* Loading Indicator */}
                    {loading && <p className="has-text-centered">Loading...</p>}

                    {/* No More Movies Message */}
                    {!hasMore && movies && movies.length > 0 && (
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
