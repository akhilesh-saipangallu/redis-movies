import "./SideBar.css";

interface SideBarProps {
    onCategorySelect: (category: string) => void;
    selectedCategory: string;
}

function SideBar({ onCategorySelect, selectedCategory }: SideBarProps) {
    const categories = [
        "Home",
        "Popular",
        "Recommended",
        "Action",
        "Comedy",
        "Romance",
        "Drama",
        "English",
        "Italian",
        "Latin",
        "German",
        "Hindi",
    ];

    return (
        <aside className="menu has-background-light">
            <ul className="menu-list">
                {categories.map((category) => {
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
