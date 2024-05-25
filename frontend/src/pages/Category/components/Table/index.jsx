const Table = (props) => {
  const {
    categories,
    onClickCategoryName,
    changeCategoryName,
    changeCategoryAlive,
    selectedCategory,
    onShowSubcategoriesModal,
    onShowConfirmationModal,
    saveCategory,
  } = props;

  return (
    <div className='section'>
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
                  <div className='edit-input'>
                    <input
                      type='text'
                      className='form-control'
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
                    className='edit-button'
                    onClick={() => onClickCategoryName(category.id)}
                  >
                    {category.name}
                  </div>
                )}
              </td>
              <td>
                <div
                  className='text-center edit-button'
                  onClick={() => onShowSubcategoriesModal(category.id)}
                >
                  {category.subcategories.length}
                </div>
              </td>
              <td>
                <div>
                  <input
                    className='me-2'
                    type='radio'
                    checked={category.isAlive}
                    onChange={() => {
                      changeCategoryAlive(category.id, true);
                      saveCategory(category.id);
                    }}
                  />
                  <label className='me-2' htmlFor='status'>
                    Enable
                  </label>
                  <input
                    className='me-2'
                    type='radio'
                    checked={!category.isAlive}
                    onChange={() => {
                      changeCategoryAlive(category.id, false);
                      saveCategory(category.id);
                    }}
                  />
                  <label className='me-2' htmlFor='status'>
                    Disable
                  </label>
                </div>
              </td>
              <td className='table-action-td'>
                <div>
                  <button
                    className='button delete-button delete-category'
                    disabled={category.subcategories.length > 0}
                    onClick={() => onShowConfirmationModal(category.id)}
                  >
                    Delete
                  </button>
                </div>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Table;
