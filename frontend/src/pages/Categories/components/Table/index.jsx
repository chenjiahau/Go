// Component
import Input from "@/components/Input";
import Button from "@/components/Button";

const Table = (props) => {
  const {
    categories,
    onClickCategoryName,
    changeCategoryName,
    changeCategoryAlive,
    selectedCategory,
    onShowConfirmationModal,
    saveCategory,
  } = props;

  return (
    <div className='section'>
      <div className='table-response'>
        <table className='table table-hover'>
          <thead>
            <tr>
              <th>Name</th>
              <th className='text-center' width='160'>
                Subcategory
              </th>
              <th width='160'>Status</th>
              <th width='160'>Action</th>
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
                    {category.subcategories.length}
                  </div>
                </td>
                <td>
                  <div className='height'>
                    <input
                      className='space-r-2'
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
                    <input
                      className='space-r-2'
                      type='radio'
                      checked={!category.isAlive}
                      onChange={() => {
                        changeCategoryAlive(category.id, false);
                        saveCategory(category.id);
                      }}
                    />
                    <label htmlFor='status'>Disable</label>
                  </div>
                </td>
                <td className='table-action-td'>
                  <div>
                    <Button
                      extraClasses={["delete-button", "delete-category"]}
                      disabled={category.subcategories.length > 0}
                      onClick={() => onShowConfirmationModal(category.id)}
                    >
                      Delete
                    </Button>
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
