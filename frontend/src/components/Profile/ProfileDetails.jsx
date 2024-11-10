import React, { useEffect, useState } from 'react';
import api from "../../services/api";


const ProfileDetails = () => {
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
                if (response.status === 'success' && response.payload) {
                    const profileData = response.payload.user; // Extract user data from payload
                    setProfile(profileData);
                    console.log('Profile:', profileData);
                } else {
                    setError("Failed to load profile data.");
                }
            } catch (err) {
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

    // Destructure the profile data for easy access
    const { firstName, lastName, nickname, aboutMe, avatar } = profile || {};

    return (
        <div className="flex items-center space-x-4">
            <img src={avatar} alt={`${firstName} ${lastName}`} className="w-24 h-24 rounded-full" />
            <div>
                <h2 className="text-2xl font-bold">{`${firstName} ${lastName}`}</h2>
                {nickname && <p className="text-sm text-grey-800">{nickname}</p>}
                <p className="mt-2">{aboutMe}</p>
            </div>
        </div>
    );
};

export default ProfileDetails;
