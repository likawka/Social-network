import React, { useEffect, useState } from "react";
import ProfileHeader from "../components/Profile/ProfileHeader";
import ProfileDetails from "../components/Profile/ProfileDetails";
import ProfileActivityFeed from "../components/Profile/ProfileActivityFeed";
import ProfileFollowers from "../components/Profile/ProfileFollowers";
import ProfileSettings from "../components/Profile/ProfileSettings";
import Modal from "../components/Common/Modal";
import Button from "../components/Common/Button";
import api from "../services/api";

const Profile = () => {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [posts, setPosts] = useState([]);
  const [userDetails, setUserDetails] = useState({
    firstName: "John",
    lastName: "Doe",
    nickname: "johndoe",
    aboutMe: "Lover of tech and coffee.",
    avatar: "/path/to/avatar.jpg",
    bannerColor: "#2c3e50",
  });

  const [followers] = useState([
    { username: "Alice", avatar: "/path/to/avatar1.jpg" },
    { username: "Bob", avatar: "/path/to/avatar2.jpg" },
  ]);

  const [following] = useState([
    { username: "Charlie", avatar: "/path/to/avatar3.jpg" },
    { username: "David", avatar: "/path/to/avatar4.jpg" },
  ]);

  const [isModalOpen, setModalOpen] = useState(false);
  const [modalType, setModalType] = useState(null);
  const isCurrentUser = true; // Change this based on whether it's the user's profile or someone else's

  useEffect(() => {
    const fetchProfile = async () => {
      const userId = localStorage.getItem('userId');
      if (!userId) {
        setError("User ID not found.");
        setLoading(false);
        return;
      }
  
      try {
        const profileResponse = await api.getProfile();
        console.log('API Response for profile:', profileResponse);
  
        if (profileResponse && profileResponse.status === 'success' && profileResponse.payload) {
          // Map posts and initialize comments as empty for now
          const fetchedPosts = profileResponse.payload.personalPosts.map(post => ({
            id: post.id,
            title: post.title,
            content: post.content,
            privacy: post.privacy,
            timestamp: post.createdAt,
            comments: [], // Initialize comments as empty
          }));
  
          // Fetch comments for each post
          const postsWithComments = await Promise.all(
            fetchedPosts.map(async (post) => {
              const commentsResponse = await api.getPostsByPostId(post.id);
              const comments = commentsResponse?.payload?.comments || [];
              
              // Ensure comments are mapped with proper structure
              const formattedComments = comments.map(comment => ({
                id: comment.id,
                user: comment.user,  // Adjust user structure if needed
                content: comment.content,
                reactionInfo: comment.reactionInfo, // likes, dislikes, etc.
              }));

              console.log('Formatted comments:', formattedComments);
  
              // Return the post with comments
              return { ...post, comments: formattedComments };
            })
          );
  
          setPosts(postsWithComments);
        } else {
          setError("Failed to load profile data.");
        }
      } catch (err) {
        console.error('API Error:', err);
        setError("Failed to load profile data.");
      } finally {
        setLoading(false);
      }
    };
  
    fetchProfile();
  }, []);
  

  const openFollowersModal = () => {
    setModalType("followers");
    setModalOpen(true);
  };

  const openFollowingModal = () => {
    setModalType("following");
    setModalOpen(true);
  };

  const openEditProfileModal = () => {
    setModalType("editProfile");
    setModalOpen(true);
  };

  const handleSaveSettings = (newSettings) => {
    setUserDetails((prevDetails) => ({
      ...prevDetails,
      avatar: newSettings.newAvatar,
      aboutMe: newSettings.newBio,
      bannerColor: newSettings.newBannerColor,
    }));
    setModalOpen(false); // Close the modal after saving settings
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  return (
    <div>
      <ProfileHeader
        avatar={userDetails.avatar}
        username={userDetails.nickname}
        bannerColor={userDetails.bannerColor}
        followersCount={followers.length}
        followingCount={following.length}
        onFollowersClick={openFollowersModal}
        onFollowingClick={openFollowingModal}
        isCurrentUser={isCurrentUser}
        onFollowClick={() => alert("Follow button clicked!")}
      />
      <div className="container mx-auto mt-6 flex">
        <div className="w-1/3">
          <ProfileDetails {...userDetails} />
          {isCurrentUser && (
            <div className="mt-6">
              <Button
                onClick={openEditProfileModal}
                className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
              >
                Edit Profile
              </Button>
            </div>
          )}
        </div>
        <div className="w-2/3 ml-6">
          {posts.length === 0 ? (
            <div>No posts!</div>
          ) : (
            <ProfileActivityFeed posts={posts} />
          )}
        </div>
      </div>

      {/* Modal */}
      <Modal isOpen={isModalOpen} onClose={() => setModalOpen(false)}>
        {modalType === "followers" && (
          <ProfileFollowers
            followers={followers}
            following={[]}
            type="followers"
          />
        )}
        {modalType === "following" && (
          <ProfileFollowers
            followers={[]}
            following={following}
            type="following"
          />
        )}
        {modalType === "editProfile" && (
          <ProfileSettings
            currentAvatar={userDetails.avatar}
            currentBio={userDetails.aboutMe}
            currentBannerColor={userDetails.bannerColor}
            onSaveSettings={handleSaveSettings}
          />
        )}
      </Modal>
    </div>
  );
};

export default Profile;
