import vape from "../assets/vape.png";
import basket from "../assets/basket.png";
import chat from "../assets/chat.png";
import { useAuthContext } from "../context/useContext";
import { useEffect } from "react";
import axios from "axios";

export const ProtectedPage = ({ children }: { children: React.ReactNode }) => {
  const { user, setUser } = useAuthContext();
  useEffect(() => {
    if (!user) {
      (async function () {
        const result = await axios.get("/api/auth/ping");

        console.log(result.data);
      })();
    }
  }, [user]);
  return <>{user ? <>{children}</> : null}</>;
};
