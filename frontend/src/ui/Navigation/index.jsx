import { Link } from "react-router-dom";
import routerConfig from "@/const/config/router";

const Navigation = () => {
  return (
    <div className='navigation'>
      <Link className='button nav_link' to={routerConfig.routes.DASHBOARD}>
        <i className='fa-solid fa-chart-simple'></i>
        <span>Dashboard</span>
      </Link>
      <Link className='button nav_link' to={routerConfig.routes.CATEGORY}>
        <i className='fa-solid fa-list'></i>
        <span>Category</span>
      </Link>
      <Link className='button nav_link'>
        <i className='fa-solid fa-record-vinyl'></i>
        <span>Records</span>
      </Link>
      <Link className='button nav_link'>
        <i className='fa-solid fa-bars'></i>
        <span>Setting</span>
      </Link>
    </div>
  );
};

export default Navigation;
