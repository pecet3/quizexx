import { MainWrapper } from "../components/MainWrapper";
import { RoomsList } from "../components/RoomsList";
import { PaperWrapper } from "../components/PaperWrapper";

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
