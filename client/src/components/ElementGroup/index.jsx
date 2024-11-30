import "./module.css";

import PropTypes from "prop-types";
import { generateExtraClass } from "../util";

const ElementGroup = ({ extraClasses, children }) => {
  let extraClassName = generateExtraClass(extraClasses);
  return <div className={`element-group ${extraClassName}`}>{children}</div>;
};

ElementGroup.propTypes = {
  extraClasses: PropTypes.array,
  children: PropTypes.node,
};

export default ElementGroup;
