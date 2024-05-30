import "./module.scss";

import { useState, useEffect, useCallback } from "react";
import { Link } from "react-router-dom";
import { cloneDeep, orderBy } from "lodash";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Component
import { pageSizeDefinition } from "@/components/Pagination";
import ConfirmationModal from "@/components/ConfirmationModal";
import Add from "./components/Add";
import Table from "./components/Table";
import Page from "./components/Page";
import Subcategories from "./components/Subcategories";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const Category = () => {
  // State
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(pageSizeDefinition[1]);
  const [totalCategoryCount, setTotalCategoryCount] = useState(0);
  const [categories, setCategories] = useState([]);
  const [selectedCategory, setSelectedCategory] = useState({});
  const [isOpenSubcategoriesModal, setIsOpenSubcategoriesModal] =
    useState(false);
  const [isOpenConfirmationModal, setIsOpenConfirmationModal] = useState(false);

  // Method
  const handleInitialization = useCallback(
    async (resetPagin = false) => {
      if (resetPagin) {
        setCurrentPage(1);
        setPageSize(pageSizeDefinition[1]);
      }

      let response = null;
      response = await apiHandler.get(apiConfig.resource.NUMBER_OF_CATEGORIES);
      setTotalCategoryCount(response.data.data.totalCategoryNumber);

      response = await apiHandler.get(
        apiConfig.resource.CATEGORIES_BY_PAGE.replace(
          ":page",
          currentPage
        ).replace(":size", pageSize)
      );

      let updatedCategories = [];
      response.data.data.categories?.forEach((category) => {
        const subcategories = [];

        if (category.subcategories) {
          category.subcategories.forEach((subcategory) => {
            subcategories.push(subcategory);
          });
        }

        updatedCategories.push({
          ...category,
          subcategories,
          originalName: category.name,
          isEditing: false,
        });
      });
      updatedCategories = orderBy(updatedCategories, ["id"], ["asc"]);

      if (selectedCategory?.id) {
        const category = updatedCategories.find(
          (category) => category.id === selectedCategory?.id
        );

        setSelectedCategory(category);
      }

      setCategories(updatedCategories);
    },
    [currentPage, pageSize, selectedCategory?.id]
  );

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
      messageUtil.showErrorMessage(commonMessage.error);
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

      const updatedCategories = categories.filter(
        (category) => category.id !== selectedCategory.id
      );
      setCategories(updatedCategories);
      setIsOpenConfirmationModal(false);
      handleInitialization(true);
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  const showSubcategoriesModal = (id) => {
    const category = categories.find((category) => category.id === id);
    setSelectedCategory(category);
    setIsOpenSubcategoriesModal(true);
  };

  const showConfirmationModal = (id) => {
    const category = categories.find((category) => category.id === id);
    setSelectedCategory(category);
    setIsOpenConfirmationModal(true);
  };

  // Side effect
  useEffect(() => {
    handleInitialization();
  }, [currentPage, handleInitialization, pageSize]);

  return (
    <>
      <Link to={routerConfig.routes.CATEGORY}>
        <div className='breadcrumb-container'>
          <div className='breadcrumb-container-item'>Category</div>
        </div>
      </Link>
      <Add onInitialization={() => handleInitialization(true)} />
      <Table
        categories={categories}
        onClickCategoryName={clickCategoryName}
        changeCategoryName={changeCategoryName}
        changeCategoryAlive={changeCategoryAlive}
        selectedCategory={selectedCategory}
        onShowSubcategoriesModal={showSubcategoriesModal}
        onShowConfirmationModal={showConfirmationModal}
        saveCategory={saveCategory}
      />

      <Page
        currentPage={currentPage}
        setCurrentPage={setCurrentPage}
        totalCategoryCount={totalCategoryCount}
        pageSize={pageSize}
        setPageSize={setPageSize}
      />

      <Subcategories
        isOpen={isOpenSubcategoriesModal}
        onClose={() => setIsOpenSubcategoriesModal(false)}
        category={selectedCategory}
        onInitialization={handleInitialization}
      />

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
