import { useState, useEffect, useCallback } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import { orderBy } from "lodash";

// Const
import apiConfig from "@/const/config/api";
import routerConfig from "@/const/config/router";

// Component
import Editor from "@/components/Editor";
import Button from "@/components/Button";
import Dropdown from "@/components/Dropdown";

// Util
import apiHandler from "@/util/api.util";
import { getDefaultEditorData } from "@/util/editor.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const errorMessage = {
  author: "Author is required.",
};

const AddDocumentComment = () => {
  const navigate = useNavigate();
  const { id } = useParams();

  // State
  const [editorData, setEditorData] = useState(getDefaultEditorData());
  const [document, setDocument] = useState({});
  const [authors, setAuthors] = useState([]);
  const [reload, setReload] = useState(false);

  // Method
  const init = useCallback(async () => {
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

      // Author
      response = await apiHandler.get(apiConfig.resource.MEMBERS);
      let members = response?.data?.data || [];
      members = members.filter((member) => member.isAlive);
      members = orderBy(members, "name", "asc");

      const updatedAuthors = members.map((member) => ({
        id: member.id,
        name: member.name,
        selected: false,
      }));

      setAuthors(updatedAuthors);
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    }
  }, []);

  const reset = () => {
    setReload(true);
    init();
    setEditorData(getDefaultEditorData());
  };

  const add = () => async () => {
    const author = authors.find((author) => author.selected);
    if (!author) {
      messageUtil.showErrorMessage(errorMessage.author);
      return;
    }

    const content = JSON.stringify(editorData);

    const payload = {
      postMemberId: author.id,
      content: content,
    };

    try {
      await apiHandler.post(
        apiConfig.resource.ADD_DOCUMENT_COMMENT.replace(":id", id),
        payload
      );
      messageUtil.showSuccessMessage(commonMessage.success);
      reset();
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  // Side effect
  useEffect(() => {
    init();
  }, [init]);

  useEffect(() => {
    setTimeout(() => {
      setReload(false);
    }, 1000);
  }, [reload]);

  return (
    <>
      <div className='breadcrumb-container'>
        <Link to={routerConfig.routes.DOCUMENTS} className='breadcrumb--item'>
          <span className='breadcrumb--item--inner'>
            <span className='breadcrumb--item-title'>Documents</span>
          </span>
        </Link>
        <Link
          to={routerConfig.routes.DOCUMENT.replace(":id", id)}
          className='breadcrumb--item'
        >
          <span className='breadcrumb--item--inner'>
            <span className='breadcrumb--item-title'>Document</span>
          </span>
        </Link>
        <Link onClick={reset} className='breadcrumb--item'>
          <span className='breadcrumb--item--inner'>
            <span className='breadcrumb--item-title'>New</span>
          </span>
        </Link>
      </div>

      <div className='section'>
        <div className='space-t-3'></div>
        <div className='title'>{document.name}</div>
        <div className='space-t-3'></div>

        {/* Author */}
        <div className='form-group'>
          <Dropdown
            zIndex={10}
            list={authors}
            selectedItem={authors.find((member) => member.selected)}
            onChange={(item) => {
              setAuthors(
                authors.map((member) => ({
                  ...member,
                  selected: member.id === item.id,
                }))
              );
            }}
          />
        </div>
        <div className='space-t-3'></div>

        {/* Editor */}
        {!reload && (
          <div style={{ width: "100%" }}>
            <Editor
              reload={reload}
              data={editorData}
              onChange={setEditorData}
              editorBlock='editorjs-container'
            />
          </div>
        )}

        {/* Handler */}
        <div className='button-container'>
          <Button id='submitBtn' onClick={add()}>
            Add
          </Button>
          <Button extraClasses={["cancel-button"]} onClick={() => reset()}>
            Reset
          </Button>
        </div>
      </div>
    </>
  );
};

export default AddDocumentComment;
