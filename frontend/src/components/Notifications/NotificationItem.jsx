import React from "react";
import PropTypes from "prop-types";

const NotificationItem = ({ message, time, isRead }) => {
    return (
        <div className={`p-4 border-b border-gray-200 &{isRead ? 'bg-gray-100' : 'bg-white'}`}>
            <p className="text-sm">{message}</p>
            <p className="text-xs text-gray-500">{time}</p>
        </div>
    );
};

NotificationItem.propTypes = {
    message: PropTypes.string.isRequired,
    time: PropTypes.string.isRequired,
    isRead: PropTypes.bool.isRequired,
};

export default NotificationItem;