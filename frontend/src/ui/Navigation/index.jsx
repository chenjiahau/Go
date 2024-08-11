import { Link } from "react-router-dom";
import routerConfig from "@/const/config/router";

const Navigation = () => {
  return (
    <div className='navigation'>
      <Link className='button nav_link' to={routerConfig.routes.DASHBOARD}>
        <i className='fa-solid fa-chart-simple'></i>
        <span>Dashboard</span>
      </Link>
      <Link className='button nav_link' to={routerConfig.routes.MEMBERS}>
        <i className='fa-solid fa-user'></i>
        <span>Members</span>
      </Link>
      <Link className='button nav_link' to={routerConfig.routes.CATEGORIES}>
        <i className='fa-solid fa-list'></i>
        <span>Categories</span>
      </Link>
      <Link className='button nav_link' to={routerConfig.routes.TAGS}>
        <i className='fa-solid fa-tag'></i>
        <span>Tags</span>
      </Link>
      <Link className='button nav_link' to={routerConfig.routes.DOCUMENTS}>
        <i className='fa-solid fa-file'></i>
        <span>Documents</span>
      </Link>
      <Link className='button nav_link'>
        <i className='fa-solid fa-bars'></i>
        <span>Setting</span>
      </Link>
    </div>
  );
};

export default Navigation;
