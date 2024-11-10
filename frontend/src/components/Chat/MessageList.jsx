import React from "react";
import Message from "./Message";
import PropTypes from "prop-types";

const MessageList = ({ messages }) => {
    return (
        <div className="p-4 overflow-y-auto flex-grow">
            {messages.map((msg, index) => (
                <Message
                    key={index}
                    sender={msg.sender}
                    time={msg.time}
                    content={msg.content}
                />
            ))}
        </div>
    );
};

MessageList.propTypes = {
    messages: PropTypes.arrayOf(
        PropTypes.shape({
            sender: PropTypes.string.isRequired,
            time: PropTypes.string.isRequired,
            content: PropTypes.string.isRequired
        })
    ).isRequired
};

export default MessageList;
