import { isArray } from "lodash";

const RadioButton = ({ extraClasses, ...props }) => {
  let extraClassName = "";
  if (isArray(extraClasses)) {
    extraClassName = extraClasses.join(" ");
  }

  return <input type='radio' className={`${extraClassName}`} {...props} />;
};

export default RadioButton;
