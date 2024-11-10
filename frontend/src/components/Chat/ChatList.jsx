import React from 'react';
import PropTypes from 'prop-types';

const ChatList = ({ activeChats, onChatSelect }) => {
  return (
    <ul className="divide-y divide-gray-200">
      {activeChats.map((chat) => (
        <li 
          key={chat.id} 
          className="p-4 cursor-pointer hover:bg-gray-100"
          onClick={() => onChatSelect(chat.id)}
        >
          <span className={`mr-2 ${chat.online ? 'text-green-500' : 'text-gray-500'}`}>●</span>
          {chat.username}
        </li>
      ))}
    </ul>
  );
};

ChatList.propTypes = {
  activeChats: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.number.isRequired,
      username: PropTypes.string.isRequired,
      online: PropTypes.bool.isRequired,
    })
  ).isRequired,
  onChatSelect: PropTypes.func.isRequired,
};

export default ChatList;
