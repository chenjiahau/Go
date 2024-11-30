import "./module.css";

import PropTypes from "prop-types";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faLeftLong, faRightLong } from "@fortawesome/free-solid-svg-icons";
import PaginationDropdownBox from "@/components/PaginationDropdownBox";
import { usePagination, DOTS } from "@/hooks/userPagination";

export const numberOfRow = [5, 10, 20, 50];

const PaginationBox = (props) => {
  const {
    onPageChange,
    totalCount,
    siblingCount = 1,
    currentPage,
    pageSize,
    setPageSize,
  } = props;

  const paginationRange = usePagination({
    currentPage,
    totalCount,
    siblingCount,
    pageSize,
  });

  const onNext = () => {
    if (currentPage === paginationRange[paginationRange.length - 1]) return;
    onPageChange(currentPage + 1);
  };

  const onPrevious = () => {
    if (currentPage === 1) return;
    onPageChange(currentPage - 1);
  };

  let lastPage = paginationRange[paginationRange.length - 1];

  return (
    <div className='pagination-box'>
      <div
        className={`pagination-item arrow ${
          currentPage === 1 ? "disabled" : ""
        }`}
        onClick={onPrevious}
      >
        <FontAwesomeIcon icon={faLeftLong} />
        <span>Prev</span>
      </div>
      {paginationRange.map((pageNumber, index) => {
        if (pageNumber === DOTS) {
          return (
            <div key={index} className='pagination-item dots'>
              {DOTS}
            </div>
          );
        }

        const isCurrent = currentPage === pageNumber;

        return (
          <div
            className={`pagination-item ${isCurrent ? "active" : ""}`}
            key={index}
            onClick={() => !isCurrent && onPageChange(pageNumber)}
          >
            {pageNumber}
          </div>
        );
      })}
      <div
        className={`pagination-item arrow ${
          !lastPage || currentPage === lastPage ? "disabled" : ""
        }`}
        onClick={onNext}
      >
        <span>Next</span>
        <FontAwesomeIcon icon={faRightLong} />
      </div>
      <div>
        <PaginationDropdownBox
          numberOfRow={numberOfRow}
          pageSize={pageSize}
          onChangePageSize={(size) => {
            setPageSize(size);
            onPageChange(1);
          }}
        />
      </div>
    </div>
  );
};

PaginationBox.propTypes = {
  onPageChange: PropTypes.func.isRequired,
  totalCount: PropTypes.number.isRequired,
  siblingCount: PropTypes.number,
  currentPage: PropTypes.number.isRequired,
  pageSize: PropTypes.number.isRequired,
  setPageSize: PropTypes.func.isRequired,
};

export default PaginationBox;
