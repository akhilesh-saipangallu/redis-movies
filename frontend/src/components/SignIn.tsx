import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

function SignIn() {
    const [formData, setFormData] = useState({ email: "", password: "" });
    const [error, setError] = useState("");
    const navigate = useNavigate();

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setError("");

        try {
            const response = await axios.post(
                `http://localhost:8080/signin`,
                formData
            );
            console.log("Success:", response.data);
            navigate("/");
        } catch (err) {
            if (axios.isAxiosError(err)) {
                setError(err.response?.data?.message || "An error occurred");
            } else {
                setError("An unexpected error occurred");
            }
        }
    };

    return (
        <div className="container mt-6" style={{ maxWidth: "400px" }}>
            <div className="box">
                <h2 className="title is-4">Sign In</h2>
                {error && <p className="has-text-danger">{error}</p>}
                <form onSubmit={handleSubmit}>
                    <div className="field">
                        <label className="label">Email</label>
                        <div className="control">
                            <input
                                className="input"
                                type="email"
                                name="email"
                                value={formData.email}
                                onChange={handleChange}
                                required
                            />
                        </div>
                    </div>
                    <div className="field">
                        <label className="label">Password</label>
                        <div className="control">
                            <input
                                className="input"
                                type="password"
                                name="password"
                                value={formData.password}
                                onChange={handleChange}
                                required
                            />
                        </div>
                    </div>
                    <div className="field">
                        <button
                            className="button is-primary is-fullwidth"
                            type="submit"
                        >
                            Sign In
                        </button>
                    </div>
                </form>
                <p className="has-text-centered">
                    Don't have an account?{" "}
                    <a href="" onClick={() => navigate("/signup")}>
                        Sign Up
                    </a>
                </p>
            </div>
        </div>
    );
}

export { SignIn };
