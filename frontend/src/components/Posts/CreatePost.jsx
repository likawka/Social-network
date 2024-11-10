import React from "react";
import PostForm from "./PostForm";

const CreatePost = () => {
    const handlePostSubmit = (postData) => {
        console.log(postData);
    }

    return (
        <div className="p-4">
            <h1 className="text-2xl font-bold mb-6">Create a New Post</h1>
            <PostForm onSubmit={handlePostSubmit} />
        </div>
    )
};

export default CreatePost;