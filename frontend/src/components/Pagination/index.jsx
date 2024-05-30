import PaginationDropdown from "@/components/PaginationDropdown";
import { usePagination, DOTS } from "@/hooks/userPagination";

export const pageSizeDefinition = [5, 10, 20, 50];

const Pagination = (props) => {
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
    <div className='pagination'>
      <div
        className={`pagination-item arrow ${
          currentPage === 1 ? "disabled" : ""
        }`}
        onClick={onPrevious}
      >
        <i className='fa-solid fa-caret-left'></i>
        <span>Previous</span>
      </div>
      {paginationRange.map((pageNumber, index) => {
        if (pageNumber === DOTS) {
          return (
            <div key={index} className='pagination-item arrow dots'>
              {DOTS}
            </div>
          );
        }

        return (
          <div
            className={`pagination-item ${
              currentPage === pageNumber ? "active" : ""
            }`}
            key={index}
            onClick={() => onPageChange(pageNumber)}
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
        <i className='fa-solid fa-caret-right'></i>
      </div>
      <PaginationDropdown
        pageSizeDefinition={pageSizeDefinition}
        pageSize={pageSize}
        onChangePageSize={(size) => {
          setPageSize(size);
          onPageChange(1);
        }}
      />
    </div>
  );
};

export default Pagination;
