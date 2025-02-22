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
            <div className="columns is-flex-grow-1">
                <div className="column is-2">
                    <div className="navbar-brand">
                        <a className="navbar-item" href="/">
                            Redis Movies
                        </a>
                    </div>
                </div>
                <div className="column is-9">
                    <div id="navbarBasicExample" className="navbar-menu">
                        <div className="navbar-center">
                            {!hideSearch && (
                                <form id="search_form" onSubmit={handleSearch}>
                                    <div className="navbar-item field has-addons">
                                        <p className="control is-expanded">
                                            <input
                                                className="input is-rounded"
                                                type="text"
                                                placeholder="I'm looking for..."
                                                value={searchText}
                                                onChange={(e) =>
                                                    setSearchText(
                                                        e.target.value
                                                    )
                                                }
                                            />
                                        </p>
                                        <p className="control">
                                            <button
                                                className="button is-rounded"
                                                type="submit"
                                            >
                                                Search
                                            </button>
                                        </p>
                                    </div>
                                </form>
                            )}
                        </div>
                    </div>
                </div>
                <div className="column is-1">
                    <div className="navbar-end">
                        <div className="navbar-item">
                            <div className="buttons">
                                <button
                                    className="button is-primary is-rounded"
                                    onClick={handleLogout}
                                >
                                    <strong>Logout</strong>
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </nav>
    );
}

export default TopNavBar;
