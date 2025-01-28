import "./App.css";
import { Route, Routes, useNavigate } from "react-router-dom";
import { Home } from "./pages/Home";
import { Navbar } from "./components/Navbar";
import { Auth } from "./pages/Auth";
import { CreateRoom } from "./pages/CreateRoom";
import { ProtectedPage } from "./components/Protected";
import { Quiz } from "./pages/Quiz";
import { useProtectedContext } from "./context/protectedContext";
import { useEffect } from "react";
import { Welcome } from "./pages/Welcome";

function App() {
  const { user, setUser, funFact, setFunFact } = useProtectedContext();
  const navigate = useNavigate();
  useEffect(() => {
    if (!user) {
      (async function () {
        try {
          const result = await fetch("/api/auth/ping");
          const data = await result.json();
          if (result.ok) {
            setUser(data);
          }
        } catch (err: any) {
          navigate("/auth");
        }
      })();
    }
  }, []);

  return (
    <>
      <Routes>
        <Route path="/" element={<Welcome />} />
        <Route
          path="/home"
          element={
            <ProtectedPage>
              <Navbar />

              <Home />
            </ProtectedPage>
          }
        />
        <Route
          path="/create-room"
          element={
            <ProtectedPage>
              <Navbar />

              <CreateRoom />
            </ProtectedPage>
          }
        />
        <Route
          path="/quiz/:roomName"
          element={
            <ProtectedPage>
              <Navbar />

              <Quiz />
            </ProtectedPage>
          }
        />

        <Route
          path="/auth"
          element={
            <>
              <Navbar />

              <Auth />
            </>
          }
        />
      </Routes>
    </>
  );
}

export default App;
