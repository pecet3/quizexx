export function MainView() {
  return (
    <>
      <div className="text-center m-auto bg-gray-400 ">
        <header className="p-1 my-4 relative">
          <h1
            className="text-8xl font-black	flex justify-center items-end text-center text-black font-mono 
            underline decoration-wavy decoration-4 decoration-teal-400 "
          >
            <span className="text-3xl absolute -top-0.5 right-1/2 mr-2 ">
              🎲
            </span>
            Quizex<span className="text-3xl"></span>
          </h1>
        </header>
        <div className="paper paper-yellow p-4 pt-8 shadow-md shadow-gray-700">
          <div className="top-tape"></div>
          <div
            id="entry"
            className="flex justify-center flex-col text-white text-xl gap-4"
          >
            <input
              type="text"
              id="joinRoomInput"
              className="ibm-regular p-0.5 text-2xl rounded-sm font m-auto border border-black
            bg-yellow-00 placeholder:text-gray-400 placeholder:text-center text-black text-center "
              placeholder="Nazwa pokoju"
            />
            <button
              id="connectButton"
              className="bg-purple-200 
                hover:shadow-none hover:rounded-xl border border-black hover:scale-[0.995]
                 font-mono font-semibold px-2 text-2xl duration-300 text-black rounded-lg m-auto py-1"
            >
              Dołącz
            </button>
            <p className="text-black text-center font-bold">lub...</p>
            <a
              href="/create"
              className="bg-blue-200 
            hover:shadow-none hover:rounded-xl border border-black hover:scale-[0.995]
            font-mono font-semibold px-2 text-2xl duration-300 text-black rounded-lg m-auto py-1"
            >
              Utwórz pokój
            </a>
          </div>
        </div>
      </div>
    </>
  );
}
