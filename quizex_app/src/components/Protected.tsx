import { useAuthContext } from "../context/authContext";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

export const ProtectedPage = ({ children }: { children: React.ReactNode }) => {
  const { user, setUser } = useAuthContext();

  const navigate = useNavigate();
  useEffect(() => {
    if (!user) {
      (async function () {
        try {
          const result = await fetch("/api/auth/ping");
          const data = await result.json;
          if (result.ok) {
            setUser(data);
          }
        } catch (err: any) {
          navigate("/auth");
        }
      })();
    }
  }, []);
  return <>{user ? <>{children}</> : null}</>;
};
