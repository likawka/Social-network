import React, { useState } from 'react';
import GroupHeader from '../components/Groups/GroupHeader';
import GroupPostFeed from '../components/Groups/GroupPostFeed';
import GroupEvents from '../components/Groups/GroupEvents';
import Modal from '../components/Common/Modal';
import GroupPostForm from '../components/Groups/GroupPostForm';
import GroupEventForm from '../components/Groups/GroupEventForm';
import Button from '../components/Common/Button';

const Group = () => {
  const [posts] = useState([
    {
      id: 1,
      title: 'Favorite Books',
      content: "Let's discuss our favorite books.",
      timestamp: '2024-08-10',
      comments: [
        { id: 1, user: 'Alice', content: 'I love Harry Potter!' },
        { id: 2, user: 'Bob', content: 'Lord of the Rings is the best!' },
      ],
    },
  ]);

  const [events, setEvents] = useState([
    {
      title: 'Monthly Book Club Meeting',
      description: 'Discussing our favorite books.',
      date: '2024-09-01',
    },
  ]);

  const [isMember, setIsMember] = useState(true);
  const [showPostForm, setShowPostForm] = useState(false);
  const [showEventForm, setShowEventForm] = useState(false);
  const [isMembersModalOpen, setMembersModalOpen] = useState(false);

  const handleJoinLeave = () => {
    setIsMember((prev) => !prev);
  };

  const handleCreateEvent = (newEvent) => {
    setEvents((prevEvents) => [...prevEvents, newEvent]);
    setShowEventForm(false); // Close modal after event is created
  };

  return (
    <div>
      <GroupHeader
        bannerImage="https://via.placeholder.com/1500x300"
        title="Book Club"
        description="A group for book lovers."
        membersCount={10}
        onMembersClick={() => setMembersModalOpen(true)}
        isMember={isMember}
        onJoinLeave={handleJoinLeave}
      />
      <div className="container mx-auto mt-6 flex">
        <div className="w-1/4">
          <GroupEvents events={events} onCreateEvent={() => setShowEventForm(true)} />
        </div>
        <div className="flex-grow">
          <GroupPostFeed posts={posts} />
          <Button className="mt-4" onClick={() => setShowPostForm(true)}>
            Create a new post
          </Button>
        </div>
      </div>

      {showPostForm && (
        <Modal isOpen={showPostForm} onClose={() => setShowPostForm(false)}>
          <GroupPostForm onSubmit={(post) => console.log(post)} />
        </Modal>
      )}

      {showEventForm && (
        <Modal isOpen={showEventForm} onClose={() => setShowEventForm(false)}>
          <h3 className="text-xl font-bold mb-4">Create a New Event</h3>
          <GroupEventForm onSubmit={handleCreateEvent} />
        </Modal>
      )}

      <Modal isOpen={isMembersModalOpen} onClose={() => setMembersModalOpen(false)}>
        <h3 className="text-xl font-bold mb-4">Group Members</h3>
        <ul>
          <li>Alice</li>
          <li>Bob</li>
          {/* Add more members */}
        </ul>
      </Modal>
    </div>
  );
};

export default Group;
