import { useState } from "react";
import { LittlePaperWrapper } from "./Wrappers";
import { useNavigate } from "react-router-dom";

export const RoomJoiner = () => {
  const [roomName, setRoomName] = useState("");
  const navigate = useNavigate();
  return (
    <LittlePaperWrapper>
      <p className="text-base text-center">Enter a room via name</p>
      <div className="flex items-center">
        <input
          className="rounded-sm text-sm mr-2 bg-white p-1"
          type="text"
          placeholder="Room Name"
          onChange={(e) => setRoomName(e.currentTarget.value)}
        ></input>
        <button
          className="btn bg-teal-400 text-xs"
          onClick={() => navigate(`/quiz/${roomName}`)}
        >
          Join
        </button>
      </div>
    </LittlePaperWrapper>
  );
};
