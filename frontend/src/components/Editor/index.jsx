import { memo, useEffect, useRef } from "react";
import EditorJS from "@editorjs/editorjs";

import apiHandler from "@/util/api.util";
import { getEditConfig } from "@/util/editor.util";

const url = `${apiHandler.axios.defaults.baseURL}/record/upload-image`;
const token = apiHandler.token;

const Editor = ({ data, onChange, editorBlock }) => {
  const ref = useRef();

  useEffect(() => {
    if (!ref.current) {
      const editor = new EditorJS({
        minHeight: 10,
        holder: editorBlock,
        tools: getEditConfig(url, token),
        data: data,
        async onChange(api, event) {
          const data = await api.saver.save();
          onChange(data);
        },
      });
      ref.current = editor;
    }

    return () => {
      if (ref.current && ref.current.destroy) {
        ref.current.destroy();
      }
    };
  }, []); // Don't add onChange to the dependency array

  return <div id={editorBlock}></div>;
};

const MemoizedEditor = memo(Editor);
export default MemoizedEditor;
