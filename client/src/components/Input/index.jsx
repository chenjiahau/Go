import { isArray } from "lodash";

const Input = ({ extraClasses, type, children, ...props }) => {
  if (children) {
    return (
      <div className='input-with-icon'>
        <input type={type} {...props} />
        {children}
      </div>
    );
  }

  let extraClassName = "";
  if (isArray(extraClasses)) {
    extraClassName = extraClasses.join(" ");
  }

  return <input className={`${extraClassName}`} {...props} />;
};

export default Input;
