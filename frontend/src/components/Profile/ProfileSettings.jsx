import React, { useState } from "react";
import Button from '../Common/Button';
import TextInput from '../Common/TextInput';
import TextArea from '../Common/TextArea';
import ImageUploader from '../Common/ImageUploader';
import PropTypes from 'prop-types';

const ProfileSettings = ({ currentAvatar, currentBio, currentBannerColor, onSaveSettings }) => {
    const [newAvatar, setNewAvatar] = useState(currentAvatar);
    const [newBio, setNewBio] = useState(currentBio);
    const [newBannerColor, setNewBannerColor] = useState(currentBannerColor);

    const handleSaveSettings = () => {
        onSaveSettings({ newAvatar, newBio, newBannerColor });
    };

    return (
        <div className="p-4">
            <h2 className="text-xl font-bold mb-4">Profile Settings</h2>

            <ImageUploader
                image={newAvatar}
                setImage={setNewAvatar}
                label="Change Avatar"
                maxSizeMB={1}
                // width={40}
                // height={40}
            />

            <TextArea
                label='Bio'
                value={newBio}
                onChange={(e) => setNewBio(e.target.value)}
                placeholder='Tell us about yourself...'
            />

            <TextInput
                label='Banner Color'
                type="color"
                value={newBannerColor}
                onChange={(e) => setNewBannerColor(e.target.value)}
                placeholder='Choose another color'
            />

            <Button onClick={handleSaveSettings}>Save</Button>
        </div>
    );
};

ProfileSettings.propTypes = {
    currentAvatar: PropTypes.string.isRequired,
    currentBio: PropTypes.string.isRequired,
    currentBannerColor: PropTypes.string.isRequired,
    onSaveSettings: PropTypes.func.isRequired,
};

export default ProfileSettings;