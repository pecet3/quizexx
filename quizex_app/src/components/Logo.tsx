import { Link } from "react-router-dom";

export const Logo = () => {
  return (
    <h1 className="flex-none font-mono text-4xl underline decoration-teal-300 decoration-wavy font-bold">
      <Link to={"/home"}> Quizex</Link>
    </h1>
  );
};
