import { useState, useEffect } from "react";
import PropTypes from "prop-types";

import ModalBox from "@/components/ModalBox";
import FormGroup from "@/components/FormGroup";
import FormLabel from "@/components/FormLabel";
import DropdownBox from "@/components/DropdownBox";

// Const
import apiConfig from "@/const/config/api";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const RoleModal = ({ openModal, selectedMember, roles, onClose, onSubmit }) => {
  const [roleOptions, setRoleOptions] = useState([]);

  // Method
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
    onClose();
  };

  const handleSubmit = async () => {
    const selectedRole = roleOptions.find((roleOption) => roleOption.active);

    if (!selectedRole) {
      return;
    }

    try {
      const apiURL = apiConfig.resource.EDIT_MEMBER.replace(
        ":id",
        selectedMember.id
      );
      const payload = {
        ...selectedMember,
        memberRoleId: selectedRole.value,
      };

      await apiHandler.put(apiURL, payload);
      messageUtil.showSuccessMessage(commonMessage.success);
      onSubmit();
    } catch (error) {
      messageUtil.showErrorMessage(error.response.data.message);
    }
  };

  // Side effect
  useEffect(() => {
    if (roles.length === 0) {
      return;
    }

    const updatedRoleOptions = roles.map((role) => ({
      value: role.id,
      label: role.title,
      active: false,
    }));

    if (selectedMember) {
      const selectedOption = updatedRoleOptions.find(
        (option) => option.value === selectedMember.memberRoleId
      );

      selectedOption.active = true;
    } else {
      updatedRoleOptions[0].active = true;
    }

    setRoleOptions(updatedRoleOptions);
  }, [selectedMember, roles]);

  if (!openModal || roleOptions.length === 0) {
    return null;
  }

  return (
    <ModalBox
      title='Edit Role'
      onClose={() => handleClose()}
      onSubmit={() => handleSubmit()}
    >
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

RoleModal.propTypes = {
  openModal: PropTypes.bool,
  selectedMember: PropTypes.object,
  roles: PropTypes.array,
  onClose: PropTypes.func,
  onSubmit: PropTypes.func,
};

export default RoleModal;
