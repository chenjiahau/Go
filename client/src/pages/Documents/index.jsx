import { useState, useEffect, useCallback } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faTrash } from "@fortawesome/free-solid-svg-icons";
import { cloneDeep } from "lodash";

import Breadcrumbs from "@/components/Breadcrumbs";
import Form from "@/components/Form";
import LinkButton from "@/components/LinkButton";
import IconButton from "@/components/IconButton";
import ToolbarBox from "@/components/ToolbarBox";
import TableBox from "@/components/TableBox";
import PaginationBox, { numberOfRow } from "@/components/PaginationBox";
import Spacer from "@/components/Spacer";
import FloatingButton from "@/components/FloatingButton";
import SearchInputBox from "@/components/SearchInputBox";
import LoadingBox from "@/components/LoadingBox";

// Const
import apiConfig from "@/const/config/api";

// Component
import DocumentModal from "./DocumentModal";
import DeleteDocumentModal from "./DeleteDocumentModal";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";
import { formatDateTime } from "@/util/datetime.util";

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
    label: "Document",
    isSortable: true,
    sort: "",
  },
  {
    key: "category",
    label: "Category",
    isSortable: true,
    isCenter: true,
    width: "200",
    sort: "",
  },
  {
    key: "subcategory",
    label: "Subcategory",
    isSortable: true,
    isCenter: true,
    width: "200",
    sort: "",
  },
  {
    key: "author",
    label: "Author",
    isSortable: true,
    isCenter: true,
    width: "200",
    sort: "",
  },
  {
    key: "createdDate",
    label: "Created Date",
    isSortable: true,
    isCenter: true,
    width: "200",
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
  notFound: "Documents not found",
};

const Documents = () => {
  const linkList = [
    { to: "/", label: "Home" },
    { to: "/documents", label: "Documents" },
  ];

  // State
  const [initialized, setInitialized] = useState(false);
  const [orderBy, setOrderBy] = useState("id");
  const [order, setOrder] = useState("asc");
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(numberOfRow[1]);
  const [documents, setDocuments] = useState([]);
  const [tableHeaders, setTableHeader] = useState(
    cloneDeep(defaultTableHeader)
  );
  const [tableData, setTableData] = useState([]);
  const [totalDataCount, setTotalDataCount] = useState(0);
  const [openDocumentModal, setOpenDocumentModal] = useState(false);
  const [openDeleteModal, setOpenDeleteModal] = useState(false);
  const [selectedDocument, setSelectedDocument] = useState(null);
  const [loading, setLoading] = useState(false);

  // Method
  const handleInitialization = useCallback(async () => {
    let response = null;

    try {
      setLoading(true);
      response = await apiHandler.get(apiConfig.resource.NUMBER_OF_DOCUMENTS);
      const totalDocumentNumber = response.data.data.totalDocumentNumber;
      setTotalDataCount(totalDocumentNumber);

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
          id: document.id,
          name: document.name,
          category: document.category.name,
          subcategory: document.subCategory.name,
          author: document.postMember.name,
          ref: document,
        });
      });

      setDocuments(updatedDocuments);
      setInitialized(true);
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    } finally {
      setLoading(false);
    }
  }, [orderBy, order, currentPage, pageSize]);

  const handleSearch = useCallback(
    async (dom) => {
      const keyword = dom?.target?.value;

      if (!initialized || documents.length === 0) {
        return;
      }

      if (!keyword || keyword.length < 3) {
        handleInitialization();
        return;
      }

      try {
        setLoading(true);
        const queryString = `?keyword=${keyword}`;
        const response = await apiHandler.get(
          `${apiConfig.resource.SEARCH_DOCUMENTS}${queryString}`
        );

        let updatedDocuments = [];
        updatedDocuments = response.data.data.map((document) => {
          return {
            id: document.id,
            name: document.name,
            category: document.category.name,
            subcategory: document.subCategory.name,
            author: document.postMember.name,
            ref: document,
          };
        });

        setDocuments(updatedDocuments);
      } catch (error) {
        messageUtil.showErrorMessage(errorMessage.notFound);
      } finally {
        setLoading(false);
      }
    },
    [documents.length, handleInitialization, initialized]
  );

  const handleChangeHeader = (newHeader, column, order) => {
    setTableHeader(newHeader);
    setOrderBy(column);
    setOrder(order);
  };

  const handleCloseModal = () => {
    setOpenDocumentModal(false);
    setOpenDeleteModal(false);
    setSelectedDocument(null);
  };

  const handleAddDocument = async () => {
    handleCloseModal();
    await handleInitialization();
  };

  const handleOpenDeleteModal = (document) => {
    setSelectedDocument(document);
    setOpenDeleteModal(true);
  };

  const handleDelete = async () => {
    try {
      const apiURL = apiConfig.resource.EDIT_DOCUMENT.replace(
        ":id",
        selectedDocument.id
      );
      await apiHandler.delete(apiURL);
      messageUtil.showSuccessMessage(commonMessage.success);

      handleCloseModal();
      handleInitialization();
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  // Side effect
  useEffect(() => {
    handleInitialization();
  }, [handleInitialization, orderBy, order, currentPage, pageSize]);

  useEffect(() => {
    const updatedTableData = documents.map((document, index) => {
      return {
        index: (currentPage - 1) * pageSize + index + 1,
        isEdit: false,
        name: (
          <LinkButton to={`/document/${document.id}`} title={document.name} />
        ),
        documentName: document.name,
        category: document.category,
        subcategory: document.subcategory,
        author: document.author,
        createdDate: formatDateTime(document.createdAt),
        action: (
          <div className='flex items-center gap-4'>
            <IconButton onClick={() => handleOpenDeleteModal(document)}>
              <FontAwesomeIcon icon={faTrash} />
            </IconButton>
          </div>
        ),
      };
    });

    setTableData(updatedTableData);
  }, [currentPage, documents, pageSize]);

  return (
    <>
      <Breadcrumbs linkList={linkList} />
      <div className='custom-container primary-bg'>
        <Form>
          <ToolbarBox>
            <SearchInputBox onChange={(dom) => handleSearch(dom)} />
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

      <FloatingButton handleExecution={() => setOpenDocumentModal(true)} />

      <DocumentModal
        openModal={openDocumentModal}
        selectedDocument={selectedDocument}
        onClose={handleCloseModal}
        onSubmit={handleAddDocument}
      />

      <DeleteDocumentModal
        deleteMode={true}
        openModal={openDeleteModal}
        selectedDocument={selectedDocument}
        onClose={() => handleCloseModal()}
        onSubmit={() => handleDelete()}
      />

      <LoadingBox visible={loading} />
    </>
  );
};

export default Documents;
