import { isArray } from "lodash";

const Button = ({ children, extraClasses, ...props }) => {
  let extraClassName = "";
  if (isArray(extraClasses)) {
    extraClassName = extraClasses.join(" ");
  }

  return (
    <div className={`general-button ${extraClassName}`} {...props}>
      {children}
    </div>
  );
};

export default Button;
