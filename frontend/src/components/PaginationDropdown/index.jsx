import { useRef } from "react";

const PaginationDropdown = (props) => {
  const { pageSizeDefinition, pageSize, onChangePageSize } = props;
  const dropdownRef = useRef(null);
  const body = document.querySelector("body");

  const handleClick = () => {
    const options = dropdownRef.current.querySelector(".options");
    const arrowUp = dropdownRef.current.querySelector(".arrow-up");
    const arrowDown = dropdownRef.current.querySelector(".arrow-down");

    if (options.classList.contains("active")) {
      options.classList.remove("active");
      arrowUp.classList.add("hidden");
      arrowDown.classList.remove("hidden");
    } else {
      options.classList.add("active");
      arrowUp.classList.remove("hidden");
      arrowDown.classList.add("hidden");
    }
  };

  body.addEventListener("click", (e) => {
    if (!dropdownRef.current) return;

    if (!dropdownRef.current.contains(e.target)) {
      const options = dropdownRef.current.querySelector(".options");
      const arrowUp = dropdownRef.current.querySelector(".arrow-up");
      const arrowDown = dropdownRef.current.querySelector(".arrow-down");

      options.classList.remove("active");
      arrowUp.classList.add("hidden");
      arrowDown.classList.remove("hidden");
    }
  });

  return (
    <div className='dropdown' onClick={handleClick} ref={dropdownRef}>
      <div className='item'>
        <span className='text'>{pageSize}</span>
        <span>
          <i className='fa-solid fa-sort-up arrow-up hidden'></i>
          <i className='arrow fa-solid fa-sort-down arrow-down'></i>
        </span>
      </div>
      <div className='options'>
        {pageSizeDefinition.map((size) => {
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

export default PaginationDropdown;
