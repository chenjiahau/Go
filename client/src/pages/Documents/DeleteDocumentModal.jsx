import React from "react";
import PropTypes from "prop-types";

import ModalBox from "@/components/ModalBox";

const DeleteDocumentModal = ({
  openModal,
  deleteMode,
  selectedDocument,
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
      title='Delete Document'
      onClose={() => onClose()}
      onSubmit={() => onSubmit()}
    >
      <p>Do you want to delete this document({selectedDocument.name})?</p>
    </ModalBox>
  );
};

DeleteDocumentModal.propTypes = {
  openModal: PropTypes.bool,
  deleteMode: PropTypes.bool,
  selectedDocument: PropTypes.object,
  onClose: PropTypes.func,
  onSubmit: PropTypes.func,
};

export default DeleteDocumentModal;
