import "./module.css";

import { useRef } from "react";
import PropTypes from "prop-types";

const RadioBox = ({ name, checked, onChange, children }) => {
  const inputRef = useRef(null);
  const handleClick = () => {
    inputRef.current.click();
  };

  return (
    <div className='radio-container'>
      <input
        type='radio'
        name={name}
        checked={checked}
        onChange={onChange}
        ref={inputRef}
      />
      <label htmlFor={name} onClick={handleClick}>
        <div className='circle'></div>
      </label>
      <div>{children}</div>
    </div>
  );
};

RadioBox.propTypes = {
  shape: PropTypes.string,
  name: PropTypes.string,
  value: PropTypes.bool,
  checked: PropTypes.bool,
  onChange: PropTypes.func,
  children: PropTypes.node,
};

export default RadioBox;
