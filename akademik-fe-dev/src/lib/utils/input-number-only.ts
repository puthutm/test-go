export const handleInputNumberOnly = (
  e: React.ChangeEvent<HTMLInputElement>
) => {
  const stringValue = e.target.value.replace(/[^0-9]/g, "");
  const numberValue = stringValue ? Number(stringValue) : "";

  const fractionalValue =  e.target.value.replace(/[^0-9.]/g, "").replace(/(\..*)\./g, "")

  return { stringValue, numberValue ,fractionalValue};
};
