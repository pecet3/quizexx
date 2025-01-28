import React, {
  createContext,
  useContext,
  useState,
  ReactNode,
  useEffect,
} from "react";

type ProtectedContextType = {
  user: User | null;
  setUser: React.Dispatch<React.SetStateAction<User | null>>;
  funFact: FunFact | null;
  setFunFact: React.Dispatch<React.SetStateAction<FunFact | null>>;
};

const ProtectedContext = createContext<ProtectedContextType | undefined>(
  undefined
);

export const ProtectedProvider: React.FC<{ children: ReactNode }> = ({
  children,
}) => {
  const [user, setUser] = useState<User | null>(null);
  const [funFact, setFunFact] = useState<FunFact | null>(null);
  useEffect(() => {
    console.log(user);
  }, [user]);
  return (
    <ProtectedContext.Provider
      value={{
        user,
        setUser,
        funFact,
        setFunFact,
      }}
    >
      {children}
    </ProtectedContext.Provider>
  );
};

export const useProtectedContext = () => {
  const context = useContext(ProtectedContext);
  if (context === undefined) {
    throw new Error(
      "useProtectedContext must be used within a ProtectedContext"
    );
  }
  return context;
};
