import React, { useState } from 'react';
import ChatList from './ChatList';
import Button from '../Common/Button';

const ChatPopup = ({ onChatSelect }) => {
  const [activeChats, setActiveChats] = useState([
    // Dummy data for active chats
    { id: 1, username: 'John Doe', online: true },
    { id: 2, username: 'Jane Smith', online: false },
  ]);

  return (
    <div className="fixed bottom-0 right-0 m-4 w-full max-w-sm bg-white border border-gray-300 rounded shadow-lg">
      <div className="p-4 flex justify-between items-center">
        <h3 className="text-lg font-semibold">Chats</h3>
        <div>
          <Button className="mr-2" onClick={() => { /* Logic for New Chat */ }}>New Chat</Button>
          <Button onClick={() => { /* Logic for New Group Chat */ }}>New Group Chat</Button>
        </div>
      </div>
      <ChatList activeChats={activeChats} onChatSelect={onChatSelect} />
    </div>
  );
};

export default ChatPopup;
