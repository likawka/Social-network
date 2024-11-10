import React, { useState } from 'react';
import Button from '../Common/Button';
import TextInput from '../Common/TextInput';
import api from '../../services/api.jsx';  // Adjust the path based on your file structure

import { useNavigate } from 'react-router-dom';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');  // State for error message
  const navigate = useNavigate();

  const handleLogin = async () => {
    const userData = {
      email, 
      password,
    };
  
    console.log('Attempting login with:', userData);
  
    try {
      // Attempt to log in using the API call
      await api.login(userData);
      // If login is successful, navigate to the profile page
      navigate('/profile');
    } catch (error) {
      // If the login fails, extract and display the error message
      console.error('Login failed:', error);
      setError(error.message || 'Login failed. Please check your email and password and try again.');
    }
  };

  return (
    <div className="max-w-sm mx-auto">
      <h2 className="text-2xl font-bold mb-4">Login</h2>

      <TextInput
        label="Email"
        name="email"
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        placeholder="Email"
      />
      <TextInput
        label="Password"
        name="password"
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        placeholder="Password"
      />
      <Button onClick={handleLogin}>Login</Button>
       {/* Display the error message if it exists */}
       {error && <div className="text-red-500 mb-4">{error}</div>}
    </div>
  );
};

export default Login;
