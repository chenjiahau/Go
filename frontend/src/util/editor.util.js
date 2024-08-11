import Paragraph from "@editorjs/paragraph";
import Header from "@editorjs/header";
import List from "@editorjs/list";
import Link from "@editorjs/link";
import Delimiter from "@editorjs/delimiter";
import CheckList from "@editorjs/checklist";
import ImageTool from '@editorjs/image';
import { cloneDeep } from "lodash";

import apiHandler from "@/util/api.util";

export const getDefaultEditorData = () => {
  return cloneDeep({
    time: new Date().getTime(),
    blocks: [
      {
        type: "header",
        data: {
          text: "New Document Title",
          level: 1,
        },
      },
    ],
  });
};

export const getEditConfig = () => {
  const url = `${apiHandler.axios.defaults.baseURL}/record/upload-image`;
  const authentication = `Bearer ${apiHandler.token}`;

  return {
    paragraph: {
      class: Paragraph,
      inlineToolbar: true,
    },
    checkList: CheckList,
    list: List,
    header: Header,
    delimiter: Delimiter,
    link: Link,
    image: {
      class: ImageTool,
      config: {
        endpoints: {
          byFile: url,
        },
        field: 'image',
        types: 'image/*',
        uploader: {
          uploadByFile(file) {
            return new Promise((resolve, reject) => {
              const formData = new FormData();
              formData.append('image', file);
              fetch(url, {
                method: 'POST',
                body: formData,
                headers: {
                  "Authorization": authentication,
                },
              })
                .then((response) => response.json())
                .then((result) => {
                  resolve({
                    success: 1,
                    file: {
                      url: result.data.url,
                    },
                  });
                })
                .catch((error) => {
                  reject(error);
                });
            });
          },
        },
      }
    },
  };
}
export const editorConfig = {

};