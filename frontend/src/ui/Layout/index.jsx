import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import { Outlet } from "react-router-dom";

// Const
import routerConfig from "../../const/config/router";
import apiConfig from "@/const/config/api";

// Slice
import { userActions, selectUser } from "@/store/slices/user";

// Handler
import apiHandler from "@/util/api.util";

// Component
import Header from "../Header";
import Navigation from "../Navigation";
import Footer from "../Footer";

// Util
import messageUtil, { commonMessage } from "../../util/message.util";

const Layout = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();

  // State
  const user = useSelector(selectUser);

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
        dispatch(userActions.setUser(savedUser));
      } catch (error) {
        dispatch(userActions.cleanUser());
        navigate(routerConfig.routes.HOME);
        messageUtil.showErrorMessage(commonMessage.error);
      }
    };

    const savedUser = JSON.parse(localStorage.getItem("user"));
    if (savedUser) {
      getUser(savedUser);
    } else {
      navigate(routerConfig.routes.LOGIN);
    }
  }, []);

  return (
    <>
      {user && (
        <div className='layout'>
          <Header user={user} onCleanUser={onCleanUser} />
          <div className='layout__content'>
            <div className='content__left'>
              <Navigation />
            </div>
            <div className='content__right'>
              <Outlet />
            </div>
            <Footer />
          </div>
        </div>
      )}
    </>
  );
};

export default Layout;
