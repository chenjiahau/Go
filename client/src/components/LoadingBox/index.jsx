import "./module.css";

import React from "react";
import PropTypes from "prop-types";

const LoadingBox = ({ visible }) => {
  if (!visible) return null;

  return (
    <div className='loading-box'>
      <div className='loading-spinner'></div>
    </div>
  );
};

LoadingBox.propTypes = {
  visible: PropTypes.bool,
};

export default LoadingBox;
