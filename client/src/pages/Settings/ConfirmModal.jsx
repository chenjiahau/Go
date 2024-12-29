import React from "react";
import PropTypes from "prop-types";
import { isFunction } from "lodash";

import ModalBox from "@/components/ModalBox";

const ConfirmModal = ({ openModal, onClose, onSubmit }) => {
  if (!openModal) {
    return null;
  }

  onClose = isFunction(onClose) ? onClose : () => {};
  onSubmit = isFunction(onSubmit) ? onSubmit : () => {};

  return (
    <ModalBox
      enableMessageModal={true}
      confirmMode={true}
      title='Warning'
      onClose={() => onClose()}
      onSubmit={() => onSubmit()}
    >
      <p>You need to log out to apply the changes.</p>
    </ModalBox>
  );
};

ConfirmModal.propTypes = {
  openModal: PropTypes.bool,
  onSubmit: PropTypes.func,
};

export default ConfirmModal;
