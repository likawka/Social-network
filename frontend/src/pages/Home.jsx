import React, { useState, useEffect } from 'react';
import HomePosts from '../components/Home/HomePosts';
import HomeSortOptions from '../components/Home/HomeSortOptions';
import HomeLoginPrompt from '../components/Home/HomeLoginPrompt';

const Home = () => {
    const [isLoggedIn, setIsLoggedIn] = useState(false); // Simulating a logged-in state
    const [posts, setPosts] = useState([]);
    const [sortOption, setSortOption] = useState('mostRecent');

    useEffect(() => {
        if (isLoggedIn) {
            loadPosts();
        }
    }, [sortOption]);

    const loadPosts = async () => {
      const dummyPosts = [
          {
              id: 1,
              title: 'My Favorite Anime',
              content: 'Let me tell you about my favorite anime...',
              timestamp: '2024-08-22',
              likes: 10,
              privacy: 'public', // Add privacy field
              comments: [
                  { id: 1, user: 'User1', content: 'I love this anime too!' },
                  { id: 2, user: 'User2', content: 'Not a big fan, but it’s okay.' },
              ],
          },
          {
              id: 2,
              title: 'Anime Recommendation',
              content: 'If you haven’t watched this anime, you should!',
              timestamp: '2024-08-20',
              likes: 20,
              privacy: 'followers', // Add privacy field
              comments: [
                  { id: 1, user: 'User3', content: 'Thanks for the recommendation!' },
              ],
          },
      ];
  
      const sortedPosts = sortOption === 'mostRecent'
          ? dummyPosts.sort((a, b) => new Date(b.timestamp) - new Date(a.timestamp))
          : dummyPosts.sort((a, b) => b.likes - a.likes);
  
      setPosts(sortedPosts);
  };
  

    return (
        <div className="container mx-auto p-4">
            {isLoggedIn ? (
                <>
                    <HomeSortOptions sortOption={sortOption} handleSortChange={(e) => setSortOption(e.target.value)} />
                    <HomePosts posts={posts} />
                </>
            ) : (
                <HomeLoginPrompt />
            )}
        </div>
    );
};

export default Home;
