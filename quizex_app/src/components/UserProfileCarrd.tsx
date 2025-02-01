import React from "react";

const UserProfileCard: React.FC<{ user: User }> = ({ user }) => {
  return (
    <div className="hidden sm:block max-w-md mx-auto bg-gradient rounded-xl shadow-md overflow-hidden md:max-w-2xl">
      <div className="flex relative">
        <div className="md:flex-shrink-0">
          <img
            className="h-48 w-full object-cover md:w-48"
            src={user.image_url}
            alt={`${user.name}'s avatar`}
          />
          <progress max="100" value={user.progress} className="w-full">
            {user.progress}
          </progress>
          <p className="absolute left-20 bottom-20 border border-black font-mono font-extrabold rounded-full p-1 px-2.5 bg-white">
            {user.level}
          </p>
        </div>
        <div className="px-8 pb-8">
          <p className="font-mono font-extrabold text-teal-600 mt-2 text-center">
            QUIZEX MEMBER CARD
          </p>

          <div className="uppercase tracking-wide text-sm text-indigo-500 font-semibold">
            {user.name}
          </div>
          <p className="mt-2 text-gray-500">
            <span className="font-bold">Email:</span> {user.email}
          </p>
          <p className="mt-2 text-gray-500">
            <span className="font-bold">Level:</span> {user.level}
          </p>
          <p className="mt-2 text-gray-500">
            <span className="font-bold">Experience:</span> {user.exp} EXP
          </p>
          <p className="mt-2 text-gray-500">
            <span className="font-bold">Progress:</span> {user.progress}%
          </p>
          <p className="mt-2 text-gray-500">
            <span className="font-bold">Account Created:</span>{" "}
            {new Date(user.created_at).toLocaleDateString()}
          </p>
          <p className="mt-2 text-gray-500">
            <span className="font-bold">Status:</span>{" "}
            {user.is_draft ? "Draft" : "Active"}
          </p>
          <p className="mt-2 text-gray-500">
            <span className="font-bold">UUID:</span> {user.uuid}
          </p>
        </div>
      </div>
    </div>
  );
};

export default UserProfileCard;
