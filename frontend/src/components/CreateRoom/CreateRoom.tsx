import { useEffect, useState } from "react";
import { TRoomSettings } from "../../types/event";
import { useAppStateContext } from "../../custom-hooks/useAppContext";
import { useWebSocket } from "../../custom-hooks/useWebSocket";
import { NameInput } from "../NameInput";
import { Router, useNavigate } from "react-router-dom";

export function CreateRoom() {
  const navigate = useNavigate();


  const { appState, setAppState } = useAppStateContext();
  const { socket, createSocket } = useWebSocket();
  const settings = appState.settings
  const [displayNameInput, setDisplayNameInput] = useState(false)

  const [publicUserName, setPublicUserName] = useState('');


  const [roomSettings, setRoomSettings] = useState<TRoomSettings>({
    roomName: settings.roomName,
    category: settings.category,
    difficulty: settings.difficulty,
    maxRounds: settings.maxRounds,
  })


  const handleFormSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setAppState((prev) => ({
      ...prev,
      settings: {
        roomName: roomSettings.roomName,
        category: roomSettings.category,
        maxRounds: roomSettings.maxRounds,
        difficulty: roomSettings.difficulty,
      }
    }))
    setDisplayNameInput(true)
  };

  useEffect(() => {

    createSocket(true)

  }, [publicUserName])


  useEffect(() => {
    if (socket !== null) {
      console.log(socket)
      navigate("/room");
    }
  }, [socket]);
  return (
    <>
      {!displayNameInput ? <section
        className="paper paper-yellow max-w-sm sm:max-w-lg text-sm sm:text-lg p-4 my-4
      m-auto  flex items-center"
      >
        <div className="tape-section"></div>
        <form
          className="flex flex-col p-6 rounded-r-xl
         gap-4 items-center text-xl"
          onSubmit={handleFormSubmit}

        >

          <input
            type="text"
            name="nameInput"
            className="p-0.5 text-2xl rounded-sm font m-auto border border-black
                bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
            placeholder="Nazwa pokoju"
            required
            onChange={(e) => setRoomSettings((prev) => ({ ...prev, roomName: e.target.value }))}

          />
          <div className="italic p-2 w-80 flex flex-col items-center">
            <input
              type="text"
              name="categoryInput"
              className="p-0.5 text-2xl rounded-sm font m-auto border border-black
            bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
              placeholder="Kategoria Pytań"
              required
              onChange={(e) => setRoomSettings((prev) => ({ ...prev, category: e.target.value }))}
            />
            <p className="font-mono text-lg">
              Kategoria może być dowolona,
              <b className="font-bold underline">
                Quizex jest połączony z AI
              </b>
              . Na podstawie dostarczonej kategorii, są przygotowywane
              pytania.
            </p>
          </div>
          <label
            className="rounded-lg font-mono text-xl font-bold underline"
          >
            Liczba round:
          </label>
          <input
            type="number"
            name="maxRounds"
            min="0"
            max="10"
            value={roomSettings.maxRounds}
            className="p-0.5 text-2xl rounded-sm font m-auto border border-black
                bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
            onChange={(e) => setRoomSettings((prev) => ({ ...prev, maxRounds: e.target.value }))}
          />
          <label
            className="rounded-lg  font-mono text-xl font-bold underline"
          >
            Poziom trudności:
          </label>
          <select
            onChange={(e) => setRoomSettings((prev) => ({ ...prev, difficulty: e.target.value }))}
            value={roomSettings.difficulty}
            name="difficulty"
            className="p-0.5 text-2xl rounded-sm font m-auto border border-black
                bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
          >
            <option value="easy" className="text-center">
              Łatwy
            </option>
            <option value="medium" className="text-center">
              Średni
            </option>
            <option value="hard" className="text-center">
              Trudny
            </option>
            x
          </select>
          <button
            type="submit"
            className="bg-teal-300 my-4 hover:shadow-none hover:rounded-xl border border-black hover:scale-[0.995]
            font-mono font-semibold px-2 text-2xl duration-300 text-black rounded-lg m-auto py-1"

          >
            Utwórz pokój
          </button>
        </form>
        <div className="tape-section"></div>
      </section> : <NameInput setPublicNameInput={setPublicUserName} />}
    </>

  );
}
