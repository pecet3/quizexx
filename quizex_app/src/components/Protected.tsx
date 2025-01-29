import { useProtectedContext } from "../context/protectedContext";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

export const ProtectedPage = ({ children }: { children: React.ReactNode }) => {
  const { user, setUser, funFact, setFunFact, topUsers, setTopUsers } =
    useProtectedContext();
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
            setFunFact(data);
          }
        } catch (err: any) {
          navigate("/auth");
        }
      })();
    }
    if (!topUsers) {
      (async function () {
        try {
          const result = await fetch("/api/social/users");
          const data = await result.json();
          console.log(111, data);
          if (result.ok) {
            setTopUsers(data);
          }
        } catch (err: any) {
          navigate("/auth");
        }
      })();
    }
  }, []);
  return <>{user ? <>{children}</> : null}</>;
};
