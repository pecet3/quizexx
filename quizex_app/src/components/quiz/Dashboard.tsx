import React, { useState, FormEvent, useEffect } from "react";

import { GameState, Settings } from "../../pages/Quiz";
import { LittlePaperWrapper, PaperWrapper } from "../PaperWrapper";
export const Dashboard: React.FC<{
  settings: Settings;
  gameState: GameState;
  serverMessage: string;
  timer: number;
  onAnswer: (answer: number) => void;
}> = ({ settings, gameState, serverMessage, timer, onAnswer }) => {
  const [selectedAnswer, setSelectedAnswer] = useState<number | null>(null);
  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    if (selectedAnswer !== null) {
      onAnswer(selectedAnswer);
      setSelectedAnswer(null);
    }
  };

  return (
    <div className="m-0 sm:m-auto max-w-3xl text-lg flex flex-col">
      <div className="flex justify-between gap-2 z-10 m-0">
        <div className="text-2xl flex sm:flex-row flex-col items-center font-bold font-mono bg-gray-400 rounded-t-md p-1 border-2 border-black border-b-0">
          <span className="hidden sm:block">Category:</span>
          <p className="text-blue-800 italic">{settings.gen_content}</p>
        </div>
        <div className="text-2xl flex-col sm:flex-row flex items-center gap-1 sm:gap-2 font-black font-mono bg-gray-400 rounded-t-md p-0.5 sm:p-1 px-1 sm:px-2 border-2 border-black border-b-0">
          <div className="flex items-start m-auto gap-1">
            <p>Round: </p>
            <p className="text-blue-800">{gameState.round}</p>
          </div>
        </div>
      </div>

      <form
        onSubmit={handleSubmit}
        className="flex flex-col gap-6 m-auto  w-full "
      >
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 bg-gray-200 border-2 border-black p-2 sm:p-4 rounded-b-xl text-lg">
          <h3 className="sm:col-span-2 text-center font-bold text-2xl">
            {gameState.question}
          </h3>
          {gameState.answers.map((answer, idx) => (
            <label
              key={idx}
              className={`has-[:checked]:rounded-lg has-[:checked]:scale-[1.01] ${
                ["bg-blue-200", "bg-red-200", "bg-green-200", "bg-purple-200"][
                  idx
                ]
              } duration-300 has-[:checked]:bg-gray-700 has-[:checked]:ring hover:scale-[1.025] hover:shadow-md hover:shadow-gray-400 ring-teal-500 has-[:checked]:text-white cursor-pointer flex justify-center items-center rounded-md hover:rounded-lg p-1 border border-black`}
            >
              <input
                type="radio"
                name="answer"
                value={idx}
                checked={selectedAnswer === idx}
                onChange={() => setSelectedAnswer(idx)}
                className="hidden"
              />
              <p>{answer}</p>
            </label>
          ))}
        </div>

        <p className="font-mono text-xl sm:text-2xl font-bold italic bg-white p-2 rounded-xl bg-opacity-70">
          {serverMessage}
        </p>

        <div className="grid grid-cols-1 sm:grid-cols-2 gap-6">
          <div className="sm:order-1  flex flex-col items-center justify-center gap-0.5">
            <p className="font-bold text-2xl">Time Left: {timer}</p>
            <button
              type="submit"
              className="bg-teal-300 hover:scale-[1.025] hover:shadow-lg hover:shadow-gray-500 hover:rounded-xl border-2 border-black font-mono font-semibold px-4 text-3xl duration-300 text-black rounded-lg m-auto py-2"
            >
              Send answer
            </button>
          </div>
          <LittlePaperWrapper>
            <table className="w-full text-xl table-fixed sm:order-none order-last">
              <thead className="">
                <tr className="border-b border-black font-mono font-black">
                  <th className="m-auto">Name</th>
                  <th className="m-auto">Points</th>
                </tr>
              </thead>
              <tbody className="">
                {gameState.score && gameState.score.length > 0
                  ? gameState.score.map((user, idx) => (
                      <tr key={idx}>
                        <td>{user.user.name}</td>
                        <td>{user.points}</td>
                      </tr>
                    ))
                  : null}
              </tbody>
            </table>
          </LittlePaperWrapper>
        </div>
      </form>
    </div>
  );
};
