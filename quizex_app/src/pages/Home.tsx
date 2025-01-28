import { RoomsList } from "../components/RoomsList";
import { PaperWrapper } from "../components/PaperWrapper";
import { useProtectedContext } from "../context/protectedContext";
import UserProfileCard from "../components/UserProfileCarrd";

export const Home = () => {
  const { funFact, user } = useProtectedContext();
  console.log(funFact);
  return (
    <main className="flex pt-24 flex-col w-full m-auto justify-center gap-16">
      <section className="">
        <PaperWrapper>
          <RoomsList />
        </PaperWrapper>
      </section>
      <section className="self-center md:self-end p-8 flex flex-col lg:flex-row justify-between w-full gap-12 items-center">
        <div className="flex-1 flex justify-center">
          <UserProfileCard user={user!} />
        </div>
        <div className="max-w-md">
          <p className="italic">{funFact?.content}</p>{" "}
          <p className="text-right">{funFact?.topic}</p>
        </div>
      </section>
    </main>
  );
};
