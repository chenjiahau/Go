import { useState, useEffect, useCallback } from "react";
import { Link } from "react-router-dom";
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

const Category = () => {
  // State
  const [colorCategories, setColorCategories] = useState([]);
  const [colors, setColors] = useState([]);
  const [orderBy, setOrderBy] = useState("id");
  const [order, setOrder] = useState("asc");
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(pageSizeDefinition[1]);
  const [totalTagCount, setTotalTagCount] = useState(0);
  const [tags, setTags] = useState([]);
  const [selectedTag, setSelectedTag] = useState({});
  const [isOpenConfirmationModal, setIsOpenConfirmationModal] = useState(false);

  // Method
  const handleInitialization = useCallback(async () => {
    let response = null;

    response = await apiHandler.get(apiConfig.resource.COLOR_CATEGORIES);
    setColorCategories(response.data.data.colorCategories);

    response = await apiHandler.get(apiConfig.resource.COLORS);
    setColors(response.data.data.colors);

    response = await apiHandler.get(apiConfig.resource.NUMBER_OF_TAGS);
    const totalTagNumber = response.data.data.totalTagNumber;
    setTotalTagCount(totalTagNumber);

    // Check if the current page is greater than the total page number
    const totalPageNum = Math.ceil(totalTagNumber / pageSize);
    let resetCurrentPage = false;
    if (totalPageNum < currentPage) {
      resetCurrentPage = true;
      setCurrentPage(1);
    }

    let queryString = "?orderBy=" + orderBy + "&order=" + order;
    response = await apiHandler.get(
      apiConfig.resource.TAGS_BY_PAGE.replace(
        ":page",
        resetCurrentPage ? 1 : currentPage
      ).replace(":size", pageSize),
      queryString
    );

    let updatedTags = [];
    response.data.data.tags?.forEach((tag) => {
      updatedTags.push({
        ...tag,
        originalName: tag.name,
        isEditing: false,
      });
    });

    if (selectedTag?.id) {
      const tag = updatedTags.find((tag) => tag.id === selectedTag?.id);

      setSelectedTag(tag);
    }

    setTags(updatedTags);
  }, [currentPage, order, orderBy, pageSize, selectedTag?.id]);

  const changeOrder = (newOrderBy) => {
    if (newOrderBy === orderBy) {
      setOrder(order === "asc" ? "desc" : "asc");
    } else {
      setOrderBy(newOrderBy);
      setOrder("asc");
    }
  };

  const clickTagName = (id) => {
    const updatedTags = tags.map((tag) => {
      if (tag.id === id) {
        tag.isEditing = !tag.isEditing;
      }

      return tag;
    });

    setTags(updatedTags);
  };

  const changeTagName = (id, name) => {
    const updatedTags = tags.map((tag) => {
      if (tag.id === id) {
        tag.name = name;
      }

      return tag;
    });

    setTags(updatedTags);
  };

  const saveTag = async (id) => {
    const tagIndex = tags.findIndex((tag) => tag.id === id);
    const updatedTags = cloneDeep(tags);

    if (updatedTags[tagIndex].name === "") {
      return;
    }

    try {
      const tag = tags.find((tag) => tag.id === id);
      const apiURL = apiConfig.resource.EDIT_TAG.replace(":id", id);
      const payload = {
        colorId: tag.colorId,
        name: tag.name,
      };
      await apiHandler.put(apiURL, payload);
      messageUtil.showSuccessMessage(commonMessage.success);

      updatedTags[tagIndex].originalName = tag.name;
      updatedTags[tagIndex].isEditing = false;
      setTags(updatedTags);
    } catch (error) {
      messageUtil.showErrorMessage(error.response.data.error.message);
    }
  };

  const deleteTag = async () => {
    try {
      const apiURL = apiConfig.resource.EDIT_TAG.replace(":id", selectedTag.id);
      await apiHandler.delete(apiURL);
      messageUtil.showSuccessMessage(commonMessage.success);
      setIsOpenConfirmationModal(false);
      handleInitialization();
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  const showConfirmationModal = (id) => {
    const tag = tags.find((tag) => tag.id === id);
    setSelectedTag(tag);
    setIsOpenConfirmationModal(true);
  };

  // Side effect
  useEffect(() => {
    handleInitialization();
  }, [currentPage, handleInitialization, pageSize, orderBy, order]);

  return (
    <>
      <div className='breadcrumb-container'>
        <Link to={routerConfig.routes.TAGS} className='breadcrumb--item'>
          <span className='breadcrumb--item--inner'>
            <span className='breadcrumb--item-title'>Tags</span>
          </span>
        </Link>
      </div>

      <Add
        onInitialization={handleInitialization}
        colorCategories={colorCategories}
        colors={colors}
      />

      <Table
        currentPage={currentPage}
        pageSize={pageSize}
        orderBy={orderBy}
        order={order}
        onChangeOrder={changeOrder}
        tags={tags}
        onClickTagName={clickTagName}
        changeTagName={changeTagName}
        selectedTag={selectedTag}
        onShowConfirmationModal={showConfirmationModal}
        saveTag={saveTag}
        onInitialization={handleInitialization}
        colorCategories={colorCategories}
        colors={colors}
      />

      <Page
        currentPage={currentPage}
        setCurrentPage={setCurrentPage}
        totalTagCount={totalTagCount}
        pageSize={pageSize}
        setPageSize={setPageSize}
      />

      <ConfirmationModal
        isOpen={isOpenConfirmationModal}
        onClose={() => {
          setSelectedTag(null);
          setIsOpenConfirmationModal(false);
        }}
        onConfirm={deleteTag}
      />
    </>
  );
};

export default Category;
