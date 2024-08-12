// Component
import Input from "@/components/Input";
import RadioButton from "@/components/RadioButton";

const Table = (props) => {
  const {
    currentPage,
    pageSize,
    orderBy,
    order,
    onChangeOrder,
    subcategories,
    onClickSubcategoryName,
    onChangeSubcategoryName,
    onChangeSubcategoryAlive,
    selectedSubcategory,
    onShowConfirmationModal,
    onsSaveSubcategory,
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
                  <span>Name</span>
                  {orderBy === "name" && (
                    <i
                      className={`fa-solid fa-arrow-${
                        order === "asc" ? "down" : "up"
                      }`}
                    />
                  )}
                </div>
              </th>
              <th width='160' onClick={() => onChangeOrder("status")}>
                <div className='order'>
                  <span>Status</span>
                  {orderBy === "status" && (
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
            {subcategories.map((subcategory, index) => (
              <tr
                key={index}
                className={`${
                  selectedSubcategory?.id &&
                  selectedSubcategory?.id === subcategory.id
                    ? "selected"
                    : ""
                }`}
              >
                <td>
                  <div>{index + 1 + (currentPage - 1) * pageSize}</div>
                </td>
                <td>
                  {subcategory.isEditing ? (
                    <div className='input-group edit-input'>
                      <Input
                        id={`subcategoryName-${subcategory.id}`}
                        type='text'
                        extraClasses={["no-border"]}
                        value={subcategory.name}
                        onChange={(e) =>
                          onChangeSubcategoryName(
                            subcategory.id,
                            e.target.value
                          )
                        }
                        onKeyDown={(e) => {
                          if (e.key === "Enter") {
                            onsSaveSubcategory(subcategory.id);
                          }

                          if (e.key === "Escape") {
                            onChangeSubcategoryName(
                              subcategory.id,
                              subcategory.originalName
                            );
                            onClickSubcategoryName(subcategory.id);
                          }
                        }}
                      />
                      <div className='edit-input-icon'>
                        <div>
                          <i
                            className='fa-solid fa-check'
                            onClick={() => onsSaveSubcategory(subcategory.id)}
                          />
                        </div>
                        <div>
                          <i
                            className='fa-solid fa-xmark'
                            onClick={() => {
                              onChangeSubcategoryName(
                                subcategory.id,
                                subcategory.originalName
                              );
                              onClickSubcategoryName(subcategory.id);
                            }}
                          />
                        </div>
                      </div>
                    </div>
                  ) : (
                    <div
                      className='edit-button height'
                      onClick={() => onClickSubcategoryName(subcategory.id)}
                    >
                      {subcategory.name}
                    </div>
                  )}
                </td>
                <td>
                  <div className='height'>
                    <RadioButton
                      extraClasses={["space-r-2"]}
                      type='radio'
                      checked={subcategory.isAlive}
                      onChange={() => {
                        onChangeSubcategoryAlive(subcategory.id, true);
                        onsSaveSubcategory(subcategory.id);
                      }}
                    />
                    <label className='space-r-3' htmlFor='status'>
                      Enable
                    </label>
                    <RadioButton
                      type='radio'
                      extraClasses={["space-r-2"]}
                      checked={!subcategory.isAlive}
                      onChange={() => {
                        onChangeSubcategoryAlive(subcategory.id, false);
                        onsSaveSubcategory(subcategory.id);
                      }}
                    />
                    <label htmlFor='status'>Disable</label>
                  </div>
                </td>
                <td>
                  <div className='action'>
                    <i
                      className='fa-solid fa-trash'
                      onClick={() => onShowConfirmationModal(subcategory.id)}
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
