import "./module.css";

import { useRef } from "react";
import PropTypes from "prop-types";

const CheckBox = ({ name, checked, onChange, children }) => {
  const inputRef = useRef(null);
  const handleClick = () => {
    inputRef.current.click();
  };

  return (
    <div className='checkbox-container'>
      <input
        type='checkbox'
        name={name}
        checked={checked}
        onChange={onChange}
        ref={inputRef}
      />
      <div className='square' onClick={handleClick}></div>
      <div>{children}</div>
    </div>
  );
};

CheckBox.propTypes = {
  shape: PropTypes.string,
  name: PropTypes.string,
  value: PropTypes.bool,
  checked: PropTypes.bool,
  onChange: PropTypes.func,
  children: PropTypes.node,
};

export default CheckBox;
