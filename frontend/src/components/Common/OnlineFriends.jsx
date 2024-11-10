import React from "react";

const OnlineFriends = () => {
    // This is a dummy component
    const friends = [
        { name: "Friend", id: 1},
        { name: "Friend2", id: 2},
        { name: "Friend3", id: 3},
        { name: "Friend4", id: 4},
        { name: "Friend5", id: 5},
        { name: "Friend6", id: 6},
        { name: "Friend7", id: 7},
        { name: "Friend8", id: 8},
        { name: "Friend9", id: 9},
        { name: "Friend10", id: 10},
    ]

    return (
        <div className="mt-auto">
            <h3 className="text-gray-400 text-sm uppercase mb-4">Online:</h3>
            <ul>
                {friends.map((friend) => {
                    <li key={friend.id} className="flex items-center p-2 hover:bg-gray-800 rounded cursor-pointer">
                        <div className="h-8 w-8 bg-gray-700 rounded-full mr-2"></div>
                        <span>{friend.name}</span>
                    </li>
                })}
            </ul>
        </div>
    );
};

export default OnlineFriends;