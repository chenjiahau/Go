import "./module.css";

import PropTypes from "prop-types";
import { generateExtraClass } from "../util";

const Form = ({ extraClasses, children }) => {
  let extraClassName = generateExtraClass(extraClasses);
  return <div className={`form-container ${extraClassName}`}>{children}</div>;
};

Form.propTypes = {
  extraClasses: PropTypes.array,
  children: PropTypes.node.isRequired,
};

export default Form;
