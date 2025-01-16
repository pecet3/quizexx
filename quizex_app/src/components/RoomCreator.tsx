import React from "react";
import { PaperWrapper } from "./PaperWrapper";

export const RoomCreator: React.FC<{
  onSubmit: (event: any) => void;
}> = ({ onSubmit }) => {
  return (
    <PaperWrapper>
      <form
        id="settingsForm"
        className="flex flex-col gap-4 items-center text-xl p-4"
        onSubmit={onSubmit}
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
  );
};
