import "./module.css";

import PropTypes from "prop-types";
import { generateExtraClass } from "../util";

const Spacer = ({ extraClasses }) => {
  let extraClassName = generateExtraClass(extraClasses);
  return <div className={`spacer ${extraClassName}`} />;
};

Spacer.propTypes = {
  extraClasses: PropTypes.array,
};

export default Spacer;
