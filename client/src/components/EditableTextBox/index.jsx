import "./module.css";

import { useRef, useEffect } from "react";
import PropTypes from "prop-types";

const EditableTextBox = ({ text, onClick }) => {
  const inputRef = useRef(null);

  useEffect(() => {
    inputRef.current.focus();
  }, []);

  return (
    <div className='editable-text-box' onClick={onClick} ref={inputRef}>
      <div className='label'>{text}</div>
    </div>
  );
};

EditableTextBox.propTypes = {
  text: PropTypes.string.isRequired,
  onClick: PropTypes.func.isRequired,
};

export default EditableTextBox;
