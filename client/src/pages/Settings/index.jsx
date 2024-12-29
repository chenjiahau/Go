import "./module.css";

import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useDispatch } from "react-redux";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faEye, faEyeSlash } from "@fortawesome/free-solid-svg-icons";

import Breadcrumbs from "@/components/Breadcrumbs";
import FormGroup from "@/components/FormGroup";
import FormLabel from "@/components/FormLabel";
import InputBox from "@/components/InputBox";
import ButtonBox from "@/components/ButtonBox";
import CardBox from "@/components/CardBox";
import LoadingBox from "@/components/LoadingBox";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Slice
import { userActions } from "@/store/slices/user";

// Component
import ConfirmModal from "./ConfirmModal";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";
import { set } from "lodash";

const successMessage = {
  changePassword: "Password changed successfully, please login again.",
};

const errorMessage = {
  fieldsNotFill: "Please fill in all fields.",
  newAndConfirmNotMatch: "New password and confirm password do not match.",
  allTheSame: "All the passwords are the same.",
  changePassword: "Failed to change password.",
};

const Settings = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const linkList = [
    { to: "/", label: "Home" },
    { to: "/settings", label: "Settings" },
  ];

  // State
  const [currentPassword, setCurrentPassword] = useState("");
  const [isCurrentPasswordVisible, setIsCurrentPasswordVisible] =
    useState(false);
  const [newPassword, setNewPassword] = useState("");
  const [isNewPasswordVisible, setIsNewPasswordVisible] = useState(false);
  const [confirmPassword, setConfirmPassword] = useState("");
  const [isConfirmPasswordVisible, setIsConfirmPasswordVisible] =
    useState(false);
  const [openConfirmModal, setOpenConfirmModal] = useState(false);
  const [loading, setLoading] = useState(false);

  // Method
  const handleChangePassword = () => {
    if (!currentPassword || !newPassword || !confirmPassword) {
      messageUtil.showErrorMessage(errorMessage.fieldsNotFill);
      return;
    }

    if (newPassword !== confirmPassword) {
      messageUtil.showErrorMessage(errorMessage.newAndConfirmNotMatch);
      return;
    }

    if (currentPassword === newPassword && newPassword === confirmPassword) {
      messageUtil.showErrorMessage(errorMessage.allTheSame);
      return;
    }

    handleOpenConfirmModal();
  };

  const handleResetPassword = () => {
    setCurrentPassword("");
    setNewPassword("");
    setConfirmPassword("");
  };

  const handleOpenConfirmModal = () => {
    setOpenConfirmModal(true);
  };

  const handleCloseConfirmModal = () => {
    setOpenConfirmModal(false);
  };

  const handleLogout = async () => {
    // Change password
    try {
      setLoading(true);

      const payload = {
        originalPassword: currentPassword,
        newPassword,
        confirmPassword,
      };

      await apiHandler.post(
        apiConfig.resource.SETTINGS_CHANGE_PASSWORD,
        payload
      );

      messageUtil.showSuccessMessage(commonMessage.success);
    } catch (error) {
      messageUtil.showErrorMessage(errorMessage.changePassword);
      setLoading(false);
      return;
    }

    // Logout if success
    try {
      await apiHandler.get("/auth/sign-out");
      messageUtil.showSuccessMessage(successMessage.changePassword);
      dispatch(userActions.cleanUser());
      localStorage.removeItem("user");
      navigate(routerConfig.routes.LOGIN);
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    } finally {
      setTimeout(() => {
        setLoading(false);
      }, 1000);
    }
  };

  return (
    <>
      <Breadcrumbs linkList={linkList} />
      <div className='settings-container'>
        <CardBox title='Password'>
          <FormGroup>
            <FormLabel forName='password'>Password</FormLabel>
            <InputBox
              type={isCurrentPasswordVisible ? "text" : "password"}
              id='password'
              name='password'
              placeholder='Password'
              value={currentPassword}
              onChange={(value) => setCurrentPassword(value)}
            >
              {
                <FontAwesomeIcon
                  icon={isCurrentPasswordVisible ? faEyeSlash : faEye}
                  onClick={() =>
                    setIsCurrentPasswordVisible(!isCurrentPasswordVisible)
                  }
                />
              }
            </InputBox>
          </FormGroup>
          <FormGroup>
            <FormLabel forName='new-password'>New Password</FormLabel>
            <InputBox
              type={isNewPasswordVisible ? "text" : "password"}
              id='new-password'
              name='new-password'
              placeholder='New Password'
              value={newPassword}
              onChange={(value) => setNewPassword(value)}
            >
              {
                <FontAwesomeIcon
                  icon={isNewPasswordVisible ? faEyeSlash : faEye}
                  onClick={() => setIsNewPasswordVisible(!isNewPasswordVisible)}
                />
              }
            </InputBox>
          </FormGroup>
          <FormGroup>
            <FormLabel forName='confirm-password'>Confirm Password</FormLabel>
            <InputBox
              type={isConfirmPasswordVisible ? "text" : "password"}
              id='confirm-password'
              name='confirm-password'
              placeholder='Confirm Password'
              value={confirmPassword}
              onChange={(value) => setConfirmPassword(value)}
            >
              {
                <FontAwesomeIcon
                  icon={isConfirmPasswordVisible ? faEyeSlash : faEye}
                  onClick={() =>
                    setIsConfirmPasswordVisible(!isConfirmPasswordVisible)
                  }
                />
              }
            </InputBox>
          </FormGroup>
          <div className='footer'>
            <ButtonBox
              onClick={handleChangePassword}
              isSave={true}
              extraClasses={["primary-shadow"]}
            >
              Save
            </ButtonBox>
            <ButtonBox
              onClick={handleResetPassword}
              isClose={true}
              extraClasses={["cancel-shadow"]}
            >
              Cancel
            </ButtonBox>
          </div>
        </CardBox>
      </div>

      <ConfirmModal
        openModal={openConfirmModal}
        onClose={handleCloseConfirmModal}
        onSubmit={handleLogout}
      />

      <LoadingBox visible={loading} />
    </>
  );
};

export default Settings;
