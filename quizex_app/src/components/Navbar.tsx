import { Link } from "react-router-dom";
import { useAuthContext } from "../context/authContext";
import { Logo } from "./Logo";

export const Navbar = () => {
  const { user } = useAuthContext();
  return (
    <>
      {user ? (
        <nav className="flex justify-between items-center w-full px-4 py-2">
          <Logo />
          <div className="justify-end flex-1 flex">
            <div className="text-right">
              <p> Hello {user.name} </p>[
              <Link to={"/edit"} className="font-mono text-blue-700 font-bold">
                edit
              </Link>
              ]
            </div>
          </div>
          <img src={user.image_url} className=" h-16 w-16 rounded-full" />
        </nav>
      ) : (
        <nav className=" flex-1  w-full px-4 py-2">
          <Logo />
        </nav>
      )}
    </>
  );
};
