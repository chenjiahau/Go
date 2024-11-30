import { useEffect, useCallback } from "react";
import { useNavigate, useLocation } from "react-router-dom";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Util
import apiHandler from "@/util/api.util";
import messageUtil from "@/util/message.util";

// Components
import LoadingBox from "@/components/LoadingBox";

const errorMessage = {
  success: "Account is activated.",
  error: "Account is not activated.",
};

const ActiveAccount = () => {
  const navigate = useNavigate();
  const token = new URLSearchParams(useLocation().search).get("token");

  const handleInitialization = useCallback(async () => {
    if (!token) {
      navigate(routerConfig.routes.LOGIN);
    }

    try {
      await apiHandler.get(
        apiConfig.resource.ACTIVE_ACCOUNT.replace(":token", token)
      );
      messageUtil.showSuccessMessage(errorMessage.success);

      setTimeout(() => {
        navigate(routerConfig.routes.LOGIN);
      }, 3000);
    } catch (error) {
      messageUtil.showErrorMessage(errorMessage.error);

      setTimeout(() => {
        navigate(routerConfig.routes.LOGIN);
      }, 3000);
    }
  }, [navigate, token]);

  useEffect(() => {
    handleInitialization();
  }, [handleInitialization, token]);

  return <LoadingBox visible={true} />;
};

export default ActiveAccount;
