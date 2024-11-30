import "./module.css";

import { useRef, useEffect } from "react";
import PropTypes from "prop-types";
import { isFunction } from "lodash";

import { generateExtraClass } from "../util";

const InputBox = ({
  type,
  extraClasses,
  disabled,
  onChange,
  onClose,
  children,
  ...props
}) => {
  type = type || "text";
  let extraClassName = generateExtraClass(extraClasses);
  let disabledClass = disabled ? "disabled" : "";
  const inputRef = useRef(null);

  const handleClick = (e) => {
    if (e.key === "Enter" && isFunction(onClose)) {
      e.preventDefault();
      e.stopPropagation();
      onClose(inputRef.current.value);
    }
  };

  useEffect(() => {
    inputRef.current.focus();
  }, []);

  return (
    <div
      className={`
        primary-shadow 
        ${children ? "input-box-with-icon" : "input-box"}
        ${disabledClass}
      `}
    >
      <input
        ref={inputRef}
        type={type}
        className={`${extraClassName} ${disabledClass}`}
        {...props}
        onChange={(e) => onChange(e.target.value)}
        onKeyDown={handleClick}
      />
      {children}
    </div>
  );
};

InputBox.propTypes = {
  type: PropTypes.string,
  extraClasses: PropTypes.array,
  disabled: PropTypes.bool,
  onChange: PropTypes.func,
  onClose: PropTypes.func,
  children: PropTypes.node,
};

export default InputBox;
