import React from "react";
import { PaperWrapper } from "./PaperWrapper";

export const RoomCreator: React.FC<{
  onSubmit: (event: any) => void;
}> = ({ onSubmit }) => {
  return (
    <PaperWrapper>
      <form
        id="settingsForm"
        className="flex flex-col gap-4 items-center text-xl p-4 text-center"
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
        <p className="italic text-xl max-w-sm">Category can be anything... </p>
        <input
          type="text"
          id="categoryInput"
          name="category"
          className="p-0.5 text-2xl rounded-sm font m-auto border border-black bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
          placeholder="Quiz Category"
          required
        />
        <button id="randomCategory" className="text-sm" type="button">
          [Get random category]
        </button>
        <div className="flex font-mono sm:flex-row flex-col gap-4">
          <div className="flex flex-col">
            <label className="rounded-lg font-mono text-lg  font-bold underline">
              Rounds:
            </label>
            <select
              id="maxRounds"
              name="maxRounds"
              className="p-0.5 text-2xl rounded-sm font m-auto border border-black bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
            >
              <option value="4" className="text-center">
                3
              </option>
              <option value="5" className="text-center">
                4
              </option>
              <option value="6" className="text-center">
                5
              </option>
              <option value="6" className="text-center">
                6
              </option>
              <option value="6" className="text-center">
                7
              </option>
              <option value="6" className="text-center">
                8
              </option>
              <option value="6" className="text-center">
                9
              </option>
              <option value="6" className="text-center">
                10
              </option>
            </select>
          </div>
          <div className="flex flex-col items-center">
            <label className="rounded-lg text-sm font-bold underline w-24 break-words">
              Seconds for answer:
            </label>
            <select
              id="sec_for_answer"
              name="sec_for_answer"
              className="p-0.5 text-2xl rounded-sm font m-auto border border-black bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
            >
              <option value="10" className="text-center">
                10
              </option>
              <option value="15" className="text-center">
                15
              </option>
              <option value="20" className="text-center">
                20
              </option>
              <option value="30" className="text-center">
                30
              </option>
              <option value="45" className="text-center">
                45
              </option>
              <option value="60" className="text-center">
                60
              </option>
            </select>
          </div>
        </div>
        <div className="flex font-mono sm:flex-row flex-col gap-4 sm:items-end">
          <div className="flex flex-col">
            <label className="rounded-lg text-lg font-bold underline">
              Difficulty:
            </label>
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
              <option value="veryhard" className="text-center">
                Very Hard
              </option>
            </select>
          </div>
        </div>
        <div className="flex flex-col">
          <label className="rounded-lg font-mono text-lg font-bold underline">
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
        <button type="submit" className="btn bg-teal-300">
          Create Room
        </button>
      </form>
    </PaperWrapper>
  );
};
