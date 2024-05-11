import Logo from "@/assets/img/brand.png";

import PropTypes from "prop-types";
import { useNavigate } from "react-router-dom";

import messageUtil from "../../util/message.util";

const message = {
  LOGOUT_SUCCESS: "Logout successfully",
};

const Header = (props) => {
  const { user, onCleanUser } = props;
  const navigate = useNavigate();

  const handleLogout = () => {
    onCleanUser();
    navigate("/");
    messageUtil.showSuccessMessage(message.LOGOUT_SUCCESS);
  };

  return (
    <div className='layout__header'>
      <div className='header'>
        <div className='header_left'>
          <img src={Logo} alt='logo' className='logo' />
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

Header.propTypes = {
  user: PropTypes.object,
  onCleanUser: PropTypes.func.isRequired,
};

export default Header;
