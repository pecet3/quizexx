
import { useAppStateContext } from "../../custom-hooks/useAppContext";


export function MainView() {
  const { appState } = useAppStateContext();


  const handleJoinRoom = () => {
    console.log("Room Name:", appState.settings.name);
    console.log(appState)
  };

  console.log(appState.settings)

  return (

    <div className="text-center m-auto bg-gray-400  w-[20rem] ">

      <div className="paper paper-yellow p-4 pt-8 shadow-md shadow-gray-700">
        <div className="top-tape"></div>
        <form
          onSubmit={handleJoinRoom}
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
        </form>
      </div>
    </div>

  );
}
