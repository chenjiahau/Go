import { useState, useEffect } from "react";
import PropTypes from "prop-types";
import { orderBy } from "lodash";

import ModalBox from "@/components/ModalBox";
import FormGroup from "@/components/FormGroup";
import FormLabel from "@/components/FormLabel";
import DropdownBox from "@/components/DropdownBox";

// Const
import apiConfig from "@/const/config/api";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const ColorModal = ({
  colorCategories,
  colors,
  openModal,
  selectedTag,
  onClose,
  onSubmit,
}) => {
  // State
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
    const firstCategoryOption = categoryOptions[0];
    handleCategoryChange(firstCategoryOption);
    onClose();
  };

  const handleSubmit = async () => {
    try {
      await apiHandler.put(
        apiConfig.resource.EDIT_TAG.replace(":id", selectedTag.id),
        {
          name: selectedTag.name,
          colorId: colorOptions.find((color) => color.active).value,
        }
      );

      handleCategoryChange(categoryOptions[0]);
      messageUtil.showSuccessMessage(commonMessage.success);
      onSubmit();
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  // Side effect
  useEffect(() => {
    if (!colorCategories.length || !colors.length || !selectedTag) {
      return;
    }

    const updatedCategoryOptions = colorCategories.map((colorCategory) => {
      return {
        id: colorCategory.id,
        value: colorCategory.id,
        label: colorCategory.name,
        active: colorCategory.id === selectedTag.colorCategoryId,
      };
    });
    orderBy(updatedCategoryOptions, "value", "asc");
    setCategoryOptions(updatedCategoryOptions);

    const selectedCategoryOptions = updatedCategoryOptions.find(
      (category) => category.active
    );

    let updatedColorOptions = colors
      .filter((color) => color.categoryId === selectedCategoryOptions.value)
      .map((color) => {
        return {
          value: color.id,
          label: color.name,
          hashCode: color.hexCode,
          active: color.id === selectedTag.colorId,
        };
      });

    setColorOptions(updatedColorOptions);
  }, [
    colorCategories,
    colorCategories.length,
    colors,
    colors.length,
    selectedTag,
  ]);

  if (!openModal) {
    return null;
  }

  return (
    <ModalBox
      title='Edit Color'
      onClose={() => handleClose()}
      onSubmit={() => handleSubmit()}
    >
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

ColorModal.propTypes = {
  openModal: PropTypes.bool,
  selectedTag: PropTypes.object,
  onClose: PropTypes.func,
  onSubmit: PropTypes.func,
};

export default ColorModal;
