import { useState } from "react";
import PropTypes from "prop-types";

import ModalBox from "@/components/ModalBox";
import FormGroup from "@/components/FormGroup";
import FormLabel from "@/components/FormLabel";
import ElementGroup from "@/components/ElementGroup";
import InputBox from "@/components/InputBox";
import RadioBox from "@/components/RadioBox";

// Const
import apiConfig from "@/const/config/api";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const errorMessage = {
  category: "Name is required.",
  duplicated: "Name is duplicated.",
};

const SubcategoryModal = ({ categoryId, openModal, onClose, onSubmit }) => {
  const [name, setName] = useState("");
  const [status, setStatus] = useState(false);

  // Method
  const handleClose = () => {
    setName("");
    setStatus(false);
    onClose();
  };

  const handleSubmit = async () => {
    if (name === "") {
      messageUtil.showErrorMessage(errorMessage.name);
      return;
    }

    try {
      await apiHandler.post(
        apiConfig.resource.ADD_SUBCATEGORY.replace(":id", categoryId),
        {
          name,
          isAlive: status,
        }
      );

      setName("");
      setStatus(false);
      messageUtil.showSuccessMessage(commonMessage.success);
      onSubmit();
    } catch (error) {
      if (error.response.data.code === 4402) {
        messageUtil.showErrorMessage(errorMessage.duplicated);
        return;
      }
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  if (!openModal) {
    return null;
  }

  return (
    <ModalBox
      title='Add Subcategory'
      onClose={() => handleClose()}
      onSubmit={() => handleSubmit()}
    >
      <FormGroup>
        <FormGroup>
          <FormLabel forName='name'>Name</FormLabel>
          <InputBox
            type='text'
            id='name'
            name='name'
            placeholder='Category Name'
            value={name}
            onChange={(value) => setName(value)}
          />
        </FormGroup>
        <FormGroup>
          <FormLabel forName='name'>Status</FormLabel>
          <ElementGroup>
            <RadioBox checked={status} onChange={() => setStatus(true)}>
              Enable
            </RadioBox>
            <RadioBox checked={!status} onChange={() => setStatus(false)}>
              Disable
            </RadioBox>
          </ElementGroup>
        </FormGroup>
      </FormGroup>
    </ModalBox>
  );
};

SubcategoryModal.propTypes = {
  openModal: PropTypes.bool,
  onClose: PropTypes.func,
  onSubmit: PropTypes.func,
};

export default SubcategoryModal;
