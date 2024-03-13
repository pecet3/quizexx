export function EntryDashboard() {
    return (

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
                className="bg-purple-300 hover:shadow-none hover:rounded-xl border
                 border-black hover:scale-[0.995] font-mono px-4 text-2xl duration-300 text-black rounded-lg m-auto py-1.5"
                id="connectButton"
            >
                Dołącz
            </button>

        </div>
    )
}