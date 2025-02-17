import 'bulma/css/bulma.min.css';
import './TopNavBar.css';

function TopNavBar() {
    return (
        <nav className="navbar" role="navigation" aria-label="main navigation">
            <div className="navbar-brand">
                <a className="navbar-item" href="">Redis Movies</a>
            </div>
            <div id="navbarBasicExample" className="navbar-menu">
                <div className="navbar-center">
                    <form id="search_form">
                        <div className="navbar-item field has-addons">
                            <p className="control is-expanded">
                                <input className="input" name="query" type="text" placeholder="I'm looking for..." required />
                            </p>
                            <p className="control">
                                <button className="button" type="submit">
                                    Search
                                </button>
                            </p>
                        </div>
                    </form>
                </div>
            </div>
            <div className="navbar-end">
                <div className="navbar-item">
                    <div className="buttons">
                        <a className="button is-primary">
                            <strong>Logout</strong>
                        </a>
                    </div>
                </div>
            </div>
        </nav>
    )
}

export default TopNavBar;

{/* <nav className="navbar" role="navigation" aria-label="main navigation">
  <div className="navbar-brand">
    <!-- navbar items, navbar burger... -->
  </div>
</nav> */}
