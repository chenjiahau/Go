import "./module.css";

import PropTypes from "prop-types";
import { generateExtraClass } from "../util";

const ToolbarBox = ({ extraClasses, children }) => {
  let extraClassName = generateExtraClass(extraClasses);

  return <div className={`toolbar-box ${extraClassName}`}>{children}</div>;
};

ToolbarBox.propTypes = {
  extraClasses: PropTypes.array,
  children: PropTypes.node,
};

export default ToolbarBox;
