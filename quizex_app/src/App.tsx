// src/App.tsx

import "./App.css";
import { Route, Routes, useNavigate } from "react-router-dom";
import { Home } from "./pages/Home";
import { Navbar } from "./components/Navbar";
import { Auth } from "./pages/Auth";
import { CreateRoom } from "./pages/CreateRoom";
import { ProtectedPage } from "./components/Protected";
import { Quiz } from "./pages/Quiz";
import { useAuthContext } from "./context/authContext";
import { useEffect } from "react";
import axios from "axios";

function App() {
  const { setUser } = useAuthContext();

  const navigate = useNavigate();
  useEffect(() => {
    (async function () {
      try {
        const result = await axios.get("/api/auth/ping");
        if (result.data) {
          console.log(result.data);
          setUser(result.data);
        }
      } catch (err: any) {
        navigate("/auth");
      }
    })();
  }, []);
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
