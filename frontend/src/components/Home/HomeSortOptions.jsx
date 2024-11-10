import React from 'react';
import PropTypes from 'prop-types';

const HomeSortOptions = ({ sortOption, handleSortChange }) => {
  return (
    <div className="flex justify-end">
      <label htmlFor="sort" className="mr-2 text-gray-700">
        Sort by:
      </label>
      <select
        id="sort"
        value={sortOption}
        onChange={handleSortChange}
        className="border rounded p-2"
      >
        <option value="mostRecent">Most Recent</option>
        <option value="mostLikes">Most Likes</option>
        <option value="mostComments">Most Comments</option>
      </select>
    </div>
  );
};

HomeSortOptions.propTypes = {
  sortOption: PropTypes.string.isRequired,
  handleSortChange: PropTypes.func.isRequired,
};

export default HomeSortOptions;
