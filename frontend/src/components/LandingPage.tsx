import MainLayout from "./MainLayout";
import TopNavBar from "./TopNavBar";

import "bulma/css/bulma.min.css";

function LandingPage() {
    return (
        <div className="main-layout-height">
            <TopNavBar />
            <MainLayout />
        </div>
    );
}

export default LandingPage;
