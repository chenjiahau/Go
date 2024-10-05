import { Link, useLocation } from "react-router-dom";
import routerConfig from "@/const/config/router";

const Navigation = () => {
  const location = useLocation();

  const checkActive = (route) => {
    const ary = location.pathname.split("/");
    let category = ary[1];
    if (category === "category") {
      category = "categories";
    }

    if (category === "tag") {
      category = "tags";
    }

    if (category === "document") {
      category = "documents";
    }

    const realPath = `/${category}`;
    return realPath.indexOf(route) !== -1 ? "nav_active" : "";
  };

  return (
    <div className='navigation'>
      <Link
        className={`button nav_link ${checkActive(
          routerConfig.routes.DASHBOARD
        )}`}
        to={routerConfig.routes.DASHBOARD}
      >
        <i className='fa-solid fa-chart-simple'></i>
        <span>Dashboard</span>
      </Link>
      <Link
        className={`button nav_link ${checkActive(
          routerConfig.routes.MEMBERS
        )}`}
        to={routerConfig.routes.MEMBERS}
      >
        <i className='fa-solid fa-user'></i>
        <span>Members</span>
      </Link>
      <Link
        className={`button nav_link ${checkActive(
          routerConfig.routes.CATEGORIES
        )}`}
        to={routerConfig.routes.CATEGORIES}
      >
        <i className='fa-solid fa-list'></i>
        <span>Categories</span>
      </Link>
      <Link
        className={`button nav_link ${checkActive(routerConfig.routes.TAGS)}`}
        to={routerConfig.routes.TAGS}
      >
        <i className='fa-solid fa-tag'></i>
        <span>Tags</span>
      </Link>
      <Link
        className={`button nav_link ${checkActive(
          routerConfig.routes.DOCUMENTS
        )}`}
        to={routerConfig.routes.DOCUMENTS}
      >
        <i className='fa-solid fa-file'></i>
        <span>Documents</span>
      </Link>
      <Link
        className={`button nav_link ${checkActive(
          routerConfig.routes.SETTING
        )}`}
        to={routerConfig.routes.SETTINGS}
      >
        <i className='fa-solid fa-bars'></i>
        <span>Settings</span>
      </Link>
    </div>
  );
};

export default Navigation;
