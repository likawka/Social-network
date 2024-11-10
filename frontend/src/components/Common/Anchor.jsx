import React from "react";
import { Link } from "react-router-dom";
import PropTypes from "prop-types";

const Anchor = ({ to, children, className, external}) => {
    if (external) {
        return (
            <a
                href={to}
                target="_blank"
                rel="noopener noreferrer"
                className={`text-blue-500 hover:text-blue-700 underline ${className}`}
            >
                {children}
            </a>
        )
    }
    
    return (
        <Link
            to={to}
            className={`text-blue-500 hover:text-blue-700 underline ${className}`}
        >
            {children}
        </Link>
        );
};

Anchor.propTypes = {
    to: PropTypes.string.isRequired,
    children: PropTypes.node.isRequired,
    className: PropTypes.string,
    external: PropTypes.bool,
};

export default Anchor;