import React, { useState } from 'react';
import ChatHeader from './ChatHeader';
import MessageList from './MessageList';
import ChatInput from './ChatInput';

const ChatWindow = ({ username, isOnline, onBack }) => {
  const [messages, setMessages] = useState([]);

  const handleSend = (messageContent) => {
    const newMessage = {
      sender: 'You',
      time: new Date().toLocaleTimeString(),
      content: messageContent,
    };
    setMessages([...messages, newMessage]);
  };

  return (
    <div className="flex flex-col flex-grow">
      <ChatHeader username={username} isOnline={isOnline} onBack={onBack} />
      <div className="flex-grow overflow-y-auto bg-white border-l border-r border-gray-200">
        <MessageList messages={messages} />
      </div>
      <div className="border-t border-gray-200">
        <ChatInput onSend={handleSend} />
      </div>
    </div>
  );
};

export default ChatWindow;
