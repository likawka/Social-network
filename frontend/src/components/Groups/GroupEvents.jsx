import React from 'react';
import PropTypes from 'prop-types';

const GroupEvents = ({ events, onCreateEvent }) => {
  return (
    <div className="bg-white p-4 rounded shadow-md mt-4">
      <h2 className="text-xl font-bold mb-4">Events</h2>
      {events.length > 0 ? (
        events.map((event, index) => (
          <div key={index} className="mb-4">
            <h3 className="text-lg font-bold">{event.title}</h3>
            <p>{event.description}</p>
            <p className="text-gray-500 text-sm">Event Date: {event.date}</p>
          </div>
        ))
      ) : (
        <p>No upcoming events.</p>
      )}
      <button className="mt-4 bg-blue-500 text-white px-4 py-2 rounded" onClick={onCreateEvent}>
        Schedule a new event
      </button>
    </div>
  );
};

GroupEvents.propTypes = {
  events: PropTypes.arrayOf(
    PropTypes.shape({
      title: PropTypes.string.isRequired,
      description: PropTypes.string.isRequired,
      date: PropTypes.string.isRequired,
    })
  ).isRequired,
  onCreateEvent: PropTypes.func.isRequired,
};

export default GroupEvents;
