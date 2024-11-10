import React from 'react';
import PropTypes from 'prop-types';

const ProfileFollowers = ({ followers, following, type }) => {
  return (
    <div>
      <h3 className="text-xl font-bold mb-4">{type === 'followers' ? 'Followers' : 'Following'}</h3>
      <ul className="space-y-4">
        {type === 'followers' && followers.map((follower, index) => (
          <li key={index} className="flex items-center space-x-4">
            <img src={follower.avatar} alt={follower.username} className="w-10 h-10 rounded-full" />
            <p>{follower.username}</p>
          </li>
        ))}
        {type === 'following' && following.map((follow, index) => (
          <li key={index} className="flex items-center space-x-4">
            <img src={follow.avatar} alt={follow.username} className="w-10 h-10 rounded-full" />
            <p>{follow.username}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

ProfileFollowers.propTypes = {
  followers: PropTypes.arrayOf(
    PropTypes.shape({
      username: PropTypes.string.isRequired,
      avatar: PropTypes.string.isRequired,
    })
  ).isRequired,
  following: PropTypes.arrayOf(
    PropTypes.shape({
      username: PropTypes.string.isRequired,
      avatar: PropTypes.string.isRequired,
    })
  ).isRequired,
  type: PropTypes.oneOf(['followers', 'following']).isRequired,
};

export default ProfileFollowers;
