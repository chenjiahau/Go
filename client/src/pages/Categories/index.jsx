import "./module.scss";

import { useState, useEffect, useCallback } from "react";
import { Link, useNavigate } from "react-router-dom";
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
  delete: "Category is used or not found.",
};

const Category = () => {
  const navigate = useNavigate();

  // State
  const [orderBy, setOrderBy] = useState("id");
  const [order, setOrder] = useState("asc");
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(pageSizeDefinition[1]);
  const [totalCategoryCount, setTotalCategoryCount] = useState(0);
  const [categories, setCategories] = useState([]);
  const [selectedCategory, setSelectedCategory] = useState({});
  const [isOpenConfirmationModal, setIsOpenConfirmationModal] = useState(false);

  // Method
  const handleInitialization = useCallback(async () => {
    let response = null;
    response = await apiHandler.get(apiConfig.resource.NUMBER_OF_CATEGORIES);
    const totalCategoryNumber = response.data.data.totalCategoryNumber;
    setTotalCategoryCount(totalCategoryNumber);

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
      updatedCategories.push({
        ...category,
        originalName: category.name,
        isEditing: false,
      });
    });

    if (selectedCategory?.id) {
      const category = updatedCategories.find(
        (category) => category.id === selectedCategory?.id
      );

      setSelectedCategory(category);
    }

    setCategories(updatedCategories);
  }, [currentPage, order, orderBy, pageSize, selectedCategory?.id]);

  const changeOrder = (newOrderBy) => {
    if (newOrderBy === orderBy) {
      setOrder(order === "asc" ? "desc" : "asc");
    } else {
      setOrderBy(newOrderBy);
      setOrder("asc");
    }
  };

  const clickCategoryName = (id) => {
    const updatedCategories = categories.map((category) => {
      if (category.id === id) {
        category.isEditing = !category.isEditing;
      }

      return category;
    });

    setCategories(updatedCategories);
  };

  const changeCategoryAlive = async (id, alive) => {
    const updatedCategories = categories.map((category) => {
      if (category.id === id) {
        category.isAlive = alive;
      }

      return category;
    });

    setCategories(updatedCategories);
  };

  const changeCategoryName = (id, name) => {
    const updatedCategories = categories.map((category) => {
      if (category.id === id) {
        category.name = name;
      }

      return category;
    });

    setCategories(updatedCategories);
  };

  const saveCategory = async (id) => {
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

      updatedCategories[categoryIndex].originalName = category.name;
      updatedCategories[categoryIndex].isEditing = false;
      setCategories(updatedCategories);
    } catch (error) {
      messageUtil.showErrorMessage(error.response.data.message);
    }
  };

  const deleteCategory = async () => {
    try {
      const apiURL = apiConfig.resource.EDIT_CATEGORY.replace(
        ":id",
        selectedCategory.id
      );
      await apiHandler.delete(apiURL);
      messageUtil.showSuccessMessage(commonMessage.success);
      setIsOpenConfirmationModal(false);
      handleInitialization();
    } catch (error) {
      messageUtil.showErrorMessage(errorMessage.delete);
    }
  };

  const showConfirmationModal = (id) => {
    const category = categories.find((category) => category.id === id);
    if (category.subcategoryCount > 0) return;
    setSelectedCategory(category);
    setIsOpenConfirmationModal(true);
  };

  const linkToEditSubcategory = (category) => {
    navigate(routerConfig.routes.CATEGORY.replace(":id", category.id));
  };

  // Side effect
  useEffect(() => {
    handleInitialization();
  }, [currentPage, handleInitialization, pageSize, orderBy, order]);

  return (
    <>
      <div className='breadcrumb-container'>
        <Link to={routerConfig.routes.CATEGORIES} className='breadcrumb--item'>
          <span className='breadcrumb--item--inner'>
            <span className='breadcrumb--item-title'>Categories</span>
          </span>
        </Link>
      </div>

      <Add onInitialization={handleInitialization} />

      <Table
        currentPage={currentPage}
        pageSize={pageSize}
        orderBy={orderBy}
        order={order}
        onChangeOrder={changeOrder}
        categories={categories}
        onClickCategoryName={clickCategoryName}
        changeCategoryName={changeCategoryName}
        changeCategoryAlive={changeCategoryAlive}
        selectedCategory={selectedCategory}
        onShowConfirmationModal={showConfirmationModal}
        onLinkToEditCategory={linkToEditSubcategory}
        saveCategory={saveCategory}
      />

      <Page
        currentPage={currentPage}
        setCurrentPage={setCurrentPage}
        totalCategoryCount={totalCategoryCount}
        pageSize={pageSize}
        setPageSize={setPageSize}
      />

      {/* <Subcategories
        isOpen={isOpenSubcategoriesModal}
        onClose={() => setIsOpenSubcategoriesModal(false)}
        category={selectedCategory}
        onInitialization={handleInitialization}
      /> */}

      <ConfirmationModal
        isOpen={isOpenConfirmationModal}
        onClose={() => {
          setSelectedCategory(null);
          setIsOpenConfirmationModal(false);
        }}
        onConfirm={deleteCategory}
      />
    </>
  );
};

export default Category;
