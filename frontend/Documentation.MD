# Anime Network Frontend Documentation

## 1. Overview

This is the frontend documentation for the Anime Network project. The application is built with React and styled using Tailwind CSS. The application features user authentication, profile management, posts with privacy settings, group management, and follower management.

### Technology Stack:
- **React**: A JavaScript library for building user interfaces.
- **Tailwind CSS**: A utility-first CSS framework for styling.
- **ESLint**: A tool for identifying and reporting on patterns in JavaScript, enforcing coding standards.

---

## 2. Components

### 2.1. Auth Components

#### `Login.jsx`
- **Description**: Provides a user interface for logging into the application.
- **Props**: None
- **State**: 
  - `email` (string): Stores the user's email.
  - `password` (string): Stores the user's password.
- **Functions**:
  - `handleLogin`: Handles the form submission and triggers the login process.

#### `Register.jsx`
- **Description**: Provides a user interface for registering a new user.
- **Props**: None
- **State**:
  - `email`, `password`, `firstName`, `lastName`, `dateOfBirth`, `selectedAvatar`, `uploadedAvatar`, `nickname`, `aboutMe`
- **Functions**:
  - `handleRegister`: Handles the form submission and triggers the registration process.
- **Related Components**:
  - `TextInput`: Used for various text input fields.
  - `TextArea`: Used for the "About Me" field.
  - `AvatarSelector`: Allows users to select a default avatar.
  - `ImageUploader`: Allows users to upload their own avatar.

### 2.2. Profile Components

#### `Profile.jsx`
- **Description**: Displays the user's profile, including their posts, followers, and following.
- **State**:
  - `posts`: Stores a list of posts made by the user.
  - `userDetails`: Stores the user's details like name, avatar, and bio.
  - `followers`: Stores a list of followers.
  - `following`: Stores a list of users the user is following.
  - `isModalOpen`: Controls the visibility of the modal for followers/following.
  - `modalType`: Determines whether the modal shows followers or following.
  - `isCurrentUser`: Determines if the profile being viewed is the current user's.
- **Functions**:
  - `openFollowersModal`, `openFollowingModal`: Opens the modal for followers or following.
- **Related Components**:
  - `ProfileHeader`: Displays the profile banner with avatar and follow counts.
  - `ProfileDetails`: Shows detailed user information like name, bio, etc.
  - `ProfileActivityFeed`: Lists the user's posts.
  - `ProfileFollowers`: Lists the user's followers or following users.
  - `Modal`: Displays a modal window for followers/following.

#### `ProfileHeader.jsx`
- **Description**: Displays the user's banner with their avatar, username, and follow counts.
- **Props**:
  - `avatar` (string): The path to the user's avatar image.
  - `username` (string): The user's nickname or username.
  - `bannerColor` (string): The color of the profile banner.
  - `followersCount` (number): Number of followers the user has.
  - `followingCount` (number): Number of users the user is following.
  - `onFollowersClick` (function): Function triggered when the followers count is clicked.
  - `onFollowingClick` (function): Function triggered when the following count is clicked.
  - `isCurrentUser` (boolean): Determines if the profile is the current user's.

#### `ProfileDetails.jsx`
- **Description**: Displays detailed information about the user, such as their name, nickname, bio, etc.
- **Props**:
  - `firstName` (string): The user's first name.
  - `lastName` (string): The user's last name.
  - `nickname` (string): The user's nickname.
  - `aboutMe` (string): A brief description or bio of the user.
  - `avatar` (string): The path to the user's avatar image.

#### `ProfileActivityFeed.jsx`
- **Description**: Displays a list of posts made by the user.
- **Props**:
  - `posts` (array of objects): List of posts made by the user.
- **Related Components**:
  - `PostItem`: Displays individual post items with content, timestamp, and privacy settings.

#### `ProfileFollowers.jsx`
- **Description**: Displays a list of followers or users the profile owner is following.
- **Props**:
  - `followers` (array of objects): List of users following the profile owner.
  - `following` (array of objects): List of users the profile owner is following.

#### `ProfileSettings.jsx`
- **Description**: Provides an interface for users to change their profile settings, such as avatar, bio, and banner color.
- **Props**:
  - `currentAvatar` (string): The user's current avatar.
  - `currentBio` (string): The user's current bio.
  - `currentBannerColor` (string): The user's current banner color.
  - `onSaveSettings` (function): Function triggered when the user saves their settings.
- **State**:
  - `newAvatar` (string): The updated avatar image.
  - `newBio` (string): The updated bio.
  - `newBannerColor` (string): The updated banner color.
- **Related Components**:
  - `ImageUploader`: Allows users to upload a new avatar.
  - `TextArea`: Used for the bio field.
  - `TextInput`: Used for the banner color field.

### 2.3. Post Components

#### `PostForm.jsx`
- **Description**: Provides a form for creating new posts with options to set content, image, and privacy.
- **Props**:
  - `onSubmit` (function): Function triggered when the form is submitted.
- **State**:
  - `title` (string): The title of the post.
  - `content` (string): The content of the post.
  - `image` (object): The image attached to the post.
  - `privacy` (string): The privacy setting of the post (`public`, `followers`, `private`).
- **Functions**:
  - `handleSubmit`: Handles form submission and calls the `onSubmit` function.
- **Related Components**:
  - `TextInput`: Used for the title field.
  - `TextArea`: Used for the content field.
  - `ImageUploader`: Allows users to upload an image for the post.
  - `Select`: Used for selecting the privacy setting.

#### `PostItem.jsx`
- **Description**: Displays an individual post, including its content, title, timestamp, privacy setting, and a section for comments.
- **Props**:
  - `post` (object): The post object containing details like content, title, timestamp, privacy, and comments.
- **State**:
  - `comments` (array of objects): List of comments on the post.
  - `isCommentSectionOpen` (boolean): Controls the visibility of the comment section.
- **Functions**:
  - `handleCommentSubmit`: Handles the submission of a new comment.
  - `toggleCommentSection`: Toggles the visibility of the comment section.

### 2.4. Group Components

#### `Group.jsx`
- **Description**: Displays a group page with its posts, events, and member management.
- **State**:
  - `posts`: Stores a list of posts in the group.
  - `events`: Stores a list of events in the group.
  - `isMember`: Determines if the current user is a member of the group.
  - `showPostForm`: Controls the visibility of the "Create Post" modal.
  - `showEventForm`: Controls the visibility of the "Create Event" modal.
  - `isMembersModalOpen`: Controls the visibility of the members list modal.
- **Functions**:
  - `handleJoinLeave`: Handles joining or leaving the group.
  - `handleCreateEvent`: Adds a new event to the group and closes the event modal.
- **Related Components**:
  - `GroupHeader`: Displays the group banner with title, description, and member count.
  - `GroupPostFeed`: Displays the list of posts in the group.
  - `GroupEvents`: Displays the list of events in the group.
  - `GroupPostForm`: Modal form for creating a new post in the group.
  - `GroupEventForm`: Modal form for creating a new event in the group.
  - `Modal`: Reusable modal component for various forms.

#### `GroupHeader.jsx`
- **Description**: Displays the group's banner with title, description, and member count.
- **Props**:
  - `bannerImage` (string): The URL of the banner image.
  - `title` (string): The title of the group.
  - `description` (string): A short description of the group.
  - `membersCount` (number): The number of members in the group.
  - `isMember` (boolean): Indicates whether the current user is a member of the group.
  - `onJoinLeave` (function): Function triggered when the user clicks the join/leave button.
  - `onMembersClick` (function): Function triggered when the member count is clicked.

#### `GroupPostFeed.jsx`
- **Description**: Displays a feed of posts within the group.
- **Props**:
  - `posts` (array of objects): The list of posts in the group.
- **Related Components**:
  - `PostItem`: Displays individual post items with content, title, timestamp, and privacy settings.

#### `GroupEvents.jsx`
- **Description**: Displays the list of events scheduled within the group.
- **Props**:
  - `events` (array of objects): The list of events in the group.
  - `onCreateEvent` (function): Function triggered when the "Schedule a new event" button is clicked.

#### `GroupPostForm.jsx`
- **Description**: Provides a form for creating a new post within the group.
- **Props**:
  - `onSubmit` (function): Function triggered when the form is submitted.
- **State**:
  - `title` (string): The title of the post.
  - `content` (string): The content of the post.
  - `image` (object): The image attached to the post.
  - `privacy` (string): The privacy setting of the post (`public`, `followers`, `private`).

#### `GroupEventForm.jsx`
- **Description**: Provides a form for scheduling a new event within the group.
- **Props**:
  - `onSubmit` (function): Function triggered when the form is submitted.
- **State**:
  - `title` (string): The title of the event.
  - `description` (string): The description of the event.
  - `date` (string): The date of the event.

### 2.5. Common Components

#### `Button.jsx`
- **Description**: A reusable button component.
- **Props**:
  - `type` (string): The type of the button (`button`, `submit`, etc.).
  - `onClick` (function): The function triggered when the button is clicked.

#### `TextInput.jsx`
- **Description**: A reusable text input component.
- **Props**:
  - `label` (string): The label for the input field.
  - `name` (string): The name attribute of the input field.
  - `type` (string): The type of the input (`text`, `password`, `email`, etc.).
  - `value` (string): The value of the input field.
  - `onChange` (function): The function triggered when the input value changes.

#### `TextArea.jsx`
- **Description**: A reusable text area component.
- **Props**:
  - `label` (string): The label for the text area.
  - `name` (string): The name attribute of the text area.
  - `value` (string): The value of the text area.
  - `onChange` (function): The function triggered when the text area value changes.

#### `ImageUploader.jsx`
- **Description**: A component for uploading images.
- **Props**:
  - `image` (object): The current image selected for upload.
  - `setImage` (function): Function to update the selected image.
  - `label` (string): The label for the file input.
  - `maxSizeMB` (number): The maximum size of the image file in megabytes.
  - `width` (number): The expected width of the uploaded image.
  - `height` (number): The expected height of the uploaded image.

#### `Select.jsx`
- **Description**: A reusable select dropdown component.
- **Props**:
  - `label` (string): The label for the select field.
  - `name` (string): The name attribute of the select field.
  - `value` (string): The current selected value.
  - `onChange` (function): The function triggered when the selected value changes.
  - `options` (array): List of options for the select dropdown.

#### `Modal.jsx`
- **Description**: A reusable modal component used to display content in a modal window.
- **Props**:
  - `isOpen` (boolean): Controls the visibility of the modal.
  - `onClose` (function): Function triggered to close the modal.
  - `children` (node): The content to display inside the modal.
- **Functions**:
  - `handleClickOutside`: Handles closing the modal when clicking outside of it.
  - `useEffect`: Adds and removes event listeners based on modal visibility.

#### `DropDownMenu.jsx`
- **Description**: A component that displays a dropdown menu when clicked.
- **Props**:
  - `label` (string): The text that will be displayed on the dropdown button.
  - `items` (array of objects): The list of items to display in the dropdown.
  - `onSelect` (function): The function to call when an item is selected.

#### `Anchor.jsx`
- **Description**: A reusable anchor link component.
- **Props**:
  - `href` (string): The URL the link points to.
  - `label` (string): The text to display for the link.
  - `className` (string): Additional classes for styling.

#### `Logo.jsx`
- **Description**: Displays the application's logo.
- **Props**: None

#### `OnlineFriends.jsx`
- **Description**: Displays a list of online friends.
- **Props**:
  - `friends` (array of objects): The list of online friends.

#### `UserInfo.jsx`
- **Description**: Displays the current user's information.
- **Props**:
  - `username` (string): The current user's username.
  - `avatar` (string): The path to the current user's avatar image.

---

## 3. Project Structure

- **`src/`**: Contains all source code.
  - **`components/`**: Contains all React components.
    - **`Auth/`**: Components related to authentication.
    - **`Profile/`**: Components related to user profiles.
    - **`Posts/`**: Components related to posts and post creation.
    - **`Groups/`**: Components related to groups, group posts, and group events.
    - **`Common/`**: Reusable components used throughout the application.
    - **`Chat/`**: Components related to chat functionality.
    - **`Layout/`**: Components related to the layout of the application, such as the sidebar, navbar, and footer.
    - **`Notifications/`**: Components related to notifications.
  - **`pages/`**: Contains page components like Home, Profile, and Group.
  - **`services/`**: Contains API and route management.
  - **`assets/`**: Contains static assets like images.

---

## 4. Future Work

- **Group Page**: Implement editing capabilities for group pages.
- **Group Roles**: Consider implementing roles such as admin, member, and guest within groups.
- **Integrate with Backend**: Implement API calls to connect the frontend with the backend database.
- **Improve Styling**: Further enhance the UI/UX for a more polished look.
- **Error Handling**: Add comprehensive error handling across components.
- **Testing**: Implement unit and integration tests for all components.

---

## 5. Contribution Guide

- **Code Style**: Follow ESLint guidelines and the project's coding standards.
- **Pull Requests**: Ensure all new features or fixes are done in separate branches and are accompanied by a pull request.
- **Documentation**: Update this documentation with any new components or significant changes.
