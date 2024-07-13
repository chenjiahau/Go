import { useState, useEffect, useCallback } from "react";
import Button from "@/components/Button";
import Dropdown from "@/components/Dropdown";

const MemberRoleModal = (props) => {
  const { isOpen, onClose, onConfirm, memberRoles, selectedMemberRole } = props;

  // State
  const [memberRoleDropdown, setMemberRoleDropdown] = useState([]);

  // Method
  const init = useCallback(() => {
    if (!selectedMemberRole.id || memberRoles.length === 0) {
      return;
    }

    const updatedMemberRoleDropdown = memberRoles.map((memberRole) => {
      return {
        id: memberRole.id,
        name: `${memberRole.title}(${memberRole.abbr})`,
        selected: memberRole.id === selectedMemberRole.memberRoleId,
      };
    });

    setMemberRoleDropdown(updatedMemberRoleDropdown);
  }, [memberRoles, selectedMemberRole]);

  const getSelectedMemberRole = () => {
    return memberRoleDropdown.find((memberRole) => memberRole.selected);
  };

  // Side effect
  useEffect(() => {
    init();
  }, [init]);

  return (
    <>
      <div className={`modal ${isOpen ? "display" : "hidden"}`}>
        <div className='modal-header'>
          <div className='modal-title'>Change member role</div>
          <div className='modal-close' onClick={onClose}>
            &times;
          </div>
        </div>
        <div className='modal-body'>
          <div style={{ width: "100%" }}>
            <Dropdown
              zIndex={8}
              list={memberRoleDropdown}
              selectedItem={memberRoleDropdown.find((item) => item.selected)}
              onChange={(item) => {
                const updatedMemberRoleDropdown = memberRoleDropdown.map(
                  (memberRole) => {
                    return {
                      ...memberRole,
                      selected: memberRole.id === item.id,
                    };
                  }
                );
                setMemberRoleDropdown(updatedMemberRoleDropdown);
              }}
            />
          </div>
          <div className='space-t-2'></div>
        </div>
        <div className='modal-footer'>
          <Button extraClasses={["cancel-button"]} onClick={onClose}>
            Close
          </Button>
          <div className='space-r-2'></div>
          <Button onClick={() => onConfirm(getSelectedMemberRole())}>
            Save
          </Button>
        </div>
      </div>
      <div className={`overlay ${isOpen ? "display" : "hidden"}`}></div>
    </>
  );
};

export default MemberRoleModal;
