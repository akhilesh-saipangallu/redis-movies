import 'bulma/css/bulma.min.css';
import './SideBar.css'

type SidebarProps = {
    onCategorySelect: (category: string) => void;
};

function SideBar({ onCategorySelect }: SidebarProps) {
    return (
        <aside className="menu">
            <ul className="menu-list">
                <li><a className="has-background-light redis-font-color" onClick={() => onCategorySelect("Popular")}>Popular</a></li>
                <li><a className="has-background-light redis-font-color" onClick={() => onCategorySelect("Recommended")}>Recommended</a></li>
            </ul>
            <p className="menu-label">Genres</p>
            <ul className="menu-list">
                <li><a className="has-background-light redis-font-color" onClick={() => onCategorySelect("Action")}>Action</a></li>
                <li><a className="has-background-light redis-font-color" onClick={() => onCategorySelect("Comedy")}>Comedy</a></li>
                <li><a className="has-background-light redis-font-color" onClick={() => onCategorySelect("Romance")}>Romance</a></li>
                <li><a className="has-background-light redis-font-color" onClick={() => onCategorySelect("Drama")}>Drama</a></li>
            </ul>
            <p className="menu-label">Languages</p>
            <ul className="menu-list">
                <li><a className="has-background-light redis-font-color" onClick={() => onCategorySelect("English")}>English</a></li>
                <li><a className="has-background-light redis-font-color" onClick={() => onCategorySelect("French")}>French</a></li>
                <li><a className="has-background-light redis-font-color" onClick={() => onCategorySelect("Korean")}>Korean</a></li>
                <li><a className="has-background-light redis-font-color" onClick={() => onCategorySelect("Hindi")}>Hindi</a></li>
            </ul>
        </aside>
    )
}

export default SideBar;
