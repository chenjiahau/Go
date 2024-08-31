import { useState, useEffect, useCallback, Fragment } from "react";
import { Link, useNavigate, useParams, useLocation } from "react-router-dom";
import { cloneDeep } from "lodash";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Component
import ContentViewer from "./components/ContentViewer";
import ConfirmationModal from "@/components/ConfirmationModal";

// Util
import apiHandler from "@/util/api.util";
import { formatDateTime } from "@/util/datetime.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const Document = () => {
  const navigate = useNavigate();
  const { search } = useLocation();
  const { id } = useParams();
  const keyword = new URLSearchParams(search).get("keyword");

  // State
  const [showDocument, setShowDocument] = useState(true);
  const [document, setDocument] = useState({});
  const [documentComments, setDocumentComments] = useState([]);
  const [selectedComment, setSelectedComment] = useState({});
  const [isOpenConfirmationModal, setIsOpenConfirmationModal] = useState(false);

  // Method
  const handleInitialization = useCallback(async () => {
    if (!id) {
      navigate(routerConfig.routes.DOCUMENTS);
    }

    let response = null;

    try {
      response = await apiHandler.get(
        apiConfig.resource.EDIT_DOCUMENT.replace(":id", id)
      );
      const document = response.data.data;
      setDocument(document);

      response = await apiHandler.get(
        apiConfig.resource.DOCUMENT_COMMENTS.replace(":id", id)
      );

      if (response.data.data) {
        setDocumentComments(
          response.data.data.map((comment) => {
            return {
              ...comment,
              showComment: keyword ? true : false,
            };
          })
        );
      } else {
        setDocumentComments([]);
      }
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    }
  }, [id, navigate]);

  const formatRelatedMembers = (members) => {
    const length = members.length;
    return members.map((member, index) => {
      return (
        <div key={member.id} className='tag author'>
          {member.name}
          {index < length - 1}
        </div>
      );
    });
  };

  const formatTags = (tags) => {
    return tags.map((tag) => {
      return (
        <div
          key={tag.id}
          className='tag'
          style={{ backgroundColor: tag.colorHexCode }}
        >
          {tag.tagName}
        </div>
      );
    });
  };

  const showConfirmationModal = (commentId) => {
    const comment = documentComments.find(
      (comment) => comment.id === commentId
    );
    setSelectedComment(comment);
    setIsOpenConfirmationModal(true);
  };

  const deleteComment = async () => {
    try {
      const apiURL = apiConfig.resource.EDIT_DOCUMENT_COMMENT.replace(
        ":id",
        id
      ).replace(":commentId", selectedComment.id);
      await apiHandler.delete(apiURL);
      messageUtil.showSuccessMessage(commonMessage.success);
      setIsOpenConfirmationModal(false);
      handleInitialization();
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  // Side effect
  useEffect(() => {
    handleInitialization();
  }, [handleInitialization]);

  if (!document.id) {
    return null;
  }

  return (
    <>
      <div className='breadcrumb-container'>
        <Link to={routerConfig.routes.DOCUMENTS} className='breadcrumb--item'>
          <span className='breadcrumb--item--inner'>
            <span className='breadcrumb--item-title'>Documents</span>
          </span>
        </Link>
      </div>

      {keyword ? (
        <div className='floating-button'>
          <i
            className='fa-solid fa-arrow-right-from-bracket'
            onClick={() => window.close()}
          />
        </div>
      ) : (
        <div className='floating-button'>
          <i
            className='fa-solid fa-plus'
            onClick={() =>
              navigate(
                routerConfig.routes.ADD_DOCUMENT_COMMENT.replace(":id", id)
              )
            }
          />
        </div>
      )}

      {/* Document */}
      <div className='document'>
        {keyword && (
          <h1 className='search-keyword'>
            <span>{`Search: ${keyword}`}</span>
          </h1>
        )}
        <div className='title'>{document.name}</div>
        <div className='box'>
          <div className='box-author'>
            <div className='left'>{document.postMember.name}</div>
            <div className='right'>
              <div className='data'>{formatDateTime(document.createdAt)}</div>
              <div className='action'>
                {!keyword && (
                  <>
                    <i
                      className='fa-solid fa-pen'
                      onClick={() =>
                        navigate(
                          routerConfig.routes.EDIT_DOCUMENT.replace(":id", id)
                        )
                      }
                    />
                    <i
                      className='fa-solid fa-expand'
                      onClick={() => setShowDocument(!showDocument)}
                    />
                  </>
                )}
              </div>
            </div>
          </div>
          {showDocument && (
            <>
              <div className='tag-container box-category'>
                <div className='tag'>{document.category?.name}</div>
                <div className='tag'>{document.subCategory?.name}</div>
              </div>
              {document.relationMembers?.length > 0 && (
                <div className='tag-container box-related-members'>
                  {formatRelatedMembers(document.relationMembers)}
                </div>
              )}

              <div className='box-content'>
                <ContentViewer content={document.content} />
              </div>
              {document.tags?.length > 0 && (
                <div className='box-tags'>
                  <div className='tag-container' style={{ maxHeight: "auto" }}>
                    {formatTags(document.tags)}
                  </div>
                </div>
              )}
            </>
          )}
        </div>
      </div>

      {/* Comments */}
      {documentComments.length > 0 && (
        <>
          {documentComments.map((comment, index) => {
            return (
              <Fragment key={index}>
                <div className='document' key={index}>
                  <div className='box'>
                    <div key={comment.id} className='box-comment'>
                      <div className='box-author'>
                        <div className='left'>{comment.postMemberName}</div>
                        <div className='right'>
                          <div className='data'>
                            {formatDateTime(comment.createdAt)}
                          </div>
                          <div className='action'>
                            {!keyword && (
                              <>
                                <i
                                  className='fa-solid fa-pen'
                                  onClick={() =>
                                    navigate(
                                      routerConfig.routes.EDIT_DOCUMENT_COMMENT.replace(
                                        ":id",
                                        id
                                      ).replace(":commentId", comment.id)
                                    )
                                  }
                                />
                                <i
                                  className='fa-solid fa-trash'
                                  onClick={() =>
                                    showConfirmationModal(comment.id)
                                  }
                                />
                                <i
                                  className='fa-solid fa-expand'
                                  onClick={() => {
                                    {
                                      const updatedDocumentComments =
                                        cloneDeep(documentComments);
                                      updatedDocumentComments[
                                        index
                                      ].showComment = !comment.showComment;
                                      setDocumentComments(
                                        updatedDocumentComments
                                      );
                                    }
                                  }}
                                />
                              </>
                            )}
                          </div>
                        </div>
                      </div>
                      {comment.showComment && (
                        <div className='box-content'>
                          <ContentViewer content={comment.content} />
                        </div>
                      )}
                    </div>
                  </div>
                </div>
                {documentComments.length - 1 === index && (
                  <div className='space-t-4'></div>
                )}
              </Fragment>
            );
          })}
        </>
      )}

      <ConfirmationModal
        isOpen={isOpenConfirmationModal}
        onClose={() => {
          setSelectedComment(null);
          setIsOpenConfirmationModal(false);
        }}
        onConfirm={deleteComment}
      />
    </>
  );
};

export default Document;
