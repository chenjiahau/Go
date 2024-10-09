import Paragraph from "@editorjs/paragraph";
import Header from "@editorjs/header";
import List from "@editorjs/list";
import Link from "@editorjs/link";
import Checklist from "@editorjs/checklist";
import ImageTool from '@editorjs/image';
import Quote from '@editorjs/quote';
import Code from '@editorjs/code';
import Table from '@editorjs/table';
import Marker from '@editorjs/marker';
import Delimiter from '@editorjs/delimiter';
import { cloneDeep } from "lodash";

import apiHandler from "@/util/api.util";

export const getDefaultEditorData = () => {
  return cloneDeep({
    time: new Date().getTime(),
    blocks: [
      {
        type: "paragraph",
        data: {
          text: "Write something here...",
        },
      },
    ],
  });
};

export const getEditConfig = () => {
  const url = `${apiHandler.axios.defaults.baseURL}/auth/record/upload-image-v2`;
  const authentication = `Bearer ${apiHandler.token}`;

  return {
    header: {
      class: Header,
      inlineToolbar: ['link'],
    },
    paragraph: {
      class: Paragraph,
      inlineToolbar: true,
    },
    checklist: {
      class: Checklist,
      inlineToolbar: true,
      config: {
        defaultStyle: 'checkbox',
      },
    },
    list: {
      class: List,
      inlineToolbar: true,
      config: {
        defaultStyle: 'unordered',
      },
    },
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
    quote: {
      class: Quote,
      inlineToolbar: true,
      config: {
        quotePlaceholder: 'Enter a quote',
        captionPlaceholder: '',
      },
    },
    code: {
      class: Code,
      config: {
        placeholder: 'Enter your code here...',
      },
    },
    table: {
      class: Table,
      inlineToolbar: true,
      config: {
        rows: 2,
        cols: 3,
      },
    },
    marker: {
      class: Marker,
      shortcut: 'CMD+SHIFT+M',
    },
    delimiter: {
      class: Delimiter,
    },
  };
}
export const editorConfig = {

};