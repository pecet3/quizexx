import React, { useState, FormEvent, useEffect } from "react";
import { useParams } from "react-router-dom";
import { GameState, Settings, User, WaitingState } from "../pages/Quiz";

export const Chat: React.FC = () => {
  const [message, setMessage] = useState("");
  const [messages, setMessages] = useState<string[]>([]);

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    if (message.trim()) {
      setMessages([...messages, message]);
      setMessage("");
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="flex-col justify-between sm:right-2 gap-1 p-1 flex bg-second-paper rounded-b-lg rounded-r-lg border-2 border-black bg-slate-200 w-96 sm:w-[30rem] z-50 m-auto text-sm my-16"
    >
      <div className="flex justify-between text-xl absolute">
        <div className="text-2xl bg-gray-400 border-2 border-black py-0.5 rounded-t-md px-2 font-mono relative bottom-[2.75rem] right-[0.35rem] font-bold">
          Chat<span className="text-lg text-teal-500">💬</span>
        </div>
      </div>
      <ul className="flex flex-col gap-1 h-64 sm:h-80 break-words overflow-y-scroll text-sm sm:text-base p-0.5 border-b border-gray-400">
        {messages.map((msg, idx) => (
          <li key={idx}>{msg}</li>
        ))}
      </ul>
      <div className="flex gap-2 m-auto justify-between">
        <input
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          className="p-1 rounded-md border border-black sm:w-[24rem] w-72"
        />
        <button
          type="submit"
          className="bg-teal-300 hover:rounded-xl border-2 border-black font-mono font-semibold px-2 text-xl duration-300 text-black rounded-lg m-auto py-1"
        >
          Send
        </button>
      </div>
    </form>
  );
};

// Waiting Room Component
export const WaitingRoom: React.FC<{
  waitingState: WaitingState;
  onReady: () => void;
}> = ({ waitingState, onReady }) => {
  return (
    <div className="flex flex-col justify-center items-center my-6">
      <div className="paper paper-yellow max-w-md text-lg m-auto p-4 pt-8 shadow-md flex flex-col items-center">
        <div className="top-tape"></div>
        <ul className="grid grid-cols-2 text-xl">
          {waitingState.players.map((user, idx) => (
            <li key={idx}>{user.name}</li>
          ))}
        </ul>
        <div className="flex gap-0.5">
          <span className="text-2xl font-sans font-bold">[</span>
          <p className="text-2xl font-sans font-bold">
            {waitingState.players.length}
          </p>
          <span className="text-2xl font-sans font-bold">]</span>
        </div>
        <button
          onClick={onReady}
          className="bg-teal-300 hover:shadow-none hover:rounded-xl border border-black hover:scale-[0.995] font-mono px-4 text-2xl duration-300 text-black rounded-lg m-auto py-2"
        >
          I am ready
        </button>
      </div>
    </div>
  );
};

// Game Dashboard Component
export const GameDashboard: React.FC<{
  settings: Settings;
  gameState: GameState;
  users: User[];
  onAnswer: (answer: number) => void;
}> = ({ settings, gameState, users, onAnswer }) => {
  const [selectedAnswer, setSelectedAnswer] = useState<number | null>(null);

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    if (selectedAnswer !== null) {
      onAnswer(selectedAnswer);
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
              } duration-300 has-[:checked]:bg-gray-700 has-[:checked]:ring ring-teal-500 has-[:checked]:text-white cursor-pointer flex justify-center items-center rounded-md hover:rounded-lg p-1 border border-black`}
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
          Have a good game!
        </p>

        <div className="grid grid-cols-1 sm:grid-cols-2 gap-2">
          <table className="text-xl table-fixed sm:order-none order-last bg-yellow-200 shadow-md shadow-gray-600">
            <thead className="m-auto">
              <tr className="flex justify-center border-b border-black font-mono font-black">
                <th className="m-auto">Name</th>
                <th className="m-auto">Points</th>
              </tr>
            </thead>
            <tbody className="flex flex-col [&_tr]:py-2 [&_tr]:gap-4 [&_tr]:flex [&_tr]:justify-between [&_td]:m-auto">
              {users.map((user, idx) => (
                <tr key={idx}>
                  <td>{user.name}</td>
                  <td>{user.points}</td>
                </tr>
              ))}
            </tbody>
          </table>
          <div className="flex flex-col items-center justify-center gap-0.5">
            <button
              type="submit"
              className="bg-teal-300 hover:rounded-xl border-2 border-black font-mono font-semibold px-4 text-3xl duration-300 text-black rounded-lg m-auto py-2"
            >
              Send the answer
            </button>
          </div>
        </div>
      </form>
    </div>
  );
};
