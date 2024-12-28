import "./module.css";
import logo from "@/assets/img/brand.png";

import { useState } from "react";
import PropTypes from "prop-types";
import { Link } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faChartSimple,
  faPerson,
  faLayerGroup,
  faTag,
  faFile,
  faGear,
} from "@fortawesome/free-solid-svg-icons";

// Const
import routerConfig from "@/const/config/router";

const Navbar = ({ user, onCleanUser }) => {
  const [menuOpen, setMenuOpen] = useState(false);

  return (
    <>
      <nav className='bg-darkPrimary custom-container !py-[0.6rem]'>
        <div className='flex items-center justify-between'>
          <div className='flex items-center gap-2'>
            <div className='text-white text-2xl font-bold bg-primary rounded-[50%] p-2'>
              <a href='#'>
                <img src={logo} className='w-4' />
              </a>
            </div>
            <h1 className='text-white lg:block md:block hidden'>
              Management Information System
            </h1>
            <h1 className='text-white lg:hidden md:hidden'>MIS</h1>
          </div>
          <div className='block md:hidden px-1'>
            <button id='menu-button' className='text-white focus:outline-none'>
              <svg
                className='w-6 h-6'
                fill='none'
                stroke='currentColor'
                viewBox='0 0 24 24'
                xmlns='http://www.w3.org/2000/svg'
                onClick={() => setMenuOpen(!menuOpen)}
              >
                <path
                  strokeLinecap='round'
                  strokeLinejoin='round'
                  strokeWidth='2'
                  d='M4 6h16M4 12h16m-7 6h7'
                />
              </svg>
            </button>
          </div>
          <div className='hidden md:flex gap-2'>
            <div className='text-white'>Hi, {user.email}</div>
            <div className='text-white'>|</div>
            <Link
              to='/'
              onClick={onCleanUser}
              className='text-white hover:text-primary'
            >
              Logout
            </Link>
          </div>
        </div>
        <div
          id='mobile-menu'
          className={`md:hidden ${menuOpen ? "block" : "hidden"}`}
        >
          <div className='mobile-nav'>
            <Link
              to={routerConfig.routes.DASHBOARD}
              className='mobile-nav-item'
            >
              Dashboard
            </Link>
            <Link to={routerConfig.routes.MEMBERS} className='mobile-nav-item'>
              Members
            </Link>
            <Link
              to={routerConfig.routes.CATEGORIES}
              className='mobile-nav-item'
            >
              Categories
            </Link>
            <Link to={routerConfig.routes.TAGS} className='mobile-nav-item'>
              Tags
            </Link>
            <Link
              to={routerConfig.routes.DOCUMENTS}
              className='mobile-nav-item'
            >
              DOCUMENTS
            </Link>
            <Link to={routerConfig.routes.SETTINGS} className='mobile-nav-item'>
              Settings
            </Link>
            <Link
              to='/'
              onClick={onCleanUser}
              className='text-white hover:text-primary mt-4'
            >
              Logout
            </Link>
          </div>
        </div>
      </nav>
      <nav className='nav'>
        <Link
          to={routerConfig.routes.DASHBOARD}
          className='nav-item primary-shadow'
        >
          <FontAwesomeIcon icon={faChartSimple} /> Dashboard
        </Link>
        <Link
          to={routerConfig.routes.MEMBERS}
          className='nav-item primary-shadow'
        >
          <FontAwesomeIcon icon={faPerson} /> Members
        </Link>
        <Link
          to={routerConfig.routes.CATEGORIES}
          className='nav-item primary-shadow'
        >
          <FontAwesomeIcon icon={faLayerGroup} /> Categories
        </Link>
        <Link to={routerConfig.routes.TAGS} className='nav-item primary-shadow'>
          <FontAwesomeIcon icon={faTag} /> Tags
        </Link>
        <Link
          to={routerConfig.routes.DOCUMENTS}
          className='nav-item primary-shadow'
        >
          <FontAwesomeIcon icon={faFile} /> Documents
        </Link>
        <Link
          to={routerConfig.routes.SETTINGS}
          className='nav-item primary-shadow'
        >
          <FontAwesomeIcon icon={faGear} /> Settings
        </Link>
      </nav>
    </>
  );
};

Navbar.propTypes = {
  user: PropTypes.object.isRequired,
  onCleanUser: PropTypes.func.isRequired,
};

export default Navbar;
