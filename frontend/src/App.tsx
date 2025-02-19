import "bulma/css/bulma.min.css";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
// import TopNavBar from "./components/TopNavBar";
// import MainLayout from "./components/MainLayout";
import { SignIn } from "./components/SignIn";
import { SignUp } from "./components/SignUp";

function App() {
    return (
        // <div className='main-layout-height'>
        //   <TopNavBar />
        //   <MainLayout/>
        // </div>

        <Router>
            <Routes>
                <Route path="/signin" element={<SignIn />} />
                <Route path="/signup" element={<SignUp />} />
            </Routes>
        </Router>
    );
}

export default App;
