import { useState, useEffect, useCallback } from "react";
import { Link, useParams } from "react-router-dom";
import { cloneDeep } from "lodash";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Component
import { pageSizeDefinition } from "@/components/Pagination";
import ConfirmationModal from "@/components/ConfirmationModal";
import Add from "./components/Add";
import Table from "./components/Table";
import Page from "./components/Page";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const errorMessage = {
  delete: "Subcategory is used or not found.",
};

const Category = () => {
  const { id } = useParams();

  // State
  const [orderBy, setOrderBy] = useState("id");
  const [order, setOrder] = useState("asc");
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(pageSizeDefinition[1]);
  const [totalSubcategoryCount, setTotalSubcategoryCount] = useState(0);
  const [category, setCategory] = useState({});
  const [subcategories, setSubcategories] = useState([]);
  const [selectedSubcategory, setSelectedSubcategory] = useState({});
  const [isOpenConfirmationModal, setIsOpenConfirmationModal] = useState(false);

  // Method
  const handleInitialization = useCallback(async () => {
    try {
      let response = null;
      response = await apiHandler.get(
        apiConfig.resource.EDIT_CATEGORY.replace(":id", id)
      );

      const category = response.data.data;
      const subcategoryCount = category.subcategoryCount;
      setCategory(category);
      setTotalSubcategoryCount(subcategoryCount);

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
          originalName: subcategory.name,
          isEditing: false,
        });
      });

      if (selectedSubcategory?.id) {
        const subcategory = updatedSubcategories.find(
          (subcategory) => subcategory.id === selectedSubcategory?.id
        );

        setSelectedSubcategory(subcategory);
      }

      setSubcategories(updatedSubcategories);
    } catch (error) {
      messageUtil.error(commonMessage.error);
    }
  }, [currentPage, id, order, orderBy, pageSize, selectedSubcategory?.id]);

  const changeOrder = (newOrderBy) => {
    if (newOrderBy === orderBy) {
      setOrder(order === "asc" ? "desc" : "asc");
    } else {
      setOrderBy(newOrderBy);
      setOrder("asc");
    }
  };

  const clickSubcategoryName = (subId) => {
    const updatedSubcategories = subcategories.map((subcategory) => {
      if (subcategory.id === subId) {
        subcategory.isEditing = !subcategory.isEditing;
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

  const changeSubcategoryName = (subId, name) => {
    const updatedSubcategories = subcategories.map((subcategory) => {
      if (subcategory.id === subId) {
        subcategory.name = name;
      }

      return subcategory;
    });

    setSubcategories(updatedSubcategories);
  };

  const saveSubcategory = async (subId) => {
    const subcategoryIndex = subcategories.findIndex(
      (subcategory) => subcategory.id === subId
    );

    const updatedSubcategories = cloneDeep(subcategories);
    if (updatedSubcategories[subcategoryIndex].name === "") {
      return;
    }

    try {
      const subcategory = subcategories.find(
        (subcategory) => subcategory.id === subId
      );
      const apiURL = apiConfig.resource.EDIT_SUBCATEGORY.replace(
        ":id",
        id
      ).replace(":subId", subId);
      const payload = {
        name: subcategory.name,
        isAlive: subcategory.isAlive,
      };
      await apiHandler.put(apiURL, payload);
      messageUtil.showSuccessMessage(commonMessage.success);

      updatedSubcategories[subcategoryIndex].originalName = subcategory.name;
      updatedSubcategories[subcategoryIndex].isEditing = false;
      setSubcategories(updatedSubcategories);
    } catch (error) {
      messageUtil.showErrorMessage(error.response.data.error.message);
    }
  };

  const deleteSubcategory = async () => {
    try {
      const apiURL = apiConfig.resource.EDIT_SUBCATEGORY.replace(
        ":id",
        id
      ).replace(":subId", selectedSubcategory.id);
      await apiHandler.delete(apiURL);
      messageUtil.showSuccessMessage(commonMessage.success);
      setIsOpenConfirmationModal(false);
      handleInitialization();
    } catch (error) {
      messageUtil.showErrorMessage(errorMessage.delete);
    }
  };

  const showConfirmationModal = (subId) => {
    const subcategory = subcategories.find(
      (subcategory) => subcategory.id === subId
    );
    setSelectedSubcategory(subcategory);
    setIsOpenConfirmationModal(true);
  };

  // Side effect
  useEffect(() => {
    if (!id) return;
    handleInitialization();
  }, [handleInitialization, id]);

  return (
    <>
      <div className='breadcrumb-container'>
        <Link to={routerConfig.routes.CATEGORIES} className='breadcrumb--item'>
          <span className='breadcrumb--item--inner'>
            <span className='breadcrumb--item-title'>Categories</span>
          </span>
        </Link>
      </div>

      <div className='section header-subtitle'>{category?.name}</div>

      <Add onInitialization={handleInitialization} category={category} />

      <Table
        currentPage={currentPage}
        pageSize={pageSize}
        orderBy={orderBy}
        order={order}
        onChangeOrder={changeOrder}
        subcategories={subcategories}
        onClickSubcategoryName={clickSubcategoryName}
        onChangeSubcategoryName={changeSubcategoryName}
        onChangeSubcategoryAlive={changeSubcategoryAlive}
        selectedSubcategory={selectedSubcategory}
        onShowConfirmationModal={showConfirmationModal}
        onsSaveSubcategory={saveSubcategory}
      />

      <Page
        currentPage={currentPage}
        setCurrentPage={setCurrentPage}
        totalSubcategoryCount={totalSubcategoryCount}
        pageSize={pageSize}
        setPageSize={setPageSize}
      />

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

export default Category;
