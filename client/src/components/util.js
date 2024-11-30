import { isArray } from "lodash";

const generateExtraClass = (extraClasses) => {
  let extraClassName = "";
  if (isArray(extraClasses)) {
    extraClassName = extraClasses.join(" ");
  }

  return extraClassName;
}

export { generateExtraClass };