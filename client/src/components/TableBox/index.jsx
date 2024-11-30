import "./module.css";

import { useRef } from "react";
import PropTypes from "prop-types";
import { cloneDeep, isFunction } from "lodash";

const TableBox = ({ headers, onChangeHeader, data }) => {
  const tableRef = useRef(null);

  const onClickSort = (header) => {
    if (isFunction(onChangeHeader)) {
      let column = null;
      let order = null;

      const updatedHeaders = cloneDeep(headers).map((h) => {
        if (h.key === header.key) {
          h.sort = h.sort === "asc" ? "desc" : "asc";
          column = h.key;
          order = h.sort;
        } else {
          h.sort = null;
        }

        return h;
      });

      onChangeHeader(updatedHeaders, column, order);
    }
  };

  return (
    <>
      <div className='table-box'>
        <table ref={tableRef}>
          <thead>
            <tr>
              {headers.map((header) => {
                const center = header.isCenter ? "center" : "";
                const sortable = header.isSortable ? "sortable" : "";
                const sort = header.sort ? "sort" : "";
                const sorted = header.sort ? "sorted" : "";

                return (
                  <th
                    key={header.key}
                    className={`${sortable} ${sort}`}
                    width={header.width}
                  >
                    {sortable ? (
                      <div
                        className={`sortable-box ${sorted} ${center}`}
                        onClick={() => onClickSort(header)}
                      >
                        <div className='sort-title'>{header.label}</div>
                        <span className='sort-icon'>
                          {header.sort === "asc" && (
                            <svg
                              className='w-5 h-5 ml-2 -mr-1 text-gray-400'
                              xmlns='http://www.w3.org/2000/svg'
                              fill='none'
                              viewBox='0 0 24 24'
                              stroke='currentColor'
                            >
                              <path
                                strokeLinecap='round'
                                strokeLinejoin='round'
                                strokeWidth='2'
                                d='M19 9l-7 7-7-7'
                              />
                            </svg>
                          )}
                          {header.sort === "desc" && (
                            <svg
                              className='w-5 h-5 ml-2 -mr-1 text-gray-400'
                              xmlns='http://www.w3.org/2000/svg'
                              fill='none'
                              viewBox='0 0 24 24'
                              stroke='currentColor'
                            >
                              <path
                                strokeLinecap='round'
                                strokeLinejoin='round'
                                strokeWidth='2'
                                d='M5 15l7-7 7 7'
                              />
                            </svg>
                          )}
                        </span>
                      </div>
                    ) : (
                      <div className='unsortable-box'>{header.label}</div>
                    )}
                  </th>
                );
              })}
            </tr>
          </thead>
          <tbody>
            {data.map((row, parentIndex) => {
              return (
                <tr key={parentIndex} data-key={row.index}>
                  {headers.map((header, childIndex) => {
                    const center = header.isCenter ? "center" : "";

                    if (header.key === "index") {
                      return (
                        <td
                          key={`${row.index}-${childIndex}`}
                          data-key={`${parentIndex}-${childIndex}`}
                        >
                          <div className={`${center} td-box`}>{row.index}</div>
                        </td>
                      );
                    }

                    return (
                      <td
                        key={`${row.index}-${childIndex}`}
                        data-key={`${parentIndex}-${childIndex}`}
                      >
                        <div className={`${center} td-box`}>
                          {row[header.key]}
                        </div>
                      </td>
                    );
                  })}
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>
      <div className='vertical-table'>
        {data.map((row, parentIndex) => {
          return (
            <div key={parentIndex} data-key={row.index} className='row-box'>
              {headers.map((header, childIndex) => {
                const center = header.isCenter ? "center" : "";
                const sortable = header.isSortable ? "sortable" : "";
                const sorted = header.sort ? "sorted" : "";

                if (header.key === "index") {
                  return (
                    <div
                      key={`${row.index}-${childIndex}`}
                      data-key={`${parentIndex}-${childIndex}`}
                      className='col-box'
                    >
                      <div className='th-box'>#</div>
                      <div className={`${center} td-box`}>{row.index}</div>
                    </div>
                  );
                }

                return (
                  <div
                    key={`${row.index}-${childIndex}`}
                    data-key={`${parentIndex}-${childIndex}`}
                    className='col-box'
                  >
                    <div
                      className={`${
                        sortable ? "th-sort-box" : "th-box"
                      } ${center} ${sortable && sorted}`}
                      onClick={() => onClickSort(header)}
                    >
                      <div className='sort-title'>{header.label}</div>
                      {sortable && (
                        <span className='sort-icon'>
                          {header.sort === "asc" && (
                            <svg
                              className='w-5 h-5 ml-2 -mr-1 text-gray-400'
                              xmlns='http://www.w3.org/2000/svg'
                              fill='none'
                              viewBox='0 0 24 24'
                              stroke='currentColor'
                            >
                              <path
                                strokeLinecap='round'
                                strokeLinejoin='round'
                                strokeWidth='2'
                                d='M19 9l-7 7-7-7'
                              />
                            </svg>
                          )}
                          {header.sort === "desc" && (
                            <svg
                              className='w-5 h-5 ml-2 -mr-1 text-gray-400'
                              xmlns='http://www.w3.org/2000/svg'
                              fill='none'
                              viewBox='0 0 24 24'
                              stroke='currentColor'
                            >
                              <path
                                strokeLinecap='round'
                                strokeLinejoin='round'
                                strokeWidth='2'
                                d='M5 15l7-7 7 7'
                              />
                            </svg>
                          )}
                        </span>
                      )}
                    </div>
                    <div
                      className={`${center} td-box ${
                        childIndex === 1 ? "td-max" : ""
                      }`}
                    >
                      {row[header.key]}
                    </div>
                  </div>
                );
              })}
            </div>
          );
        })}
      </div>
    </>
  );
};

TableBox.propTypes = {
  headers: PropTypes.arrayOf(
    PropTypes.shape({
      key: PropTypes.string.isRequired,
      label: PropTypes.string.isRequired,
    })
  ).isRequired,
  onChangeHeader: PropTypes.func,
  data: PropTypes.array,
};

export default TableBox;
