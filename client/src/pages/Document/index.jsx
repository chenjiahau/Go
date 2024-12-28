import "./module.css";

import { useState, useEffect, useCallback, Fragment } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faFile, faTrash } from "@fortawesome/free-solid-svg-icons";

import Breadcrumbs from "@/components/Breadcrumbs";
import MainTitle from "@/components/MainTitle";
import Hr from "@/components/Hr";
import Form from "@/components/Form";
import IconButton from "@/components/IconButton";
import TagBox from "@/components/TagBox";
import EditorJS from "@/components/Editor";
import FloatingButton from "@/components/FloatingButton";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Component
import DocumentModal from "@/pages/Documents/DocumentModal";
import CommentModal from "@/pages/Document/CommentModal";
import DeleteDocumentModal from "@/pages/Documents/DeleteDocumentModal";
import DeleteCommentModal from "./DeleteCommentModal";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";
import { formatDateTime } from "@/util/datetime.util";

const Document = () => {
  const navigate = useNavigate();
  const { id } = useParams();
  const [linkList, setLinkList] = useState([]);

  // State
  const [document, setDocument] = useState(null);
  const [comments, setComments] = useState([]);
  const [openDocumentModal, setOpenDocumentModal] = useState(false);
  const [openCommentModal, setOpenCommentModal] = useState(false);
  const [openDeleteDocumentModal, setOpenDeleteDocumentModal] = useState(false);
  const [openDeleteCommentModal, setOpenDeleteCommentModal] = useState(false);
  const [selectedComment, setSelectedComment] = useState(null);
  const [reloadContent, setReloadContent] = useState(false);

  // Method
  const handleInitialization = useCallback(async () => {
    if (!id) {
      navigate(routerConfig.routes.DOCUMENTS);
    }

    let response = null;

    try {
      setReloadContent(true);

      response = await apiHandler.get(
        apiConfig.resource.EDIT_DOCUMENT.replace(":id", id)
      );

      const data = response.data.data;

      let tags = data.tags
        ?.map((tag) => {
          return {
            label: tag.tagName,
            hashCode: tag.colorHexCode,
          };
        })
        .sort((a, b) => a.label.localeCompare(b.label));

      let members = data.relationMembers
        ?.map((member) => {
          return {
            label: member.name,
          };
        })
        .sort((a, b) => a.label.localeCompare(b.label));

      const updatedDocument = {
        id: data.id,
        name: data.name,
        author: data.postMember.name,
        category: data.category.name,
        subcategory: data.subCategory.name,
        tags,
        members,
        content: JSON.parse(data.content),
        createdDate: formatDateTime(data.createdDate),
        ref: data,
      };
      setDocument(updatedDocument);

      response = await apiHandler.get(
        apiConfig.resource.DOCUMENT_COMMENTS.replace(":id", id)
      );

      if (response.data.data) {
        const updatedComments = response.data.data.map((comment, index) => {
          return {
            index,
            id: comment.id,
            author: comment.postMemberName,
            content: JSON.parse(comment.content),
            createdDate: formatDateTime(comment.createdDate),
            ref: comment,
          };
        });

        setComments(updatedComments);
      } else {
        setComments([]);
      }
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
      navigate(routerConfig.routes.DOCUMENTS);
    } finally {
      setReloadContent(false);
    }
  }, [id, navigate]);

  const handleCloseModal = () => {
    setOpenDocumentModal(false);
    setOpenCommentModal(false);
    setOpenDeleteDocumentModal(false);
    setOpenDeleteCommentModal(false);
    setSelectedComment(null);
  };

  const handleUpdateDocument = async () => {
    handleCloseModal();
    await handleInitialization();
  };

  const handleAddCommentModal = async () => {
    handleCloseModal();
    await handleInitialization();
  };

  const handleEditCommentModal = (comment) => {
    setSelectedComment(comment);
    setOpenCommentModal(true);
  };

  const handleOpenDeleteDocumentModal = () => {
    setOpenDeleteDocumentModal(true);
  };

  const handleDeleteDocument = async () => {
    try {
      const apiURL = apiConfig.resource.EDIT_DOCUMENT.replace(
        ":id",
        document.id
      );
      await apiHandler.delete(apiURL);
      messageUtil.showSuccessMessage(commonMessage.success);

      navigate(routerConfig.routes.DOCUMENTS);
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  const handleOpenDeleteCommentModal = (comment) => {
    setSelectedComment(comment);
    setOpenDeleteCommentModal(true);
  };

  const handleDeleteComment = async () => {
    try {
      const apiURL = apiConfig.resource.EDIT_DOCUMENT_COMMENT.replace(
        ":id",
        document.id
      ).replace(":commentId", selectedComment.id);
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
  }, [handleInitialization, id]);

  useEffect(() => {
    if (!document) {
      return;
    }

    const updatedLinkList = [
      { to: "/", label: "Home" },
      { to: "/documents", label: "Documents" },
      {
        label: document.name,
      },
    ];

    setLinkList(updatedLinkList);
  }, [document, id]);

  if (reloadContent || !document) {
    return null;
  }

  return (
    <>
      <Breadcrumbs linkList={linkList} />
      <div className='custom-container primary-bg'>
        <Form>
          <div className='document-title'>
            <div>
              <MainTitle extraClasses={["text-title", "!text-primary", "mb-2"]}>
                {document.name}
              </MainTitle>
            </div>
            <div className='flex items-center gap-4'>
              <IconButton onClick={() => setOpenDocumentModal(true)}>
                <FontAwesomeIcon icon={faFile} />
              </IconButton>
              {comments.length === 0 && (
                <IconButton
                  onClick={() => handleOpenDeleteDocumentModal(document)}
                >
                  <FontAwesomeIcon icon={faTrash} />
                </IconButton>
              )}
            </div>
          </div>
          <Hr />
          <div className='document-block mt-4 mb-4'>
            <div className='document-info'>
              <div className='document-info-item'>
                <p className='title'>Author</p>
                <p className='label'>{document.author}</p>
              </div>
              <div className='document-info-item'>
                <p className='title'>Category</p>
                <p className='label'>{document.category}</p>
              </div>
              <div className='document-info-item'>
                <p className='title'>Subcategory</p>
                <p className='label'>{document.subcategory}</p>
              </div>
              <div className='document-info-item'>
                <p className='title'>Tags</p>
                <div className='label'>
                  {document.tags?.map((tag) => (
                    <TagBox key={tag.label} tag={tag} />
                  ))}
                </div>
              </div>
              <div className='document-info-item'>
                <p className='title'>Members</p>
                <div className='label'>
                  {document.members?.map((member, index) => (
                    <Fragment key={index}>
                      <TagBox tag={member} />
                    </Fragment>
                  ))}
                </div>
              </div>
            </div>
            <div className='document-content'>
              <div className='editorjs-container '>
                {!reloadContent && (
                  <EditorJS
                    readOnly={true}
                    extraClasses={["m-0"]}
                    data={document.content}
                    editorBlock={`editorjs-container-document-${document.id}`}
                  />
                )}
              </div>
            </div>
          </div>
        </Form>
      </div>

      {/* Comments */}
      {comments.map((comment, index) => (
        <div key={index} className='custom-container primary-bg !pt-0'>
          <Form>
            <div className='document-title'>
              <div>
                <MainTitle
                  extraClasses={["text-title", "!text-primary", "mb-2"]}
                >
                  Comment #{index + 1}
                </MainTitle>
              </div>
              <div className='flex items-center gap-4'>
                <IconButton onClick={() => handleEditCommentModal(comment)}>
                  <FontAwesomeIcon icon={faFile} />
                </IconButton>
                <IconButton
                  onClick={() => handleOpenDeleteCommentModal(comment)}
                >
                  <FontAwesomeIcon icon={faTrash} />
                </IconButton>
              </div>
            </div>
            <Hr extraClasses={["mb-4"]} />

            <div key={index} className='comment-block mt-4'>
              <div className='comment-info'>
                <div className='comment-info-item'>
                  <p className='title'>Author</p>
                  <p className='label'>{comment.author}</p>
                </div>
                <div className='comment-info-item'>
                  <p className='title'>Created Date</p>
                  <p className='label'>{comment.createdDate}</p>
                </div>
              </div>
              <div className='comment-content'>
                <div className='editorjs-container primary-shadow'>
                  <EditorJS
                    readOnly={true}
                    extraClasses={["m-0"]}
                    data={comment.content}
                    editorBlock={`editorjs-container-comment-${document.id}-${comment.id}`}
                  />
                </div>
              </div>
            </div>
          </Form>
        </div>
      ))}

      <FloatingButton handleExecution={() => setOpenCommentModal(true)} />

      <DocumentModal
        openModal={openDocumentModal}
        selectedDocument={document.ref}
        onClose={handleCloseModal}
        onSubmit={handleUpdateDocument}
      />

      <CommentModal
        openModal={openCommentModal}
        selectedDocument={document}
        selectedComment={selectedComment}
        onClose={handleCloseModal}
        onSubmit={() => handleAddCommentModal()}
      />

      <DeleteDocumentModal
        deleteMode={true}
        openModal={openDeleteDocumentModal}
        selectedDocument={document}
        onClose={() => handleCloseModal()}
        onSubmit={() => handleDeleteDocument()}
      />

      <DeleteCommentModal
        deleteMode={true}
        openModal={openDeleteCommentModal}
        selectedComment={selectedComment}
        onClose={() => handleCloseModal()}
        onSubmit={() => handleDeleteComment()}
      />
    </>
  );
};

export default Document;
