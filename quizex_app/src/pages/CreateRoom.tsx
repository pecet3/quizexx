import { useState } from "react";
import { MainWrapper } from "../components/MainWrapper";
import { useNavigate } from "react-router-dom";
import { Settings } from "./Quiz";
import { RoomCreator } from "../components/RoomCreator";

// to do : better msgor handling in UI

export const CreateRoom = () => {
  const nav = useNavigate();
  const [msg, setMsg] = useState("");

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    setMsg("Creating a game, please be patient");
    const formData = new FormData(event.target);
    const maxRoundsStr = formData.get("maxRounds");
    const maxRounds = parseInt(maxRoundsStr as string);
    const data: Settings = {
      name: formData.get("roomName") as string,
      gen_content: formData.get("category") as string,
      difficulty: formData.get("difficulty") as string,
      max_rounds: maxRounds,
      language: formData.get("lang") as string,
    };

    const response = await fetch("/api/quiz/rooms", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });

    if (response.status === 200) {
      const result = await response.json;
      console.log("Room created successfully:", result);
      nav(`/quiz/${data.name}`);
      setMsg("");
      return;
    }
    if (response.status === 403) {
      setMsg("Room with this name already exists!");
      return;
    } else {
      setMsg("Something went wrong...");
    }
  };

  return (
    <MainWrapper>
      <section className="section">
        <RoomCreator onSubmit={handleSubmit} />
        <p className="text-xl my-2">{msg}</p>
      </section>
    </MainWrapper>
  );
};
