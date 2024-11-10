import React from 'react';
import SidebarItem from './SidebarItem';
import UserInfo from '../Common/UserInfo';
import Logo from '../Common/Logo';
import NotificationDropdown from '../Notifications/NotificationsDropdown';

const SideBar = ({ isAuthenticated }) => {
    // Dummy data for groups
    const groupItems = [
        { label: 'Book Club', link: '/groups/book-club' },
        { label: 'Anime Lovers', link: '/groups/anime-lovers' },
        { label: 'Tech Enthusiasts', link: '/groups/tech-enthusiasts' },
    ];

    return (
        <div className="bg-gray-900 text-white w-64 flex flex-col p-4 h-full">
            {isAuthenticated && (
                <>
                    <div className='text-center mb-8'>
                        <Logo />
                    </div>
                    <div className='text-center mb-8'>
                        <UserInfo />
                    </div>
                </>

            )}
            <div className="flex-grow overflow-y-auto">
                <SidebarItem icon="home" label="Home" link="/" />
                <SidebarItem icon="user" label="Profile" link="/profile" />
                <SidebarItem icon="users" label="Groups" items={groupItems} />
                <SidebarItem icon="comments" label="Chat" link="/chat" />
                <NotificationDropdown />
            </div>
        </div>
    );
};

export default SideBar;
