import "./module.css";

import React, { Fragment, useState, useEffect, useCallback } from "react";
import PropTypes from "prop-types";

import ModalBox from "@/components/ModalBox";
import FormGroup from "@/components/FormGroup";
import FormLabel from "@/components/FormLabel";
import DropdownBox from "@/components/DropdownBox";
import InputBox from "@/components/InputBox";
import TagBox from "@/components/TagBox";
import EditorJS from "@/components/Editor";

// Const
import apiConfig from "@/const/config/api";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";
import { getDefaultEditorData } from "@/util/editor.util";

const errorMessage = {
  title: "Title is required.",
  author: "Author is required.",
  category: "Category or sub subcategory required.",
  duplicated: "Document title is duplicated.",
};

const DocumentModal = ({ openModal, selectedDocument, onClose, onSubmit }) => {
  // State
  const [title, setTitle] = useState("");
  const [authorOptions, setAuthorOptions] = useState([]);
  const [categoryOptions, setCategoryOptions] = useState([]);
  const [subcategoryOptions, setSubcategoryOptions] = useState([]);
  const [tagOptions, setTagOptions] = useState([]);
  const [selectedTags, setSelectedTags] = useState([]);
  const [memberOptions, setMemberOptions] = useState([]);
  const [selectedMembers, setSelectedMembers] = useState([]);
  const [content, setContent] = useState(getDefaultEditorData());
  const [hiddenInfo, setHiddenInfo] = useState(false);
  const [viewMode, setViewMode] = useState(false);
  const [reloadEditor, setReloadEditor] = useState(false);

  // Method
  const handleInitialization = useCallback(async () => {
    let response = null;

    try {
      setReloadEditor(true);

      // Roles
      response = await apiHandler.get(apiConfig.resource.MEMBER_ROLES);
      const roles = response?.data?.data?.memberRoles || [];

      // Title
      if (selectedDocument) {
        setTitle(selectedDocument.name);
      }

      // Author options
      response = await apiHandler.get(apiConfig.resource.MEMBERS);
      let members = response?.data?.data || [];
      members = members.sort((a, b) => a.name.localeCompare(b.name));

      const updatedAuthorOptions = members.map((member) => {
        const role = roles.find((role) => role.id === member.memberRoleId);

        return {
          value: member.id,
          label: `[${role.abbr}] ${member.name}`,
          active: false,
        };
      });

      if (selectedDocument) {
        updatedAuthorOptions.forEach((option) => {
          option.active = false;
          if (option.value === selectedDocument.postMember.id) {
            option.active = true;
          }
        });
      }

      setAuthorOptions(updatedAuthorOptions);

      // Category options
      response = await apiHandler.get(apiConfig.resource.CATEGORIES);
      let categories = response?.data?.data || [];

      categories = categories.sort((a, b) => a.name.localeCompare(b.name));
      if (categories.length > 0) {
        const categoryOptions = categories.map((category, index) => ({
          value: category.id,
          label: category.name,
          active: index === 0,
        }));

        if (selectedDocument) {
          categoryOptions.forEach((option) => {
            option.active = false;
            if (option.value === selectedDocument.category.id) {
              option.active = true;
            }
          });
        }

        setCategoryOptions(categoryOptions);

        const selectedCategory = categoryOptions.find(
          (option) => option.active
        );
        response = await apiHandler.get(
          apiConfig.resource.SUBCATEGORIES.replace(
            ":id",
            selectedCategory.value
          )
        );
        let subcategories = response?.data?.data || [];
        subcategories = subcategories.sort((a, b) =>
          a.name.localeCompare(b.name)
        );

        if (subcategories.length > 0) {
          const subcategoryOptions = subcategories.map(
            (subcategory, index) => ({
              value: subcategory.id,
              label: subcategory.name,
              active: index === 0,
            })
          );

          if (selectedDocument) {
            subcategoryOptions.forEach((option) => {
              option.active = false;
              if (option.value === selectedDocument.subCategory.id) {
                option.active = true;
              }
            });
          }

          setSubcategoryOptions(subcategoryOptions);
        }
      }

      // Tag options
      response = await apiHandler.get(apiConfig.resource.TAGS);
      let tags = response?.data?.data || [];
      tags = tags.sort((a, b) => a.name.localeCompare(b.name));

      let updatedTagOptions = tags.map((tag) => ({
        value: tag.id,
        label: tag.name,
        hashCode: tag.colorHexCode,
        active: false,
      }));

      if (selectedDocument && selectedDocument.tags) {
        updatedTagOptions = updatedTagOptions.filter((option) => {
          const found = selectedDocument.tags.find(
            (tag) => tag.tagId === option.value
          );

          if (!found) {
            return {
              ...option,
            };
          }
        });

        const updatedSelectedTags = selectedDocument.tags.map((tag) => {
          return {
            value: tag.tagId,
            label: tag.tagName,
            hashCode: tag.colorHexCode,
          };
        });

        setSelectedTags(updatedSelectedTags);
      }

      setTagOptions(updatedTagOptions);

      // Member options
      let updatedMemberOptions = members.map((member) => {
        const role = roles.find((role) => role.id === member.memberRoleId);

        return {
          value: member.id,
          label: `[${role.abbr}] ${member.name}`,
          active: false,
        };
      });

      if (selectedDocument && selectedDocument.relationMembers) {
        updatedMemberOptions = updatedMemberOptions.filter((option) => {
          const found = selectedDocument.relationMembers.find(
            (member) => member.memberId === option.value
          );

          if (!found) {
            return {
              ...option,
            };
          }
        });

        const updatedSelectedMembers = selectedDocument.relationMembers.map(
          (member) => {
            return {
              value: member.memberId,
              label: member.name,
            };
          }
        );
        setSelectedMembers(updatedSelectedMembers || []);
      }

      setMemberOptions(updatedMemberOptions);

      // Content
      if (selectedDocument) {
        const content = JSON.parse(selectedDocument.content);
        setContent(content);
      }
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    } finally {
      setReloadEditor(false);
    }
  }, [selectedDocument]);

  const handleChangeAuthor = (author) => {
    const updatedAuthorOptions = authorOptions.map((option) => {
      if (option.value === author.value) {
        return {
          ...option,
          active: true,
        };
      }

      return {
        ...option,
        active: false,
      };
    });

    setAuthorOptions(updatedAuthorOptions);
  };

  const handleChangeCategory = async (category) => {
    const updatedCategoryOptions = categoryOptions.map((option) => {
      if (option.value === category.value) {
        return {
          ...option,
          active: true,
        };
      }

      return {
        ...option,
        active: false,
      };
    });

    setCategoryOptions(updatedCategoryOptions);
    const response = await apiHandler.get(
      apiConfig.resource.SUBCATEGORIES.replace(":id", category.value)
    );
    let subcategories = response?.data?.data || [];
    subcategories = subcategories.sort((a, b) => a.name.localeCompare(b.name));

    let subcategoryOptions = [];
    if (subcategories.length > 0) {
      subcategoryOptions = subcategories.map((subcategory, index) => ({
        value: subcategory.id,
        label: subcategory.name,
        active: index === 0,
      }));
    }

    setSubcategoryOptions(subcategoryOptions);
  };

  const handleChangeSubcategory = (subcategory) => {
    const updatedSubcategoryOptions = subcategoryOptions.map((option) => {
      if (option.value === subcategory.value) {
        return {
          ...option,
          active: true,
        };
      }

      return {
        ...option,
        active: false,
      };
    });

    setSubcategoryOptions(updatedSubcategoryOptions);
  };

  const handleClickTag = (tag) => {
    const updatedTagOptions = tagOptions.filter((option) => {
      if (option.value !== tag.value) {
        return {
          ...option,
        };
      }
    });

    setTagOptions(updatedTagOptions);

    const updatedTags = [...selectedTags, tag];
    setSelectedTags(updatedTags);
  };

  const handleDeleteTag = (tag) => {
    const updatedTags = selectedTags.filter(
      (selectedTag) => selectedTag.value !== tag.value
    );
    setSelectedTags(updatedTags);

    const updatedTagOptions = [...tagOptions, tag];
    setTagOptions(updatedTagOptions);
  };

  const handleClickMember = (member) => {
    const updatedMemberOptions = memberOptions.filter((option) => {
      if (option.value !== member.value) {
        return {
          ...option,
        };
      }
    });

    setMemberOptions(updatedMemberOptions);

    const updatedMembers = [...selectedMembers, member];
    setSelectedMembers(updatedMembers);
  };

  const handleDeleteMember = (member) => {
    const updatedMembers = selectedMembers.filter(
      (selectedMember) => selectedMember.value !== member.value
    );
    setSelectedMembers(updatedMembers);

    const updatedMemberOptions = [...memberOptions, member];
    setMemberOptions(updatedMemberOptions);
  };

  const changeViewMode = () => {
    setViewMode(!viewMode);
    setReloadEditor(true);

    setTimeout(() => {
      setReloadEditor(false);
    }, 100);
  };

  const handleClose = () => {
    setTitle("");
    setAuthorOptions([]);
    setCategoryOptions([]);
    setSubcategoryOptions([]);
    setTagOptions([]);
    setMemberOptions([]);
    setSelectedTags([]);
    setSelectedMembers([]);
    setContent(getDefaultEditorData());
    setHiddenInfo(false);
    setViewMode(false);
    setReloadEditor(true);

    // Force reload editor
    setTimeout(() => {
      setReloadEditor(false);
    }, 100);

    onClose();
  };

  const handleSubmit = async () => {
    if (!title) {
      messageUtil.showErrorMessage(errorMessage.title);
      return;
    }

    const author = authorOptions.find((option) => option.active);
    if (!author) {
      messageUtil.showErrorMessage(errorMessage.author);
      return;
    }

    const category = categoryOptions.find((option) => option.active);
    const subcategory = subcategoryOptions.find((option) => option.active);
    const tagIds = selectedTags.map((tag) => tag.value);
    const memberIds = selectedMembers.map((member) => member.value);
    const jsonContent = JSON.stringify(content);

    if (!category || !subcategory) {
      messageUtil.showErrorMessage(errorMessage.category);
      return;
    }

    const payload = {
      name: title,
      categoryId: category.value,
      subcategoryId: subcategory.value,
      postMemberId: author.value,
      relationMemberIds: memberIds,
      tagIds: tagIds,
      content: jsonContent,
    };

    try {
      if (selectedDocument) {
        await apiHandler.put(
          apiConfig.resource.EDIT_DOCUMENT.replace(":id", selectedDocument.id),
          payload
        );
      } else {
        await apiHandler.post(apiConfig.resource.ADD_DOCUMENT, payload);
      }

      messageUtil.showSuccessMessage(commonMessage.success);

      setTitle("");
      setAuthorOptions([]);
      setCategoryOptions([]);
      setSubcategoryOptions([]);
      setTagOptions([]);
      setMemberOptions([]);
      setSelectedTags([]);
      setSelectedMembers([]);
      setContent(getDefaultEditorData());
      setHiddenInfo(false);
      setViewMode(false);
      setReloadEditor(true);

      onSubmit();
    } catch (error) {
      if (error.response.data.code === 7402) {
        messageUtil.showErrorMessage(errorMessage.duplicated);
        return;
      }

      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  // Side effect
  useEffect(() => {
    if (openModal) {
      handleInitialization();
    }
  }, [handleInitialization, openModal]);

  if (!openModal) {
    return null;
  }

  let editorBlockId = "editorjs-container";
  if (selectedDocument) {
    editorBlockId = `editorjs-container-document-${selectedDocument.id}-edit`;
  }

  return (
    <ModalBox
      enableScroll={true}
      title={selectedDocument ? "Edit Document" : "Add Document"}
      customWidthClass='document-modal-body'
      onClose={() => handleClose()}
      onSubmit={() => handleSubmit()}
    >
      <FormGroup>
        <FormLabel forName='name'>Title</FormLabel>
        <InputBox
          type='text'
          id='ttile'
          name='title'
          placeholder='Title'
          value={title}
          onChange={(value) => setTitle(value)}
        />
      </FormGroup>
      <FormGroup>
        <div className='flex justify-end'>
          <TagBox
            tag={{
              label: hiddenInfo ? "Show Info" : "Hide Info",
              hashCode: hiddenInfo ? "#5bbcff" : "#054673",
            }}
            extraClasses={["!h-6"]}
            onClick={() => setHiddenInfo(!hiddenInfo)}
          />
        </div>
      </FormGroup>
      {!hiddenInfo && (
        <div className='mb-6'>
          <FormGroup>
            <FormLabel forName='name'>Author</FormLabel>
            <DropdownBox
              options={authorOptions}
              onClick={(author) => handleChangeAuthor(author)}
              zIndex={1}
            />
          </FormGroup>
          <FormGroup>
            <FormLabel forName='name'>Category</FormLabel>
            <DropdownBox
              options={categoryOptions}
              onClick={(category) => handleChangeCategory(category)}
              zIndex={2}
            />
          </FormGroup>
          <FormGroup>
            <FormLabel forName='name'>Subcategory</FormLabel>
            <DropdownBox
              options={subcategoryOptions}
              onClick={(subcategory) => handleChangeSubcategory(subcategory)}
              zIndex={3}
            />
          </FormGroup>
          <FormGroup>
            <FormLabel forName='name'>Tags</FormLabel>
            <DropdownBox
              options={tagOptions}
              isColor={true}
              onClick={(tag) => handleClickTag(tag)}
              zIndex={4}
            />
          </FormGroup>
          {selectedTags.length > 0 && (
            <div className='flex flex-wrap gap-2 mb-6'>
              {selectedTags.map((tag, index) => (
                <Fragment key={index}>
                  <TagBox
                    isDelBtn={true}
                    tag={tag}
                    onClick={(tag) => handleDeleteTag(tag)}
                  />
                </Fragment>
              ))}
            </div>
          )}
          <FormGroup extraClasses={["!mb-0"]}>
            <FormLabel forName='name'>Related Members</FormLabel>
            <DropdownBox
              options={memberOptions}
              onClick={(member) => handleClickMember(member)}
              zIndex={5}
            />
          </FormGroup>
          {selectedMembers.length > 0 && (
            <div className='flex flex-wrap gap-2 mt-6'>
              {selectedMembers.map((member, index) => (
                <Fragment key={index}>
                  <TagBox
                    isDelBtn={true}
                    tag={member}
                    onClick={handleDeleteMember}
                  />
                </Fragment>
              ))}
            </div>
          )}
        </div>
      )}
      <FormGroup>
        <div className='flex justify-end'>
          <TagBox
            tag={{
              label: viewMode ? "Edit Mode" : "View Mode",
              hashCode: viewMode ? "#5bbcff" : "#054673",
            }}
            extraClasses={["!h-6"]}
            onClick={() => changeViewMode()}
          />
        </div>
      </FormGroup>
      <FormGroup>
        <div className='editorjs-container primary-shadow'>
          {!reloadEditor && (
            <EditorJS
              reload={reloadEditor}
              readOnly={viewMode}
              data={content}
              onChange={(data) => setContent(data)}
              editorBlock={editorBlockId}
            />
          )}
        </div>
      </FormGroup>
    </ModalBox>
  );
};

DocumentModal.propTypes = {
  openModal: PropTypes.bool,
  members: PropTypes.array,
  categories: PropTypes.array,
  tags: PropTypes.array,
  selectedDocument: PropTypes.object,
  onClose: PropTypes.func,
  onSubmit: PropTypes.func,
};

export default DocumentModal;
