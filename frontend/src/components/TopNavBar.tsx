import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import "bulma/css/bulma.min.css";
import "./TopNavBar.css";

interface TopNavBarProps {
    selectedCategory: string;
}

function TopNavBar({ selectedCategory }: TopNavBarProps) {
    const navigate = useNavigate();
    const [searchText, setSearchText] = useState("");

    useEffect(() => {
        const token = localStorage.getItem("authToken");
        if (!token) {
            navigate("/signin");
        }
    }, [navigate]);

    const handleLogout = () => {
        localStorage.removeItem("authToken");
        navigate("/signin");
    };

    const handleSearch = (e: React.FormEvent) => {
        e.preventDefault();
        navigate(`/?search_text=${encodeURIComponent(searchText)}`);
    };

    // Hide search bar if "Popular" or "Recommendations" is selected
    const hideSearch =
        selectedCategory === "Popular" || selectedCategory === "Recommended";

    return (
        <nav className="navbar" role="navigation" aria-label="main navigation">
            <div className="navbar-brand">
                <a className="navbar-item" href="/">
                    Redis Movies
                </a>
            </div>
            <div id="navbarBasicExample" className="navbar-menu">
                <div className="navbar-center">
                    {!hideSearch && (
                        <form id="search_form" onSubmit={handleSearch}>
                            <div className="navbar-item field has-addons">
                                <p className="control is-expanded">
                                    <input
                                        className="input"
                                        type="text"
                                        placeholder="I'm looking for..."
                                        value={searchText}
                                        onChange={(e) =>
                                            setSearchText(e.target.value)
                                        }
                                    />
                                </p>
                                <p className="control">
                                    <button className="button" type="submit">
                                        Search
                                    </button>
                                </p>
                            </div>
                        </form>
                    )}
                </div>
            </div>
            <div className="navbar-end">
                <div className="navbar-item">
                    <div className="buttons">
                        <button
                            className="button is-primary"
                            onClick={handleLogout}
                        >
                            <strong>Logout</strong>
                        </button>
                    </div>
                </div>
            </div>
        </nav>
    );
}

export default TopNavBar;
