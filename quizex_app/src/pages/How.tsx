import vape from "../assets/vape.png";
import basket from "../assets/basket.png";
import chat from "../assets/chat.png";

export const How = () => {
  return (
    <div className="flex-grow flex items-center justify-center md:py-16 -mt-6 md:-mt-10">
      <div className="w-full max-w-xs md:max-w-screen-lg px-10">
        <div className="relative bg-black bg-opacity-50 rounded-lg shadow-2xl border-4 border-black p-6">
          {/* Wewnętrzny biały border */}
          <div className="relative border-4 border-white rounded-lg p-2 md:p-8">
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
              <div className="border-4 border-purple-950 rounded-lg flex flex-col items-center gap-3 p-4">
                <img
                  src={vape}
                  className="w-16 h-16 sm:w-20 sm:h-20 object-contain"
                  alt="Step 1"
                />
                <p className="text-white text-base sm:text-lg text-center mt-2 step-title">
                  ADD TO BASKET
                </p>
              </div>
              <div className="border-4 border-purple-950 rounded-lg flex flex-col items-center gap-3 p-4">
                <img
                  src={basket}
                  className="w-16 h-16 sm:w-20 sm:h-20 object-contain"
                  alt="Step 2"
                />
                <p className="text-white text-base sm:text-lg text-center mt-2 step-title">
                  FILL UP IMPORTANT INFORMATION
                </p>
              </div>
              <div className="border-4 border-purple-950 rounded-lg flex flex-col items-center gap-3 p-4">
                <img
                  src={chat}
                  className="w-16 h-16 sm:w-20 sm:h-20 object-contain"
                  alt="Step 3"
                />
                <p className="text-white text-base sm:text-lg text-center mt-2 step-title">
                  W8 FOR YOUR JUICE
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
