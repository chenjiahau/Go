import { useParams } from "react-router-dom";

const Category = () => {
  const { id } = useParams();

  console.log(id);

  return (
    <div>
      <h1>Category</h1>
    </div>
  );
};

export default Category;
