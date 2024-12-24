import { useState, useEffect, useCallback } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faPlus,
  faTrash,
  faFloppyDisk,
  faXmark,
} from "@fortawesome/free-solid-svg-icons";
import { cloneDeep } from "lodash";

import Breadcrumbs from "@/components/Breadcrumbs";
import Form from "@/components/Form";
import InputBox from "@/components/InputBox";
import EditableTextBox from "@/components/EditableTextBox";
import TagBox from "@/components/TagBox";
import IconButton from "@/components/IconButton";
import ToolbarBox from "@/components/ToolbarBox";
import TableBox from "@/components/TableBox";
import PaginationBox, { numberOfRow } from "@/components/PaginationBox";
import Spacer from "@/components/Spacer";

// Const
import apiConfig from "@/const/config/api";

// Component
import TagModal from "./TagModal";
import ColorModal from "./ColorModal";
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
    label: "Tag",
    isSortable: true,
    sort: "",
  },
  {
    key: "color",
    label: "Color",
    isSortable: true,
    isCenter: true,
    width: "220",
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
  duplicated: "Tag is duplicated.",
};

const Tags = () => {
  const linkList = [
    { to: "/", label: "Home" },
    { to: "/tags", label: "Tags" },
  ];

  // State
  const [orderBy, setOrderBy] = useState("id");
  const [order, setOrder] = useState("asc");
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(numberOfRow[1]);
  const [colorCategories, setColorCategories] = useState([]);
  const [colors, setColors] = useState([]);
  const [tags, setTags] = useState([]);
  const [tableHeaders, setTableHeader] = useState(
    cloneDeep(defaultTableHeader)
  );
  const [tableData, setTableData] = useState([]);
  const [totalDataCount, setTotalDataCount] = useState(0);
  const [openTagModal, setOpenTagModal] = useState(false);
  const [openColorModal, setOpenColorModal] = useState(false);
  const [openDeleteModal, setOpenDeleteModal] = useState(false);
  const [selectedTag, setSelectedTag] = useState(null);

  // Method
  const handleInitialization = useCallback(async () => {
    let response = null;

    response = await apiHandler.get(apiConfig.resource.COLOR_CATEGORIES);
    setColorCategories(response.data.data);

    response = await apiHandler.get(apiConfig.resource.COLORS);
    setColors(response.data.data);

    response = await apiHandler.get(apiConfig.resource.NUMBER_OF_TAGS);
    const totalTagNumber = response.data.data.totalTagNumber;
    setTotalDataCount(totalTagNumber);

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
        color: tag.colorName,
        originalName: tag.name,
        isEditing: false,
      });
    });

    setTags(updatedTags);
  }, [currentPage, order, orderBy, pageSize]);

  const handleChangeHeader = (newHeader, column, order) => {
    setTableHeader(newHeader);
    setOrderBy(column);
    setOrder(order);
  };

  const handleCloseModal = () => {
    setOpenTagModal(false);
    setOpenColorModal(false);
    setOpenDeleteModal(false);
  };

  const handleAddTag = async () => {
    handleCloseModal();
    await handleInitialization();
  };

  const handleChangeColor = async () => {
    setSelectedTag(null);
    handleCloseModal();
    await handleInitialization();
  };

  const handleDelete = async () => {
    try {
      const apiURL = apiConfig.resource.EDIT_TAG.replace(":id", selectedTag.id);

      await apiHandler.delete(apiURL);
      messageUtil.showSuccessMessage(commonMessage.success);
      setOpenDeleteModal(false);
      handleInitialization();
    } catch (error) {
      messageUtil.showErrorMessage(
        messageUtil.showErrorMessage(error.response.data.message)
      );
    }
  };

  const handleChangeTag = useCallback(
    (id, value) => {
      const updateTags = tags.map((tag) => {
        if (tag.id === id) {
          return {
            ...tag,
            name: value,
          };
        }
        return tag;
      });

      setTags(updateTags);
    },
    [tags]
  );

  const handleSaveTag = useCallback(
    async (id) => {
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

        updatedTags[tagIndex].isEdit = false;
        updatedTags[tagIndex].originalName = tag.name;
        setTags(updatedTags);

        handleInitialization();
      } catch (error) {
        if (error.response.data.code == 5423) {
          messageUtil.showErrorMessage(errorMessage.duplicated);
          return;
        }

        messageUtil.showErrorMessage(commonMessage.error);
      }
    },
    [handleInitialization, tags]
  );

  const handleCancelEdit = useCallback(
    (id) => {
      const updatedTags = tags.map((tag) => {
        if (tag.id === id) {
          return {
            ...tag,
            isEdit: false,
            name: tag.originalName,
          };
        }
        return tag;
      });

      setTags(updatedTags);
    },
    [tags]
  );

  const handleEditTag = useCallback(
    (id) => {
      const updatedTags = tags.map((tag) => {
        if (tag.id === id) {
          return {
            ...tag,
            isEdit: true,
          };
        }
        return tag;
      });

      setTags(updatedTags);
    },
    [tags]
  );

  const handleOpenDeleteModal = (tag) => {
    setSelectedTag(tag);
    setOpenDeleteModal(true);
  };

  // Side effect
  useEffect(() => {
    handleInitialization();
  }, [handleInitialization, orderBy, order, currentPage, pageSize]);

  useEffect(() => {
    const updatedTableData = tags.map((tag, index) => {
      return {
        ...tag,
        index: (currentPage - 1) * pageSize + index + 1,
        isEdit: false,
        name: tag.isEdit ? (
          <InputBox
            type='text'
            id={`tag${tag.id}`}
            name={`tag-${tag.id}`}
            value={tag.name}
            onChange={(value) => handleChangeTag(tag.id, value)}
          >
            <FontAwesomeIcon
              icon={faFloppyDisk}
              onClick={() => handleSaveTag(tag.id)}
            />
            <FontAwesomeIcon
              icon={faXmark}
              onClick={() => handleCancelEdit(tag.id)}
            />
          </InputBox>
        ) : (
          <EditableTextBox
            text={tag.name}
            onClick={() => {
              handleEditTag(tag.id);
            }}
          />
        ),
        color: (
          <TagBox
            tag={{ label: tag.colorName, hashCode: tag.colorHexCode }}
            onClick={() => {
              setOpenColorModal(true);
              setSelectedTag(tag);
            }}
          />
        ),
        action: (
          <div className='flex items-center gap-4'>
            <IconButton onClick={() => handleOpenDeleteModal(tag)}>
              <FontAwesomeIcon icon={faTrash} />
            </IconButton>
          </div>
        ),
      };
    });

    setTableData(updatedTableData);
  }, [
    currentPage,
    handleCancelEdit,
    handleChangeTag,
    handleEditTag,
    handleSaveTag,
    pageSize,
    tags,
  ]);

  return (
    <>
      <Breadcrumbs linkList={linkList} />
      <div className='custom-container primary-bg'>
        <Form>
          <ToolbarBox>
            <IconButton rounded={true} onClick={() => setOpenTagModal(true)}>
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

      <TagModal
        openModal={openTagModal}
        colorCategories={colorCategories}
        colors={colors}
        onClose={handleCloseModal}
        onSubmit={handleAddTag}
      />

      <ColorModal
        openModal={openColorModal}
        colorCategories={colorCategories}
        colors={colors}
        selectedTag={selectedTag}
        onClose={handleCloseModal}
        onSubmit={handleChangeColor}
      />

      <DeleteModal
        deleteMode={true}
        openModal={openDeleteModal}
        selectedTag={selectedTag}
        onClose={handleCloseModal}
        onSubmit={handleDelete}
      />
    </>
  );
};

export default Tags;
