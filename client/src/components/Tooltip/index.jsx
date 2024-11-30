import "./module.css";

import { useState } from "react";
import PropTypes from "prop-types";

const Tooltip = ({ children, content }) => {
  const [visible, setVisible] = useState(false);

  return (
    <div className='tooltip-container'>
      <div
        className='trigger'
        onMouseEnter={() => setVisible(true)}
        onMouseLeave={(e) => {
          const relatedTarget = e.relatedTarget;
          if (!relatedTarget || !e.currentTarget.contains(relatedTarget)) {
            setVisible(false);
          }
        }}
      >
        {children}
      </div>
      {visible && (
        <div className='content' onMouseEnter={() => setVisible(true)}>
          <div className='triangle'></div>
          <div className='inner' onMouseLeave={() => setVisible(false)}>
            {typeof content === "string" ? <span>{content}</span> : content}
          </div>
        </div>
      )}
    </div>
  );
};

Tooltip.propTypes = {
  children: PropTypes.node.isRequired,
  content: PropTypes.oneOfType([PropTypes.string, PropTypes.node]).isRequired,
};

export default Tooltip;
