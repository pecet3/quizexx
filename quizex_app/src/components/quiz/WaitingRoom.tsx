import { FaCheck } from "react-icons/fa";
import { WaitingState } from "../../pages/Quiz";
import { IoCheckbox, IoCheckmark, IoClose } from "react-icons/io5";

export const WaitingRoom: React.FC<{
  waitingState: WaitingState;
  serverMessage: string;
  onReady: () => void;
}> = ({ waitingState, serverMessage, onReady }) => {
  return (
    <>
      <div className="flex flex-col justify-center items-center my-6 mt-16">
        <div className="paper paper-yellow max-w-xs w-full text-lg m-auto p-4 pt-8 shadow-md gap-2 flex flex-col items-center">
          <div className="top-tape"></div>
          <ul className="grid grid-cols-2 text-xl w-full">
            {waitingState.players.map((user, idx) => (
              <li key={idx} className="flex items-center gap-0.5">
                {user.is_ready ? (
                  <IoCheckmark size={26} />
                ) : (
                  <IoClose size={26} />
                )}
                {user.name}
              </li>
            ))}
          </ul>
          <div className="flex gap-0.5">
            <span className="text-2xl font-sans font-bold">[</span>
            <p className="text-2xl font-sans font-bold">
              {waitingState.players.filter((u) => u.is_ready).length} /
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
      <p className="text-2xl">{serverMessage}</p>
    </>
  );
};
