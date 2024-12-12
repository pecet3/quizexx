import { useEffect } from "react";
import { MainWrapper } from "../components/MainWrapper";
import { Link } from "react-router-dom";
import { RoomsList } from "../components/RoomsList";
import { PaperWrapper } from "../components/PaperWrapper";

export const Home = () => {
  useEffect(() => {
    // Zablokowanie przewijania
    document.body.style.overflow = "hidden";

    // Przywrócenie przewijania po odmontowaniu komponentu
    return () => {
      document.body.style.overflow = "auto";
    };
  }, []);

  return (
    <MainWrapper>
      <section className="section">
        <h2 className="text-2xl ">Available Rooms</h2>
        <PaperWrapper>
          <RoomsList />
        </PaperWrapper>
      </section>
    </MainWrapper>
  );
};
