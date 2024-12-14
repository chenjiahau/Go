import { useState, useEffect, useCallback } from "react";
import PropTypes from "prop-types";

import ModalBox from "@/components/ModalBox";
import FormGroup from "@/components/FormGroup";
import FormLabel from "@/components/FormLabel";
import DropdownBox from "@/components/DropdownBox";
import InputBox from "@/components/InputBox";

// Const
import apiConfig from "@/const/config/api";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const errorMessage = {
  name: "Name is required.",
  duplicated: "Name is duplicated.",
};

const MemberModal = ({ openModal, roles, onClose, onSubmit }) => {
  // State
  const [name, setName] = useState("");
  const [roleOptions, setRoleOptions] = useState([]);

  // Method
  const handleInitialization = useCallback(() => {
    setName("");
    const updatedRoleOptions = roles.map((role) => ({
      value: role.id,
      label: role.title,
      active: false,
    }));

    updatedRoleOptions.length > 0 && (updatedRoleOptions[0].active = true);
    setRoleOptions(updatedRoleOptions);
  }, [roles]);

  const handleRoleChange = (option) => {
    const updatedRoleOptions = roleOptions.map((roleOption) => {
      if (roleOption.value === option.value) {
        roleOption.active = true;
      } else {
        roleOption.active = false;
      }

      return roleOption;
    });

    setRoleOptions(updatedRoleOptions);
  };

  const handleClose = () => {
    setName("");
    onClose();
  };

  const handleSubmit = async () => {
    if (name === "") {
      messageUtil.showErrorMessage(errorMessage.name);
      return;
    }

    try {
      await apiHandler.post(apiConfig.resource.ADD_MEMBER, {
        memberRoleId: roleOptions.find((role) => role.active).value,
        name,
      });

      messageUtil.showSuccessMessage(commonMessage.success);
      handleInitialization();
      onSubmit();
    } catch (error) {
      if (error.response.data.code === 2402) {
        messageUtil.showErrorMessage(errorMessage.duplicated);
        return;
      }

      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  // Side effect
  useEffect(() => {
    handleInitialization();
  }, [roles, handleInitialization]);

  if (!openModal) {
    return null;
  }

  return (
    <ModalBox
      title='Add Member'
      onClose={() => handleClose()}
      onSubmit={() => handleSubmit()}
    >
      <FormGroup>
        <FormLabel forName='name'>Name</FormLabel>
        <InputBox
          type='text'
          id='name'
          name='name'
          placeholder='Name'
          value={name}
          onChange={(value) => setName(value)}
        />
      </FormGroup>
      <FormGroup>
        <FormLabel forName='name'>Role</FormLabel>
        <DropdownBox
          options={roleOptions}
          onClick={(option) => handleRoleChange(option)}
          zIndex={1}
        />
      </FormGroup>
    </ModalBox>
  );
};

MemberModal.propTypes = {
  openModal: PropTypes.bool,
  selectedMember: PropTypes.object,
  roles: PropTypes.array,
  onClose: PropTypes.func,
  onSubmit: PropTypes.func,
};

export default MemberModal;
