import React, { useState, useEffect } from "react";

const UserInfo = () => {
    const [user, setUser] = useState({
        username: '',
        firstName: '',
        lastName: '',
        avatar: '',
    });

    useEffect(() => {
        // Replace this with actual API call
        const fetchUserData = async () => {
            const mockUserData = {
                username: 'User123',
                firstName: 'John',
                lastName: 'Doe',
                avatar: 'https://via.placeholder.com/150', // Placeholder image URL
            };
            setUser(mockUserData);
        };

        fetchUserData();
    }, []);

    return (
        <div className="flex items-center space-x-4">
            <img
                src={user.avatar}
                alt="User Avatar"
                className="h-12 w-12 rounded-full"
            />
            <div>
                <p className="text-lg font-semibold">{user.username}</p>
                <p className="text-sm">{user.firstName} {user.lastName}</p>
                <p className="text-sm text-gray-400">Online</p>
            </div>
        </div>
    );
};

export default UserInfo;
