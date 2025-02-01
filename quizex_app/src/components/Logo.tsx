import { Link } from "react-router-dom";

export const Logo = () => {
  return (
    <h1 className="flex-none font-mono text-4xl underline decoration-teal-300 decoration-wavy font-bold">
      <Link to={"/home"}> Quizex</Link>
      {/* <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 300 100">
        <text
          x="20"
          y="60"
          font-family="monospace"
          font-size="48"
          font-weight="bold"
          fill="black"
        >
          Qu&#x131;zex
        </text>

        <text
          x="83"
          y="30"
          font-family="monospace"
          font-size="24"
          font-weight="bold"
          fill="black"
        >
          ?
        </text>
        <path
          d="M20 70 C 35 55, 50 85, 65 70 C 80 55, 95 85, 110 70 C 125 55, 140 85, 155 70 C 170 55, 185 85, 200 70"
          stroke="#00ffff"
          stroke-width="3"
          fill="none"
          stroke-linecap="round"
        />
      </svg> */}
    </h1>
  );
};
