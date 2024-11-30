import "./module.css";

import React from "react";
import PropTypes from "prop-types";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faFile } from "@fortawesome/free-solid-svg-icons";

const FloatingButton = ({ handleExecution }) => {
  return (
    <div className='floating-button'>
      <div className='button' onClick={handleExecution}>
        <FontAwesomeIcon icon={faFile} />
      </div>
    </div>
  );
};

FloatingButton.propTypes = {
  handleExecution: PropTypes.func,
};

export default FloatingButton;
