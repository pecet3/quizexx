import React, { useEffect, useState } from "react";
import axios from "axios";

const API_URL = "/api";
const PREFIX = "/auth";

type Step = "register" | "login" | "exchange";

interface FormData {
  name: string;
  email: string;
  code: string;
}

export const Auth: React.FC = () => {
  const [currentStep, setCurrentStep] = useState<Step>("register");
  const [formData, setFormData] = useState<FormData>({
    name: "",
    email: "",
    code: "",
  });
  const [loading, setLoading] = useState<boolean>(false);
  useEffect(() => {
    console.log(currentStep);
  }, [currentStep]);
  const handleRegister = async () => {
    try {
      setLoading(true);
      const response = await axios.post(`${API_URL}${PREFIX}/register`, {
        name: formData.name,
        email: formData.email,
      });

      if (response.status === 200) {
        alert(
          "Registration successful! Please check your email for verification code."
        );
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
        alert("Please check your email for verification code.");
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
        alert("Authentication successful!");
        // Navigate to main app or handle success accordingly
        console.log(response.data);
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
        className="bg-gray-300 w-full border border-gray-300 rounded-md p-3 mb-4"
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
        className={`w-full bg-blue-600 text-white p-3 rounded-md mb-4 ${
          loading ? "opacity-50" : ""
        }`}
        onClick={handleRegister}
        disabled={loading}
      >
        {loading ? "Registering..." : "Register"}
      </button>
      <p
        className="text-center text-blue-600 cursor-pointer"
        onClick={() => setCurrentStep("login")}
      >
        Already have an account? Login
      </p>
    </>
  );

  const renderLoginForm = () => (
    <>
      <input
        className="bg-gray-300 w-full border border-gray-300 rounded-md p-3 mb-4"
        type="email"
        placeholder="Email"
        value={formData.email}
        onChange={(e) => setFormData({ ...formData, email: e.target.value })}
      />
      <button
        className={`w-full bg-blue-600 text-white p-3 rounded-md mb-4 ${
          loading ? "opacity-50" : ""
        }`}
        onClick={handleLogin}
        disabled={loading}
      >
        {loading ? "Logging in..." : "Login"}
      </button>
      <p
        className="text-center text-blue-600 cursor-pointer"
        onClick={() => setCurrentStep("register")}
      >
        Don't have an account? Register
      </p>
    </>
  );

  const renderExchangeForm = () => (
    <>
      <input
        className="bg-gray-300 w-full border border-gray-300 rounded-md p-3 mb-4"
        type="text"
        placeholder="Verification Code"
        value={formData.code}
        onChange={(e) => setFormData({ ...formData, code: e.target.value })}
        maxLength={6}
      />
      <button
        className={`w-full bg-blue-600 text-white p-3 rounded-md ${
          loading ? "opacity-50" : ""
        }`}
        onClick={handleExchange}
        disabled={loading}
      >
        {loading ? "Verifying..." : "Verify Code"}
      </button>
    </>
  );

  return (
    <div className="flex flex-col items-center justify-center h-screen pb-64">
      <h1 className="text-2xl font-bold mb-6">
        {currentStep === "register"
          ? "Register"
          : currentStep === "login"
          ? "Login"
          : "Verify Code"}
      </h1>
      <div className="w-full max-w-md bg-white p-6 rounded-md shadow-md">
        {currentStep === "register" && renderRegisterForm()}
        {currentStep === "login" && renderLoginForm()}
        {currentStep === "exchange" && renderExchangeForm()}
      </div>
    </div>
  );
};
