import "./module.css";

import PropTypes from "prop-types";

const IconButton = ({ rounded, disabled, onClick, children }) => {
  rounded = rounded ?? false;

  if (disabled) {
    if (rounded) {
      return (
        <button
          className='icon-button icon-button__rounded icon-button__disable'
          disabled
        >
          {children}
        </button>
      );
    }

    return (
      <button className='icon-button icon-button__disable' disabled>
        {children}
      </button>
    );
  }

  if (rounded) {
    return (
      <button className='icon-button icon-button__rounded' onClick={onClick}>
        {children}
      </button>
    );
  }

  return (
    <button className={`icon-button `} onClick={onClick}>
      {children}
    </button>
  );
};

IconButton.propTypes = {
  rounded: PropTypes.bool,
  disabled: PropTypes.bool,
  onClick: PropTypes.func.isRequired,
  children: PropTypes.node.isRequired,
};

export default IconButton;
