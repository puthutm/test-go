interface FormErrorMessageProps {
  errors: any;
}

export const FormErrorMessage: React.FC<FormErrorMessageProps> = ({
  errors,
}) => {
  return (
    errors && (
      <div
        className="text-danger position-absolute fst-italic"
        style={{ fontSize: "10px" }}
      >
        {errors.message}
      </div>
    )
  );
};
