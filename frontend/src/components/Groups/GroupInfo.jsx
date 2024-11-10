import React from "react";
import PropTypes from "prop-types";

const GroupInfo = ({ membersCount, rules }) => {
    return (
        <div className="bg-white p-4 rounded shadow-md">
            <h2 className="text-xl font-bold">Group Information</h2>
            <p>Members: {membersCount}</p>
            <h3 className="mt-4 text-lg font-bold">Group Rules</h3>
            <ul className="list-disc list-inside">
                {rules.map((rule, index) => {
                    <li key={index}>{rule}</li>
                })}
            </ul>
        </div>
    );
};

GroupInfo.propTypes = {
    membersCount: PropTypes.number.isRequired,
    rules: PropTypes.arrayOf(PropTypes.string).isRequired,
};

export default GroupInfo;