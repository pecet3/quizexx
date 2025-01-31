import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";
import { BrowserRouter } from "react-router-dom";
import { Toaster } from "react-hot-toast";
import { ProtectedProvider } from "./context/protectedContext.tsx";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <ProtectedProvider>
    <Toaster
      position="bottom-center"
      toastOptions={{
        duration: 3000,
        style: {
          backgroundColor: "#3b0764",
          color: "white",
        },
      }}
    />
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </ProtectedProvider>
);
