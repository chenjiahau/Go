import "./module.css";

import { useEffect } from "react";
import PropTypes from "prop-types";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faXmark } from "@fortawesome/free-solid-svg-icons";
import { isFunction } from "lodash";

import IconButton from "@/components/IconButton";
import ButtonBox from "@/components/ButtonBox";

const ModalBox = ({
  enableMessageModal,
  enableScroll,
  customWidthClass,
  confirmMode,
  deleteMode,
  title,
  onClose,
  onSubmit,
  children,
}) => {
  enableMessageModal = enableMessageModal || false;
  customWidthClass = customWidthClass || "";
  confirmMode = confirmMode || false;
  deleteMode = deleteMode || false;
  title = title || "Modal Title";
  onClose = isFunction(onClose) ? onClose : () => {};
  onSubmit = isFunction(onSubmit) ? onSubmit : () => {};

  useEffect(() => {
    const nav = document.querySelector("nav");
    nav.style.zIndex = "0";

    const mainContent = document.querySelector(".main-content");
    mainContent.style.overflow = "hidden";

    const modalBox = document.querySelector(".modal-box");
    modalBox.style.zIndex = "1000";

    return () => {
      nav.style.display = "block";
      nav.style.zIndex = "1000";

      mainContent.style.overflow = "auto";

      modalBox.style.zIndex = "0";
    };
  }, []);

  return (
    <div className='modal-box'>
      <div
        className={`modal-body ${
          enableScroll ? "!overflow-y-auto" : "!overflow-visible"
        } ${customWidthClass}`}
      >
        <div className='modal-header'>
          <h1 className='title'>{title}</h1>
          <IconButton onClick={onClose}>
            <FontAwesomeIcon icon={faXmark} />
          </IconButton>
        </div>
        <div className='modal-content'>{children}</div>
        <div
          className={`${
            enableMessageModal ? "modal-message-footer " : "modal-footer "
          }`}
        >
          {confirmMode ? (
            <ButtonBox
              onClick={onSubmit}
              isSave={true}
              extraClasses={["primary-shadow"]}
            >
              Done
            </ButtonBox>
          ) : (
            <>
              <ButtonBox
                onClick={onSubmit}
                isSave={true}
                extraClasses={["primary-shadow"]}
              >
                {deleteMode ? "Delete" : "Save"}
              </ButtonBox>
              <ButtonBox
                onClick={onClose}
                isClose={true}
                extraClasses={["cancel-shadow"]}
              >
                Close
              </ButtonBox>
            </>
          )}
        </div>
      </div>
    </div>
  );
};

ModalBox.propTypes = {
  enableMessageModal: PropTypes.bool,
  enableScroll: PropTypes.bool,
  customWidthClass: PropTypes.string,
  confirmMode: PropTypes.bool,
  deleteMode: PropTypes.bool,
  title: PropTypes.string,
  onClose: PropTypes.func,
  onSubmit: PropTypes.func,
  children: PropTypes.node,
};

export default ModalBox;
