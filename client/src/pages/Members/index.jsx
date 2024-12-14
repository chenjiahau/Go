import { useState, useEffect, useCallback } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faPlus,
  faTrash,
  faFloppyDisk,
  faXmark,
} from "@fortawesome/free-solid-svg-icons";
import { cloneDeep } from "lodash";

import Breadcrumbs from "@/components/Breadcrumbs";
import Form from "@/components/Form";
import ElementGroup from "@/components/ElementGroup";
import InputBox from "@/components/InputBox";
import EditableTextBox from "@/components/EditableTextBox";
import RadioBox from "@/components/RadioBox";
import IconButton from "@/components/IconButton";
import ToolbarBox from "@/components/ToolbarBox";
import TableBox from "@/components/TableBox";
import PaginationBox, { numberOfRow } from "@/components/PaginationBox";
import TagBox from "@/components/TagBox";
import Spacer from "@/components/Spacer";

// Const
import apiConfig from "@/const/config/api";

// Component
import MemberModal from "./MemberModal";
import RoleModal from "./RoleModal";
import DeleteModal from "./DeleteModal";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const errorMessage = {
  member: "Member is required.",
  duplicated: "Member is duplicated.",
};

const defaultTableHeader = [
  {
    key: "index",
    label: "#",
    isSortable: false,
    isCenter: true,
    width: "50",
    sort: "",
  },
  {
    key: "member",
    label: "Member",
    isSortable: true,
    sort: "",
  },
  {
    key: "role",
    label: "Role",
    isSortable: true,
    isCenter: true,
    width: "220",
    sort: "",
  },
  {
    key: "status",
    label: "Status",
    isSortable: true,
    width: "100",
    sort: "",
  },
  {
    key: "action",
    label: "Action",
    isSortable: false,
    isCenter: true,
    width: "100",
    sort: "",
  },
];

const Members = () => {
  const linkList = [
    { to: "/", label: "Home" },
    { to: "/members", label: "Members" },
  ];

  // State
  const [orderBy, setOrderBy] = useState("id");
  const [order, setOrder] = useState("asc");
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(numberOfRow[1]);
  const [roles, setRoles] = useState([]);
  const [members, setMembers] = useState([]);
  const [tableHeaders, setTableHeader] = useState(
    cloneDeep(defaultTableHeader)
  );
  const [tableData, setTableData] = useState([]);
  const [totalDataCount, setTotalDataCount] = useState(0);
  const [openMemberModal, setOpenMemberModal] = useState(false);
  const [openRoleModal, setOpenRoleModal] = useState(false);
  const [openDeleteModal, setOpenDeleteModal] = useState(false);
  const [selectedMember, setSelectedMember] = useState(null);

  // Method
  const handleInitialization = useCallback(async () => {
    let response = null;

    response = await apiHandler.get(apiConfig.resource.MEMBER_ROLES);
    const updatedRoles = response.data.data.memberRoles;
    setRoles(updatedRoles);

    response = await apiHandler.get(apiConfig.resource.NUMBER_OF_MEMBERS);
    const totalMemberNumber = response.data.data.totalMemberNumber;
    setTotalDataCount(totalMemberNumber);

    // Check if the current page is greater than the total page number
    const totalPageNum = Math.ceil(totalMemberNumber / pageSize);
    let resetCurrentPage = false;
    if (totalPageNum < currentPage) {
      resetCurrentPage = true;
      setCurrentPage(1);
    }

    let queryString = "?orderBy=" + orderBy + "&order=" + order;
    response = await apiHandler.get(
      apiConfig.resource.MEMBERS_BY_PAGE.replace(
        ":page",
        resetCurrentPage ? 1 : currentPage
      ).replace(":size", pageSize),
      queryString
    );

    let updatedMembers = [];
    response.data.data.members?.forEach((member) => {
      updatedMembers.push({
        ...member,
        isEdit: false,
        originalName: member.name,
      });
    });

    setMembers(updatedMembers);
  }, [orderBy, order, currentPage, pageSize]);

  const handleChangeHeader = (newHeader, column, order) => {
    setTableHeader(newHeader);
    setOrderBy(column);
    setOrder(order);
  };

  const handleCloseModal = () => {
    setOpenMemberModal(false);
    setOpenRoleModal(false);
    setOpenDeleteModal(false);
    setSelectedMember(null);
  };

  const handleAddMember = async () => {
    setOpenMemberModal(false);
    setOpenRoleModal(false);
    setOpenDeleteModal(false);
    await handleInitialization();
  };

  const handleEditMember = useCallback(
    (id) => {
      const updatedMembers = members.map((member) => {
        if (member.id === id) {
          return {
            ...member,
            isEdit: !member.isEdit,
          };
        }

        return member;
      });

      setMembers(updatedMembers);
    },
    [members]
  );

  const handleChangeMember = useCallback(
    (id, value) => {
      const updatedMembers = members.map((member) => {
        if (member.id === id) {
          return {
            ...member,
            name: value,
          };
        }

        return member;
      });

      setMembers(updatedMembers);
    },
    [members]
  );

  const handleSaveMember = useCallback(
    async (id) => {
      const memberIndex = members.findIndex((member) => member.id === id);
      const updatedMembers = cloneDeep(members);

      if (updatedMembers[memberIndex].name === "") {
        return;
      }

      try {
        const member = members.find((member) => member.id === id);
        const apiURL = apiConfig.resource.EDIT_MEMBER.replace(":id", id);
        const payload = {
          ...member,
          name: member.name,
        };

        await apiHandler.put(apiURL, payload);
        messageUtil.showSuccessMessage(commonMessage.success);

        updatedMembers[memberIndex].originalName = member.name;
        updatedMembers[memberIndex].isEdit = false;
        setMembers(updatedMembers);
      } catch (error) {
        messageUtil.showErrorMessage(error.response.data.message);
      }
    },
    [members]
  );

  const handleCancelEdit = useCallback(
    (id) => {
      const updatedMembers = members.map((member) => {
        if (member.id === id) {
          return {
            ...member,
            isEdit: false,
            name: member.originalName,
          };
        }

        return member;
      });

      setMembers(updatedMembers);
    },
    [members]
  );

  const handleChangeStatus = useCallback(
    async (id, isAlive) => {
      const memberIndex = members.findIndex((member) => member.id === id);
      const updatedMembers = cloneDeep(members);

      try {
        const member = members.find((member) => member.id === id);
        const apiURL = apiConfig.resource.EDIT_MEMBER.replace(":id", id);
        const payload = {
          ...member,
          isAlive: isAlive,
        };

        await apiHandler.put(apiURL, payload);
        messageUtil.showSuccessMessage(commonMessage.success);

        updatedMembers[memberIndex].isAlive = isAlive;
        setMembers(updatedMembers);
      } catch (error) {
        messageUtil.showErrorMessage(error.response.data.message);
      }
    },
    [members]
  );

  const handleOpenRoleModal = (member) => {
    setOpenRoleModal(true);
    setSelectedMember(member);
  };

  const handleChangeRole = async () => {
    setOpenMemberModal(false);
    setOpenRoleModal(false);
    setOpenDeleteModal(false);
    await handleInitialization();
  };

  const handleOpenDeleteModal = (member) => {
    setOpenDeleteModal(true);
    setSelectedMember(member);
  };

  const handleDelete = async () => {
    try {
      const apiURL = apiConfig.resource.EDIT_MEMBER.replace(
        ":id",
        selectedMember.id
      );
      await apiHandler.delete(apiURL);
      messageUtil.showSuccessMessage(commonMessage.success);
      setOpenDeleteModal(false);
      handleInitialization();
    } catch (error) {
      messageUtil.showErrorMessage(errorMessage.delete);
    }
  };

  // Side effect
  useEffect(() => {
    handleInitialization();
  }, [handleInitialization, orderBy, order, currentPage, pageSize]);

  useEffect(() => {
    const updatedTableData = members.map((member, index) => {
      const role = roles.find((role) => role.id === member.memberRoleId);

      return {
        ref: cloneDeep(member),
        index: (currentPage - 1) * pageSize + index + 1,
        isEdit: false,
        member: member.isEdit ? (
          <InputBox
            type='text'
            id={`member${member.id}`}
            name={`member-${member.id}`}
            value={member.name}
            onChange={(value) => handleChangeMember(member.id, value)}
          >
            <FontAwesomeIcon
              icon={faFloppyDisk}
              onClick={() => handleSaveMember(member.id)}
            />
            <FontAwesomeIcon
              icon={faXmark}
              onClick={() => handleCancelEdit(member.id)}
            />
          </InputBox>
        ) : (
          <EditableTextBox
            text={member.name}
            onClick={() => {
              handleEditMember(member.id);
            }}
          />
        ),
        role: (
          <TagBox
            tag={{ label: role.title }}
            onClick={() => handleOpenRoleModal(member)}
          />
        ),
        status: (
          <ElementGroup>
            <RadioBox
              checked={member.isAlive}
              onChange={() => handleChangeStatus(member.id, true)}
            >
              Enable
            </RadioBox>
            <RadioBox
              checked={!member.isAlive}
              onChange={() => handleChangeStatus(member.id, false)}
            >
              Disable
            </RadioBox>
          </ElementGroup>
        ),
        action: (
          <div className='flex items-center gap-4'>
            <IconButton onClick={() => handleOpenDeleteModal(member)}>
              <FontAwesomeIcon icon={faTrash} />
            </IconButton>
          </div>
        ),
      };
    });

    setTableData(updatedTableData);
  }, [
    currentPage,
    handleCancelEdit,
    handleChangeMember,
    handleChangeStatus,
    handleEditMember,
    handleSaveMember,
    members,
    pageSize,
    roles,
  ]);

  return (
    <>
      <Breadcrumbs linkList={linkList} />
      <div className='custom-container primary-bg'>
        <Form>
          <ToolbarBox>
            <IconButton rounded={true} onClick={() => setOpenMemberModal(true)}>
              <FontAwesomeIcon icon={faPlus} />
              <div>Add</div>
            </IconButton>
          </ToolbarBox>

          <TableBox
            headers={tableHeaders}
            onChangeHeader={handleChangeHeader}
            data={tableData}
          />

          <Spacer />
          <PaginationBox
            currentPage={currentPage}
            totalCount={totalDataCount}
            pageSize={pageSize}
            setPageSize={setPageSize}
            onPageChange={(page) => {
              setCurrentPage(page);
            }}
          />
        </Form>
      </div>

      <MemberModal
        openModal={openMemberModal}
        roles={roles}
        onClose={handleCloseModal}
        onSubmit={handleAddMember}
      />

      <RoleModal
        openModal={openRoleModal}
        roles={roles}
        selectedMember={selectedMember}
        onClose={() => handleCloseModal()}
        onSubmit={() => handleChangeRole()}
      />

      <DeleteModal
        deleteMode={true}
        openModal={openDeleteModal}
        selectedMember={selectedMember}
        onClose={() => handleCloseModal()}
        onSubmit={() => handleDelete()}
      />
    </>
  );
};

export default Members;
