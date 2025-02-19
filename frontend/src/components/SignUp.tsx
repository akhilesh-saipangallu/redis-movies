import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

function SignUp() {
    const [formData, setFormData] = useState({
        firstName: "",
        lastName: "",
        email: "",
        password: "",
    });
    const [error, setError] = useState("");
    const navigate = useNavigate();

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setError("");

        try {
            await axios.post(`http://localhost:8080/signup`, {
                first_name: formData.firstName,
                last_name: formData.lastName,
                email: formData.email,
                password: formData.password,
            });
            navigate("/signin");
        } catch (err) {
            if (axios.isAxiosError(err)) {
                setError(err.response?.data?.error || "An error occurred");
            } else {
                setError("An unexpected error occurred");
            }
        }
    };

    return (
        <div className="container mt-6" style={{ maxWidth: "400px" }}>
            <div className="box">
                <h2 className="title is-4">Sign Up</h2>
                {error && <p className="has-text-danger">{error}</p>}
                <form onSubmit={handleSubmit}>
                    <div className="field">
                        <label className="label">First Name</label>
                        <div className="control">
                            <input
                                className="input"
                                type="text"
                                name="firstName"
                                value={formData.firstName}
                                onChange={handleChange}
                                required
                            />
                        </div>
                    </div>
                    <div className="field">
                        <label className="label">Last Name</label>
                        <div className="control">
                            <input
                                className="input"
                                type="text"
                                name="lastName"
                                value={formData.lastName}
                                onChange={handleChange}
                                required
                            />
                        </div>
                    </div>
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
                            Sign Up
                        </button>
                    </div>
                </form>
                <p className="has-text-centered">
                    Already have an account?{" "}
                    <a href="" onClick={() => navigate("/signin")}>
                        Sign In
                    </a>
                </p>
            </div>
        </div>
    );
}

export { SignUp };
