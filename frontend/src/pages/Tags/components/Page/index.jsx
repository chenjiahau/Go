import Pagination from "@/components/Pagination";

const Page = ({
  currentPage,
  setCurrentPage,
  pageSize,
  setPageSize,
  totalTagCount,
}) => {
  return (
    <div className='section section--center d-flex'>
      <Pagination
        className='pagination-bar'
        currentPage={currentPage}
        totalCount={totalTagCount}
        pageSize={pageSize}
        setPageSize={setPageSize}
        onPageChange={(page) => setCurrentPage(page)}
      />
    </div>
  );
};

export default Page;
