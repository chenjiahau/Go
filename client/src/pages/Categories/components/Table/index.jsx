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
    categories,
    onClickCategoryName,
    changeCategoryName,
    changeCategoryAlive,
    selectedCategory,
    onShowConfirmationModal,
    onLinkToEditCategory,
    saveCategory,
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
              <th
                width='160'
                className='text-center'
                onClick={() => onChangeOrder("subcategory")}
              >
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
            {categories.map((category, index) => (
              <tr
                key={index}
                className={`${
                  selectedCategory?.id && selectedCategory?.id === category.id
                    ? "selected"
                    : ""
                }`}
              >
                <td>
                  <div>{index + 1 + (currentPage - 1) * pageSize}</div>
                </td>
                <td>
                  {category.isEditing ? (
                    <div className='input-group edit-input'>
                      <Input
                        id={`categoryName-${category.id}`}
                        type='text'
                        extraClasses={["no-border"]}
                        value={category.name}
                        onChange={(e) =>
                          changeCategoryName(category.id, e.target.value)
                        }
                        onKeyDown={(e) => {
                          if (e.key === "Enter") {
                            saveCategory(category.id);
                          }

                          if (e.key === "Escape") {
                            changeCategoryName(
                              category.id,
                              category.originalName
                            );
                            onClickCategoryName(category.id);
                          }
                        }}
                      />
                      <div className='edit-input-icon'>
                        <div>
                          <i
                            className='fa-solid fa-check'
                            onClick={() => saveCategory(category.id)}
                          />
                        </div>
                        <div>
                          <i
                            className='fa-solid fa-xmark'
                            onClick={() => {
                              changeCategoryName(
                                category.id,
                                category.originalName
                              );
                              onClickCategoryName(category.id);
                            }}
                          />
                        </div>
                      </div>
                    </div>
                  ) : (
                    <div
                      className='edit-button height'
                      onClick={() => onClickCategoryName(category.id)}
                    >
                      {category.name}
                    </div>
                  )}
                </td>
                <td>
                  <div className='height text-center'>
                    {category.subcategoryCount}
                  </div>
                </td>
                <td>
                  <div className='height'>
                    <RadioButton
                      extraClasses={["space-r-2"]}
                      type='radio'
                      checked={category.isAlive}
                      onChange={() => {
                        changeCategoryAlive(category.id, true);
                        saveCategory(category.id);
                      }}
                    />
                    <label className='space-r-3' htmlFor='status'>
                      Enable
                    </label>
                    <RadioButton
                      type='radio'
                      extraClasses={["space-r-2"]}
                      checked={!category.isAlive}
                      onChange={() => {
                        changeCategoryAlive(category.id, false);
                        saveCategory(category.id);
                      }}
                    />
                    <label htmlFor='status'>Disable</label>
                  </div>
                </td>
                <td>
                  <div className='action'>
                    <i
                      className={`fa-solid fa-trash ${
                        category.subcategoryCount > 0 ? "disabled" : ""
                      }`}
                      onClick={() => onShowConfirmationModal(category.id)}
                    />
                    <i
                      className='fa-solid fa-pen'
                      onClick={() => onLinkToEditCategory(category)}
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
