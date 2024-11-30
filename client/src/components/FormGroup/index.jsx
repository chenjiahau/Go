import "./module.css";

import PropTypes from "prop-types";
import { generateExtraClass } from "../util";

const FormGroup = ({ isHorizontal, extraClasses, children }) => {
  isHorizontal = isHorizontal ? "form-group-horizontal" : "form-group";
  let extraClassName = generateExtraClass(extraClasses);
  return <div className={`${isHorizontal} ${extraClassName}`}>{children}</div>;
};

FormGroup.propTypes = {
  isHorizontal: PropTypes.bool,
  extraClasses: PropTypes.array,
  children: PropTypes.node,
};

export default FormGroup;
