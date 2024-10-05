import { Link } from "react-router-dom";

// Const
import routerConfig from "@/const/config/router";

// Component
import ChangePassword from "./components/ChangePassword";

const Settings = () => {
  return (
    <>
      <div className='breadcrumb-container'>
        <Link to={routerConfig.routes.SETTINGS} className='breadcrumb--item'>
          <span className='breadcrumb--item--inner'>
            <span className='breadcrumb--item-title'>Settings</span>
          </span>
        </Link>
      </div>

      <ChangePassword />
    </>
  );
};

export default Settings;
