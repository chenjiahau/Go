import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useDispatch } from "react-redux";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Slice
import { userActions } from "@/store/slices/user";

// Component
import Input from "@/components/Input";
import Button from "@/components/Button";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

// Message
const successMessage = {
  changePassword: "Password changed successfully, please login again.",
};

const errorMessage = {
  fieldsNotFill: "Please fill in all fields.",
  newAndConfirmNotMatch: "New password and confirm password do not match.",
  allTheSame: "All the passwords are the same.",
  changePassword: "Failed to change password.",
};

const ChangePassword = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();

  // State
  const [originalPassword, setOriginalPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [isTextTypeOP, setIsTextTypeOP] = useState(true);
  const [isTextTypeNP, setIsTextTypeNP] = useState(true);
  const [isTextTypeCP, setIsTextTypeCP] = useState(true);

  // Method
  const save = async () => {
    if (!originalPassword || !newPassword || !confirmPassword) {
      messageUtil.showErrorMessage(errorMessage.fieldsNotFill);
      return;
    }

    if (newPassword !== confirmPassword) {
      messageUtil.showErrorMessage(errorMessage.newAndConfirmNotMatch);
      return;
    }

    if (originalPassword === newPassword && newPassword === confirmPassword) {
      messageUtil.showErrorMessage(errorMessage.allTheSame);
      return;
    }

    // Change password
    try {
      const payload = {
        originalPassword,
        newPassword,
        confirmPassword,
      };
      await apiHandler.post(
        apiConfig.resource.SETTINGS_CHANGE_PASSWORD,
        payload
      );
      messageUtil.showSuccessMessage(commonMessage.success);
      reset();
    } catch (error) {
      messageUtil.showErrorMessage(errorMessage.changePassword);
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
    }
  };

  const reset = () => {
    setOriginalPassword("");
    setNewPassword("");
    setConfirmPassword("");
  };

  return (
    <div className='section'>
      <div className='input-title'>Change Password</div>
      <div className='space-t-2'></div>
      <div className='input-group'>
        <Input
          id='original-password'
          type={isTextTypeOP ? "password" : "text"}
          name='originalPassword'
          autoComplete='off'
          placeholder='Original password'
          value={originalPassword}
          onChange={(e) => setOriginalPassword(e.target.value)}
          required
        >
          {isTextTypeOP ? (
            <i
              className='fa-regular fa-eye'
              title='show password'
              onClick={() => setIsTextTypeOP(!isTextTypeOP)}
            />
          ) : (
            <i
              className='fa-regular fa-eye-slash'
              title='hide password'
              onClick={() => setIsTextTypeOP(!isTextTypeOP)}
            />
          )}
        </Input>
      </div>
      <div className='space-t-2'></div>
      <div className='input-group'>
        <Input
          id='new-password'
          type={isTextTypeNP ? "password" : "text"}
          name='newPassword'
          autoComplete='off'
          placeholder='New password'
          value={newPassword}
          onChange={(e) => setNewPassword(e.target.value)}
          required
        >
          {isTextTypeNP ? (
            <i
              className='fa-regular fa-eye'
              title='show password'
              onClick={() => setIsTextTypeNP(!isTextTypeNP)}
            />
          ) : (
            <i
              className='fa-regular fa-eye-slash'
              title='hide password'
              onClick={() => setIsTextTypeNP(!isTextTypeNP)}
            />
          )}
        </Input>
      </div>
      <div className='space-t-2'></div>
      <div className='input-group'>
        <Input
          id='confirm-password'
          type={isTextTypeCP ? "password" : "text"}
          name='confirmPassword'
          autoComplete='off'
          placeholder='Confirm password'
          value={confirmPassword}
          onChange={(e) => setConfirmPassword(e.target.value)}
          required
        >
          {isTextTypeCP ? (
            <i
              className='fa-regular fa-eye'
              title='show password'
              onClick={() => setIsTextTypeCP(!isTextTypeCP)}
            />
          ) : (
            <i
              className='fa-regular fa-eye-slash'
              title='hide password'
              onClick={() => setIsTextTypeCP(!isTextTypeCP)}
            />
          )}
        </Input>
      </div>
      {/* Handler */}
      <div className='button-container'>
        <Button id='submitBtn' onClick={save}>
          Save
        </Button>
        <Button extraClasses={["cancel-button"]} onClick={() => reset()}>
          Reset
        </Button>
      </div>
    </div>
  );
};

export default ChangePassword;
