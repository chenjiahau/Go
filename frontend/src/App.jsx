import "bootstrap/dist/css/bootstrap.css";
import "react-toastify/dist/ReactToastify.css";
import "./App.css";

import { useEffect, useState } from "react";
import { HashRouter, Routes, Route } from "react-router-dom";
import { ToastContainer } from "react-toastify";

import Layout from "@/ui/Layout";

import Login from "@/pages/Login";
import Dashboard from "@/pages/Dashboard";

function App() {
  // State
  const [user, setUser] = useState(null);

  // Method
  const handleSaveUserInfo = (data) => {
    setUser(data);
  };

  const handleCleanUserInfo = () => {
    setUser(null);
    localStorage.removeItem("user");
  };

  useEffect(() => {
    const user = JSON.parse(localStorage.getItem("user"));

    if (user) {
      setUser(user);
    }
  }, []);

  return (
    <HashRouter>
      <Routes>
        <Route
          path='/'
          element={<Login user={user} onSaveUser={handleSaveUserInfo} />}
        />
        <Route
          path='/'
          element={<Layout user={user} onCleanUser={handleCleanUserInfo} />}
        >
          <Route path='/dashboard' element={<Dashboard user={user} />} />
        </Route>
      </Routes>
      <ToastContainer />
    </HashRouter>
  );
}

export default App;
