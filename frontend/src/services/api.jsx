import CreatePost from "../components/Posts/CreatePost";

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

const api = {
  // AUTH API CALLS
  login: async (userData) => {
    const response = await fetch(`${API_BASE_URL}/auth/login`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(userData),
    });

    const data = await response.json();

    if (response.ok) {
      localStorage.setItem("userId", data.user.id);
      localStorage.setItem("userNickname", data.user.nickname);
      alert("Login successful!");
      window.location.href = "/"; // Redirect to home page after login
    } else {
      // Throw an error with the message from the server
      throw new Error(
        data.message || "Login failed. Please check your credentials."
      );
    }
  },

  register: async (userData) => {
    const response = await fetch(`${API_BASE_URL}/auth/register`, {
      method: 'POST',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(userData),
    });
  
    const data = await response.json();
  
    if (response.ok) {
      localStorage.setItem("userId", data.user.id);
      localStorage.setItem("userNickname", data.user.nickname);
      return data;
    } else {
      throw {
        details: data.error.details || [], // Default to empty array if details are undefined
      };
    }
  },
  

  // Add more API calls as needed

  logout: async () => {
    try {
      const response = await fetch(`${API_BASE_URL}/auth/logout`, {
        method: "DELETE",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
      });
      if (response.ok) {
        const responseData = await response.json();
        console.log("Logout successful!");
      } else {
        const errorData = await response.json();
        console.error("Logout failed:", errorData);
      }
    } catch (error) {
      console.error("Error during logout:", error);
      alert("Error during logout. Please try again later.");
    }
  },

  // USERS API CALLS
  getUsers: async () => {
    try {
      const response = await fetch(`${API_BASE_URL}/users`, {
        method: "GET",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
      });
      if (response.ok) {
        const responseData = await response.json();
        return responseData;
      } else {
        const errorData = await response.json();
        console.error("Get users failed:", errorData);
      }
    } catch (error) {
      console.error("Error during get users:", error);
      alert("Error during get users. Please try again later.");
    }
  },

  getProfile: async () => {
    const userId = localStorage.getItem("userId");
    try {
      const response = await fetch(`${API_BASE_URL}/users/${userId}`, {
        method: "GET",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
      });
      if (response.ok) {
        const responseData = await response.json();
        return responseData;
      } else {
        const errorData = await response.json();
        console.error("Get profile failed:", errorData);
      }
    } catch (error) {
      console.error("Error during get profile:", error);
      alert("Error during get profile. Please try again later.");
    }
  },

  // POSTS API CALLS

  createPost: async (postData) => {
    
      const response = await fetch(`${API_BASE_URL}/posts`, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(postData),
      });

      const data = await response.json();
  
    if (response.ok) {
      return data;
    } else {
      throw {
        details: data.error.details || [], // Default to empty array if details are undefined
      };
    }
    
  },

  getPost: async () => {
    try {
      const response = await fetch(`${API_BASE_URL}/posts`, {
        method: "GET",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
      });
      if (response.ok) {
        const responseData = await response.json();
        return responseData;
      } else {
        const errorData = await response.json();
        console.error("Get posts failed:", errorData);
      }
    } catch (error) {
      console.error("Error during get posts:", error);
      alert("Error during get posts. Please try again later.");
    }
  },

  getPostsByPostId: async (userId) => {
    try {
      const response = await fetch(`${API_BASE_URL}/posts/${userId}`, {
        method: "GET",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
      });
      if (response.ok) {
        const responseData = await response.json();
        return responseData; // Assuming posts are in `payload.posts`
      } else {
        const errorData = await response.json();
        console.error("Get posts by user ID failed:", errorData);
      }
    } catch (error) {
      console.error("Error during get posts by user ID:", error);
      alert("Error during get posts by user ID. Please try again later.");
    }
  },

  // REACTIONS API CALLS

  createReaction: async (reactionData) => {   
    const response = await fetch(`${API_BASE_URL}/reactions`, {
      method: "PATCH",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(reactionData),
    });

    const data = await response.json();

    if (response.ok) {
      return data;
    } else {
      throw {
        details: data.error.details || [], // Default to empty array if details are undefined
      };
    }
  }
  

  // END, do not remove this line
};

export default api;
