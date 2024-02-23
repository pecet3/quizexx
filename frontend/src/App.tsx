import { useState } from "react";
import "./App.css";
import { CreateRoom } from "./components/CreateRoom/CreateRoom";
import { Header } from "./components/Header";
import { MainView } from "./components/MainView/MainView";
import { Room } from "./components/Room/Room";
import { TGameState, TRoomSettings } from "./types/event";
import { AppStateProvider } from "./custom-hooks/useAppContext";

function App() {


  return (
    <>
      <AppStateProvider>     <Header />
        <MainView />

        <CreateRoom />
        <Room /></AppStateProvider>

    </>
  );
}

export default App;
