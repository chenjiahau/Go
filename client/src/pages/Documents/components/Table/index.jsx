// Const

// Component
import Input from "@/components/Input";

// Util
import { formatDateTime } from "@/util/datetime.util";

const Table = (props) => {
  const {
    currentPage,
    pageSize,
    orderBy,
    order,
    onChangeOrder,
    documents,
    onClickDocumentName,
    onChangeDocumentName,
    selectedDocument,
    onShowConfirmationModal,
    saveDocument,
    onLinkToDocument,
  } = props;

  // State

  // Method

  return (
    <>
      <div className='section'>
        <div className='table-response'>
          <table className='table table-hover'>
            <thead>
              <tr>
                <th width='30'>
                  <div className='n'>#</div>
                </th>
                <th onClick={() => onChangeOrder("name")}>
                  <div className='order'>
                    <span>Document</span>
                    {orderBy === "name" && (
                      <i
                        className={`fa-solid fa-arrow-${
                          order === "asc" ? "down" : "up"
                        }`}
                      />
                    )}
                  </div>
                </th>
                <th onClick={() => onChangeOrder("category")}>
                  <div className='order'>
                    <span>Category</span>
                    {orderBy === "category" && (
                      <i
                        className={`fa-solid fa-arrow-${
                          order === "asc" ? "down" : "up"
                        }`}
                      />
                    )}
                  </div>
                </th>
                <th onClick={() => onChangeOrder("subcategory")}>
                  <div className='order'>
                    <span>Subcategory</span>
                    {orderBy === "subcategory" && (
                      <i
                        className={`fa-solid fa-arrow-${
                          order === "asc" ? "down" : "up"
                        }`}
                      />
                    )}
                  </div>
                </th>
                <th onClick={() => onChangeOrder("postMember")}>
                  <div className='order'>
                    <span>Author</span>
                    {orderBy === "postMember" && (
                      <i
                        className={`fa-solid fa-arrow-${
                          order === "asc" ? "down" : "up"
                        }`}
                      />
                    )}
                  </div>
                </th>
                <th onClick={() => onChangeOrder("created")}>
                  <div className='order'>
                    <span>Created Date</span>
                    {orderBy === "created" && (
                      <i
                        className={`fa-solid fa-arrow-${
                          order === "asc" ? "down" : "up"
                        }`}
                      />
                    )}
                  </div>
                </th>
                <th width='50'>
                  <div className='n'></div>
                </th>
              </tr>
            </thead>
            <tbody>
              {documents.map((document, index) => (
                <tr
                  key={index}
                  className={`${
                    selectedDocument?.id && selectedDocument?.id === document.id
                      ? "selected"
                      : ""
                  }`}
                >
                  <td>
                    <div>{index + 1 + (currentPage - 1) * pageSize}</div>
                  </td>
                  <td>
                    {document.isEditing ? (
                      <div className='input-group edit-input'>
                        <Input
                          id={`documentName-${document.id}`}
                          type='text'
                          extraClasses={["no-border"]}
                          value={document.name}
                          onChange={(e) =>
                            onChangeDocumentName(document.id, e.target.value)
                          }
                          onKeyDown={(e) => {
                            if (e.key === "Enter") {
                              saveDocument(document.id);
                            }

                            if (e.key === "Escape") {
                              onChangeDocumentName(
                                document.id,
                                document.originalName
                              );
                              onClickDocumentName(document.id);
                            }
                          }}
                        />
                        <div className='edit-input-icon'>
                          <div>
                            <i
                              className='fa-solid fa-check'
                              onClick={() => saveDocument(document.id)}
                            />
                          </div>
                          <div>
                            <i
                              className='fa-solid fa-xmark'
                              onClick={() => {
                                onChangeDocumentName(
                                  document.id,
                                  document.originalName
                                );
                                onClickDocumentName(document.id);
                              }}
                            />
                          </div>
                        </div>
                      </div>
                    ) : (
                      <div
                        className='edit-button height'
                        onClick={() => onClickDocumentName(document.id)}
                      >
                        {document.name}
                      </div>
                    )}
                  </td>
                  <td>
                    <div>{document.category?.name}</div>
                  </td>
                  <td>
                    <div>{document.subCategory?.name}</div>
                  </td>
                  <td>
                    <div>{document.postMember?.name}</div>
                  </td>
                  <td>
                    <div>{formatDateTime(document.createdAt)}</div>
                  </td>
                  <td>
                    <div className='action'>
                      <i
                        className='fa-solid fa-trash'
                        onClick={() => onShowConfirmationModal(document.id)}
                      />
                      <i
                        className='fa-solid fa-file'
                        onClick={() => onLinkToDocument(document)}
                      />
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </>
  );
};

export default Table;
