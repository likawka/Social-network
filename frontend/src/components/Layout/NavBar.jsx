import React, { useState } from "react";
import { useNavigate } from 'react-router-dom';
import Button from "../Common/Button";
import api from '../../services/api.jsx';  // Adjust the path based on your file structure


const NavBar = () => {
    const [searchTerm, setSearchTerm] = useState("");
    const navigate = useNavigate();
    const isAuthenticated = !!localStorage.getItem('authToken'); // Check if the user is authenticated

    const handleSearch = (e) => {
        setSearchTerm(e.target.value);
        // Handle search logic here if needed
    };

    const handleLogout = () => {
        // Clear the auth token (or any other logout logic)
        localStorage.removeItem('authToken');
        navigate('/login');
    };

    const handleLogin = () => {
        navigate('/login');

    };

    return (
        <nav className="bg-gray-800 p-4 flex justify-between items-center">
            <div className="flex-grow relative mx-2">
                <input
                    type="text"
                    placeholder="Search..."
                    value={searchTerm}
                    onChange={handleSearch}
                    className="px-2 py-2 w-full rounded bg-gray-700 text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
            </div>
            <div>
                {isAuthenticated ? (
                    <Button onClick={handleLogout} className="bg-red-600 text-white">
                        Logout
                    </Button>
                ) : (
                    <Button onClick={handleLogin} className="bg-blue-600 text-white">
                        Login
                    </Button>
                )}
            </div>
        </nav>
    );
};

export default NavBar;
