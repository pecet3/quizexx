import { useEffect } from "react";
import { MainWrapper } from "../components/MainWrapper";
import { Link } from "react-router-dom";
import { RoomsList } from "../components/RoomsList";
import { PaperWrapper } from "../components/PaperWrapper";

export const Home = () => {
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
