import "./module.css";

import PropTypes from "prop-types";
import Form from "@/components/Form";

const CardBox = ({ title, children }) => {
  return (
    <div className='card-box'>
      <Form>
        <h1 className='title'>{title}</h1>
        <div className='divider'></div>
        {children}
      </Form>
    </div>
  );
};

CardBox.propTypes = {
  title: PropTypes.string,
  children: PropTypes.node,
};

export default CardBox;
