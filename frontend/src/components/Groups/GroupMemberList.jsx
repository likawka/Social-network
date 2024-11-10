import React from "react";
import PropTypes from "prop-types";

const GroupMemberList = ({ members }) => {
    return (
        <div className="bg-white p-4 rounded shadow-md mt-4">
            <h2 className="text-xl font-bold mb-4">Group Members</h2>
            <ul>
                {members.map((member, index) => {
                    <li key={index} className="flex items-center mb-2">
                        <img
                        src={member.avatar}
                        alt={member.username}
                        className="w-10 h-10 rounded-full mr-4"
                        />
                        <span>{member.username}</span>
                    </li>
                })}
            </ul>
        </div>
    );
};

GroupMemberList.propTypes = {
    members: PropTypes.arrayOf(
        PropTypes.shape({
            avatar: PropTypes.string.isRequired,
            username: PropTypes.string.isRequired,
        })
    ).isRequired,
};

export default GroupMemberList;