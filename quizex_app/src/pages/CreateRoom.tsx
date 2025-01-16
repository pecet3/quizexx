import { useState } from "react";
import { MainWrapper } from "../components/MainWrapper";
import { PaperWrapper } from "../components/PaperWrapper";
import { useNavigate } from "react-router-dom";
import { Settings } from "./Quiz";
import { Error } from "../components/Error";
import { RoomCreator } from "../components/RoomCreator";

export const CreateRoom = () => {
  const nav = useNavigate();
  const [isLoading, setIsLoading] = useState(false);
  const [err, setErr] = useState("");

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    setIsLoading(true);
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
      setIsLoading(false);
      return;
    }
    if (response.status === 403) {
      setErr("Room with this name already exists!");
      setIsLoading(false);
      return;
    } else {
      setErr("Something went wrong...");
      setIsLoading(false);
    }
  };

  return (
    <MainWrapper>
      <section className="section">
        {err != "" ? (
          <Error err={err} />
        ) : (
          <>
            {isLoading ? (
              <p className="my-64 text-2xl">
                Creating a game, please be patient
              </p>
            ) : (
              <RoomCreator onSubmit={handleSubmit} />
            )}
          </>
        )}
      </section>
    </MainWrapper>
  );
};
