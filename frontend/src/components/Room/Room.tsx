export function Room() {
  return (
    <div className="bg-gray-400 pl-0 pr-2 p-0 sm:p-8 py-2 sm:py-10 text-center m-auto">
      <div
        id="entryDashboard"
        className="paper paper-yellow p-4 pt-8 shadow-md flex flex-col gap-2 items-center"
      >
        <div className="top-tape"></div>
        <input
          id="userNameInput"
          type="text"
          className="p-0.5 text-2xl rounded-sm font m-auto border border-black bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
          placeholder="Podaj swój nick"
        />
        <button
          className="bg-purple-300 hover:shadow-none hover:rounded-xl border border-black hover:scale-[0.995] font-mono px-4 text-2xl duration-300 text-black rounded-lg m-auto py-1.5"
          id="connectButton"
        >
          Dołącz
        </button>
      </div>

      <div
        id="waitingRoomDashboard"
        className="flex flex-col justify-center items-center"
      >
        <div className="paper paper-yellow max-w-md text-lg m-auto p-4 pt-8 shadow-md flex flex-col items-center">
          <div className="top-tape"></div>

          <ul className="grid grid-cols-2 text-xl" id="readyUsersList"></ul>

          <div className="flex gap-0.5">
            <span className="text-2xl font-sans font-bold">[</span>
            <p
              id="displayReadyCount"
              className="text-2xl font-sans font-bold"
            ></p>
            <span className="text-2xl font-sans font-bold">]</span>
          </div>

          <button
            className="bg-teal-300 hover:shadow-none hover:rounded-xl border border-black hover:scale-[0.995] font-mono px-4 text-2xl duration-300 text-black rounded-lg m-auto py-2"
            id="readyButton"
          >
            Jestem gotowy
          </button>
        </div>
        <p
          id="displayServerMessageWaiting"
          className="font-mono text-lg font-bold"
        ></p>
      </div>

      <div
        id="gameDashboard"
        className="m-0 sm:m-auto max-w-3xl text-lg shadow-xl shadow-slate-700 bg-pattern h-screen p-2 pr-0 sm:pr-4 pl-0 sm:pl-12 rounded-r-2xl relative"
      >
        <p className="hidden sm:block absolute top-1/4 left-0 ml-1 sm:ml-2 text-5xl text-gray-600">
          ●
        </p>
        <p className="hidden sm:block absolute top-3/4 left-0 ml-1 sm:ml-2 text-5xl text-gray-600">
          ●
        </p>

        <header className="p-1 my-4 relative">
          <h1 className="text-8xl font-black flex justify-center items-end text-center text-black font-mono underline decoration-wavy decoration-4 decoration-teal-500 ">
            Quizex
          </h1>
        </header>

        <div className="flex justify-between gap-2 z-10">
          <div className="text-2xl flex sm:flex-row flex-col items-center font-bold font-mono bg-gray-400 rounded-t-md p-1 border-2 border-black border-b-0">
            <span className="hidden sm:block">Kategoria:</span>
            <span id="displayCategory" className="text-blue-800 italic">
              -
            </span>
          </div>
          <div className="text-2xl flex-col sm:flex-row flex items-center gap-1 sm:gap-2 font-black font-mono bg-gray-400 rounded-t-md p-0.5 sm:p-1 px-1 sm:px-2 border-2 border-black border-b-0">
            <div className="flex items-start m-auto gap-1">
              <p>Runda: </p>
              <p id="displayRound" className="text-blue-800">
                1
              </p>
            </div>
          </div>
        </div>

        <div className="flex flex-col gap-6 items-center">
          <form
            id="gameForm"
            className="flex flex-col gap-6 w-full md:w-[44rem] h-64"
          >
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 bg-gray-200 border-2 border-black p-2 sm:p-4 rounded-b-xl text-lg font-medium">
              <h3
                id="displayQuestion"
                className="sm:col-span-2 text-center font-bold text-2xl"
              ></h3>
              <label
                htmlFor="a1"
                className="answer-option bg-blue-200 cursor-pointer flex justify-center items-center rounded-md hover:rounded-lg p-1 border border-black"
              >
                <input
                  type="radio"
                  id="a1"
                  name="q1"
                  value="0"
                  className="hidden"
                />
                <p id="answerA"></p>
              </label>
              <label
                htmlFor="a2"
                className="answer-option bg-red-200 cursor-pointer flex justify-center items-center rounded-md hover:rounded-lg p-1 border border-black"
              >
                <input
                  type="radio"
                  id="a2"
                  name="q1"
                  value="1"
                  className="hidden"
                />
                <p id="answerB"></p>
              </label>
              <label
                htmlFor="a3"
                className="answer-option bg-green-200 cursor-pointer flex justify-center items-center rounded-md hover:rounded-lg p-1 border border-black"
              >
                <input
                  type="radio"
                  id="a3"
                  name="q1"
                  value="2"
                  className="hidden"
                />
                <p id="answerC"></p>
              </label>
              <label
                htmlFor="a4"
                className="answer-option bg-purple-200 cursor-pointer flex justify-center items-center rounded-md hover:rounded-lg p-1 border border-black"
              >
                <input
                  type="radio"
                  id="a4"
                  name="q1"
                  value="3"
                  className="hidden"
                />
                <p id="answerD"></p>
              </label>
            </div>
            <p
              id="displayServerMessageDashboard"
              className="font mono text-xl font-bold"
            >
              ...
            </p>
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-2 ">
              <table
                id="scoreTable"
                className="text-xl table-fixed sm:order-none order-last bg-yellow-200 shadow-md shadow-gray-600"
              >
                <thead className="m-auto">
                  <tr className="flex justify-center border-b border-black font-mono font-black">
                    <th className="m-auto">Imię</th>
                    <th className="m-auto">Punkty</th>
                  </tr>
                </thead>
                <tbody
                  id="scoreTableBody"
                  className="flex flex-col [&_tr]:py-2 [&_tr]:gap-4 [&_tr]:flex [&_tr]:justify-between [&_td]:m-auto "
                ></tbody>
              </table>
              <div className="flex flex-col items-center justify-center gap-0.5">
                <button
                  className="bg-teal-300 hover:rounded-xl border-2 border-black font-mono font-semibold px-4 text-3xl duration-300 text-black rounded-lg m-auto py-2"
                  id="submitAnswerButton"
                >
                  Wyślij Odpowiedź
                </button>
                <div className="flex gap-0.5">
                  <span className="text-lg font-sans font-bold">[</span>
                  <p
                    id="displayAnswered"
                    className="text-xl font-sans font-bold"
                  ></p>
                  <span className="text-lg font-sans font-bold">]</span>
                </div>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}
