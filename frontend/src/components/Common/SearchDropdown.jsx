import React from "react";
import PropTypes from "prop-types";

const SearchDropdown = ({ results }) => {
    // If there are no results, don't render the component
    if (!results || results.length === 0) {
        return null;
    }

    return (
        <div className="absolute mt-1 w-full bg-white shadow-lg rounded-lg z-10">
            <ul className="divide-y divide-gray-200">
                {results.map((result, index) => (
                    <li key={index} className="p-2 hover:bg-gray-100 cursor-pointer">
                        {result}
                    </li>
                ))}
            </ul>
        </div>
    );
};

SearchDropdown.propTypes = {
    results: PropTypes.arrayOf(PropTypes.string).isRequired,
};

export default SearchDropdown;