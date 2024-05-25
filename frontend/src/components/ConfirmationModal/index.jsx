import Modal from "react-bootstrap/Modal";

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
        <button
          type='button'
          className='button cancel-button'
          onClick={onClose}
        >
          Close
        </button>
        <button
          type='button'
          className='button delete-button'
          onClick={onConfirm}
        >
          Delete
        </button>
      </Modal.Footer>
    </Modal>
  );
};

export default ConfirmationModal;
