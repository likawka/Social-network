import React from 'react';
import PropTypes from 'prop-types';
import PostItem from '../Posts/PostItem';

const ProfileActivityFeed = ({ posts = [] }) => {
  return (
    <div className="mt-8">
      <h3 className="text-xl font-bold mb-4">My Posts</h3>
      {posts.length > 0 ? (
        posts.map((post, index) => (
          <PostItem key={index} post={post} />
        ))
      ) : (
        <p>No posts yet.</p>
      )}
    </div>
  );
};

ProfileActivityFeed.propTypes = {
  posts: PropTypes.arrayOf(PropTypes.object),
};

export default ProfileActivityFeed;
