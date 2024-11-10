import React from 'react';
import PropTypes from 'prop-types';
import Button from '../Common/Button';

const ChatHeader = ({ username, isOnline, onBack }) => {
  return (
    <div className="bg-gray-800 text-white p-4 shadow-md flex justify-between items-center border-t">
      <Button onClick={onBack} className='text-blue-500 hover:text-blue-300 bg-transparent'>
        Back
      </Button>
      <div className="flex-grow text-center">
        <h2 className="text-lg font-bold">{username}</h2>
      </div>
      <div className="flex items-center ml-2">
        <span className={`mr-2 ${isOnline ? 'text-green-500' : 'text-red-500'}`}>
          ●
        </span>
        <span>{isOnline ? 'Online' : 'Offline'}</span>
      </div>
    </div>
  );
};

ChatHeader.propTypes = {
  username: PropTypes.string.isRequired,
  isOnline: PropTypes.bool.isRequired,
};

export default ChatHeader;
