import { TypeAnimation } from "react-type-animation";
import { motion } from "framer-motion";
import { FaQuestion } from "react-icons/fa";
import { FaRegCirclePlay } from "react-icons/fa6";
import { Link } from "react-router-dom";
export const Welcome = () => {
  return (
    <main className="h-screen flex items-center justify-center">
      <section
        className="m-auto flex flex-col 
      gap-8 sm:gap-12 items-center justify-center
      p-2 sm:p-4 max-w-xl"
      >
        <div
          className="justify-center 
            font-bold   flex  flex-col gap-6
                items-center"
        >
          <TypeAnimation
            sequence={[300, "Quizex", 1000]}
            wrapper="div"
            speed={40}
            cursor={false}
            className="font-mono m-auto md:-top-16 font-extrabold 
            tracking-widest md:mt-auto   underline decoration-teal-300 decoration-wavy 
       text-tal-600 col-span-4 left-1.5 text-center text-7xl md:text-[12rem] "
          />

          <motion.div
            initial={{ opacity: 0, x: -40 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ delay: 1.2, duration: 0.3 }}
          >
            <p className="italic font-thin text-black text-lg sm:text-2xl">
              Multiplayer Quiz Game based on AI
            </p>
          </motion.div>
          <motion.div
            initial={{ opacity: 0, y: 40, scale: 0.5 }}
            animate={{ opacity: 1, y: 0, scale: 1 }}
            transition={{ delay: 1.5, duration: 0.3 }}
          >
            <motion.button
              whileHover={{ scale: 1.15, rotate: 0 }}
              whileTap={{ scale: 1 }}
              color={"black"}
              animate={{
                backgroundColor: ["#00EEEE", "#C1FFC1", "#00EEEE"],
                y: [0],
                scale: [1, 1.1, 1],
                rotate: [0],
                transition: {
                  duration: 2,
                  delay: 1,
                  repeat: Infinity,
                  ease: "linear",
                },
              }}
              className="flex mt-12 rounded-xl text-black py-2 px-4 text-4xl font-ibm-plex tracking-tight
                items-center  border-black border-4  duration-300 gap-3 shadow-gray-400
                 bg-teal-600 m-auto shadow-lg hover:shadow-slate-600"
            >
              <Link to={"/home"} className=" flex items-center gap-2">
                <FaRegCirclePlay size={32} color="black" /> Play
              </Link>
            </motion.button>
          </motion.div>
        </div>
      </section>
    </main>
  );
};
