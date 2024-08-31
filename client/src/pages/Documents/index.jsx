import { useState, useEffect, useCallback } from "react";
import { useNavigate, Link } from "react-router-dom";
import { cloneDeep } from "lodash";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Component
import { pageSizeDefinition } from "@/components/Pagination";
import ConfirmationModal from "@/components/ConfirmationModal";
import Search from "./components/Search";
import Table from "./components/Table";
import Page from "./components/Page";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const Documents = () => {
  const navigate = useNavigate();

  // State
  const [initialized, setInitialized] = useState(false);
  const [orderBy, setOrderBy] = useState("id");
  const [order, setOrder] = useState("asc");
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(pageSizeDefinition[1]);
  const [totalDocumentCount, setTotalDocumentCount] = useState(0);
  const [documents, setDocuments] = useState([]);
  const [selectedDocument, setSelectedDocument] = useState({});
  const [isOpenConfirmationModal, setIsOpenConfirmationModal] = useState(false);
  const [search, setSearch] = useState("");

  // Method
  const handleInitialization = useCallback(async () => {
    let response = null;

    response = await apiHandler.get(apiConfig.resource.NUMBER_OF_DOCUMENTS);
    const totalDocumentNumber = response.data.data.totalDocumentNumber;
    setTotalDocumentCount(totalDocumentNumber);

    // Check if the current page is greater than the total page number
    const totalPageNum = Math.ceil(totalDocumentNumber / pageSize);
    let resetCurrentPage = false;
    if (totalPageNum < currentPage) {
      resetCurrentPage = true;
      setCurrentPage(1);
    }

    let queryString = "?orderBy=" + orderBy + "&order=" + order;
    response = await apiHandler.get(
      apiConfig.resource.DOCUMENTS_BY_PAGE.replace(
        ":page",
        resetCurrentPage ? 1 : currentPage
      ).replace(":size", pageSize),
      queryString
    );

    let updatedDocuments = [];
    response.data.data.documents?.forEach((document) => {
      updatedDocuments.push({
        ...document,
        originalName: document.name,
        isEditing: false,
      });
    });

    if (selectedDocument?.id) {
      const document = updatedDocuments.find(
        (document) => document.id === selectedDocument?.id
      );

      setSelectedDocument(document);
    }

    setDocuments(updatedDocuments);
    setInitialized(true);
  }, [orderBy, order, currentPage, pageSize, selectedDocument?.id]);

  const searchDocuments = useCallback(async () => {
    if (!initialized) {
      return;
    }

    if (!search || search.length < 3) {
      handleInitialization();
      return;
    }

    const queryString = `?keyword=${search}`;
    const response = await apiHandler.get(
      `${apiConfig.resource.SEARCH_DOCUMENTS}${queryString}`
    );

    setDocuments(response.data.data || []);
  }, [handleInitialization, initialized, search]);

  const changeOrder = (newOrderBy) => {
    if (newOrderBy === orderBy) {
      setOrder(order === "asc" ? "desc" : "asc");
    } else {
      setOrderBy(newOrderBy);
      setOrder("asc");
    }
  };

  const clickDocumentName = (id) => {
    const updatedDocuments = documents.map((document) => {
      if (document.id === id) {
        document.isEditing = !document.isEditing;
      }

      return document;
    });

    setDocuments(updatedDocuments);
  };

  const changeDocumentName = (id, name) => {
    const updatedDocuments = documents.map((document) => {
      if (document.id === id) {
        document.name = name;
      }

      return document;
    });

    setDocuments(updatedDocuments);
  };

  const saveDocument = async (id) => {
    const documentIndex = documents.findIndex((document) => document.id === id);
    const updatedDocuments = cloneDeep(documents);

    if (updatedDocuments[documentIndex].name === "") {
      return;
    }

    try {
      const document = documents.find((document) => document.id === id);
      const apiURL = apiConfig.resource.EDIT_DOCUMENT.replace(":id", id);
      const payload = {
        name: document.name,
        categoryId: document.category.id,
        subcategoryId: document.subCategory.id,
        postMemberId: document.postMember.id,
        relationMemberIds:
          document.relationMembers?.map((member) => member.memberId) || [],
        tagIds: document.tags?.map((tag) => tag.tagId) || [],
        content: document.content,
      };

      await apiHandler.put(apiURL, payload);
      messageUtil.showSuccessMessage(commonMessage.success);

      updatedDocuments[documentIndex].originalName = document.name;
      updatedDocuments[documentIndex].isEditing = false;
      setDocuments(updatedDocuments);
    } catch (error) {
      messageUtil.showErrorMessage(error.response.data.message);
    }
  };

  const deleteDocument = async () => {
    try {
      const apiURL = apiConfig.resource.EDIT_DOCUMENT.replace(
        ":id",
        selectedDocument.id
      );
      await apiHandler.delete(apiURL);
      messageUtil.showSuccessMessage(commonMessage.success);
      setIsOpenConfirmationModal(false);
      handleInitialization();
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  const showConfirmationModal = (id) => {
    const document = documents.find((document) => document.id === id);
    setSelectedDocument(document);
    setIsOpenConfirmationModal(true);
  };

  const linkToDocument = (document) => {
    if (search) {
      window.open(
        `/#${routerConfig.routes.DOCUMENT.replace(
          ":id",
          document.id
        )}?keyword=${search}`
      );
      return;
    }

    navigate(routerConfig.routes.DOCUMENT.replace(":id", document.id));
  };

  // Side effect
  useEffect(() => {
    handleInitialization();
  }, [handleInitialization]);

  return (
    <>
      <div className='breadcrumb-container'>
        <Link to={routerConfig.routes.DOCUMENTS} className='breadcrumb--item'>
          <span className='breadcrumb--item--inner'>
            <span className='breadcrumb--item-title'>Documents</span>
          </span>
        </Link>
      </div>

      <div className='floating-button'>
        <i
          className='fa-solid fa-plus'
          onClick={() => navigate(routerConfig.routes.ADD_DOCUMENT)}
        />
      </div>

      <Search
        search={search}
        setSearch={setSearch}
        onSearchDocuments={searchDocuments}
      />

      {search ? (
        <>
          <Table
            search={search}
            currentPage={currentPage}
            pageSize={pageSize}
            orderBy={orderBy}
            order={order}
            onChangeOrder={changeOrder}
            documents={documents}
            onClickDocumentName={clickDocumentName}
            onChangeDocumentName={changeDocumentName}
            selectedDocument={selectedDocument}
            onShowConfirmationModal={showConfirmationModal}
            saveDocument={saveDocument}
            onLinkToDocument={linkToDocument}
          />
        </>
      ) : (
        <>
          <Table
            currentPage={currentPage}
            pageSize={pageSize}
            orderBy={orderBy}
            order={order}
            onChangeOrder={changeOrder}
            documents={documents}
            onClickDocumentName={clickDocumentName}
            onChangeDocumentName={changeDocumentName}
            selectedDocument={selectedDocument}
            onShowConfirmationModal={showConfirmationModal}
            saveDocument={saveDocument}
            onLinkToDocument={linkToDocument}
          />

          <Page
            currentPage={currentPage}
            setCurrentPage={setCurrentPage}
            totalDocumentCount={totalDocumentCount}
            pageSize={pageSize}
            setPageSize={setPageSize}
          />

          <ConfirmationModal
            isOpen={isOpenConfirmationModal}
            onClose={() => {
              setSelectedDocument(null);
              setIsOpenConfirmationModal(false);
            }}
            onConfirm={deleteDocument}
          />
        </>
      )}
    </>
  );
};

export default Documents;
