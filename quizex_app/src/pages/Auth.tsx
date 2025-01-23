import React, { useEffect, useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const API_URL = "/api";
const PREFIX = "/auth";

type Step = "register" | "login" | "exchange";

interface FormData {
  name: string;
  email: string;
  code: string;
}

export const Auth: React.FC = () => {
  const navigate = useNavigate();

  const [currentStep, setCurrentStep] = useState<Step>("register");
  const [formData, setFormData] = useState<FormData>({
    name: "",
    email: "",
    code: "",
  });
  const [loading, setLoading] = useState<boolean>(false);
  const [message, setMessage] = useState("");
  useEffect(() => {
    document.body.style.overflow = "hidden";

    return () => {
      document.body.style.overflow = "auto";
    };
  }, []);

  const handleRegister = async () => {
    try {
      setLoading(true);
      const response = await axios.post(`${API_URL}${PREFIX}/register`, {
        name: formData.name,
        email: formData.email,
      });

      if (response.status === 200) {
        setMessage("Please check your email for verification code.");
        setCurrentStep("exchange");
      }
    } catch (error: any) {
      alert(error.response?.data?.message || "Registration failed");
    } finally {
      setLoading(false);
    }
  };

  const handleLogin = async () => {
    try {
      setLoading(true);
      const response = await axios.post(`${API_URL}${PREFIX}/login`, {
        email: formData.email,
      });

      if (response.status === 200) {
        setMessage("Please check your email for verification code.");
        setCurrentStep("exchange");
      }
    } catch (error: any) {
      alert(error.response?.data?.message || "Login failed");
    } finally {
      setLoading(false);
    }
  };

  const handleExchange = async () => {
    try {
      setLoading(true);
      const response = await axios.post(`${API_URL}${PREFIX}/exchange`, {
        email: formData.email,
        code: formData.code,
      });

      if (response.data) {
        navigate("/");
      }
    } catch (error: any) {
      alert(error.response?.data?.message || "Code verification failed");
    } finally {
      setLoading(false);
    }
  };

  const renderRegisterForm = () => (
    <>
      <input
        className="bg-gray-300 w-full  border border-gray-300 rounded-md p-3 mb-4"
        type="text"
        placeholder="Name"
        value={formData.name}
        onChange={(e) => setFormData({ ...formData, name: e.target.value })}
      />
      <input
        className="bg-gray-300 w-full border border-gray-300 rounded-md p-3 mb-4"
        type="email"
        placeholder="Email"
        value={formData.email}
        onChange={(e) => setFormData({ ...formData, email: e.target.value })}
      />
      <button
        className={`w-full bg-teal-600 text-white p-3 rounded-md mb-4 ${
          loading ? "opacity-50" : ""
        }`}
        onClick={handleRegister}
        disabled={loading}
      >
        {loading ? "Loading..." : "Submit"}
      </button>
      <p
        className="text-center text-blue-600 cursor-pointer"
        onClick={() => setCurrentStep("login")}
      >
        Already have an account? Sign in
      </p>
    </>
  );

  const renderLoginForm = () => (
    <>
      <input
        className="bg-gray-300 w-full placeholder:text-center border border-gray-300 rounded-md p-3 mb-4"
        type="email"
        placeholder="Email"
        value={formData.email}
        onChange={(e) => setFormData({ ...formData, email: e.target.value })}
      />
      <button
        className={`w-full bg-teal-600 text-white p-3 rounded-md mb-4 ${
          loading ? "opacity-50" : ""
        }`}
        onClick={handleLogin}
        disabled={loading}
      >
        {loading ? "Loading..." : "Submit"}
      </button>
      <p
        className="text-center text-blue-600 cursor-pointer"
        onClick={() => setCurrentStep("register")}
      >
        Don't have an account? Sign up
      </p>
    </>
  );

  const renderExchangeForm = () => (
    <>
      <input
        className="bg-gray-300 w-full border placeholder:text-center  border-gray-300 rounded-md p-3 mb-4"
        type="text"
        placeholder="Code"
        value={formData.code}
        onChange={(e) => setFormData({ ...formData, code: e.target.value })}
        maxLength={6}
      />
      <button
        className={`w-full bg-teal-600 text-white p-3 rounded-md ${
          loading ? "opacity-50" : ""
        }`}
        onClick={handleExchange}
        disabled={loading}
      >
        {loading ? "Loading..." : "Submit"}
      </button>
    </>
  );

  return (
    <div className="flex flex-col items-center gap-4 justify-center  h-screen pb-64">
      <h1 className="text-2xl font-bold ">
        {currentStep === "register"
          ? "Sign Up"
          : currentStep === "login"
          ? "Sign In"
          : null}
      </h1>
      <p> {message}</p>
      <div className="w-full max-w-md bg-white p-6 rounded-md shadow-md">
        {currentStep === "register" && renderRegisterForm()}
        {currentStep === "login" && renderLoginForm()}
        {currentStep === "exchange" && renderExchangeForm()}
      </div>
    </div>
  );
};
