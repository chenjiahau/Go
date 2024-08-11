import { useState, useEffect } from "react";
import { useNavigate, Link } from "react-router-dom";

// Const
import routerConfig from "@/const/config/router";

// Component

const Documents = () => {
  const navigate = useNavigate();

  // State

  // Method
  const handleInitialization = async () => {};

  // Side effect
  useEffect(() => {
    handleInitialization();
  }, []);

  return (
    <>
      <div className='breadcrumb-container'>
        <Link to={routerConfig.routes.DOCUMENTS} className='breadcrumb--item'>
          <span className='breadcrumb--item--inner'>
            <span className='breadcrumb--item-title'>Documents</span>
          </span>
        </Link>
      </div>

      <div className='floating-button'>
        <i
          className='fa-solid fa-plus'
          onClick={() => navigate(routerConfig.routes.ADD_DOCUMENT)}
        />
      </div>
    </>
  );
};

export default Documents;
