import { useEffect, useState } from "react";
import { TRoomSettings } from "../../types/event";
import { IAppStateProps, TAppState } from "../../App";

export function CreateRoom({ appState, setAppState }: IAppStateProps) {
  const settings = appState.settings
  const [roomSettings, setRoomSettings] = useState<TRoomSettings>({
    name: settings.name,
    category: settings.category,
    difficulty: settings.difficulty,
    maxRounds: settings.maxRounds,
  })

  const handleFormSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    setRoomSettings({
      name: (event.currentTarget.elements.namedItem("nameInput") as HTMLInputElement).value,
      category: (event.currentTarget.elements.namedItem("categoryInput") as HTMLInputElement).value,
      maxRounds: (event.currentTarget.elements.namedItem("maxRounds") as HTMLInputElement).value,
      difficulty: (event.currentTarget.elements.namedItem("difficulty") as HTMLSelectElement).value,
    });

    setAppState({
      name: (event.currentTarget.elements.namedItem("nameInput") as HTMLInputElement).value,
      category: (event.currentTarget.elements.namedItem("categoryInput") as HTMLInputElement).value,
      maxRounds: (event.currentTarget.elements.namedItem("maxRounds") as HTMLInputElement).value,
      difficulty: (event.currentTarget.elements.namedItem("difficulty") as HTMLSelectElement).value,
    });
  };


  useEffect(() => {
    console.log(roomSettings)
    console
  }, [roomSettings])
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
        />
        <div className="italic p-2 w-80 flex flex-col items-center">
          <input
            type="text"
            name="categoryInput"
            className="p-0.5 text-2xl rounded-sm font m-auto border border-black
            bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
            placeholder="Kategoria Pytań"
            required
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
          value="5"
          className="p-0.5 text-2xl rounded-sm font m-auto border border-black
                bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
        />
        <label
          className="rounded-lg  font-mono text-xl font-bold underline"
        >
          Poziom trudności:
        </label>
        <select
          id="difficulty"
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
