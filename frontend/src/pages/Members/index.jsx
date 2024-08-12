import { useState, useEffect, useCallback } from "react";
import { Link } from "react-router-dom";
import { cloneDeep } from "lodash";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Component
import { pageSizeDefinition } from "@/components/Pagination";
import ConfirmationModal from "@/components/ConfirmationModal";
import Add from "./components/Add";
import Table from "./components/Table";
import Page from "./components/Page";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const errorMessage = {
  delete: "Member is used or not found.",
};

const Category = () => {
  // State
  const [memberRoles, setMemberRoles] = useState([]);
  const [orderBy, setOrderBy] = useState("id");
  const [order, setOrder] = useState("asc");
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(pageSizeDefinition[1]);
  const [totalMemberCount, setTotalMemberCount] = useState(0);
  const [members, setMembers] = useState([]);
  const [selectedMember, setSelectedMember] = useState({});
  const [isOpenConfirmationModal, setIsOpenConfirmationModal] = useState(false);

  // Method
  const handleInitialization = useCallback(async () => {
    let response = null;

    response = await apiHandler.get(apiConfig.resource.MEMBER_ROLES);
    setMemberRoles(response.data.data.memberRoles);

    response = await apiHandler.get(apiConfig.resource.NUMBER_OF_MEMBERS);
    const totalMemberNumber = response.data.data.totalMemberNumber;
    setTotalMemberCount(totalMemberNumber);

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
        originalName: member.name,
        isEditing: false,
      });
    });

    if (selectedMember?.id) {
      const member = updatedMembers.find(
        (member) => member.id === selectedMember?.id
      );

      setSelectedMember(member);
    }

    setMembers(updatedMembers);
  }, [currentPage, order, orderBy, pageSize, selectedMember?.id]);

  const changeOrder = (newOrderBy) => {
    if (newOrderBy === orderBy) {
      setOrder(order === "asc" ? "desc" : "asc");
    } else {
      setOrderBy(newOrderBy);
      setOrder("asc");
    }
  };

  const clickMemberName = (id) => {
    const updatedMembers = members.map((member) => {
      if (member.id === id) {
        member.isEditing = !member.isEditing;
      }

      return member;
    });

    setMembers(updatedMembers);
  };

  const changeMemberName = (id, name) => {
    const updatedMembers = members.map((member) => {
      if (member.id === id) {
        member.name = name;
      }

      return member;
    });

    setMembers(updatedMembers);
  };

  const changeMemberAlive = async (id, alive) => {
    const updatedMembers = members.map((member) => {
      if (member.id === id) {
        member.isAlive = alive;
      }

      return member;
    });

    setMembers(updatedMembers);
  };

  const saveMember = async (id) => {
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
        memberRoleId: member.memberRoleId,
        name: member.name,
        isAlive: member.isAlive,
      };
      await apiHandler.put(apiURL, payload);
      messageUtil.showSuccessMessage(commonMessage.success);

      updatedMembers[memberIndex].originalName = member.name;
      updatedMembers[memberIndex].isEditing = false;
      setMembers(updatedMembers);
    } catch (error) {
      messageUtil.showErrorMessage(error.response.data.error.message);
    }
  };

  const deleteMember = async () => {
    try {
      const apiURL = apiConfig.resource.EDIT_MEMBER.replace(
        ":id",
        selectedMember.id
      );
      await apiHandler.delete(apiURL);
      messageUtil.showSuccessMessage(commonMessage.success);
      setIsOpenConfirmationModal(false);
      handleInitialization();
    } catch (error) {
      messageUtil.showErrorMessage(errorMessage.delete);
    }
  };

  const showConfirmationModal = (id) => {
    const member = members.find((member) => member.id === id);
    setSelectedMember(member);
    setIsOpenConfirmationModal(true);
  };

  // Side effect
  useEffect(() => {
    handleInitialization();
  }, [currentPage, handleInitialization, pageSize, orderBy, order]);

  return (
    <>
      <div className='breadcrumb-container'>
        <Link to={routerConfig.routes.MEMBERS} className='breadcrumb--item'>
          <span className='breadcrumb--item--inner'>
            <span className='breadcrumb--item-title'>Members</span>
          </span>
        </Link>
      </div>

      <Add onInitialization={handleInitialization} memberRoles={memberRoles} />

      <Table
        currentPage={currentPage}
        pageSize={pageSize}
        orderBy={orderBy}
        order={order}
        onChangeOrder={changeOrder}
        members={members}
        onClickMemberName={clickMemberName}
        onChangeMemberName={changeMemberName}
        onChangeMemberAlive={changeMemberAlive}
        selectedMember={selectedMember}
        onShowConfirmationModal={showConfirmationModal}
        saveMember={saveMember}
        onInitialization={handleInitialization}
        memberRoles={memberRoles}
      />

      <Page
        currentPage={currentPage}
        setCurrentPage={setCurrentPage}
        totalMemberCount={totalMemberCount}
        pageSize={pageSize}
        setPageSize={setPageSize}
      />

      <ConfirmationModal
        isOpen={isOpenConfirmationModal}
        onClose={() => {
          setSelectedMember(null);
          setIsOpenConfirmationModal(false);
        }}
        onConfirm={deleteMember}
      />
    </>
  );
};

export default Category;
