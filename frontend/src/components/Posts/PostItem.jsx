import React, { useState } from 'react';
import PropTypes from 'prop-types';
import api from '../../services/api'; // Assuming api is imported

const PostItem = ({ post, privacy }) => {
  const [likes, setLikes] = useState(post.reactionInfo?.likes || 0);
  const [dislikes, setDislikes] = useState(post.reactionInfo?.dislikes || 0);
  const [visibleComments, setVisibleComments] = useState(0);
  const [showComments, setShowComments] = useState(false);

  const handleReaction = async (objectId, objectType, reactionType) => {
    try {
      const reactionData = {
        objectId,
        objectType,
        reactionType,
      };

      const response = await api.createReaction(reactionData);
      if (response && response.payload) {
        const { likeCount, dislikeCount } = response.payload.reactionInfo;

        if (objectType === 'post') {
          setLikes(likeCount);
          setDislikes(dislikeCount);
        }
      }
    } catch (error) {
      console.error('Error updating reaction:', error);
    }
  };

  const handleLike = () => {
    handleReaction(post.id, 'post', 'like');
  };

  const handleDislike = () => {
    handleReaction(post.id, 'post', 'dislike');
  };

  const handleCommentLike = (commentId) => {
    handleReaction(commentId, 'comment', 'like');
  };

  const handleCommentDislike = (commentId) => {
    handleReaction(commentId, 'comment', 'dislike');
  };

  const toggleComments = () => {
    setShowComments(!showComments);
    if (!showComments) {
      setVisibleComments(3); // Initially show 3 comments
    }
  };

  const showMoreComments = () => {
    setVisibleComments((prev) => prev + 3); // Show 3 more comments
  };

  const comments = Array.isArray(post.comments) ? post.comments : [];

  return (
    <div className="bg-white shadow-md rounded-md p-4 mb-4">
      <h2 className="text-xl font-bold mb-2">{post.title}</h2>
      <p className="mb-4">{post.content}</p>

      <div className="flex gap-4 mb-4">
        <button onClick={handleLike} className="text-green-500">👍 {likes}</button>
        <button onClick={handleDislike} className="text-red-500">👎 {dislikes}</button>
      </div>

      <p className="text-gray-500 text-sm mb-4">Posted on: {post.timestamp}</p>
      <p className="text-gray-500 text-sm mb-4">Privacy: {privacy}</p>

      <button
        onClick={toggleComments}
        className="text-blue-500 hover:underline"
      >
        {showComments ? 'Hide comments' : 'Show comments'}
      </button>

      {showComments && (
        <div>
          <h3 className="text-lg font-bold mb-2">Comments</h3>
          {comments.slice(0, visibleComments).map((comment) => (
            <div key={comment.id} className="mb-2">
              <strong>{comment.user?.nickname || 'Anonymous'}: </strong>{comment.content}
              <div className="flex gap-2 mt-1">
                <button onClick={() => handleCommentLike(comment.id)} className="text-blue-500">
                  👍 {comment.reactionInfo?.likeCount || 0}
                </button>
                <button onClick={() => handleCommentDislike(comment.id)} className="text-red-500">
                  👎 {comment.reactionInfo?.dislikeCount || 0}
                </button>
              </div>
            </div>
          ))}
          {comments.length > visibleComments && (
            <button
              onClick={showMoreComments}
              className="text-blue-500 hover:underline"
            >
              Show more comments
            </button>
          )}
        </div>
      )}
    </div>
  );
};

PostItem.propTypes = {
  post: PropTypes.shape({
    id: PropTypes.number.isRequired, // Add id for API calls
    title: PropTypes.string.isRequired,
    content: PropTypes.string.isRequired,
    timestamp: PropTypes.string.isRequired,
    comments: PropTypes.array,
    reactionInfo: PropTypes.shape({
      likes: PropTypes.number,
      dislikes: PropTypes.number,
    }),
  }).isRequired,
  privacy: PropTypes.string.isRequired,
};

export default PostItem;
