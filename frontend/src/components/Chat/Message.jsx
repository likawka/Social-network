import React from "react";
import PropTypes from "prop-types";

const Message = ({ sender, time, content }) => {
    return (
        <div className="mb-4">
            <div className="text-sm text-grey-500">
                <span className="font-semibold">{sender}</span> at {time}
            </div>
            <div className="bg-grey-200 p-2 rounded">
                {content}
            </div>
        </div>
    );
};

Message.PropTypes = {
    sender: PropTypes.string.isRequired,
    time: PropTypes.string.isRequired,
    content: PropTypes.string.isRequired
};

export default Message;