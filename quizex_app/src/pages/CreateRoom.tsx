import { useEffect } from "react";
import { MainWrapper } from "../components/MainWrapper";
import { PaperWrapper } from "../components/PaperWrapper";
import { QuizSettings } from "../types";
import { useNavigate } from "react-router-dom";

export const CreateRoom = () => {
  const nav = useNavigate();

  useEffect(() => {
    document.body.style.overflow = "hidden";

    return () => {
      document.body.style.overflow = "auto";
    };
  }, []);

  const handleSubmit = async (event: any) => {
    event.preventDefault();

    const formData = new FormData(event.target);
    const maxRoundsStr = formData.get("maxRounds");
    const maxRounds = parseInt(maxRoundsStr as string);
    const data: QuizSettings = {
      name: formData.get("roomName") as string,
      gen_content: formData.get("category") as string,
      difficulty: formData.get("difficulty") as string,
      max_rounds: maxRounds,
      language: formData.get("lang") as string,
    };
    try {
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
      } else {
        console.error("Failed to create room", response.statusText);
      }
    } catch (error) {
      console.error("Error creating room:", error);
    }
  };

  return (
    <MainWrapper>
      <section className="section">
        <PaperWrapper>
          <form
            id="settingsForm"
            className="flex flex-col gap-4 items-center text-xl p-4"
            onSubmit={handleSubmit}
          >
            <input
              type="text"
              id="nameInput"
              name="roomName"
              className="p-0.5 text-2xl rounded-sm font m-auto border border-black bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
              placeholder="Room Name"
              required
            />
            <input
              type="text"
              id="categoryInput"
              name="category"
              className="p-0.5 text-2xl rounded-sm font m-auto border border-black bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
              placeholder="Category of Questions"
              required
            />
            <button id="randomCategory" className="text-sm py-2" type="button">
              [Get random category]
            </button>
            <p className="font-mono text-lg max-w-sm">
              Category can be anything,{" "}
              <b className="font-bold underline">
                Quizex is connected with Chat-GPT-3.5
              </b>
              . Based on the provided category, questions are prepared.
            </p>
            <label className="rounded-lg font-mono text-xl font-bold underline">
              Difficulty Level:
            </label>{" "}
            <select
              id="difficulty"
              name="difficulty"
              className="p-0.5 text-2xl rounded-sm font m-auto border border-black bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
            >
              <option value="easy" className="text-center">
                Easy
              </option>
              <option value="medium" className="text-center">
                Medium
              </option>
              <option value="hard" className="text-center">
                Hard
              </option>
            </select>
            <div className="flex sm:flex-row flex-col gap-4">
              <div className="flex flex-col">
                <label className="rounded-lg font-mono text-xl font-bold underline">
                  Rounds:
                </label>
                <select
                  id="maxRounds"
                  name="maxRounds"
                  className="p-0.5 text-2xl rounded-sm font m-auto border border-black bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
                >
                  <option value="4" className="text-center">
                    4
                  </option>
                  <option value="5" className="text-center">
                    5
                  </option>
                  <option value="6" className="text-center">
                    6
                  </option>
                </select>
              </div>
              <div className="flex flex-col">
                <label className="rounded-lg font-mono text-xl font-bold underline">
                  Language:
                </label>
                <select
                  id="lang"
                  name="lang"
                  className="p-0.5 text-2xl rounded-sm font m-auto border border-black bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
                >
                  <option value="polish" className="text-center">
                    Polish
                  </option>
                  <option value="english" className="text-center">
                    English
                  </option>
                </select>
              </div>
            </div>
            <button type="submit" className="btn bg-teal-300">
              Create Room
            </button>
          </form>
        </PaperWrapper>
      </section>
    </MainWrapper>
  );
};
