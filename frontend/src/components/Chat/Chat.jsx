import React, { useState, useEffect } from "react";
import ChatWindow from "./ChatWindow";
import Button from "../Common/Button";
import Modal from "../Common/Modal";
import TextInput from "../Common/TextInput";

const Chat = () => {
    const [activeChats, setActiveChats] = useState([]);
    const [onlineFriends, setOnlineFriends] = useState([]);
    const [selectedChat, setSelectedChat] = useState(null);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [modalType, setModalType] = useState(null);
    const [chatName, setChatName] = useState('');
    const [groupMembers, setGroupMembers] = useState('');

    useEffect(() => {
        // Simulate fetching active chats
        const fetchActiveChats = () => {
            const dummyActiveChats = [
                { id: 1, username: "John Doe", isOnline: true, type: "private" },
                { id: 2, username: "Jane Doe", isOnline: false, type: "private" },
                { id: 3, username: "Book Club", isOnline: true, type: "group" }, // Group chat
            ];
            setActiveChats(dummyActiveChats);
        };

        // Simulate fetching online friends
        const fetchOnlineFriends = () => {
            const dummyOnlineFriends = [
                { id: 4, username: "Alice", isOnline: true },
                { id: 5, username: "Bob", isOnline: true },
            ];
            setOnlineFriends(dummyOnlineFriends);
        };

        fetchActiveChats();
        fetchOnlineFriends();
    }, []);

    const handleChatClick = (chat) => {
        setSelectedChat(chat);
    };

    const handleBackClick = () => {
        setSelectedChat(null); // Reset selected chat to show chat list again
    };

    const handleOpenModal = (type) => {
        setModalType(type);
        setIsModalOpen(true);
    };

    const handleCloseModal = () => {
        setIsModalOpen(false);
        setChatName('');
        setGroupMembers('');
    };

    const handleCreateChat = () => {
        const newChat = {
            id: activeChats.length + 1,
            username: chatName,
            isOnline: true,
            type: modalType === 'group' ? 'group' : 'private'
        };
        setActiveChats([...activeChats, newChat]);
        setSelectedChat(newChat); // Automatically open the new chat
        handleCloseModal();
    };

    return (
        <div className="flex h-full">
            <div className="w-1/4 bg-gray-100 border-r border-gray-200">
                <div className="flex justify-between p-4">
                    <h2 className="text-xl border border-red-800 rounded">Active Chats</h2>
                    <div className="flex space-x-2">
                        <Button onClick={() => handleOpenModal('private')} className="text-blue-500 hover:text-blue-300">New Chat</Button>
                        <Button onClick={() => handleOpenModal('group')} className="text-blue-500 hover:text-blue-300">New Group Chat</Button>
                    </div>
                </div>
                <ul>
                    {activeChats.map((chat) => (
                        <li
                            key={chat.id}
                            className={`p-4 cursor-pointer hover:bg-gray-200 ${
                                selectedChat && selectedChat.id === chat.id
                                    ? 'border border-dotted border-gray-200 rounded'
                                    : ''
                            }`}
                            onClick={() => handleChatClick(chat)}
                        >
                            {chat.username} {chat.isOnline ? "(Online)" : "(Offline)"}
                        </li>
                    ))}
                </ul>
            </div>
            <div className="flex-grow flex flex-col">
                {selectedChat ? (
                    <ChatWindow 
                        username={selectedChat.username} 
                        isOnline={selectedChat.isOnline}
                        onBack={handleBackClick}  // Pass the back click handler
                    />
                ) : (
                    <div className="p-4">
                        <h3 className="text-lg font-semibold border border-red-800 rounded">Online Friends</h3>
                        <ul>
                            {onlineFriends.map((friend) => (
                                <li
                                    key={friend.id}
                                    className={`p-4 cursor-pointer hover:bg-gray-200 ${
                                        selectedChat && selectedChat.id === friend.id
                                            ? 'border border-dotted border-gray-500 rounded'
                                            : ''
                                    }`}
                                    onClick={() => handleChatClick(friend)}
                                >
                                    <span className="ml-2 h-3 w-3 rounded-full bg-green-500 inline-block mr-4"></span>
                                    {friend.username}
                                </li>
                            ))}
                        </ul>
                    </div>
                )}
            </div>

            <Modal isOpen={isModalOpen} onClose={handleCloseModal}>
                <h3 className="text-xl font-bold mb-4">{modalType === 'group' ? 'Create a New Group Chat' : 'Create a New Chat'}</h3>
                <TextInput
                    label={modalType === 'group' ? 'Group Name' : 'Username'}
                    value={chatName}
                    onChange={(e) => setChatName(e.target.value)}
                />
                {modalType === 'group' && (
                    <TextInput
                        label="Add Group Members (comma separated)"
                        value={groupMembers}
                        onChange={(e) => setGroupMembers(e.target.value)}
                    />
                )}
                <Button onClick={handleCreateChat} className="mt-4">
                    {modalType === 'group' ? 'Create Group Chat' : 'Create Chat'}
                </Button>
            </Modal>
        </div>
    );
};

export default Chat;
