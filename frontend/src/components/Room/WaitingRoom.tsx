export const WaitingRoom = () => {
    return (
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

    )
}