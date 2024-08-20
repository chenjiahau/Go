import { useState } from "react";

// Const
import apiConfig from "@/const/config/api";

// Component
import Input from "@/components/Input";
import Button from "@/components/Button";
import RadioButton from "@/components/RadioButton";
import MemberRoleModal from "./MemberRoleModal";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const Table = (props) => {
  const {
    currentPage,
    pageSize,
    orderBy,
    order,
    onChangeOrder,
    members,
    onClickMemberName,
    onChangeMemberName,
    onChangeMemberAlive,
    selectedMember,
    onShowConfirmationModal,
    saveMember,
    onInitialization,
    memberRoles,
  } = props;

  // State
  const [selectedMemberRole, setSelectedMemberRole] = useState({});
  const [isOpenMemberRoleModal, setIsOpenMemberRoleModal] = useState(false);

  // Method
  const changeMemberRole = async (memberRole) => {
    try {
      const apiURL = apiConfig.resource.EDIT_MEMBER.replace(
        ":id",
        selectedMemberRole.id
      );
      const payload = {
        ...selectedMemberRole,
        memberRoleId: memberRole.id,
      };

      await apiHandler.put(apiURL, payload);
      messageUtil.showSuccessMessage(commonMessage.success);

      onInitialization();
      setIsOpenMemberRoleModal(false);
    } catch (error) {
      messageUtil.showErrorMessage(error.response.data.message);
    }
  };

  return (
    <>
      <div className='section'>
        <div className='table-response'>
          <table className='table table-hover'>
            <thead>
              <tr>
                <th width='30'>
                  <div className='n'>#</div>
                </th>
                <th onClick={() => onChangeOrder("name")}>
                  <div className='order'>
                    <span>Member</span>
                    {orderBy === "name" && (
                      <i
                        className={`fa-solid fa-arrow-${
                          order === "asc" ? "down" : "up"
                        }`}
                      />
                    )}
                  </div>
                </th>
                <th width='360' onClick={() => onChangeOrder("memberRole")}>
                  <div className='order'>
                    <span>Role</span>
                    {orderBy === "memberRole" && (
                      <i
                        className={`fa-solid fa-arrow-${
                          order === "asc" ? "down" : "up"
                        }`}
                      />
                    )}
                  </div>
                </th>
                <th width='160' onClick={() => onChangeOrder("status")}>
                  <div className='order'>
                    <span>Status</span>
                    {orderBy === "status" && (
                      <i
                        className={`fa-solid fa-arrow-${
                          order === "asc" ? "down" : "up"
                        }`}
                      />
                    )}
                  </div>
                </th>
                <th width='50'>
                  <div className='n'></div>
                </th>
              </tr>
            </thead>
            <tbody>
              {members.map((member, index) => (
                <tr
                  key={index}
                  className={`${
                    selectedMember?.id && selectedMember?.id === member.id
                      ? "selected"
                      : ""
                  }`}
                >
                  <td>
                    <div>{index + 1 + (currentPage - 1) * pageSize}</div>
                  </td>
                  <td>
                    {member.isEditing ? (
                      <div className='input-group edit-input'>
                        <Input
                          id={`memberName-${member.id}`}
                          type='text'
                          extraClasses={["no-border"]}
                          value={member.name}
                          onChange={(e) =>
                            onChangeMemberName(member.id, e.target.value)
                          }
                          onKeyDown={(e) => {
                            if (e.key === "Enter") {
                              saveMember(member.id);
                            }

                            if (e.key === "Escape") {
                              onChangeMemberName(
                                member.id,
                                member.originalName
                              );
                              onClickMemberName(member.id);
                            }
                          }}
                        />
                        <div className='edit-input-icon'>
                          <div>
                            <i
                              className='fa-solid fa-check'
                              onClick={() => saveMember(member.id)}
                            />
                          </div>
                          <div>
                            <i
                              className='fa-solid fa-xmark'
                              onClick={() => {
                                onChangeMemberName(
                                  member.id,
                                  member.originalName
                                );
                                onClickMemberName(member.id);
                              }}
                            />
                          </div>
                        </div>
                      </div>
                    ) : (
                      <div
                        className='edit-button height'
                        onClick={() => onClickMemberName(member.id)}
                      >
                        {member.name}
                      </div>
                    )}
                  </td>
                  <td>
                    <Button
                      onClick={() => {
                        setSelectedMemberRole(member);
                        setIsOpenMemberRoleModal(true);
                      }}
                    >
                      <>
                        {member.memberRoleTitle}({member.memberRoleAbbr})
                      </>
                    </Button>
                  </td>
                  <td>
                    <div className='height'>
                      <RadioButton
                        extraClasses={["space-r-2"]}
                        type='radio'
                        checked={member.isAlive}
                        onChange={() => {
                          onChangeMemberAlive(member.id, true);
                          saveMember(member.id);
                        }}
                      />
                      <label className='space-r-3' htmlFor='status'>
                        Enable
                      </label>
                      <RadioButton
                        type='radio'
                        extraClasses={["space-r-2"]}
                        checked={!member.isAlive}
                        onChange={() => {
                          onChangeMemberAlive(member.id, false);
                          saveMember(member.id);
                        }}
                      />
                      <label htmlFor='status'>Disable</label>
                    </div>
                  </td>
                  <td>
                    <div className='action'>
                      <i
                        className='fa-solid fa-trash'
                        onClick={() => onShowConfirmationModal(member.id)}
                      />
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>

      <MemberRoleModal
        memberRoles={memberRoles}
        selectedMemberRole={selectedMemberRole}
        isOpen={isOpenMemberRoleModal}
        onClose={() => {
          setSelectedMemberRole({});
          setIsOpenMemberRoleModal(false);
        }}
        onConfirm={changeMemberRole}
      />
    </>
  );
};

export default Table;
