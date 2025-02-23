import "./SideBar.css";

interface SideBarProps {
    onCategorySelect: (category: string) => void;
    selectedCategory: string;
}

function SideBar({ onCategorySelect, selectedCategory }: SideBarProps) {
    const generalCategories = ["Home", "Popular", "Recommended"];
    const genres = ["Action", "Comedy", "Romance", "Drama"];
    const languages = ["English", "Italian", "Latin", "German", "Hindi"];

    return (
        <aside className="menu has-background-light">
            <ul className="menu-list">
                {generalCategories.map((category) => {
                    const isActive = selectedCategory === category;
                    return (
                        <li key={category}>
                            <a
                                className={`has-background-light ${
                                    isActive ? "active" : ""
                                }`}
                                onClick={() => onCategorySelect(category)}
                            >
                                {category}
                            </a>
                        </li>
                    );
                })}
            </ul>

            <p className="menu-label">Genres</p>
            <ul className="menu-list">
                {genres.map((category) => {
                    const isActive = selectedCategory === category;
                    return (
                        <li key={category}>
                            <a
                                className={`has-background-light ${
                                    isActive ? "active" : ""
                                }`}
                                onClick={() => onCategorySelect(category)}
                            >
                                {category}
                            </a>
                        </li>
                    );
                })}
            </ul>

            <p className="menu-label">Languages</p>
            <ul className="menu-list">
                {languages.map((category) => {
                    const isActive = selectedCategory === category;
                    return (
                        <li key={category}>
                            <a
                                className={`has-background-light ${
                                    isActive ? "active" : ""
                                }`}
                                onClick={() => onCategorySelect(category)}
                            >
                                {category}
                            </a>
                        </li>
                    );
                })}
            </ul>
        </aside>
    );
}

export default SideBar;
