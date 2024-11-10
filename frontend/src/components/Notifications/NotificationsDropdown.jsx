import React from 'react';
import DropdownMenu from '../Common/DropDownMenu';

const NotificationDropdown = () => {
    const notifications = [
        { label: 'New message from John', link: '#', time: '2m ago', isRead: false },
        { label: 'Alice commented on your post', link: '#', time: '10m ago', isRead: true },
        { label: 'Group event reminder', link: '#', time: '1h ago', isRead: false },
    ];

    return (
        <DropdownMenu
            label="Notifications"
            items={notifications.map(notification => ({
                label: `${notification.label} - ${notification.time}`,
                link: notification.link,
            }))}
            icon="bell"
        />
    );
};

export default NotificationDropdown;
