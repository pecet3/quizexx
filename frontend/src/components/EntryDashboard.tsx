import React, { useState } from 'react';

export function EntryDashboard() {
    const [userName, setUserName] = useState('');

    const handleUserNameChange = (event) => {
        setUserName(event.target.value);
    };

    const handleSubmit = (event: HTMLFormElement) => {
        event.preventDefault(); // Prevent default form submission behavior
        // You can perform further actions here, such as sending data to the server
        console.log('User name submitted:', userName);
        // Clear the input field after submission
        setUserName('');
    };

    return (
        <div
            id="entryDashboard"
            className="paper paper-yellow p-4 pt-8 shadow-md flex flex-col gap-2 items-center"
        >
            <div className="top-tape"></div>
            <form onSubmit={handleSubmit}>
                <input
                    id="userNameInput"
                    type="text"
                    className="p-0.5 text-2xl rounded-sm font m-auto border border-black bg-white placeholder:text-gray-400 placeholder:text-center text-black text-center"
                    placeholder="Podaj swój nick"
                    value={userName}
                    onChange={(e) => setUserName(e.target.value)
                    }
                />
                <button
                    type="submit"
                    className="bg-purple-300 hover:shadow-none hover:rounded-xl border
                     border-black hover:scale-[0.995] font-mono px-4 text-2xl duration-300 text-black rounded-lg m-auto py-1.5"
                    id="connectButton"
                >
                    Dołącz
                </button>
            </form>
        </div>
    );
}
