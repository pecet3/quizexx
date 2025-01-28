import { useProtectedContext } from "../context/protectedContext";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

export const ProtectedPage = ({ children }: { children: React.ReactNode }) => {
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
    if (!funFact) {
      (async function () {
        try {
          const result = await fetch("/api/social/fun-facts/latest");
          const data = await result.json();
          if (result.ok) {
            console.log(data);
            setFunFact(data);
            console.log(funFact);
          }
        } catch (err: any) {}
      })();
    }
  }, []);
  return <>{user ? <>{children}</> : null}</>;
};
