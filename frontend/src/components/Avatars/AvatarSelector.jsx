import React, { useState, useEffect } from 'react';
import ShapeAvatar from './ShapeAvatar';
import PropTypes from 'prop-types';

const AvatarSelector = ({ selectedAvatar, onSelectAvatar }) => {
    const [avatars, setAvatars] = useState([]);

    useEffect(() => {
        // Replace this with actual API call
        const fetchAvatars = async () => {
            const mockAvatars = [
                { shape: 'circle', color: 'bg-red-500' },
                { shape: 'square', color: 'bg-green-500' },
                { shape: 'triangle', color: 'bg-blue-500' },
                { shape: 'hexagon', color: 'bg-purple-500' },
                { shape: 'pentagon', color: 'bg-orange-500' }
            ];
            setAvatars(mockAvatars);
        };

        fetchAvatars();
    }, []);

    return (
        <div className="mb-4">
            <label className="block text-gray-700 text-sm font-bold mb-2">
                Select a Default Avatar
            </label>
            <div className="flex space-x-2">
                {avatars.map((avatar, index) => (
                    <ShapeAvatar
                        key={index}
                        shape={avatar.shape}
                        color={avatar.color}
                        isSelected={
                            selectedAvatar?.shape === avatar.shape &&
                            selectedAvatar?.color === avatar.color
                        }
                        onClick={() => onSelectAvatar(avatar)}
                    />
                ))}
            </div>
        </div>
    );
};

AvatarSelector.propTypes = {
    selectedAvatar: PropTypes.object,
    onSelectAvatar: PropTypes.func.isRequired,
};

export default AvatarSelector;
