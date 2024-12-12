// src/App.tsx

import "./App.css";
import { Route, Routes } from "react-router-dom";
import { Home } from "./pages/Home";
import { Navbar } from "./components/Navbar";
import { Auth } from "./pages/Auth";
import { useAuthContext } from "./context/useContext";
import { CreateRoom } from "./pages/CreateRoom";
import { ProtectedPage } from "./components/Protected";

function App() {
  const { user, setUser } = useAuthContext();

  return (
    <>
      {true ? (
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
                  <Home />
                </ProtectedPage>
              }
            />
            <Route path="/auth" element={<Auth />} />
          </Routes>
        </>
      ) : (
        <Auth />
      )}
    </>
  );
}

export default App;
