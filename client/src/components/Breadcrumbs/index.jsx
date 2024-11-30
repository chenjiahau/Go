import "./module.css";

import { Link } from "react-router-dom";

const Breadcrumbs = (prop) => {
  const { linkList } = prop;

  return (
    <ul className='breadcrumb'>
      {linkList.map((link, index) => {
        return (
          <li key={index}>
            {link.to ? (
              <Link to={link.to}>{link.label}</Link>
            ) : (
              <span>{link.label}</span>
            )}
          </li>
        );
      })}
    </ul>
  );
};

export default Breadcrumbs;
