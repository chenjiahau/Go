import PropTypes from "prop-types";
import { Outlet } from "react-router-dom";

import Header from "../Header";
import Navigation from "../Navigation";
import Footer from "../Footer";

const Layout = (props) => {
  const { user, onCleanUser } = props;

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
