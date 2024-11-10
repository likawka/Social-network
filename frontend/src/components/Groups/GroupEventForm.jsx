import React, { useState } from "react";
import PropTypes from "prop-types";
import Button from "../Common/Button";
import TextInput from "../Common/TextInput";
import TextArea from "../Common/TextArea";

const GroupEventForm = ({ onSubmit }) => {
    const [title, setTitle] = useState("");
    const [description, setDescription] = useState("");
    const [date, setDate] = useState("");

    const handleSubmit = (e) => {
        e.preventDefault();
        if (title && date ){
            onSubmit({ title, description, date });
            setTitle("");
            setDescription("");
            setDate("");
        }
    };

    return (
        <form onSubmit={handleSubmit} className="space-y-4">
            <TextInput
                label='Event Title'
                name="title"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                placeholder='Enter event title'
                required
            />

            <TextArea
                label='Event Description'
                name="description"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                placeholder='Describe the event'
            />

            <TextInput
                label='Event Date'
                name="date"
                type="date"
                value={date}
                onChange={(e) => setDate(e.target.value)}
                required
            />

            <Button type="submit">Create Event</Button>
        </form>
    );
};

GroupEventForm.propTypes = {
    onSubmit: PropTypes.func.isRequired,
};

export default GroupEventForm;