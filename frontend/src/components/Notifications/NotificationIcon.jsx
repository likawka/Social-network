import React from 'react';
import PropTypes from 'prop-types';

const NotificationIcon = ({ count }) => {
    return (
        <div className="relative">
            <svg
                className="w-6 h-6 text-white"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
            >
                <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14V11a6.002 6.002 0 00-5-5.917V5a3 3 0 10-6 0v.083A6.002 6.002 0 002 11v3c0 .217-.02.43-.058.638L1 17h5m7 0v1a3 3 0 01-6 0v-1m7 0H9"
                />
            </svg>
            {count > 0 && (
                <span className="absolute top-0 right-0 inline-flex items-center justify-center px-2 py-1 text-xs font-bold leading-none text-white bg-red-600 rounded-full">
                    {count > 9 ? '9+' : count}
                </span>
            )}
        </div>
    );
};

NotificationIcon.propTypes = {
    count: PropTypes.number.isRequired,
};

export default NotificationIcon;
