import React from "react";
import PropTypes from "prop-types";

import ModalBox from "@/components/ModalBox";

const DeleteCommentModal = ({
  openModal,
  deleteMode,
  selectedComment,
  onClose,
  onSubmit,
}) => {
  if (!openModal) {
    return null;
  }

  return (
    <ModalBox
      enableMessageModal={true}
      deleteMode={deleteMode}
      title='Delete Comment'
      onClose={() => onClose()}
      onSubmit={() => onSubmit()}
    >
      <p>Do you want to delete this comment(#{selectedComment.index + 1})?</p>
    </ModalBox>
  );
};

DeleteCommentModal.propTypes = {
  openModal: PropTypes.bool,
  deleteMode: PropTypes.bool,
  selectedDocument: PropTypes.object,
  onClose: PropTypes.func,
  onSubmit: PropTypes.func,
};

export default DeleteCommentModal;
