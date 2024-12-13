// src/App.tsx

import "./App.css";
import { Route, Routes } from "react-router-dom";
import { Home } from "./pages/Home";
import { Navbar } from "./components/Navbar";
import { Auth } from "./pages/Auth";
import { CreateRoom } from "./pages/CreateRoom";
import { ProtectedPage } from "./components/Protected";
import { Quiz } from "./pages/Quiz";

function App() {
  return (
    <>
      <Navbar />
      <Routes>
        <Route
          path="/"
          element={
            <ProtectedPage>
              <Home />
            </ProtectedPage>
          }
        />
        <Route
          path="/create-room"
          element={
            <ProtectedPage>
              <CreateRoom />
            </ProtectedPage>
          }
        />
        <Route
          path="/quiz"
          element={
            <ProtectedPage>
              <Quiz />
            </ProtectedPage>
          }
        />
        <Route path="/auth" element={<Auth />} />
      </Routes>
    </>
  );
}

export default App;
