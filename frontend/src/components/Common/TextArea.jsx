import React from 'react';
import PropTypes from 'prop-types';

const TextArea = ({ label, name, value, onChange, placeholder, rows = 4, className }) => {
  return (
    <div className={`mb-4 ${className}`}>
      {label && <label className="block text-sm font-medium mb-2" htmlFor={name}>{label}</label>}
      <textarea
        name={name}
        value={value}
        onChange={onChange}
        placeholder={placeholder}
        rows={rows}
        className="w-full px-3 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
    </div>
  );
};

TextArea.propTypes = {
  label: PropTypes.string,
  name: PropTypes.string.isRequired,
  value: PropTypes.string.isRequired,
  onChange: PropTypes.func.isRequired,
  placeholder: PropTypes.string,
  rows: PropTypes.number,
  className: PropTypes.string,
};

export default TextArea;
