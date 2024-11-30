import "./module.css";

import PropTypes from "prop-types";
import { generateExtraClass } from "../util";

const MainTitle = ({ tag, extraClasses, children, ...props }) => {
  tag = tag || "h1";
  const Tag = tag;
  let extraClassName = generateExtraClass(extraClasses);

  return (
    <Tag className={`main-title ${extraClassName}`} {...props}>
      {children}
    </Tag>
  );
};

MainTitle.propTypes = {
  tag: PropTypes.string,
  extraClasses: PropTypes.array,
  children: PropTypes.node.isRequired,
};

export default MainTitle;
