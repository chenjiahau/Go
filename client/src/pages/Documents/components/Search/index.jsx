import { useEffect } from "react";

const Search = ({ search, setSearch, onSearchDocuments }) => {
  // Side effect
  useEffect(() => {
    const handler = setTimeout(() => {
      onSearchDocuments();
    }, 1000);

    return () => {
      clearTimeout(handler);
    };
  }, [onSearchDocuments, search]);

  return (
    <div className='section'>
      <div className='input-group'>
        <input
          value={search}
          type='text'
          placeholder='Search'
          className='search-input'
          onChange={(e) => setSearch(e.target.value)}
        />
      </div>
    </div>
  );
};

export default Search;
