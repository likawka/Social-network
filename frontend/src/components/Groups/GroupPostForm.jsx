import React, { useState } from 'react';
import Button from '../Common/Button';
import TextArea from '../Common/TextArea';
import ImageUploader from '../Common/ImageUploader';
import Select from '../Common/Select';
import TextInput from '../Common/TextInput';

const GroupPostForm = ({ onSubmit }) => {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [image, setImage] = useState(null);
  const [privacy, setPrivacy] = useState('public');

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit({ title, content, image, privacy });
    setTitle('');
    setContent('');
    setImage(null);
    setPrivacy('public');
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <TextInput
        label="Title"
        name="title"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="Post Title"
        className="w-full"
      />

      <TextArea
        label="What's on your mind?"
        name="content"
        value={content}
        onChange={(e) => setContent(e.target.value)}
        placeholder="Share something..."
      />

      <ImageUploader
        image={image}
        setImage={setImage}
        label="Add an image"
        maxSizeMB={5}
        width={600}
        height={400}
      />

      <Select
        label="Privacy"
        name="privacy"
        value={privacy}
        onChange={(e) => setPrivacy(e.target.value)}
        options={[
          { value: 'public', label: 'Public' },
          { value: 'followers', label: 'Followers' },
          { value: 'private', label: 'Private' },
        ]}
      />

      <Button type="submit">Post</Button>
    </form>
  );
};

export default GroupPostForm;
