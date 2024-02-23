import { useState } from "react";
import "./App.css";
import { CreateRoom } from "./components/CreateRoom/CreateRoom";
import { Header } from "./components/Header";
import { MainView } from "./components/MainView/MainView";
import { Room } from "./components/Room/Room";
import { TGameState, TRoomSettings } from "./types/event";
export type TAppState = {

  ui: any
  settings: TRoomSettings,
  gameState: TGameState,

};
export interface IAppStateProps {
  appState: TAppState;
  setAppState: React.Dispatch<React.SetStateAction<TAppState | any>>; // Adjust the type accordingly
}
function App() {
  const [appState, setAppState] = useState({
    ui: {/* initial UI state */ },
    settings: {
      name: "Error Room",
      category: "",
      difficulty: "easy",
      maxRounds: "5",
    },
    gameState: {
      round: 1,
      question: "",
      answers: [],
      actions: [],
      score: [],
      playersFinished: [],
    },
  });

  return (
    <>
      <Header />
      <MainView appState={appState} setAppState={setAppState} />

      <CreateRoom appState={appState} setAppState={setAppState} />
      <Room appState={appState} setAppState={setAppState} />
    </>
  );
}

export default App;
