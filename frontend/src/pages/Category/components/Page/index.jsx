import Pagination from "@/components/Pagination";

const Page = ({
  currentPage,
  setCurrentPage,
  pageSize,
  setPageSize,
  totalCategoryCount,
}) => {
  return (
    <div className='section section--center d-flex'>
      <Pagination
        className='pagination-bar'
        currentPage={currentPage}
        totalCount={totalCategoryCount}
        pageSize={pageSize}
        setPageSize={setPageSize}
        onPageChange={(page) => setCurrentPage(page)}
      />
    </div>
  );
};

export default Page;
