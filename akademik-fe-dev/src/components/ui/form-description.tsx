export const FormDescription = ({ message }: { message: string }) => {
  return (
    <div
      className="position-absolute fst-italic"
      style={{ fontSize: "10px", color: "#495057" }}
    >
      {message}
    </div>
  );
};
