// src/App.tsx

import "./App.css";
import { Route, Routes, useNavigate } from "react-router-dom";
import { Home } from "./pages/Home";
import { Navbar } from "./components/Navbar";
import { How } from "./pages/How";
import { Auth } from "./pages/Auth";
import { useAuthContext } from "./context/useContext";

function App() {
  const { user, setUser } = useAuthContext();

  return (
    <>
      {!user ? (
        <>
          <Navbar />
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/how" element={<How />} />
          </Routes>
        </>
      ) : (
        <Auth />
      )}
    </>
  );
}

export default App;
