import Logo from "@/assets/img/brand.png";

import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "../../util/message.util";

const message = {
  LOGOUT_SUCCESS: "Logout successfully",
};

const Header = (props) => {
  const { user, onCleanUser } = props;

  const handleLogout = async () => {
    try {
      await apiHandler.get("/auth/sign-out");
      messageUtil.showSuccessMessage(message.LOGOUT_SUCCESS);
      onCleanUser();
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  const redirectToHome = () => {
    window.location.href = "/";
  };

  return (
    <div className='layout__header'>
      <div className='header'>
        <div className='header_left'>
          <img
            src={Logo}
            alt='logo'
            className='logo'
            onClick={redirectToHome}
          />
          <h1>Management Information System</h1>
        </div>
        <div className='header_right'>
          Hi! <span className='user'>{user?.name}</span> |{" "}
          <span className='link-button' onClick={handleLogout}>
            Logout
          </span>
        </div>
      </div>
    </div>
  );
};

export default Header;
