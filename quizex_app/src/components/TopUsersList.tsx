import React from "react";

interface TopUsersListProps {
  users: User[];
}

export const TopUsersList: React.FC<TopUsersListProps> = ({ users }) => {
  return (
    <div className=" bg-pattern shadow-lg w-full h-full p-10">
      <h2 className="text-xl font-bold mb-4">Top Users</h2>
      <ul>
        {users.map((user) => (
          <li
            key={user.uuid}
            className="flex items-center gap-4 p-2 border-b last:border-none"
          >
            <img
              src={user.image_url}
              alt={user.name}
              className="w-10 h-10 rounded-full object-cover"
            />
            <div>
              <p className="font-medium">{user.name}</p>
              <p className="text-sm text-gray-500">Level: {user.level}</p>
            </div>
            <span className="ml-auto text-sm font-semibold text-blue-500">
              Exp: {user.exp}
            </span>
          </li>
        ))}
      </ul>
    </div>
  );
};
