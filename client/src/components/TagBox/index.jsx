import "./module.css";

import PropTypes from "prop-types";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faTrash } from "@fortawesome/free-solid-svg-icons";
import { isFunction } from "lodash";
import { generateExtraClass } from "../util";

const TagBox = ({ isDelBtn, extraClasses, tag, onClick }) => {
  isDelBtn = isDelBtn ? isDelBtn : false;
  tag.hashCode = tag.hashCode ? tag.hashCode : "#000000";
  let extraClassName = generateExtraClass(extraClasses);

  const clickable = isFunction(onClick) ? "clickable" : "unclickable";
  const onClickTag = (tag) => {
    if (isFunction(onClick)) {
      onClick(tag);
    }
  };

  return (
    <div
      className={`tag-box ${clickable} ${extraClassName}`}
      style={{ backgroundColor: tag.hashCode }}
      onClick={() => onClickTag(tag)}
    >
      <div key={tag}>{tag.label}</div>
      {isDelBtn && (
        <div className='delete-btn'>
          <FontAwesomeIcon icon={faTrash} />
        </div>
      )}
    </div>
  );
};

TagBox.propTypes = {
  isDelBtn: PropTypes.bool,
  extraClasses: PropTypes.array,
  tag: PropTypes.object,
  onClick: PropTypes.func,
};

export default TagBox;
