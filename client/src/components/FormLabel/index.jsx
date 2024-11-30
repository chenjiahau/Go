import "./module.css";

import PropTypes from "prop-types";
import { generateExtraClass } from "../util";

const FormLabel = ({ extraClasses, forName, children }) => {
  forName = forName || "";
  let extraClassName = generateExtraClass(extraClasses);

  return (
    <label className={`form-label ${extraClassName}`} htmlFor={forName}>
      {children}
    </label>
  );
};

FormLabel.propTypes = {
  extraClasses: PropTypes.array,
  forName: PropTypes.string,
  children: PropTypes.node,
};

export default FormLabel;
