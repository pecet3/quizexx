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
      <section className="mt-24 flex flex-col gap-2 items-end max-w-2xl m-auto">
        <div className="">
          <Link to={"/test"} className="btn bg-red-300">
            Create a Room
          </Link>
        </div>
        <PaperWrapper>
          <RoomsList />
        </PaperWrapper>
      </section>
    </MainWrapper>
  );
};
