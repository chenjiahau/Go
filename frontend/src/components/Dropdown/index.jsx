import { has } from "lodash";
import { useRef } from "react";

const Dropdown = (props) => {
  const { hasBackground, list, selectedItem, onChange, zIndex } = props;
  const dropdownRef = useRef(null);
  const body = document.querySelector("body");

  const handleClick = () => {
    const contentRight = document.querySelector(".content__right");
    const options = dropdownRef.current.querySelector(".options");
    const arrowUp = dropdownRef.current.querySelector(".arrow-up");
    const arrowDown = dropdownRef.current.querySelector(".arrow-down");

    if (options.classList.contains("active")) {
      options.classList.remove("active");
      arrowUp.classList.add("hidden");
      arrowDown.classList.remove("hidden");
      contentRight.classList.remove("space-b-6");
    } else {
      options.classList.add("active");
      arrowUp.classList.remove("hidden");
      arrowDown.classList.add("hidden");
      contentRight.classList.add("space-b-6");
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
    <div
      className='dropdown'
      style={{ zIndex: zIndex || 3 }}
      onClick={handleClick}
      ref={dropdownRef}
    >
      <div
        className='item'
        style={{
          width: "100%",
          backgroundColor: hasBackground ? selectedItem?.bgcolor : "",
        }}
      >
        <span
          className='text'
          style={{
            color: hasBackground ? "white" : "",
            backgroundColor: hasBackground ? selectedItem?.bgcolor : "",
          }}
        >
          {selectedItem?.name}
        </span>
        <span>
          <i className='fa-solid fa-sort-up arrow-up hidden'></i>
          <i className='arrow fa-solid fa-sort-down arrow-down'></i>
        </span>
      </div>
      <div
        className='options'
        style={{ width: "100%", maxHeight: "320px", overflowY: "auto" }}
      >
        {list.map((item) => {
          return (
            <div
              key={item.id}
              className={`option ${
                item.id === selectedItem?.id ? "active" : ""
              }`}
              style={{
                color: hasBackground ? "white" : "",
                backgroundColor: hasBackground ? item.bgcolor : "",
              }}
              onClick={() => onChange(item)}
            >
              <span
                style={{ backgroundColor: hasBackground ? item.bgcolor : "" }}
              >
                {item.name}
              </span>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default Dropdown;
