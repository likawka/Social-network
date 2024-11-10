import React from 'react';
import PropTypes from 'prop-types';

const Select = ({ label, name, value, onChange, options }) => {
    return (
        <div className="mb-4">
            {label && <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor={name}>{label}</label>}
            <select
                id={name}
                name={name}
                value={value}
                onChange={onChange}
                className="w-full px-3 py-2 border rounded"
            >
                {options.map((option) => (
                    <option key={option.value} value={option.value}>
                        {option.label}
                    </option>
                ))}
            </select>
        </div>
    );
};

Select.propTypes = {
    label: PropTypes.string,
    name: PropTypes.string.isRequired,
    value: PropTypes.string.isRequired,
    onChange: PropTypes.func.isRequired,
    options: PropTypes.arrayOf(
        PropTypes.shape({
            value: PropTypes.string.isRequired,
            label: PropTypes.string.isRequired
        })
    ).isRequired
};

export default Select;
