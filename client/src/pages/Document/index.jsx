import { useState, useEffect, useCallback } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import { cloneDeep } from "lodash";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Component
import ContentViewer from "./components/ContentViewer";

// Util
import apiHandler from "@/util/api.util";
import { formatDateTime } from "@/util/datetime.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const Document = () => {
  const navigate = useNavigate();
  const { id } = useParams();

  // State
  const [showDocument, setShowDocument] = useState(true);
  const [document, setDocument] = useState({});
  const [documentComments, setDocumentComments] = useState([]);

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
      const document = response.data.data.document;
      setDocument(document);

      response = await apiHandler.get(
        apiConfig.resource.DOCUMENT_COMMENTS.replace(":id", id)
      );

      if (response.data.data.documentComments) {
        setDocumentComments(
          response.data.data.documentComments.map((comment) => {
            return {
              ...comment,
              showComment: false,
            };
          })
        );
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

      {/* Document */}
      <div className='document'>
        <div className='title'>{document.name}</div>
        <div className='box'>
          <div className='box-author'>
            <div className='left'>{document.postMember.name}</div>
            <div className='right'>
              <div className='data'>{formatDateTime(document.createdAt)}</div>
              <div className='action'>
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
                            className='fa-solid fa-expand'
                            onClick={() => {
                              {
                                const updatedDocumentComments =
                                  cloneDeep(documentComments);
                                updatedDocumentComments[index].showComment =
                                  !comment.showComment;
                                setDocumentComments(updatedDocumentComments);
                              }
                            }}
                          />
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
            );
          })}
        </>
      )}
    </>
  );
};

export default Document;
