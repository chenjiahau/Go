import Button from "../Button";

const ConfirmationModal = (props) => {
  const { isOpen, onClose, onConfirm } = props;

  return (
    <>
      <div className={`overlay ${isOpen ? "display" : "hidden"}`}>
        <div className={`modal ${isOpen ? "display" : "hidden"}`}>
          <div className='modal-header'>
            <div className='modal-title'>Confirmation</div>
            <div className='modal-close' onClick={onClose}>
              &times;
            </div>
          </div>
          <div className='modal-body'>
            <p>Do you want to delete it?</p>
          </div>
          <div className='modal-footer'>
            <Button extraClasses={["cancel-button"]} onClick={onClose}>
              Close
            </Button>
            <div className='space-r-2'></div>
            <Button extraClasses={["delete-button"]} onClick={onConfirm}>
              Delete
            </Button>
          </div>
        </div>
      </div>
    </>
  );
};

export default ConfirmationModal;
