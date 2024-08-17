import Editor from "@/components/Editor";
import { random } from "lodash";

const ContentViewer = ({ content }) => {
  const randomKey = random(0, 100000);

  return (
    <Editor
      reload={false}
      readOnly={true}
      data={JSON.parse(content)}
      editorBlock={`editorjs-container-${randomKey}`}
    />
  );
};

export default ContentViewer;
