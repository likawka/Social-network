import React from 'react';
import PropTypes from 'prop-types';

const Button = ({ children, onClick, className, type = 'button' }) => {
    return (
        <button
            type={type}
            className={`px-4 py-2 font-semibold text-white bg-blue-500 rounded hover:bg-blue-700 focus:outline-none focus:shadow-outline ${className}`}
            onClick={onClick}
        >
            {children}
        </button>
    );
};

Button.propTypes = {
    children: PropTypes.node.isRequired,
    onClick: PropTypes.func,
    className: PropTypes.string,
    type: PropTypes.oneOf(['button', 'submit', 'reset']),
};

export default Button;
