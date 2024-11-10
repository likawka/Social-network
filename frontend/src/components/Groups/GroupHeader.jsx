import React from 'react';
import PropTypes from 'prop-types';
import Button from '../Common/Button';

const GroupHeader = ({ bannerImage, title, description, membersCount, onMembersClick, isMember, onJoinLeave }) => {
  return (
    <div className="relative" style={{ backgroundImage: `url(${bannerImage})`, backgroundSize: 'cover' }}>
      <div className="h-48 bg-gray-800 bg-opacity-50">
        <div className="container mx-auto flex justify-between items-center h-full px-4">
          <div>
            <h1 className="text-4xl text-white font-bold">{title}</h1>
            <p className="text-xl text-white mt-2">{description}</p>
          </div>
          <div className="text-right">
            <p
              className="text-white cursor-pointer"
              onClick={onMembersClick}
            >
              Members: {membersCount}
            </p>
            <Button
              className={`${isMember ? 'bg-red-500' : 'bg-blue-500'} mt-4`}
              onClick={onJoinLeave}
            >
              {isMember ? 'Leave Group' : 'Join Group'}
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
};

GroupHeader.propTypes = {
  bannerImage: PropTypes.string.isRequired,
  title: PropTypes.string.isRequired,
  description: PropTypes.string.isRequired,
  membersCount: PropTypes.number.isRequired,
  onMembersClick: PropTypes.func.isRequired,
  isMember: PropTypes.bool.isRequired,
  onJoinLeave: PropTypes.func.isRequired,
};

export default GroupHeader;
