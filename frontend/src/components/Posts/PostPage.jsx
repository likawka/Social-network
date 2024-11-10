import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import PostItem from "./PostItem";
import api from "../../services/api"; // Assuming you have an api module for fetching data

const PostPage = () => {
    const { postId } = useParams();
    const [posts, setPosts] = useState([]);
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchProfile = async () => {
            const userId = localStorage.getItem('userId'); // Retrieve userId from localStorage
            if (!userId) {
                setError("User ID not found.");
                setLoading(false);
                return;
            }

            try {
                const profileResponse = await api.getProfile();
                console.log('API Response for profile:', profileResponse); // Log the response for debugging

                if (profileResponse && profileResponse.status === 'success' && profileResponse.payload) {
                    const fetchedPosts = profileResponse.payload.personalPosts.map(post => ({
                        id: post.id,
                        title: post.title,
                        content: post.content,
                        privacy: post.privacy,
                        timestamp: post.createdAt,
                        comments: [], // Initialize comments as empty
                    }));
                    setPosts(fetchedPosts); // Set the posts in state
                } else {
                    setError("Failed to load profile data.");
                }
            } catch (err) {
                console.error('API Error:', err); // Log the error for debugging
                setError("Failed to load profile data.");
            } finally {
                setLoading(false);
            }
        };

        fetchProfile();
    }, []);

    if (loading) {
        return <h1>Loading...</h1>;
    }

    if (error) {
        return <h1>{error}</h1>;
    }

    const post = posts.find(p => p.id === parseInt(postId));



    if (!post) {
        return <h1>Post not found!</h1>;
    }

    return (
        <div>
            <PostItem post={post} privacy={post.privacy} />
        </div>
    );
};

export default PostPage;
