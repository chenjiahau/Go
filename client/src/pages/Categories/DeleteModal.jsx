import React from "react";
import PropTypes from "prop-types";

import ModalBox from "@/components/ModalBox";

const DeleteModal = ({
  openModal,
  deleteMode,
  selectedCategory,
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
      title='Delete Category'
      onClose={() => onClose()}
      onSubmit={() => onSubmit()}
    >
      <p>Do you want to delete this category({selectedCategory.name})?</p>
    </ModalBox>
  );
};

DeleteModal.propTypes = {
  openModal: PropTypes.bool,
  deleteMode: PropTypes.bool,
  selectedCategory: PropTypes.object,
  onClose: PropTypes.func,
  onSubmit: PropTypes.func,
};

export default DeleteModal;