import React from 'react';
import { BrowserRouter as Router } from 'react-router-dom';
import SideBar from './components/Layout/SideBar.jsx';
import NavBar from './components/Layout/NavBar';
import Footer from './components/Layout/Footer';
import AppRoutes from './services/routes.jsx';

function App() {
    // Authenticate user
    const isAuthenticated = localStorage.getItem('authToken'); // For example


    return (
        <Router>
            <div className='flex flex-col h-screen'>
                <NavBar />
                <div className='flex flex-grow'>
                    <SideBar isAuthenticated={isAuthenticated} />
                    <div className='flex-grow flex flex-col'>
                        <AppRoutes />
                    </div>
                </div>
                <Footer />
            </div>
        </Router>
    );
}

export default App;
