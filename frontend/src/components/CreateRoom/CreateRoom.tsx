export function CreateRoom() {
  return (
    <>
      <body className="bg-gray-400 p-0.5 pr-2 sm:p-2 py-2 sm:py-4 m-auto my-10">
        {/* my for development */}
        <header className="p-1 my-4 relative">
          <h1
            className="text-8xl font-black	flex justify-center items-end text-center text-black font-mono 
            underline decoration-wavy decoration-4 decoration-teal-400 "
          >
            <span className="text-3xl absolute -top-0.5 right-1/2 mr-2">
              🎲
            </span>
            Quizex<span className="text-3xl"></span>
          </h1>
        </header>
        <div
          id="entryDashboard"
          className="bg-blue-300 max-w-sm sm:max-w-lg text-sm sm:text-lg
     m-auto  border-r-4 border-b-4 border-r-white
     border-b-white outline outline-1
        border-black rounded-r-2xl flex items-center"
        >
          <form
            id="settingsForm"
            className="flex flex-col py-8 px-16 border border-black rounded-r-xl
         gap-4 items-center text-xl"
          >
            <input
              type="text"
              id="nameInput"
              className="p-0.5 text-2xl rounded-sm font m-auto border border-black
                bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
              placeholder="Nazwa pokoju"
              required
            />
            <div className="italic p-2 w-80 flex flex-col items-center">
              <input
                type="text"
                id="categoryInput"
                className="p-0.5 text-2xl rounded-sm font m-auto border border-black
            bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
                placeholder="Kategoria Pytań"
                required
              />
              <p className="font-mono text-lg">
                Kategoria może być dowolona,{" "}
                <b className="font-bold underline">
                  {" "}
                  Quizex jest połączony z AI
                </b>
                . Na podstawie dostarczonej kategorii, są przygotowywane
                pytania.
              </p>
            </div>
            <label
              for="maxRounds"
              className="rounded-lg font-mono text-xl font-bold underline"
            >
              Liczba round:
            </label>
            <input
              type="number"
              id="maxRounds"
              min="0"
              max="10"
              value="5"
              className="p-0.5 text-2xl rounded-sm font m-auto border border-black
                bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
            />
            <label
              for="difficulty"
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
              id="readyButton"
              id="createRoom"
            >
              Utwórz pokój
            </button>
          </form>
        </div>
      </body>
    </>
  );
}
