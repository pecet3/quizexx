import { useState } from "react";
import "./App.css";
import { CreateRoom } from "./components/CreateRoom/CreateRoom";
import { Header } from "./components/Header";
import { MainView } from "./components/MainView/MainView";
import { Room } from "./components/Room/Room";
import { TGameState, TRoomSettings } from "./types/event";
import { AppStateProvider } from "./custom-hooks/useAppContext";
import {
  createBrowserRouter,
  RouterProvider,

} from "react-router-dom";
const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <>
        <Header />
        <MainView />
      </>
    ),
  },
  {
    path: "/create",
    element: (
      <>
        <Header />
        <CreateRoom />
      </>
    ),
  },
  {
    path: "room",
    element: (
      <>
        <Header />
        <Room />
      </>
    ),
  },
]);

function App() {


  return (
    <>
      <AppStateProvider>
        <RouterProvider router={router} />

      </AppStateProvider>

    </>
  );
}

export default App;
