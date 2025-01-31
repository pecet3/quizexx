import { RoomsList } from "../components/RoomsList";
import { PaperWrapper, PatternWrapper } from "../components/Wrappers";
import { useProtectedContext } from "../context/protectedContext";
import UserProfileCard from "../components/UserProfileCarrd";
import { TopUsersList } from "../components/TopUsersList";
import { RoomJoiner } from "../components/RoomJoiner";

export const Home = () => {
  const { funFact, user, topUsers } = useProtectedContext();
  console.log(funFact);
  return (
    <main className="flex pt-24 flex-col w-full m-auto justify-center">
      <section className="self-center md:self-end p-8 flex flex-col lg:flex-row justify-center w-full gap-12 items-center">
        <PaperWrapper>
          <RoomsList />
        </PaperWrapper>
        <div className="flex flex-col gap-8">
          <RoomJoiner />
          <UserProfileCard user={user!} />
        </div>
      </section>
      <section className="self-center md:self-end p-8 flex flex-col lg:flex-row justify-center w-full gap-8 items-center">
        <div className="max-w-md self-center m-auto">
          <p className="italic sm:text-lg">{funFact?.content}</p>{" "}
          <span className="text-sm sm:text-base flex justify-end gap-2">
            Fun fact about:
            <p className="text-right font-bold"> {funFact?.topic}</p>
          </span>
        </div>
        <PatternWrapper>
          <TopUsersList users={topUsers!} />
        </PatternWrapper>
      </section>
    </main>
  );
};
