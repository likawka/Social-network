import React, { useState } from "react";
import Button from "../Common/Button";
import TextArea from "../Common/TextArea";
import PropTypes from "prop-types";

const ChatInput = ({ onSend }) => {
    const [message, setMessage] = useState("");

    const handleSend = () => {
        if (message.trim()) {
            onSend(message);
            setMessage("");
        }
    };

    return (
        <div className="bg-gray-100 p-4 border-t">
            <div className="flex items-center">
                <TextArea
                    value={message}
                    onChange={(e) => setMessage(e.target.value)}
                    placeholder="Type a message..."
                    rows={2}
                    className="flex-grow"
                />
                <Button onClick={handleSend} className="ml-4">
                    Send
                </Button>
            </div>
        </div>
    );
};

ChatInput.propTypes = {
    onSend: PropTypes.func.isRequired
};

export default ChatInput;