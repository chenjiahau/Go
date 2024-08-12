import { isArray } from "lodash";

const Input = ({ extraClasses, ...props }) => {
  let extraClassName = "";
  if (isArray(extraClasses)) {
    extraClassName = extraClasses.join(" ");
  }

  return <input className={`${extraClassName}`} {...props} />;
};

export default Input;
