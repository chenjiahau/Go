import { useEffect, useState } from "react";
import Modal from "react-bootstrap/Modal";

// Const
import apiConfig from "@/const/config/api";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

// Component
import ConfirmationModal from "@/components/ConfirmationModal";
import { orderBy } from "lodash";

const errorMessage = {
  category: "Subcategory is required.",
};

const Subcategories = (props) => {
  const { isOpen, onClose, category, onInitialization } = props;

  // State
  const [subcategory, setSubcategory] = useState("");
  const [subcategories, setSubcategories] = useState([]);
  const [selectedSubcategory, setSelectedSubcategory] = useState({});
  const [isOpenConfirmationModal, setIsOpenConfirmationModal] = useState(false);

  // Method
  const handleInitialization = (category) => {
    if (category?.subcategories) {
      let updatedSubcategories = [];
      category.subcategories.forEach((subcategory) => {
        updatedSubcategories.push({
          ...subcategory,
          originalName: subcategory.name,
          isEditing: false,
        });
      });
      updatedSubcategories = orderBy(updatedSubcategories, ["id"], ["asc"]);

      setSubcategories(updatedSubcategories);
    }
  };

  const addSubcategory = async () => {
    if (subcategory === "") {
      messageUtil.showErrorMessage(errorMessage.category);
      return;
    }

    try {
      await apiHandler.post(
        apiConfig.resource.ADD_SUBCATEGORY.replace(":id", category.id),
        {
          name: subcategory,
        }
      );
      setSubcategory("");
      messageUtil.showSuccessMessage(commonMessage.success);
      onInitialization();
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  const clickSubcategoryName = (id) => {
    const updatedSubcategories = subcategories.map((subcategory) => {
      if (subcategory.id === id) {
        subcategory.isEditing = !subcategory.isEditing;
      }

      return subcategory;
    });

    setSubcategories(updatedSubcategories);
  };

  const changeSubcategoryName = (id, name) => {
    const updatedSubcategories = subcategories.map((subcategory) => {
      if (subcategory.id === id) {
        subcategory.name = name;
      }

      return subcategory;
    });

    setSubcategories(updatedSubcategories);
  };

  const changeSubcategoryAlive = async (id, alive) => {
    const updatedSubcategories = subcategories.map((subcategory) => {
      if (subcategory.id === id) {
        subcategory.isAlive = alive;
      }

      return subcategory;
    });

    setSubcategories(updatedSubcategories);
  };

  const saveSubcategory = async (id) => {
    const subcategoryIndex = subcategories.findIndex(
      (subcategory) => subcategory.id === id
    );

    const updatedSubcategories = [...subcategories];

    if (updatedSubcategories[subcategoryIndex].name === "") {
      return;
    }

    try {
      const subcategory = subcategories.find(
        (subcategory) => subcategory.id === id
      );
      const apiURL = apiConfig.resource.EDIT_SUBCATEGORY.replace(
        ":id",
        category.id
      ).replace(":subId", id);
      const payload = {
        name: subcategory.name,
        isAlive: subcategory.isAlive,
      };
      await apiHandler.put(apiURL, payload);
      messageUtil.showSuccessMessage(commonMessage.success);

      updatedSubcategories[subcategoryIndex].originalName = subcategory.name;
      updatedSubcategories[subcategoryIndex].isEditing = false;
      setSubcategories(updatedSubcategories);
      onInitialization();
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  const deleteSubcategory = async () => {
    try {
      const apiURL = apiConfig.resource.EDIT_SUBCATEGORY.replace(
        ":id",
        category.id
      ).replace(":subId", selectedSubcategory.id);
      await apiHandler.delete(apiURL);
      messageUtil.showSuccessMessage(commonMessage.success);
      setIsOpenConfirmationModal(false);
      onInitialization();
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  const showConfirmationModal = (id) => {
    setSelectedSubcategory(
      subcategories.find((subcategory) => subcategory.id === id)
    );
    setIsOpenConfirmationModal(true);
  };

  // Side effect
  useEffect(() => {
    handleInitialization(category);
  }, [category]);

  return (
    <>
      <Modal fullscreen={true} scrollable={true} show={isOpen} onHide={onClose}>
        <Modal.Header closeButton>
          <Modal.Title>Subcategory</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <div className='section mb-2'>
            <div className='mb-2'>
              <input
                type='text'
                name='subcategory'
                id='subcategory'
                className='form-control'
                autoComplete='off'
                placeholder='New subcategory'
                value={subcategory}
                onChange={(e) => setSubcategory(e.target.value)}
                onKeyDown={(e) => {
                  if (e.key === "Enter") {
                    addSubcategory();
                  }
                }}
                required
              />
            </div>
            <div className='mb-2 text-right'>
              <button className='button w-100' onClick={addSubcategory}>
                <i className='fa-solid fa-plus'></i>
                <span className='ms-1'>Add</span>
              </button>
            </div>
          </div>

          <div className='section'>
            <table className='table table-hover'>
              <thead>
                <tr>
                  <th>Name</th>
                  <th width='160'>Status</th>
                  <th width='160'>Action</th>
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
                      {subcategory.isEditing ? (
                        <div className='edit-input'>
                          <input
                            type='text'
                            className='form-control'
                            value={subcategory.name}
                            onChange={(e) =>
                              changeSubcategoryName(
                                subcategory.id,
                                e.target.value
                              )
                            }
                            onKeyDown={(e) => {
                              if (e.key === "Enter") {
                                saveSubcategory(subcategory.id);
                              }

                              if (e.key === "Escape") {
                                changeSubcategoryName(
                                  subcategory.id,
                                  subcategory.originalName
                                );
                                clickSubcategoryName(subcategory.id);
                              }
                            }}
                          />
                          <div className='edit-input-icon'>
                            <div>
                              <i
                                className='fa-solid fa-check'
                                onClick={() => saveSubcategory(subcategory.id)}
                              />
                            </div>
                            <div>
                              <i
                                className='fa-solid fa-xmark'
                                onClick={() => {
                                  changeSubcategoryName(
                                    subcategory.id,
                                    subcategory.originalName
                                  );
                                  clickSubcategoryName(subcategory.id);
                                }}
                              />
                            </div>
                          </div>
                        </div>
                      ) : (
                        <div
                          className='edit-button'
                          onClick={() => clickSubcategoryName(subcategory.id)}
                        >
                          {subcategory.name}
                        </div>
                      )}
                    </td>
                    <td>
                      <div>
                        <input
                          className='me-2'
                          type='radio'
                          checked={subcategory.isAlive}
                          onChange={() => {
                            changeSubcategoryAlive(subcategory.id, true);
                            saveSubcategory(subcategory.id);
                          }}
                        />
                        <label className='me-2' htmlFor='status'>
                          Enable
                        </label>
                        <input
                          className='me-2'
                          type='radio'
                          checked={!subcategory.isAlive}
                          onChange={() => {
                            changeSubcategoryAlive(subcategory.id, false);
                            saveSubcategory(subcategory.id);
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
                          onClick={() => showConfirmationModal(subcategory.id)}
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
        </Modal.Body>
        <Modal.Footer>
          <button
            type='button'
            className='button cancel-button'
            onClick={onClose}
          >
            Close
          </button>
        </Modal.Footer>
      </Modal>

      <ConfirmationModal
        isOpen={isOpenConfirmationModal}
        onClose={() => {
          setSelectedSubcategory(null);
          setIsOpenConfirmationModal(false);
        }}
        onConfirm={deleteSubcategory}
      />
    </>
  );
};

export default Subcategories;
