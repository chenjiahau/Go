import "./module.scss";

import { useState, useEffect, useCallback } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import { orderBy } from "lodash";

// Const
import apiConfig from "@/const/config/api";
import routerConfig from "@/const/config/router";

// Component
import Editor from "@/components/Editor";
import Input from "@/components/Input";
import Button from "@/components/Button";
import Dropdown from "@/components/Dropdown";
import TitleButton from "@/components/TitleButton";

// Util
import apiHandler from "@/util/api.util";
import { getDefaultEditorData } from "@/util/editor.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const errorMessage = {
  title: "Title is required.",
  author: "Author is required.",
  category: "Category or sub subcategory required.",
  duplicated: "Document name is duplicated.",
};

const modeType = {
  edit: "edit",
  view: "view",
};

const EditDocument = () => {
  const navigate = useNavigate();
  const { id } = useParams();

  // State
  const [mode, setMode] = useState(modeType.edit);
  const [editorData, setEditorData] = useState(getDefaultEditorData());
  const [title, setTitle] = useState("");
  const [authors, setAuthors] = useState([]);
  const [categories, setCategories] = useState([]);
  const [subCategories, setSubCategories] = useState([]);
  const [tags, setTags] = useState([]);
  const [selectedTags, setSelectedTags] = useState([]);
  const [relatedMembers, setRelatedMembers] = useState([]);
  const [selectedMembers, setSelectedMembers] = useState([]);
  const [reload, setReload] = useState(true);

  // Method
  const init = useCallback(async () => {
    if (!id) {
      navigate(routerConfig.routes.DOCUMENTS);
      return;
    }

    let response = null;

    try {
      response = await apiHandler.get(
        apiConfig.resource.EDIT_DOCUMENT.replace(":id", id)
      );
      const document = response.data.data;

      // Title
      setTitle(document.name);

      // Author
      response = await apiHandler.get(apiConfig.resource.MEMBERS);
      let members = response?.data?.data || [];

      members = members.filter((member) => member.isAlive);
      members = orderBy(members, "name", "asc");

      const updatedAuthors = members.map((member) => ({
        ...member,
        selected: member.id === document.postMember.id,
      }));
      setAuthors(updatedAuthors);

      // Related members
      const updatedRelatedMembers = members.map((member) => ({
        ...member,
        selected: document.relationMembers?.find(
          (relationMember) => member.id === relationMember.memberId
        ),
      }));
      setSelectedMembers(
        updatedRelatedMembers.filter((member) => member.selected)
      );
      setRelatedMembers(
        updatedRelatedMembers.filter((member) => !member.selected)
      );

      // Category
      response = await apiHandler.get(apiConfig.resource.CATEGORIES);
      let categories = response?.data?.data || [];
      categories = categories.filter((category) => category.isAlive);
      categories = orderBy(categories, "name", "asc");

      let selectedCategory = null;
      const updatedCategories = categories.map((category) => {
        if (category.id === document.category.id) {
          selectedCategory = category;
        }

        return {
          ...category,
          selected: category.id === document.category.id,
        };
      });
      setCategories(updatedCategories);

      if (updatedCategories.length > 0) {
        response = await apiHandler.get(
          apiConfig.resource.SUBCATEGORIES.replace(":id", selectedCategory.id)
        );
        let subCategories = response?.data?.data || [];
        subCategories = orderBy(subCategories, "name", "asc");
        setSubCategories(
          subCategories.map((subCategory) => ({
            id: subCategory.id,
            name: subCategory.name,
            selected: subCategory.id === document.subCategory.id,
          }))
        );
      }

      // Tags
      response = await apiHandler.get(apiConfig.resource.TAGS);
      let tags = response?.data?.data || [];
      tags = orderBy(tags, "name", "asc");

      let updatedTags = tags.map((tag) => ({
        ...tag,
        bgcolor: tag.colorHexCode,
        id: tag.id,
        name: tag.name,
        selected: document.tags?.find(
          (selectedTag) => tag.id === selectedTag.tagId
        ),
      }));
      setSelectedTags(updatedTags.filter((tag) => tag.selected));
      setTags(updatedTags.filter((tag) => !tag.selected));

      // Content
      setEditorData(JSON.parse(document.content));
    } catch (error) {
      console.log(error);
      messageUtil.showErrorMessage(commonMessage.error);
    }
  }, [id, navigate]);

  const changeSubCategories = (category) => async () => {
    try {
      const response = await apiHandler.get(
        apiConfig.resource.SUBCATEGORIES.replace(":id", category.id)
      );
      let subCategories = response?.data?.data || [];
      subCategories = orderBy(subCategories, "name", "asc");
      setSubCategories(
        subCategories.map((subCategory, index) => ({
          id: subCategory.id,
          name: subCategory.name,
          selected: index === 0,
        }))
      );
    } catch (error) {
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  const changeMode = (mode) => {
    setMode(mode);
    setReload(true);

    setTimeout(() => {
      setReload(false);
    }, 1000);
  };

  const reset = () => {
    setReload(true);
    init();
  };

  const save = () => async () => {
    if (!title) {
      messageUtil.showErrorMessage(errorMessage.title);
      return;
    }

    const author = authors.find((author) => author.selected);
    if (!author) {
      messageUtil.showErrorMessage(errorMessage.author);
      return;
    }

    const category = categories.find((category) => category.selected);
    const subCategory = subCategories.find(
      (subCategory) => subCategory.selected
    );
    const tagIds = selectedTags.map((tag) => tag.id);
    const memberIds = selectedMembers.map((member) => member.id);
    const content = JSON.stringify(editorData);

    if (!category || !subCategory) {
      messageUtil.showErrorMessage(errorMessage.category);
      return;
    }

    const payload = {
      name: title,
      categoryId: category.id,
      subcategoryId: subCategory.id,
      postMemberId: author.id,
      relationMemberIds: memberIds,
      tagIds: tagIds,
      content: content,
    };

    try {
      await apiHandler.put(
        apiConfig.resource.EDIT_DOCUMENT.replace(":id", id),
        payload
      );
      messageUtil.showSuccessMessage(commonMessage.success);
      reset();
    } catch (error) {
      if (error.response.data.error.code === 409) {
        messageUtil.showErrorMessage(errorMessage.duplicated);
        return;
      }

      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  // Side effect
  useEffect(() => {
    init();
  }, [init]);

  useEffect(() => {
    setTimeout(() => {
      setReload(false);
    }, 1000);
  }, [reload]);

  return (
    <>
      <div className='breadcrumb-container'>
        <Link to={routerConfig.routes.DOCUMENTS} className='breadcrumb--item'>
          <span className='breadcrumb--item--inner'>
            <span className='breadcrumb--item-title'>Documents</span>
          </span>
        </Link>
        <Link
          to={routerConfig.routes.DOCUMENT.replace(":id", id)}
          className='breadcrumb--item'
        >
          <span className='breadcrumb--item--inner'>
            <span className='breadcrumb--item-title'>Document</span>
          </span>
        </Link>
        <Link onClick={reset} className='breadcrumb--item'>
          <span className='breadcrumb--item--inner'>
            <span className='breadcrumb--item-title'>Edit</span>
          </span>
        </Link>
      </div>

      <div className='section'>
        {/* Title */}
        <div className='input-title'>Title</div>
        <div className='space-t-2'></div>
        <div className='input-group'>
          <Input
            id='title'
            type='text'
            name='title'
            autoComplete='off'
            placeholder='New document title'
            minLength={1}
            maxLength={512}
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
          />
        </div>
        <div className='space-t-3'></div>

        {/* Author */}
        <div className='author-container'>
          <div className='input-title'>Author</div>
          <div className='input-title'>Category</div>
          <div className='input-title'></div>
          <Dropdown
            zIndex={10}
            list={authors}
            selectedItem={authors.find((member) => member.selected)}
            onChange={(item) => {
              setAuthors(
                authors.map((member) => ({
                  ...member,
                  selected: member.id === item.id,
                }))
              );
            }}
          />
          <Dropdown
            zIndex={9}
            list={categories}
            selectedItem={categories.find((category) => category.selected)}
            onChange={(item) => {
              setCategories(
                categories.map((category) => ({
                  ...category,
                  selected: category.id === item.id,
                }))
              );
              changeSubCategories(item)();
            }}
          />
          <Dropdown
            zIndex={8}
            list={subCategories}
            selectedItem={subCategories.find(
              (subCategory) => subCategory.selected
            )}
            onChange={(item) => {
              setSubCategories(
                subCategories.map((subCategory) => ({
                  ...subCategory,
                  selected: subCategory.id === item.id,
                }))
              );
            }}
          />
        </div>
        <div className='space-t-3'></div>

        {/* Tags */}
        <div className='input-title'>Tag</div>
        <div className='space-t-2'></div>
        {selectedTags.length > 0 && (
          <div className='tag-container'>
            {selectedTags.map((tag) => (
              <div
                key={tag.id}
                className='tag'
                style={{ backgroundColor: tag.colorHexCode }}
              >
                {tag.name}
                <i
                  className='fa-solid fa-x'
                  onClick={() => {
                    setSelectedTags(
                      selectedTags.filter((selectedTag) => {
                        if (selectedTag.id === tag.id) {
                          return;
                        }

                        return selectedTag;
                      })
                    );
                    setTags([...tags, tag]);
                  }}
                />
              </div>
            ))}
          </div>
        )}
        <Dropdown
          zIndex={7}
          hasBackground={true}
          list={tags}
          onChange={(item) => {
            setSelectedTags([...selectedTags, item]);
            setTags(
              tags.filter((tag) => {
                if (tag.id === item.id) {
                  return;
                }

                return tag;
              })
            );
          }}
        />
        <div className='space-t-3'></div>

        {/* Related members */}
        <div className='input-title'>Related Member</div>
        <div className='space-t-2'></div>
        {selectedMembers.length > 0 && (
          <div className='tag-container'>
            {selectedMembers.map((member) => (
              <div key={member.id} className='tag'>
                {member.name}
                <i
                  className='fa-solid fa-x'
                  onClick={() => {
                    setSelectedMembers(
                      selectedMembers.filter((selectedMember) => {
                        if (selectedMember.id === member.id) {
                          return;
                        }

                        return selectedMember;
                      })
                    );
                    setRelatedMembers([...relatedMembers, member]);
                  }}
                />
              </div>
            ))}
          </div>
        )}
        <Dropdown
          zIndex={6}
          list={relatedMembers}
          onChange={(item) => {
            setSelectedMembers([...selectedMembers, item]);
            setRelatedMembers(
              relatedMembers.filter((member) => {
                if (member.id === item.id) {
                  return;
                }

                return member;
              })
            );
          }}
        />
        <div className='space-t-4'></div>

        <div className='edit-container'>
          {mode === modeType.edit ? (
            <TitleButton
              onClick={() => changeMode(modeType.view)}
              title='View mode'
            />
          ) : (
            <TitleButton
              onClick={() => changeMode(modeType.edit)}
              title='Edit mode'
            />
          )}
        </div>
        <div className='space-t-1'></div>

        {/* Editor */}
        {!reload && (
          <div style={{ width: "100%" }}>
            <Editor
              reload={reload}
              readOnly={mode === modeType.view}
              data={editorData}
              onChange={setEditorData}
              editorBlock='editorjs-container'
            />
          </div>
        )}

        {/* Handler */}
        <div className='button-container'>
          <Button id='submitBtn' onClick={save()}>
            Save
          </Button>
          <Button extraClasses={["cancel-button"]} onClick={() => reset()}>
            Reset
          </Button>
        </div>
      </div>
    </>
  );
};

export default EditDocument;
