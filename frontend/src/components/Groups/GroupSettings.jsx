import React, { useState } from 'react';
import PropTypes from 'prop-types';
import Button from '../Common/Button';
import TextArea from '../Common/TextArea';
import TextInput from '../Common/TextInput';

const GroupSettings = ({ currentTitle, currentDescription, onSaveSettings }) => {
  const [newTitle, setNewTitle] = useState(currentTitle);
  const [newDescription, setNewDescription] = useState(currentDescription);

  const handleSaveSettings = () => {
    onSaveSettings({ newTitle, newDescription });
  };

  return (
    <div className="bg-white p-4 rounded shadow-md mt-4">
      <h2 className="text-xl font-bold mb-4">Group Settings</h2>

      <TextInput
        label="Group Title"
        name="title"
        value={newTitle}
        onChange={(e) => setNewTitle(e.target.value)}
        placeholder="Enter new title"
      />

      <TextArea
        label="Group Description"
        name="description"
        value={newDescription}
        onChange={(e) => setNewDescription(e.target.value)}
        placeholder="Enter new description"
      />

      <Button onClick={handleSaveSettings}>Save Settings</Button>
    </div>
  );
};

GroupSettings.propTypes = {
  currentTitle: PropTypes.string.isRequired,
  currentDescription: PropTypes.string.isRequired,
  onSaveSettings: PropTypes.func.isRequired,
};

export default GroupSettings;
