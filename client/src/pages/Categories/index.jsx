import { useState, useEffect, useCallback } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useNavigate } from "react-router-dom";
import {
  faPlus,
  faPenToSquare,
  faTrash,
  faFloppyDisk,
  faXmark,
} from "@fortawesome/free-solid-svg-icons";
import { cloneDeep } from "lodash";

import Breadcrumbs from "@/components/Breadcrumbs";
import Form from "@/components/Form";
import ElementGroup from "@/components/ElementGroup";
import InputBox from "@/components/InputBox";
import EditableTextBox from "@/components/EditableTextBox";
import RadioBox from "@/components/RadioBox";
import IconButton from "@/components/IconButton";
import ToolbarBox from "@/components/ToolbarBox";
import TableBox from "@/components/TableBox";
import PaginationBox, { numberOfRow } from "@/components/PaginationBox";
import Spacer from "@/components/Spacer";
import Tooltip from "@/components/Tooltip";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Component
import CategoryModal from "./CategoryModal";
import DeleteModal from "./DeleteModal";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const errorMessage = {
  delete: "Category is used or not found.",
};

const defaultTableHeader = [
  {
    key: "index",
    label: "#",
    isSortable: false,
    isCenter: true,
    width: "50",
    sort: "",
  },
  {
    key: "name",
    label: "Category",
    isSortable: true,
    sort: "",
  },
  {
    key: "subcategories",
    label: "Subcategories",
    isSortable: true,
    isCenter: true,
    width: "50",
    sort: "",
  },
  {
    key: "status",
    label: "Status",
    isSortable: true,
    width: "100",
    sort: "",
  },
  {
    key: "action",
    label: "Action",
    isSortable: false,
    isCenter: true,
    width: "100",
    sort: "",
  },
];

const Categories = () => {
  const navigate = useNavigate();
  const linkList = [
    { to: "/", label: "Home" },
    { to: "/categories", label: "Categories" },
  ];

  const [orderBy, setOrderBy] = useState("id");
  const [order, setOrder] = useState("asc");
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(numberOfRow[1]);
  const [categories, setCategories] = useState([]);
  const [tableHeaders, setTableHeader] = useState(
    cloneDeep(defaultTableHeader)
  );
  const [tableData, setTableData] = useState([]);
  const [totalDataCount, setTotalDataCount] = useState(0);
  const [openCategoryModal, setOpenCategoryModal] = useState(false);
  const [openDeleteModal, setOpenDeleteModal] = useState(false);
  const [selectedCategory, setSelectedCategory] = useState(null);

  // Method
  const handleInitialization = useCallback(async () => {
    let response = null;
    response = await apiHandler.get(apiConfig.resource.NUMBER_OF_CATEGORIES);
    const totalCategoryNumber = response.data.data.totalCategoryNumber;
    setTotalDataCount(totalCategoryNumber);

    // Check if the current page is greater than the total page number
    const totalPageNum = Math.ceil(totalCategoryNumber / pageSize);
    let resetCurrentPage = false;
    if (totalPageNum < currentPage) {
      resetCurrentPage = true;
      setCurrentPage(1);
    }

    let queryString = "?orderBy=" + orderBy + "&order=" + order;
    response = await apiHandler.get(
      apiConfig.resource.CATEGORIES_BY_PAGE.replace(
        ":page",
        resetCurrentPage ? 1 : currentPage
      ).replace(":size", pageSize),
      queryString
    );

    let updatedCategories = [];
    response.data.data.categories?.forEach((category) => {
      let subcategories = [];
      if (category.subCategories) {
        subcategories = category.subCategories.sort((a, b) => {
          return a.name.localeCompare(b.name);
        });
      }

      updatedCategories.push({
        ...category,
        subcategoryCount: subcategories.length || 0,
        subcategories: subcategories,
        isEdit: false,
        originalName: category.name,
      });

      return;
    });

    setCategories(updatedCategories);
  }, [orderBy, order, currentPage, pageSize]);

  const handleChangeHeader = (newHeader, column, order) => {
    setTableHeader(newHeader);
    setOrderBy(column);
    setOrder(order);
  };

  const handleCloseModal = () => {
    setOpenCategoryModal(false);
    setOpenDeleteModal(false);
    setSelectedCategory(null);
  };

  const handleAddCategory = async () => {
    setOpenCategoryModal(false);
    setOpenDeleteModal(false);
    await handleInitialization();
  };

  const handleChangeCategory = useCallback(
    (id, value) => {
      const updatedCategories = categories.map((category) => {
        if (category.id === id) {
          return {
            ...category,
            name: value,
          };
        }
        return category;
      });

      setCategories(updatedCategories);
    },
    [categories]
  );

  const handleSaveCategory = useCallback(
    async (id) => {
      const categoryIndex = categories.findIndex(
        (category) => category.id === id
      );
      const updatedCategories = cloneDeep(categories);

      if (updatedCategories[categoryIndex].name === "") {
        return;
      }

      try {
        const category = categories.find((category) => category.id === id);
        const apiURL = apiConfig.resource.EDIT_CATEGORY.replace(":id", id);
        const payload = {
          name: category.name,
          isAlive: category.isAlive,
        };

        await apiHandler.put(apiURL, payload);
        messageUtil.showSuccessMessage(commonMessage.success);

        updatedCategories[categoryIndex].isEdit = false;
        updatedCategories[categoryIndex].originalName = category.name;
        setCategories(updatedCategories);

        handleInitialization();
      } catch (error) {
        messageUtil.showErrorMessage(commonMessage.error);
      }
    },
    [categories, handleInitialization]
  );

  const handleCancelEdit = useCallback(
    (id) => {
      const updatedCategories = categories.map((category) => {
        if (category.id === id) {
          return {
            ...category,
            isEdit: false,
            name: category.originalName,
          };
        }
        return category;
      });

      setCategories(updatedCategories);
    },
    [categories]
  );

  const handleEditCategory = useCallback(
    (id) => {
      const updatedCategories = categories.map((category) => {
        if (category.id === id) {
          return {
            ...category,
            isEdit: true,
          };
        }
        return category;
      });

      setCategories(updatedCategories);
    },
    [categories]
  );

  const handleChangeStatus = useCallback(
    async (id, status) => {
      const categoryIndex = categories.findIndex(
        (category) => category.id === id
      );
      const apiURL = apiConfig.resource.EDIT_CATEGORY.replace(":id", id);
      const payload = {
        name: categories[categoryIndex].name,
        isAlive: status,
      };

      try {
        await apiHandler.put(apiURL, payload);
        messageUtil.showSuccessMessage(commonMessage.success);

        const updatedCategories = cloneDeep(categories);
        updatedCategories[categoryIndex].isAlive = status;
        setCategories(updatedCategories);
      } catch (error) {
        messageUtil.showErrorMessage(commonMessage.error);
      }
    },
    [categories]
  );

  const handleSubcategory = useCallback(
    (category) => {
      navigate(routerConfig.routes.CATEGORY.replace(":id", category.id));
    },
    [navigate]
  );

  const handleOpenDeleteModal = (category) => {
    setSelectedCategory(category);
    setOpenDeleteModal(true);
  };

  const handleDelete = async () => {
    try {
      const apiURL = apiConfig.resource.EDIT_CATEGORY.replace(
        ":id",
        selectedCategory.id
      );

      await apiHandler.delete(apiURL);
      messageUtil.showSuccessMessage(commonMessage.success);
      setOpenDeleteModal(false);
      handleInitialization();
    } catch (error) {
      messageUtil.showErrorMessage(errorMessage.delete);
    }
  };

  // Side effect
  useEffect(() => {
    handleInitialization();
  }, [handleInitialization, orderBy, order, currentPage, pageSize]);

  useEffect(() => {
    const updatedTableData = categories.map((category, index) => {
      return {
        ...category,
        index: index + 1,
        isEdit: false,
        name: category.isEdit ? (
          <InputBox
            type='text'
            id={`category${category.id}`}
            name={`category-${category.id}`}
            value={category.name}
            onChange={(value) => handleChangeCategory(category.id, value)}
          >
            <FontAwesomeIcon
              icon={faFloppyDisk}
              onClick={() => handleSaveCategory(category.id)}
            />
            <FontAwesomeIcon
              icon={faXmark}
              onClick={() => handleCancelEdit(category.id)}
            />
          </InputBox>
        ) : (
          <EditableTextBox
            text={category.name}
            onClick={() => {
              handleEditCategory(category.id);
            }}
          />
        ),
        subcategories:
          category.subcategoryCount === 0 ? (
            <div>0</div>
          ) : (
            <Tooltip
              content={
                <ul>
                  {category.subcategories.map((subcategory) => (
                    <li key={subcategory.id}>{subcategory.name}</li>
                  ))}
                </ul>
              }
            >
              {category.subcategoryCount}
            </Tooltip>
          ),
        status: (
          <ElementGroup>
            <RadioBox
              checked={category.isAlive}
              onChange={() => handleChangeStatus(category.id, true)}
            >
              Enable
            </RadioBox>
            <RadioBox
              checked={!category.isAlive}
              onChange={() => handleChangeStatus(category.id, false)}
            >
              Disable
            </RadioBox>
          </ElementGroup>
        ),
        action: (
          <div className='flex items-center gap-4'>
            <IconButton onClick={() => handleSubcategory(category)}>
              <FontAwesomeIcon icon={faPenToSquare} />
            </IconButton>
            <IconButton onClick={() => handleOpenDeleteModal(category)}>
              <FontAwesomeIcon icon={faTrash} />
            </IconButton>
          </div>
        ),
      };
    });

    setTableData(updatedTableData);
  }, [
    currentPage,
    categories,
    pageSize,
    handleChangeCategory,
    handleSaveCategory,
    handleCancelEdit,
    handleEditCategory,
    handleChangeStatus,
    handleSubcategory,
  ]);

  return (
    <>
      <Breadcrumbs linkList={linkList} />
      <div className='custom-container primary-bg'>
        <Form>
          <ToolbarBox>
            <IconButton
              rounded={true}
              onClick={() => setOpenCategoryModal(true)}
            >
              <FontAwesomeIcon icon={faPlus} />
              <div>Add</div>
            </IconButton>
          </ToolbarBox>

          <TableBox
            headers={tableHeaders}
            onChangeHeader={handleChangeHeader}
            data={tableData}
          />

          <Spacer />
          <PaginationBox
            currentPage={currentPage}
            totalCount={totalDataCount}
            pageSize={pageSize}
            setPageSize={setPageSize}
            onPageChange={(page) => {
              setCurrentPage(page);
            }}
          />
        </Form>
      </div>

      <CategoryModal
        openModal={openCategoryModal}
        selectedCategory={selectedCategory}
        onClose={handleCloseModal}
        onSubmit={handleAddCategory}
      />

      <DeleteModal
        deleteMode={true}
        openModal={openDeleteModal}
        selectedCategory={selectedCategory}
        onClose={() => handleCloseModal()}
        onSubmit={() => handleDelete()}
      />
    </>
  );
};

export default Categories;
