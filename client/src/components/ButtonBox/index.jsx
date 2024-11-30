import "./module.css";
import PropTypes from "prop-types";
import { generateExtraClass } from "../util";

const ButtonBox = ({
  extraClasses,
  isDisabled,
  isClose,
  isSave,
  title,
  onClick,
  children,
}) => {
  let extraClassName = generateExtraClass(extraClasses);
  let disabledClassName = isDisabled ? "disabled-button" : "";
  let saveClassName = isSave ? "save-button" : "";
  let closeClassName = isClose ? "close-button" : "";

  const handleClick = () => {
    if (isDisabled) return;
    onClick();
  };

  return (
    <div
      className={`button-box dark-primary-shadow ${extraClassName} ${disabledClassName} ${closeClassName} ${saveClassName}`}
      onClick={handleClick}
    >
      <button className={`${disabledClassName}`}>{children || title}</button>
    </div>
  );
};

ButtonBox.propTypes = {
  extraClasses: PropTypes.array,
  isDisabled: PropTypes.bool,
  isSave: PropTypes.bool,
  isClose: PropTypes.bool,
  title: PropTypes.string,
  onClick: PropTypes.func,
  children: PropTypes.node,
};

export default ButtonBox;
