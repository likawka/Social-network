import React from "react";
import PropTypes from "prop-types";
import PostItem from '../Posts/PostItem';

const HomePosts = ({ posts }) => {
    return (
        <div>
            {posts.length > 0 ? (
                posts.map(post =>
                    <PostItem
                    key={post.id}
                    post={post}
                    privacy={post.privacy}
                     />)
            ) : (
                <p>No posts available. Follow some users to see their posts here</p>
            )}
        </div>
    );
};

HomePosts.propTypes = {
    posts: PropTypes.arrayOf(PropTypes.shape({
        id: PropTypes.number.isRequired,
        title: PropTypes.string.isRequired,
        content: PropTypes.string.isRequired,
        privacy: PropTypes.string.isRequired,
        timestamp: PropTypes.string.isRequired,
        comments: PropTypes.arrayOf(
          PropTypes.shape({
            id: PropTypes.number.isRequired,
            user: PropTypes.string.isRequired,
            content: PropTypes.string.isRequired,
          })
        ).isRequired,
    })).isRequired,
};

export default HomePosts;