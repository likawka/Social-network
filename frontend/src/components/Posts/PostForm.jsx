import React, { useState } from 'react';
import Button from '../Common/Button';
import TextArea from '../Common/TextArea';
import ImageUploader from '../Common/ImageUploader';
import Select from '../Common/Select';
import TextInput from '../Common/TextInput';
import api from '../../services/api.jsx'; // Import the createPost function
import { useNavigate } from "react-router-dom";


const PostForm = ({ onSubmit }) => {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [image, setImage] = useState(null);
  const [privacy, setPrivacy] = useState('public'); // Default privacy setting
  const [errors, setErrors] = useState({}); // State to store validation errors

  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Assuming the image is in Base64 format or needs to be converted
    const formattedImage = image ? await convertToBase64(image) : null;

    const postData = {
      title,
      content,
      image: formattedImage,
      privacy,
    };

    console.log('Post data:', postData);

    try {
      const response = await api.createPost(postData);
      console.log('Post created successfully:', response);
      navigate('/');
    } catch (error) {
      console.error('Failed to create post:', error);

      const newErrors = {
        title: '',
        content: '',
      };

      if (error.details && Array.isArray(error.details)) {
        error.details.forEach(err => {
          if (newErrors.hasOwnProperty(err.field)) {
            newErrors[err.field] = err.message;
          }
        });
      } else {
        console.error('Error details are missing or not an array:', error.details);
      }
  
      setErrors(newErrors);
    }

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
        placeholder="Title of your post"
      />
      {errors.title && (
        <p className="text-red-500 text-sm">{errors.title}</p>
      )}

      <TextArea
        label="What's on your mind?"
        name="content"
        value={content}
        onChange={(e) => setContent(e.target.value)}
        placeholder="Share something..."
      />
      {errors.content && (
        <p className="text-red-500 text-sm">{errors.content}</p>
      )}

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

// Example function to convert image file to Base64
const convertToBase64 = (file) => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = (error) => reject(error);
  });
};

export default PostForm;
