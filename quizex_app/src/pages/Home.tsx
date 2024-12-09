import { useEffect } from "react";
import { MainWrapper } from "../components/MainWrapper";
import { Link } from "react-router-dom";

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
      <div className="flex flex-col items-center justify-center min-h-screen">
        <div className="mt-0 md:mt-10 mb-80 text-center">
          <div className="text-5xl text-7xl md:text-8xl font-bold leading-tight text-white step-title">
            PROJECT JUICE
          </div>
          <span className="font-semibold text-purple-950 text-4xl md:text-4xl block gradient-text step-title mt-2">
            Level Up Your Vaping
          </span>
          <p className="text-lg sm:text-xl md:text-xl text-white font-semibold mt-2 step-title max-w-xs sm:max-w-md md:max-w-lg mx-auto">
            premium vape oils,
          </p>
          <p className="text-lg sm:text-xl md:text-xl text-white font-semibold mt-2 step-title max-w-xs sm:max-w-md md:max-w-lg mx-auto">
            top-quality flavors
          </p>

          <div className="mt-4 text-white">
            <Link
              to="/ProductsTypes"
              className="rounded-3xl py-2 md:py-3 px-4 sm:px-6 md:px-8 font-medium inline-block 
                hover:bg-transparent hover:border-white hover:text-white duration-300 hover:border-4 border-4 border-purple-950 step-title"
            >
              Order Now
            </Link>
          </div>
        </div>
      </div>
    </MainWrapper>
  );
};
