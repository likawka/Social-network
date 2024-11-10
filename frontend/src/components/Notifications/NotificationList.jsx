import React from 'react';
import PropTypes from 'prop-types';
import NotificationItem from './NotificationItem';

const NotificationList = ({ notifications }) => {
    return (
        <div className="max-w-sm w-full bg-white shadow-lg rounded-lg overflow-hidden z-50">
            {notifications.length === 0 ? (
                <div className="p-4 text-gray-500">No notifications</div>
            ) : (
                notifications.map((notification, index) => (
                    <NotificationItem
                        key={index}
                        message={notification.message}
                        time={notification.time}
                        isRead={notification.isRead}
                    />
                ))
            )}
        </div>
    );
};

NotificationList.propTypes = {
    notifications: PropTypes.arrayOf(
        PropTypes.shape({
            message: PropTypes.string.isRequired,
            time: PropTypes.string.isRequired,
            isRead: PropTypes.bool.isRequired,
        })
    ).isRequired,
};

export default NotificationList;
