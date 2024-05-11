import PropTypes from "prop-types";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Outlet } from "react-router-dom";

import apiHandler from "@/util/api.util";

import Header from "../Header";
import Navigation from "../Navigation";
import Footer from "../Footer";
import messageUtil, { commonMessage } from "../../util/message.util";

const Layout = (props) => {
  const { user, onCleanUser } = props;
  const navigate = useNavigate();

  useEffect(() => {
    const getUser = async (savedUser) => {
      try {
        apiHandler.setToken(savedUser.token);
        await apiHandler.get("/auth/verify-token");
      } catch (error) {
        onCleanUser();
        navigate("/");
        messageUtil.showErrorMessage(commonMessage.error);
      }
    };

    const savedUser = JSON.parse(localStorage.getItem("user"));
    if (savedUser) {
      getUser(savedUser);
    }
  }, []);

  return (
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
  );
};

Layout.propTypes = {
  user: PropTypes.object,
  onCleanUser: PropTypes.func.isRequired,
};

export default Layout;
