import { useEffect } from "react";
import { MainWrapper } from "../components/MainWrapper";
import { Link } from "react-router-dom";
import { RoomsList } from "../components/RoomsList";
import { PaperWrapper } from "../components/PaperWrapper";
import { useAuthContext } from "../context/useContext";

export const Home = () => {
  return (
    <MainWrapper>
      <section className="section">
        <h2 className="text-2xl ml-2">Available Rooms</h2>
        <PaperWrapper>
          <RoomsList />
        </PaperWrapper>
      </section>
    </MainWrapper>
  );
};
