import { memo, useEffect, useRef } from "react";
import EditorJS from "@editorjs/editorjs";

import apiHandler from "@/util/api.util";
import { getEditConfig } from "@/util/editor.util";

const url = `${apiHandler.axios.defaults.baseURL}/record/upload-image`;
const token = apiHandler.token;

const Editor = ({ data, onChange, editorBlock, readOnly = false }) => {
  const ref = useRef();

  useEffect(() => {
    if (!ref.current) {
      const editor = new EditorJS({
        readOnly,
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

  const style = {};
  if (readOnly) {
    style["padding"] = "0 0";
    style["border"] = "0";
  } else {
    style["height"] = "60vh";
    style["overflowY"] = "auto";
  }

  return <div id={editorBlock} style={{ ...style }}></div>;
};

const MemoizedEditor = memo(Editor);
export default MemoizedEditor;
