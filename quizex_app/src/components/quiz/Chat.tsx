import { FormEvent, useState } from "react";
import { ChatMessage } from "../../pages/Quiz";

export const Chat: React.FC<{
  onMessage: (message: string) => void;
  messages: ChatMessage[];
}> = ({ onMessage, messages }) => {
  const [message, setMessage] = useState("");
  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    onMessage(message);
    setMessage("");
  };

  return (
    <div className="flex-col justify-between sm:right-2 gap-1 p-1 flex bg-second-paper rounded-b-lg rounded-r-lg border-2 border-black bg-slate-200 w-96 sm:w-[30rem] z-50 m-auto text-sm my-16">
      <div className="flex justify-between text-xl absolute">
        <div className="text-2xl bg-gray-400 border-2 border-black py-0.5 rounded-t-md px-2 font-mono relative bottom-[2.75rem] right-[0.35rem] font-bold">
          Chat<span className="text-lg text-teal-500">💬</span>
        </div>
      </div>
      <ul className="flex flex-col gap-1 h-64 sm:h-80 break-words overflow-y-scroll text-sm sm:text-base p-0.5 border-b border-gray-400">
        {messages.map((msg, idx) => (
          <li
            key={idx}
            className="flex flex-col justify-center w-full bg-white p-0.5 rounded-lg"
          >
            <span className="flex items-center justify-between font-mono">
              <p className="self-start font-bold">{msg.name}</p>
              <p className="self-start italic">{msg.date.toString()}</p>
            </span>
            <p>{msg.message}</p>
          </li>
        ))}
      </ul>
      <div className="flex gap-2 m-auto justify-between">
        <input
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          className="p-1 rounded-md border bg-white border-black sm:w-[24rem] w-72"
        />
        <button
          type="submit"
          onClick={handleSubmit}
          className="bg-teal-300 hover:rounded-xl border-2 border-black font-mono font-semibold px-2 text-xl duration-300 text-black rounded-lg m-auto py-1"
        >
          Send
        </button>
      </div>
    </div>
  );
};
