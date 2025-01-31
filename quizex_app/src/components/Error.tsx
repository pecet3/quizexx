import { Link } from "react-router-dom";

export const Error = ({ err }: { err: string }) => {
  return (
    <div className="my-64 flex flex-col gap-4 items-center">
      <p className="text-2xl">{err}</p>
      <Link to="/" className="">
        Go back
      </Link>
    </div>
  );
};

export const ErrorSomethingWentWrong = () => {
  return (
    <div className="my-64 flex flex-col gap-4 items-center">
      <p className="text-2xl">Ups...Something went wrong</p>
      <Link to="/" className="">
        Go back
      </Link>
    </div>
  );
};
