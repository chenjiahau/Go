import { useState, useCallback, useEffect } from "react";

// Const
import apiConfig from "@/const/config/api";

// Component
import Input from "@/components/Input";
import Button from "@/components/Button";
import Dropdown from "@/components/Dropdown";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const errorMessage = {
  member: "Member is required.",
  duplicated: "Member is duplicated.",
};

const Add = (props) => {
  const { onInitialization, memberRoles } = props;

  // State
  const [memberRoleDropdown, setMemberRoleDropdown] = useState([]);
  const [member, setMember] = useState("");

  // Method
  const init = useCallback(() => {
    if (memberRoles.length === 0) {
      return;
    }

    const updatedMemberRoleDropdown = memberRoles.map((memberRole) => {
      return {
        id: memberRole.id,
        name: `${memberRole.title}(${memberRole.abbr})`,
        selected: false,
      };
    });
    updatedMemberRoleDropdown[0].selected = true;

    setMemberRoleDropdown(updatedMemberRoleDropdown);
  }, [memberRoles]);

  const addMember = async () => {
    if (member === "") {
      messageUtil.showErrorMessage(errorMessage.member);
      return;
    }

    try {
      await apiHandler.post(apiConfig.resource.ADD_MEMBER, {
        memberRoleId: memberRoleDropdown.find(
          (memberRole) => memberRole.selected
        ).id,
        name: member,
      });
      setMember("");
      init();
      messageUtil.showSuccessMessage(commonMessage.success);
      onInitialization();
    } catch (error) {
      if (error.response.data.error.code === 500) {
        messageUtil.showErrorMessage(errorMessage.duplicated);
        return;
      }
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  // Side effect
  useEffect(() => {
    init();
  }, [init]);

  return (
    <div className='section'>
      <div className='input-group'>
        <Input
          id='member'
          type='text'
          name='member'
          autoComplete='off'
          placeholder='New member'
          value={member}
          onChange={(e) => setMember(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              addMember();
            }
          }}
          required
        />
      </div>
      <div className='space-t-2'></div>
      <div>
        <Dropdown
          zIndex={6}
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
      <div className='text-right'>
        <Button onClick={addMember}>
          <i className='fa-solid fa-plus'></i>
          <span className='space-l-1'>Add</span>
        </Button>
      </div>
    </div>
  );
};

export default Add;
