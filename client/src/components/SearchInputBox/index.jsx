import "./module.css";

import searchIcon from "@/assets/img/search.svg";
import PropTypes from "prop-types";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faXmark } from "@fortawesome/free-solid-svg-icons";
import { isFunction } from "lodash";

const SearchInputBox = ({ delay, onChange, onBlur }) => {
  !delay && (delay = 500);
  !isFunction(onChange) && (onChange = () => {});
  !isFunction(onBlur) && (onBlur = () => {});

  const debounce = (func) => {
    let debounceTimer;

    return function () {
      const context = this;
      const args = arguments;
      clearTimeout(debounceTimer);
      debounceTimer = setTimeout(() => func.apply(context, args), delay);
    };
  };

  const clearInput = () => {
    let lastValue = document.querySelector(".search-input-box input").value;
    document.querySelector(".search-input-box input").value = "";

    if (lastValue) {
      onChange("");
    }
  };

  return (
    <div className='search-input-box primary-shadow'>
      <div className='icon'>
        <img src={searchIcon} alt='search' />
      </div>
      <input
        type='text'
        placeholder='Search...'
        onChange={debounce(onChange)}
        onBlur={onBlur}
      />
      <div className='icon cancel-icon' onClick={clearInput}>
        <FontAwesomeIcon icon={faXmark} />
      </div>
    </div>
  );
};

SearchInputBox.propTypes = {
  delay: PropTypes.number,
  onChange: PropTypes.func,
  onBlur: PropTypes.func,
};

export default SearchInputBox;
