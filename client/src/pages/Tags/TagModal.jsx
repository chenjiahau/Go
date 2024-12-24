import { useState, useEffect } from "react";
import PropTypes from "prop-types";

import ModalBox from "@/components/ModalBox";
import FormGroup from "@/components/FormGroup";
import FormLabel from "@/components/FormLabel";
import InputBox from "@/components/InputBox";
import DropdownBox from "@/components/DropdownBox";
import { orderBy } from "lodash";

// Const
import apiConfig from "@/const/config/api";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const errorMessage = {
  name: "Tag is required.",
  duplicated: "Tag is duplicated.",
};

const TagModal = ({
  colorCategories,
  colors,
  openModal,
  onClose,
  onSubmit,
}) => {
  const [name, setName] = useState("");
  const [categoryOptions, setCategoryOptions] = useState([]);
  const [colorOptions, setColorOptions] = useState([]);

  // Method
  const handleCategoryChange = (categoryOption) => {
    const updatedCategoryOptions = categoryOptions.map((category) => {
      return {
        ...category,
        active: category.value === categoryOption.value,
      };
    });

    setCategoryOptions(updatedCategoryOptions);

    const selectedCategoryOptions = updatedCategoryOptions.find(
      (category) => category.active
    );

    let updatedColorOptions = colors
      .filter((color) => color.categoryId === selectedCategoryOptions.value)
      .map((color, index) => {
        return {
          value: color.id,
          label: color.name,
          hashCode: color.hexCode,
          active: index === 0,
        };
      });

    setColorOptions(updatedColorOptions);
  };

  const handleColorChange = (colorOption) => {
    const updatedColorOptions = colorOptions.map((color) => {
      return {
        ...color,
        active: color.value === colorOption.value,
      };
    });

    setColorOptions(updatedColorOptions);
  };

  const handleClose = () => {
    setName("");
    const firstCategoryOption = categoryOptions[0];
    handleCategoryChange(firstCategoryOption);
    onClose();
  };

  const handleSubmit = async () => {
    if (name === "") {
      messageUtil.showErrorMessage(errorMessage.name);
      return;
    }

    try {
      await apiHandler.post(apiConfig.resource.ADD_TAG, {
        name,
        colorId: colorOptions.find((color) => color.active).value,
      });

      setName("");
      handleCategoryChange(categoryOptions[0]);
      messageUtil.showSuccessMessage(commonMessage.success);
      onSubmit();
    } catch (error) {
      if (error.response.data.code === 3402) {
        messageUtil.showErrorMessage(errorMessage.duplicated);
        return;
      }
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  // Side effect
  useEffect(() => {
    if (!colorCategories.length || !colors.length) {
      return;
    }

    const updatedCategoryOptions = colorCategories.map(
      (colorCategory, index) => {
        return {
          id: colorCategory.id,
          value: colorCategory.id,
          label: colorCategory.name,
          active: index === 0,
        };
      }
    );
    orderBy(updatedCategoryOptions, "value", "asc");
    setCategoryOptions(updatedCategoryOptions);

    const selectedCategoryOptions = updatedCategoryOptions.find(
      (category) => category.active
    );

    let updatedColorOptions = colors
      .filter((color) => color.categoryId === selectedCategoryOptions.value)
      .map((color, index) => {
        return {
          value: color.id,
          label: color.name,
          hashCode: color.hexCode,
          active: index === 0,
        };
      });

    setColorOptions(updatedColorOptions);
  }, [colorCategories, colors]);

  if (!openModal) {
    return null;
  }

  return (
    <ModalBox
      title='Add Tag'
      onClose={() => handleClose()}
      onSubmit={() => handleSubmit()}
    >
      <FormGroup>
        <FormLabel forName='search'>Tag</FormLabel>
        <InputBox
          type='text'
          id='name'
          name='name'
          placeholder='Enter tag name'
          value={name}
          onChange={(value) => setName(value)}
        />
      </FormGroup>
      <FormGroup>
        <FormLabel forName='category'>Category</FormLabel>
        <DropdownBox
          options={categoryOptions}
          onClick={(option) => handleCategoryChange(option)}
        />
      </FormGroup>
      <FormGroup>
        <FormLabel forName='color'>Color</FormLabel>
        <DropdownBox
          isColor={true}
          options={colorOptions}
          onClick={(option) => handleColorChange(option)}
        />
      </FormGroup>
    </ModalBox>
  );
};

TagModal.propTypes = {
  colorCategories: PropTypes.array.isRequired,
  openModal: PropTypes.bool.isRequired,
  onClose: PropTypes.func.isRequired,
  onSubmit: PropTypes.func.isRequired,
};

export default TagModal;
