import "./module.css";

import { useRef } from "react";
import PropTypes from "prop-types";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faDownLong, faUpLong } from "@fortawesome/free-solid-svg-icons";

const PaginationDropdownBox = (props) => {
  const { numberOfRow, pageSize, onChangePageSize } = props;
  const dropdownRef = useRef(null);
  const body = document.querySelector("body");

  const handleClick = () => {
    const mainContent = document.querySelector(".main-content");
    const options = dropdownRef.current.querySelector(".options");
    const arrowUp = dropdownRef.current.querySelector(".arrow-up");
    const arrowDown = dropdownRef.current.querySelector(".arrow-down");

    if (options.classList.contains("active")) {
      options.classList.remove("active");
      arrowUp.classList.add("hidden");
      arrowDown.classList.remove("hidden");
      mainContent.classList.remove("space");
    } else {
      options.classList.add("active");
      arrowUp.classList.remove("hidden");
      arrowDown.classList.add("hidden");
      mainContent.classList.add("space");
    }
  };

  body.addEventListener("click", (e) => {
    if (!dropdownRef.current) return;

    if (!dropdownRef.current.contains(e.target)) {
      const mainContent = document.querySelector(".main-content");
      const options = dropdownRef.current.querySelector(".options");
      const arrowUp = dropdownRef.current.querySelector(".arrow-up");
      const arrowDown = dropdownRef.current.querySelector(".arrow-down");

      options.classList.remove("active");
      arrowUp.classList.add("hidden");
      arrowDown.classList.remove("hidden");
      mainContent.classList.remove("space");
    }
  });

  return (
    <div
      className='pagination-dropdown-box'
      onClick={handleClick}
      ref={dropdownRef}
    >
      <div className='item'>
        <span className='text'>{pageSize}</span>
        <span>
          <FontAwesomeIcon icon={faUpLong} className='arrow-up hidden' />
          <FontAwesomeIcon icon={faDownLong} className='arrow-down' />
        </span>
      </div>
      <div className='options'>
        {numberOfRow.map((size) => {
          return (
            <div
              key={size}
              className={`option ${size === pageSize ? "active" : ""}`}
              onClick={() => onChangePageSize(size)}
            >
              <span>{size}</span>
            </div>
          );
        })}
      </div>
    </div>
  );
};

PaginationDropdownBox.propTypes = {
  numberOfRow: PropTypes.arrayOf(PropTypes.number).isRequired,
  pageSize: PropTypes.number.isRequired,
  onChangePageSize: PropTypes.func.isRequired,
};

export default PaginationDropdownBox;
