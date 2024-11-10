import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Home from '../pages/Home';
import Register from '../components/Auth/Register';
import Login from '../components/Auth/Login';
import Profile from '../pages/Profile';
import CreatePost from '../components/Posts/CreatePost';
import Group from '../pages/Group';
import Chat from '../components/Chat/Chat';
import PostPage from '../components/Posts/PostPage';

const AppRoutes = () => {
    return (
        <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/register" element={<Register />} />
            <Route path="/login" element={<Login />} />
            <Route path="/profile" element={<Profile />} />
            <Route path='/create-post' element={<CreatePost />} />
            <Route path='/chat' element={<Chat />} />
            <Route path='createpost' element={<CreatePost />} />
            <Route path="/posts/:postId" element={<PostPage />} />  {/* Post-specific route */}
            <Route path='/groups/:groupId' element={<Group />} /> {/* Dynamic Group Route */}
        </Routes>
    );
};

export default AppRoutes;
