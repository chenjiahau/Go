import "../module.scss";
import Logo from "@/assets/img/brand.png";

import { useState } from "react";
import { Link } from "react-router-dom";
import _ from "lodash";

// Const
import apiConfig from "@/const/config/api";

// Util
import apiHandler from "@/util/api.util";
import messageUtil from "@/util/message.util";

// Components
import Input from "@/components/Input";
import Button from "@/components/Button";

const ForgotPassword = ({
  successMessage,
  errorMessage,
  stageType,
  onChangeStage,
}) => {
  // State
  const [form, setForm] = useState({
    email: "",
  });

  // Method
  const changeForm = (key) => (e) => {
    let value = "";
    if (_.isObject(e)) {
      value = e.target.value;
    }

    setForm({ ...form, [key]: value });
  };

  const emptyForm = () => {
    setForm({
      email: "",
    });
  };

  const handleForgotPassword = async () => {
    if (!form.email) {
      messageUtil.showErrorMessage(errorMessage.fieldsNotFill);
      return;
    }

    const payload = {
      email: form.email,
    };

    // Call API
    try {
      await apiHandler.post(apiConfig.resource.CREATE_FORGOT_PASSWORD, payload);
      emptyForm();
      messageUtil.showSuccessMessage(successMessage.forgotPassword);
    } catch (error) {
      messageUtil.showErrorMessage(apiHandler.extractErrorMessage(error));
    }
  };

  return (
    <div className='login-section'>
      <div className='login-block'>
        <div className='login-block__logo'>
          <img src={Logo} alt='logo' />
        </div>
        <div className='header-title login-block__title'>
          <h2>Forgot Password</h2>
        </div>
        <div className='login-block__body'>
          <div className='input-group'>
            <label htmlFor='username'>E-mail</label>
            <Input
              id='email'
              type='text'
              name='email'
              placeholder='Your email'
              value={form.email}
              onChange={(e) => {
                changeForm("email")(e);
              }}
            />
          </div>
          <div className='space-b-4'></div>
          <div className='button-container'>
            <Button id='submitBtn' onClick={handleForgotPassword}>
              Submit
            </Button>
            <Button extraClasses={["cancel-button"]} onClick={emptyForm}>
              Reset
            </Button>
          </div>
          <div className='button-container button-container--right space-t-3'>
            <div>Do you have an account?</div>
            <Link
              className='link-button'
              onClick={onChangeStage(stageType.LOGIN)}
            >
              Sign in
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ForgotPassword;
