import Modal from "react-bootstrap/Modal";

import Button from "../Button";

const ConfirmationModal = (props) => {
  const { isOpen, onClose, onConfirm } = props;

  return (
    <Modal centered={true} show={isOpen} onHide={onClose}>
      <Modal.Header closeButton>
        <Modal.Title>Confirmation</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <p>Do you want to delete it?</p>
      </Modal.Body>
      <Modal.Footer>
        <Button extraClasses={["cancel-button"]} onClick={onClose}>
          Close
        </Button>
        <Button extraClasses={["delete-button"]} onClick={onConfirm}>
          Delete
        </Button>
      </Modal.Footer>
    </Modal>
  );
};

export default ConfirmationModal;
