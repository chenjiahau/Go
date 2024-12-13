import "./module.css";

import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import { Outlet } from "react-router-dom";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Slice
import { userActions, selectUser } from "@/store/slices/user";

// Handler
import apiHandler from "@/util/api.util";

// Component
import Navbar from "@/ui/Navbar";
import Footer from "@/ui/Footer";

function Layout() {
  const dispatch = useDispatch();
  const navigate = useNavigate();

  // State
  const user = useSelector(selectUser);
  const [validToken, setValidToken] = useState(false);

  // Method
  const onCleanUser = () => {
    dispatch(userActions.cleanUser());
    localStorage.removeItem("user");
    navigate(routerConfig.routes.LOGIN);
  };

  // Side effect
  useEffect(() => {
    const getUser = async (savedUser) => {
      try {
        apiHandler.setToken(savedUser.token);
        await apiHandler.get(apiConfig.resource.VERIFY_TOKEN);
        setValidToken(true);
        dispatch(userActions.setUser(savedUser));
      } catch (error) {
        localStorage.removeItem("user");
        setValidToken(false);
        dispatch(userActions.cleanUser());
        navigate(routerConfig.routes.LOGIN);
      }
    };

    const savedUser = JSON.parse(localStorage.getItem("user"));
    if (savedUser) {
      getUser(savedUser);
    } else {
      navigate(routerConfig.routes.LOGIN);
    }
  }, [dispatch, navigate]);

  if (!user || !validToken) {
    return null;
  }

  return (
    <>
      <div className='layout'>
        <Navbar user={user} onCleanUser={onCleanUser} />
        <div className='main-content'>
          <Outlet />
        </div>
        <Footer />
      </div>
    </>
  );
}

export default Layout;
