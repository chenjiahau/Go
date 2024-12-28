import "@/pages/Documents/DocumentModal/module.css";

import React, { useState, useEffect, useCallback } from "react";
import PropTypes from "prop-types";

import ModalBox from "@/components/ModalBox";
import FormGroup from "@/components/FormGroup";
import FormLabel from "@/components/FormLabel";
import DropdownBox from "@/components/DropdownBox";
import InputBox from "@/components/InputBox";
import TagBox from "@/components/TagBox";
import EditorJS from "@/components/Editor";

// Const
import apiConfig from "@/const/config/api";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";
import { getDefaultEditorData } from "@/util/editor.util";

const errorMessage = {
  author: "Author is required.",
};

const CommentModal = ({
  openModal,
  selectedDocument,
  selectedComment,
  onClose,
  onSubmit,
}) => {
  // State
  const [title, setTitle] = useState("");
  const [authorOptions, setAuthorOptions] = useState([]);
  const [content, setContent] = useState(getDefaultEditorData());
  const [viewMode, setViewMode] = useState(false);
  const [reloadEditor, setReloadEditor] = useState(false);

  // Method
  const handleInitialization = useCallback(async () => {
    let response = null;

    try {
      setReloadEditor(true);

      // Roles
      response = await apiHandler.get(apiConfig.resource.MEMBER_ROLES);
      const roles = response?.data?.data?.memberRoles || [];

      // Title
      setTitle(selectedDocument?.name || "");

      // Author options
      response = await apiHandler.get(apiConfig.resource.MEMBERS);
      let members = response?.data?.data || [];
      members = members.sort((a, b) => a.name.localeCompare(b.name));

      const updatedAuthorOptions = members.map((member) => {
        const role = roles.find((role) => role.id === member.memberRoleId);

        return {
          value: member.id,
          label: `[${role.abbr}] ${member.name}`,
          active: false,
        };
      });

      if (selectedComment) {
        updatedAuthorOptions.forEach((option) => {
          option.active = option.value === selectedComment.ref.postMemberId;
        });
      }

      setAuthorOptions(updatedAuthorOptions);

      // Content
      if (selectedComment) {
        const data = JSON.parse(selectedComment.ref.content);
        setContent(data);
      }
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    } finally {
      setReloadEditor(false);
    }
  }, [selectedComment, selectedDocument?.name]);

  const handleChangeAuthor = (author) => {
    const updatedAuthorOptions = authorOptions.map((option) => {
      return {
        ...option,
        active: option.value === author.value,
      };
    });

    setAuthorOptions(updatedAuthorOptions);
  };

  const changeViewMode = () => {
    setViewMode(!viewMode);
    setReloadEditor(true);

    setTimeout(() => {
      setReloadEditor(false);
    }, 100);
  };

  const handleClose = () => {
    setTitle("");
    setAuthorOptions([]);
    onClose();
  };

  const handleSubmit = async () => {
    const author = authorOptions.find((author) => author.active);

    if (!author) {
      messageUtil.showErrorMessage(errorMessage.author);
      return;
    }

    const payload = {
      postMemberId: author.value,
      content: JSON.stringify(content),
    };

    try {
      if (selectedComment) {
        await apiHandler.put(
          apiConfig.resource.EDIT_DOCUMENT_COMMENT.replace(
            ":id",
            selectedDocument.id
          ).replace(":commentId", selectedComment.ref.id),
          payload
        );
      } else {
        await apiHandler.post(
          apiConfig.resource.ADD_DOCUMENT_COMMENT.replace(
            ":id",
            selectedDocument.ref.id
          ),
          payload
        );
      }

      messageUtil.showSuccessMessage(commonMessage.success);

      setTitle("");
      setAuthorOptions([]);
      onSubmit();
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  // Side effect
  useEffect(() => {
    if (openModal) {
      handleInitialization();
    }
  }, [handleInitialization, openModal]);

  if (!openModal) {
    return null;
  }

  let editorBlockId = "editorjs-container";
  if (selectedDocument) {
    editorBlockId = `editorjs-container-comment-${selectedDocument.id}-edit`;
  }

  return (
    <ModalBox
      enableScroll={true}
      title={selectedComment ? "Edit Comment" : "Add Comment"}
      customWidthClass='document-modal-body'
      onClose={() => handleClose()}
      onSubmit={() => handleSubmit()}
    >
      <FormGroup>
        <FormLabel forName='name'>Title</FormLabel>
        <InputBox
          disabled={true}
          type='text'
          id='ttile'
          name='title'
          placeholder='Title'
          value={title}
          onChange={(value) => setTitle(value)}
        />
      </FormGroup>
      <FormGroup>
        <FormLabel forName='name'>Author</FormLabel>
        <DropdownBox
          options={authorOptions}
          onClick={(author) => handleChangeAuthor(author)}
          zIndex={2}
        />
      </FormGroup>
      <FormGroup>
        <div className='w-32'>
          <TagBox
            tag={{
              label: viewMode ? "Edit Mode" : "View Mode",
            }}
            onClick={() => changeViewMode()}
          />
        </div>
      </FormGroup>
      <FormGroup>
        <div className='editorjs-container primary-shadow'>
          {!reloadEditor && (
            <EditorJS
              reload={reloadEditor}
              readOnly={viewMode}
              data={content}
              onChange={(data) => setContent(data)}
              editorBlock={editorBlockId}
            />
          )}
        </div>
      </FormGroup>
    </ModalBox>
  );
};

CommentModal.propTypes = {
  openModal: PropTypes.bool,
  members: PropTypes.array,
  selectedDocument: PropTypes.object,
  selectedComment: PropTypes.object,
  onClose: PropTypes.func,
  onSubmit: PropTypes.func,
};

export default CommentModal;
