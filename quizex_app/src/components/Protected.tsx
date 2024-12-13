import vape from "../assets/vape.png";
import basket from "../assets/basket.png";
import chat from "../assets/chat.png";
import { useAuthContext } from "../context/useContext";
import { useEffect } from "react";
import axios from "axios";
import { redirect, useNavigate } from "react-router-dom";

export const ProtectedPage = ({ children }: { children: React.ReactNode }) => {
  const { user, setUser } = useAuthContext();

  const navigate = useNavigate();
  useEffect(() => {
    if (!user) {
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
    }
  }, []);
  return <>{user ? <>{children}</> : null}</>;
};
