// Component
import Input from "@/components/Input";

const Table = (props) => {
  const {
    currentPage,
    pageSize,
    orderBy,
    order,
    onChangeOrder,
    tags,
    onClickTagName,
    changeTagName,
    selectedTag,
    onShowConfirmationModal,
    saveTag,
  } = props;

  return (
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
                  <span>Tag</span>
                  {orderBy === "name" && (
                    <i
                      className={`fa-solid fa-arrow-${
                        order === "asc" ? "down" : "up"
                      }`}
                    />
                  )}
                </div>
              </th>
              <th width='160' onClick={() => onChangeOrder("color")}>
                <div className='order'>
                  <span>Color</span>
                  {orderBy === "color" && (
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
            {tags.map((tag, index) => (
              <tr
                key={index}
                className={`${
                  selectedTag?.id && selectedTag?.id === tag.id
                    ? "selected"
                    : ""
                }`}
              >
                <td>
                  <div>{index + 1 + (currentPage - 1) * pageSize}</div>
                </td>
                <td>
                  {tag.isEditing ? (
                    <div className='input-group edit-input'>
                      <Input
                        id={`tagName-${tag.id}`}
                        type='text'
                        extraClasses={["no-border"]}
                        value={tag.name}
                        onChange={(e) => changeTagName(tag.id, e.target.value)}
                        onKeyDown={(e) => {
                          if (e.key === "Enter") {
                            saveTag(tag.id);
                          }

                          if (e.key === "Escape") {
                            changeTagName(tag.id, tag.originalName);
                            onClickTagName(tag.id);
                          }
                        }}
                      />
                      <div className='edit-input-icon'>
                        <div>
                          <i
                            className='fa-solid fa-check'
                            onClick={() => saveTag(tag.id)}
                          />
                        </div>
                        <div>
                          <i
                            className='fa-solid fa-xmark'
                            onClick={() => {
                              changeTagName(tag.id, tag.originalName);
                              onClickTagName(tag.id);
                            }}
                          />
                        </div>
                      </div>
                    </div>
                  ) : (
                    <div
                      className='edit-button height'
                      onClick={() => onClickTagName(tag.id)}
                    >
                      {tag.name}
                    </div>
                  )}
                </td>
                <td>
                  <div
                    className='tag'
                    style={{ backgroundColor: tag.colorHexCode }}
                  >
                    {tag.colorName}
                  </div>
                </td>
                <td>
                  <div className='action'>
                    <i
                      className='fa-solid fa-trash'
                      onClick={() => onShowConfirmationModal(tag.id)}
                    />
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default Table;
