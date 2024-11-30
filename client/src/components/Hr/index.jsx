import "./module.css";
import PropTypes from "prop-types";
import { generateExtraClass } from "../util";

const Hr = ({ extraClasses }) => {
  let extraClassName = generateExtraClass(extraClasses);

  return <hr className={`hr ${extraClassName}`} />;
};

Hr.propTypes = {
  extraClasses: PropTypes.array,
};

export default Hr;
