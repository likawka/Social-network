import React, { useState } from 'react';
import PropTypes from 'prop-types';
import PostItem from '../Posts/PostItem';
import Button from '../Common/Button';

const GroupPostFeed = ({ posts }) => {
  const [visiblePosts, setVisiblePosts] = useState(3);

  const showMorePosts = () => {
    setVisiblePosts((prev) => prev + 3); // Load 3 more posts
  };

  return (
    <div className="bg-white p-4 rounded shadow-md mt-4">
      <h2 className="text-xl font-bold mb-4">Group Posts</h2>
      {posts.slice(0, visiblePosts).map((post, index) => (
        <PostItem key={index} post={post} />
      ))}
      {posts.length > visiblePosts && (
        <Button className="mt-4" onClick={showMorePosts}>
          Load more posts
        </Button>
      )}
    </div>
  );
};

GroupPostFeed.propTypes = {
  posts: PropTypes.arrayOf(PropTypes.object).isRequired,
};

export default GroupPostFeed;
