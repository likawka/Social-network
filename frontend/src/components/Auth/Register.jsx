import React, { useState } from "react";
import Button from "../Common/Button";
import TextInput from "../Common/TextInput";
import TextArea from "../Common/TextArea";
import AvatarSelector from "../Avatars/AvatarSelector";
import ImageUploader from "../Common/ImageUploader";
import api from "../../services/api.jsx";
import { useNavigate } from "react-router-dom";

const Register = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [dateOfBirth, setDateOfBirth] = useState("");
  const [selectedAvatar, setSelectedAvatar] = useState(null);
  const [uploadedAvatar, setUploadedAvatar] = useState(null); // State for uploaded image
  const [nickname, setNickname] = useState("");
  const [aboutMe, setAboutMe] = useState("");
  const [errors, setErrors] = useState({}); // State to store validation errors
  const navigate = useNavigate();

  const handleRegister = async () => {
    // Determine which avatar to use
    const avatar = uploadedAvatar || selectedAvatar;
  
    // Prepare user data
    const userData = {
      email,
      password,   
      firstName,
      lastName,
      dateOfBirth,
      avatar, // Use the determined avatar
      nickname,
      aboutMe,
    };
  
    console.log('Register to fetch:', userData);
  
    try {
      const response = await api.register(userData);
      console.log('Registration successful:', response);
      navigate('/');
    } catch (error) {
      console.error('Registration failed:', error);
  
      // Initialize errors state
      const newErrors = {
        email: '',
        password: '',
        firstName: '',
        lastName: '',
        dateOfBirth: '',
        avatar: '', // Changed to avatar
        nickname: '',
        aboutMe: '',
      };
  
      // Safely handle error details
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
  };
  
  

  return (
    <div className="max-w-sm mx-auto">
      <h2 className="text-2xl font-bold mb-4">Register</h2>

      <TextInput
        label="First Name"
        name="firstName"
        value={firstName}
        onChange={(e) => setFirstName(e.target.value)}
        placeholder="First Name"
      />
      {errors.firstName && (
        <p className="text-red-500 text-sm">{errors.firstName}</p>
      )}

      <TextInput
        label="Last Name"
        name="lastName"
        value={lastName}
        onChange={(e) => setLastName(e.target.value)}
        placeholder="Last Name"
      />
      {errors.lastName && (
        <p className="text-red-500 text-sm">{errors.lastName}</p>
      )}

      <TextInput
        label="Email"
        name="email"
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        placeholder="Email"
      />
      {errors.email && <p className="text-red-500 text-sm">{errors.email}</p>}

      <TextInput
        label="Password"
        name="password"
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        placeholder="Password"
      />
      {errors.password && (
        <p className="text-red-500 text-sm">{errors.password}</p>
      )}

      <TextInput
        label="Date of Birth"
        name="dateOfBirth"
        type="date"
        value={dateOfBirth}
        onChange={(e) => setDateOfBirth(e.target.value)}
        placeholder="Date of Birth"
      />
      {errors.dateOfBirth && (
        <p className="text-red-500 text-sm">{errors.dateOfBirth}</p>
      )}

      <AvatarSelector
        selectedAvatar={selectedAvatar}
        onSelectAvatar={setSelectedAvatar}
      />
      {errors.selectedAvatar && (
        <p className="text-red-500 text-sm">{errors.selectedAvatar}</p>
      )}

      <ImageUploader
        image={uploadedAvatar}
        setImage={setUploadedAvatar}
        label="Upload Your Own Avatar"
      />
      {errors.uploadedAvatar && (
        <p className="text-red-500 text-sm">{errors.uploadedAvatar}</p>
      )}

      <TextInput
        label="Nickname (Optional)"
        name="nickname"
        value={nickname}
        onChange={(e) => setNickname(e.target.value)}
        placeholder="Nickname"
      />
      {errors.nickname && (
        <p className="text-red-500 text-sm">{errors.nickname}</p>
      )}

      <TextArea
        label="About Me (Optional)"
        name="aboutMe"
        value={aboutMe}
        onChange={(e) => setAboutMe(e.target.value)}
        placeholder="Tell us about yourself"
      />
      {errors.aboutMe && (
        <p className="text-red-500 text-sm">{errors.aboutMe}</p>
      )}

      <Button onClick={handleRegister}>Register</Button>
    </div>
  );
};

export default Register;
