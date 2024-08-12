import { useEffect, useState } from "react";
import Modal from "react-bootstrap/Modal";
import { orderBy } from "lodash";

// Const
import apiConfig from "@/const/config/api";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

// Component
import ConfirmationModal from "@/components/ConfirmationModal";
import Input from "@/components/Input";
import Button from "@/components/Button";

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
    if (isOpen) {
      handleInitialization(category);
    }
  }, [category, isOpen]);

  if (!category) {
    return null;
  }

  return (
    <>
      <Modal size='xl' scrollable={true} show={isOpen} onHide={onClose}>
        <Modal.Header closeButton>
          <Modal.Title>
            {category.name}({subcategories.length})
          </Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <div className='section mb-2'>
            <div className='input-group'>
              <Input
                id='subcategory'
                name='subcategory'
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
            <div className='space-t-2'></div>
            <div className='text-right'>
              <Button onClick={addSubcategory}>
                <i className='fa-solid fa-plus'></i>
                <span className='ms-1'>Add</span>
              </Button>
            </div>
          </div>

          <div className='section'>
            <div className='table-response table-response--inner'>
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
                          <div className='input-group edit-input'>
                            <Input
                              id={`subcategoryName-${subcategory.id}`}
                              value={subcategory.name}
                              extraClasses={["no-border"]}
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
                                  onClick={() =>
                                    saveSubcategory(subcategory.id)
                                  }
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
                            className='edit-button height'
                            onClick={() => clickSubcategoryName(subcategory.id)}
                          >
                            {subcategory.name}
                          </div>
                        )}
                      </td>
                      <td>
                        <div className='height'>
                          <input
                            className='space-r-2'
                            type='radio'
                            checked={subcategory.isAlive}
                            onChange={() => {
                              changeSubcategoryAlive(subcategory.id, true);
                              saveSubcategory(subcategory.id);
                            }}
                          />
                          <label className='space-r-3' htmlFor='status'>
                            Enable
                          </label>
                          <input
                            className='space-r-2'
                            type='radio'
                            checked={!subcategory.isAlive}
                            onChange={() => {
                              changeSubcategoryAlive(subcategory.id, false);
                              saveSubcategory(subcategory.id);
                            }}
                          />
                          <label htmlFor='status'>Disable</label>
                        </div>
                      </td>
                      <td className='table-action-td'>
                        <div>
                          <Button
                            extraClasses={["delete-button", "delete-category"]}
                            onClick={() =>
                              showConfirmationModal(subcategory.id)
                            }
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
        </Modal.Body>
        <Modal.Footer>
          <Button extraClasses={["cancel-button"]} onClick={onClose}>
            Close
          </Button>
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
