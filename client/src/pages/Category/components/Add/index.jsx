import { useState } from "react";

// Const
import apiConfig from "@/const/config/api";

// Component
import Input from "@/components/Input";
import Button from "@/components/Button";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const errorMessage = {
  subcategory: "Subcategory is required.",
  duplicated: "Subcategory is duplicated.",
};

const Add = (props) => {
  const { onInitialization, category } = props;

  // State
  const [subcategory, setSubcategory] = useState("");

  // Method
  const addSubcategory = async () => {
    if (subcategory === "") {
      messageUtil.showErrorMessage(errorMessage.subcategory);
      return;
    }

    try {
      await apiHandler.post(
        apiConfig.resource.ADD_SUBCATEGORY.replace(":id", category.id),
        {
          name: subcategory,
        }
      );
      setSubcategory("");
      messageUtil.showSuccessMessage(commonMessage.success);
      onInitialization();
    } catch (error) {
      if (error.response.data.code === 4402) {
        messageUtil.showErrorMessage(errorMessage.duplicated);
        return;
      }
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  return (
    <div className='section'>
      <div className='input-group'>
        <Input
          id='category'
          type='text'
          name='subcategory'
          autoComplete='off'
          placeholder='New subcategory'
          value={subcategory}
          onChange={(e) => setSubcategory(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              addSubcategory();
            }
          }}
          required
        />
      </div>
      <div className='space-t-2'></div>
      <div className='text-right'>
        <Button onClick={addSubcategory}>
          <i className='fa-solid fa-plus'></i>
          <span className='space-l-1'>Add</span>
        </Button>
      </div>
    </div>
  );
};

export default Add;
