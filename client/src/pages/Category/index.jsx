import { useState, useEffect, useCallback } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useParams } from "react-router-dom";
import {
  faPlus,
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

// Const
import apiConfig from "@/const/config/api";

// Component
import SubcategoryModal from "./SubcategoryModal";
import DeleteModal from "./DeleteModal";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

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
    label: "Subcategory",
    isSortable: true,
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

const errorMessage = {
  duplicated: "Name is duplicated.",
};

const Category = () => {
  const { id: categoryId } = useParams();

  // State
  const [category, setCategory] = useState(null);
  const [linkList, setLinkList] = useState([]);
  const [orderBy, setOrderBy] = useState("id");
  const [order, setOrder] = useState("asc");
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(numberOfRow[1]);
  const [subcategories, setSubcategories] = useState([]);
  const [tableHeaders, setTableHeader] = useState(
    cloneDeep(defaultTableHeader)
  );
  const [tableData, setTableData] = useState([]);
  const [totalDataCount, setTotalDataCount] = useState(0);
  const [openSubcategoryModal, setOpenSubcategoryModal] = useState(false);
  const [openDeleteModal, setOpenDeleteModal] = useState(false);
  const [selectedSubcategory, setSelectedSubcategory] = useState(null);

  // Method
  const handleInitialization = useCallback(async () => {
    try {
      let response = null;
      response = await apiHandler.get(
        apiConfig.resource.EDIT_CATEGORY.replace(":id", categoryId)
      );

      const category = response.data.data;
      setCategory(category);
      const subcategoryCount = category.subcategoryCount;
      setTotalDataCount(subcategoryCount);

      // Check if the current page is greater than the total page number
      const totalPageNum = Math.ceil(subcategoryCount / pageSize);
      let resetCurrentPage = false;
      if (totalPageNum < currentPage) {
        resetCurrentPage = true;
        setCurrentPage(1);
      }

      let queryString = "?orderBy=" + orderBy + "&order=" + order;
      response = await apiHandler.get(
        apiConfig.resource.SUBCATEGORIES_BY_PAGE.replace(":id", category.id)
          .replace(":page", resetCurrentPage ? 1 : currentPage)
          .replace(":size", pageSize),
        queryString
      );

      let updatedSubcategories = [];
      response.data.data.subcategories?.forEach((subcategory) => {
        updatedSubcategories.push({
          ...subcategory,
          categoryId: category.id,
          originalName: subcategory.name,
          isEditing: false,
        });
      });

      setSubcategories(updatedSubcategories);
    } catch (error) {
      messageUtil.showErrorMessage(error.response.data.message);
    }
  }, [categoryId, currentPage, pageSize, orderBy, order]);

  const handleCloseModal = () => {
    setOpenSubcategoryModal(false);
    setOpenDeleteModal(false);
    setSelectedSubcategory(null);
  };

  const handleAddSubcategory = async () => {
    setOpenSubcategoryModal(false);
    setOpenDeleteModal(false);
    await handleInitialization();
  };

  const handleChangeHeader = (newHeader, column, order) => {
    setTableHeader(newHeader);
    setOrderBy(column);
    setOrder(order);
  };

  const handleChangeSubcategory = useCallback(
    (id, value) => {
      const updatedSubcategories = subcategories.map((subcategory) => {
        if (subcategory.id === id) {
          return {
            ...subcategory,
            name: value,
          };
        }

        return subcategory;
      });

      setSubcategories(updatedSubcategories);
    },
    [subcategories]
  );

  const handleSaveSubcategory = useCallback(
    async (id) => {
      const subcategoryIndex = subcategories.findIndex(
        (subcategory) => subcategory.id === id
      );
      const updatedSubcategories = cloneDeep(subcategories);

      if (updatedSubcategories[subcategoryIndex].name === "") {
        return;
      }

      try {
        const subcategory = subcategories.find(
          (subcategory) => subcategory.id === id
        );
        const apiURL = apiConfig.resource.EDIT_SUBCATEGORY.replace(
          ":id",
          categoryId
        ).replace(":subId", id);
        const payload = {
          name: subcategory.name,
          isAlive: subcategory.isAlive,
        };

        await apiHandler.put(apiURL, payload);
        messageUtil.showSuccessMessage(commonMessage.success);

        updatedSubcategories[subcategoryIndex].isEdit = false;
        updatedSubcategories[subcategoryIndex].originalName = subcategory.name;
        setSubcategories(updatedSubcategories);

        handleInitialization();
      } catch (error) {
        if (error.response.data.code === 4423) {
          messageUtil.showErrorMessage(errorMessage.duplicated);
          return;
        }

        messageUtil.showErrorMessage(commonMessage.error);
      }
    },
    [handleInitialization, subcategories]
  );

  const handleCancelEdit = useCallback(
    (id) => {
      const updatedSubcategories = subcategories.map((subcategory) => {
        if (subcategory.id === id) {
          return {
            ...subcategory,
            isEdit: false,
            name: subcategory.originalName,
          };
        }

        return subcategory;
      });

      setSubcategories(updatedSubcategories);
    },
    [subcategories]
  );

  const handleEditSubcategory = useCallback(
    (id) => {
      const updatedSubcategories = subcategories.map((subcategory) => {
        if (subcategory.id === id) {
          return {
            ...subcategory,
            isEdit: true,
          };
        }

        return subcategory;
      });

      setSubcategories(updatedSubcategories);
    },
    [subcategories]
  );

  const handleChangeStatus = useCallback(
    async (id, status) => {
      const subcategoryIndex = subcategories.findIndex(
        (subcategory) => subcategory.id === id
      );
      const apiURL = apiConfig.resource.EDIT_SUBCATEGORY.replace(
        ":id",
        categoryId
      ).replace(":subId", id);
      const payload = {
        name: subcategories[subcategoryIndex].name,
        isAlive: status,
      };

      try {
        await apiHandler.put(apiURL, payload);
        messageUtil.showSuccessMessage(commonMessage.success);

        const updatedSubcategories = cloneDeep(subcategories);
        updatedSubcategories[subcategoryIndex].isAlive = status;
        setSubcategories(updatedSubcategories);
      } catch (error) {
        messageUtil.showErrorMessage(commonMessage.error);
      }
    },
    [categoryId, subcategories]
  );

  const handleOpenDeleteModal = (subcategory) => {
    setSelectedSubcategory(subcategory);
    setOpenDeleteModal(true);
  };

  const handleDelete = async () => {
    try {
      const apiURL = apiConfig.resource.EDIT_SUBCATEGORY.replace(
        ":id",
        categoryId
      ).replace(":subId", selectedSubcategory.id);

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
    const updatedTableData = subcategories.map((subcategory, index) => {
      return {
        ...subcategory,
        index: index + 1,
        isEdit: false,
        name: subcategory.isEdit ? (
          <InputBox
            type='text'
            id={`subcategory-${subcategory.id}`}
            name={`subcategory-${subcategory.id}`}
            value={subcategory.name}
            onChange={(value) => handleChangeSubcategory(subcategory.id, value)}
          >
            <FontAwesomeIcon
              icon={faFloppyDisk}
              onClick={() => handleSaveSubcategory(subcategory.id)}
            />
            <FontAwesomeIcon
              icon={faXmark}
              onClick={() => handleCancelEdit(subcategory.id)}
            />
          </InputBox>
        ) : (
          <EditableTextBox
            text={subcategory.name}
            onClick={() => {
              handleEditSubcategory(subcategory.id);
            }}
          />
        ),
        status: (
          <ElementGroup>
            <RadioBox
              checked={subcategory.isAlive}
              onChange={() => handleChangeStatus(subcategory.id, true)}
            >
              Enable
            </RadioBox>
            <RadioBox
              checked={!subcategory.isAlive}
              onChange={() => handleChangeStatus(subcategory.id, false)}
            >
              Disable
            </RadioBox>
          </ElementGroup>
        ),
        action: (
          <div className='flex items-center gap-4'>
            <IconButton onClick={() => handleOpenDeleteModal(subcategory)}>
              <FontAwesomeIcon icon={faTrash} />
            </IconButton>
          </div>
        ),
      };
    });

    setTableData(updatedTableData);
  }, [
    handleCancelEdit,
    handleChangeStatus,
    handleChangeSubcategory,
    handleEditSubcategory,
    handleSaveSubcategory,
    subcategories,
  ]);

  useEffect(() => {
    if (!category) return;

    const updatedLinkList = [
      { to: "/", label: "Home" },
      { to: "/categories", label: "Categories" },
      {
        to: `/categories/${category.id}`,
        label: category.name,
      },
    ];
    setLinkList(updatedLinkList);
  }, [category]);

  return (
    <>
      <Breadcrumbs linkList={linkList} />
      <div className='custom-container primary-bg'>
        <Form>
          <ToolbarBox>
            <IconButton
              rounded={true}
              onClick={() => setOpenSubcategoryModal(true)}
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

      <SubcategoryModal
        categoryId={categoryId}
        openModal={openSubcategoryModal}
        onClose={handleCloseModal}
        onSubmit={handleAddSubcategory}
      />

      <DeleteModal
        deleteMode={true}
        openModal={openDeleteModal}
        selectedSubcategory={selectedSubcategory}
        onClose={() => handleCloseModal()}
        onSubmit={() => handleDelete()}
      />
    </>
  );
};

export default Category;
