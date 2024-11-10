import React from "react";
import PropTypes from "prop-types";

const TextInput = ({ label, name, value, onChange, placeholder, type = "text", className }) => {
    return (
        <div className={`mb-4 ${className}`}>
            {label && <label className="block text-sm font-medium mb-2" htmlFor={name}>{label}</label>}
            <input
                type={type}
                name={name}
                value={value}
                onChange={onChange}
                placeholder={placeholder}
                className="w-full px-3 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
        </div>
    );
};

TextInput.propTypes = {
    label: PropTypes.string,
    name: PropTypes.string.isRequired,
    value: PropTypes.string.isRequired,
    onChange: PropTypes.func.isRequired,
    placeholder: PropTypes.string,
    type: PropTypes.string,
    className: PropTypes.string,
};

export default TextInput;