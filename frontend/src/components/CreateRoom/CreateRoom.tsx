import { useState } from "react";
import { TRoomSettings } from "../../types/event";
import { useAppStateContext } from "../../custom-hooks/useAppContext";

export function CreateRoom() {
  const { appState, setAppState } = useAppStateContext();

  const settings = appState.settings
  const [roomSettings, setRoomSettings] = useState<TRoomSettings>({
    name: settings.name,
    category: settings.category,
    difficulty: settings.difficulty,
    maxRounds: settings.maxRounds,
  })

  const handleFormSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    setAppState((prev) => ({
      ...prev,
      settings: {
        name: roomSettings.name,
        category: roomSettings.category,
        maxRounds: roomSettings.maxRounds,
        difficulty: roomSettings.difficulty,
      }
    }))

  };




  return (
    <section
      className="bg-blue-300 w-[22rem] sm:w-[26rem]   text-sm sm:text-lg
     m-auto  border-r-4 border-b-4 border-r-white
     border-b-white outline outline-1 rounded-r-xl
        "
    >
      <form
        className="flex flex-col p-6 border border-black rounded-r-xl
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
          onChange={(e) => setRoomSettings((prev) => ({ ...prev, name: e.target.value }))}

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
    </section>
  );
}
