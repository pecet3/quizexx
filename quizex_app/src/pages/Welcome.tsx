import { TypeAnimation } from "react-type-animation";
import { motion } from "framer-motion";
import { FaQuestion } from "react-icons/fa";
import { FaRegCirclePlay } from "react-icons/fa6";
import { Link } from "react-router-dom";
export const Welcome = () => {
  return (
    <main>
      <section
        className="m-auto h-screen flex flex-col 
      gap-8 sm:gap-12 items-center justify-start mt-72
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
            tracking-widest md:mt-auto 
       text-tal-600 col-span-4 left-1.5 text-center text-7xl md:text-[12rem] "
          />

          <motion.div
            initial={{ opacity: 0, x: -40 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ delay: 1.2, duration: 0.3 }}
          >
            <p className="italic text-teal-600 text-lg">
              Multiplayer Quiz Game based on AI
            </p>
          </motion.div>
          <motion.div
            initial={{ opacity: 0, x: 40 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ delay: 1.5, duration: 0.3 }}
          >
            <motion.button
              whileHover={{ scale: 1.02, rotate: 0 }}
              whileTap={{ scale: 1 }}
              color={"black"}
              animate={{
                backgroundColor: ["#0d9488", "#0f53b6", "#0d9488"],
                y: [0],
                scale: [1, 1.05, 1],
                rotate: [0],
                transition: {
                  duration: 2,
                  delay: 1,
                  repeat: Infinity,
                  ease: "linear",
                },
              }}
              className="flex mt-12 rounded-xl text-white py-2 px-4 text-4xl font-ibm-plex tracking-tight
                items-center  border-white border-4  duration-300 gap-3 shadow-gray-400
                 bg-teal-600 m-auto shadow-lg hover:shadow-slate-900"
            >
              <Link to={"/home"} className=" flex items-center gap-2">
                <FaRegCirclePlay size={32} /> Play
              </Link>
            </motion.button>
          </motion.div>
        </div>
      </section>
    </main>
  );
};
