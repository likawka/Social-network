import React from 'react';
import { useNavigate } from 'react-router-dom';
import Login from '../Auth/Login';
import Button from '../Common/Button';

const HomeLoginPrompt = () => {
    const navigate = useNavigate();

    return (
        <div className="flex flex-col items-center border rounded">
            <h1 className="text-2xl mb-4">Welcome to Anime Network</h1>
            <Login />
            <p className="mt-4">Don't have an account? <Button onClick={() => navigate('/register')}>Register</Button></p>
        </div>
    );
};

export default HomeLoginPrompt;
