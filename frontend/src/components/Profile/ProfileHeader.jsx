import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import api from "../../services/api";


const ProfileHeader = () => {

  const [profile, setProfile] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchProfile = async () => {
        const userId = localStorage.getItem('userId'); // Retrieve userId from localStorage
        if (!userId) {
            setError("User ID not found.");
            setLoading(false);
            return;
        }

        try {
            const response = await api.getProfile();
            console.log('API Response:', response); // Log the response for debugging

            if (response.status === 'success' && response.payload) {
                const profileData = response.payload.user; // Extract user data from payload
                setProfile(profileData);
                console.log('Profile data:', profileData);
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
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  const { avatar,
    nickname,
    bannerColor,
    followerCount,
    followingCount,
    onFollowersClick,
    onFollowingClick,
    isCurrentUser,
    onFollowClick, } = profile || {};

    console.log('Profile data:', profile);

  return (
    <div className="relative" style={{ backgroundColor: "#2c3e50" }}>
      <div className="h-32"></div>
      <div className="absolute top-1/2 transform -translate-y-1/2 left-4 flex items-center">
        <img
          src={avatar}
          alt={`${nickname}'s avatar`}
          className="w-20 h-20 rounded-full border-4 border-white"
        />
        <div className="ml-4 text-white">
          <h1 className="text-2xl font-bold">{nickname}</h1>
        </div>
      </div>
      <div className="absolute top-1/2 transform -translate-y-1/2 right-4 text-white text-right">
        <div>
          <span className="cursor-pointer mr-4" onClick={onFollowersClick}>
            Followers: {followerCount}
          </span>
          <span className="cursor-pointer" onClick={onFollowingClick}>
            Following: {followingCount}
          </span>
        </div>
        {!isCurrentUser && (
          <button className="mt-2 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" onClick={onFollowClick}>
            Follow
          </button>
        )}
      </div>
    </div>
  );
};

ProfileHeader.propTypes = {
  avatar: PropTypes.string.isRequired,
  username: PropTypes.string.isRequired,
  bannerColor: PropTypes.string.isRequired,
  followersCount: PropTypes.number.isRequired,
  followingCount: PropTypes.number.isRequired,
  onFollowersClick: PropTypes.func.isRequired,
  onFollowingClick: PropTypes.func.isRequired,
  isCurrentUser: PropTypes.bool.isRequired,
  onFollowClick: PropTypes.func,
};

export default ProfileHeader;
