export const createSearchParams = (paramName: string, value: string) => {
  const params = new URLSearchParams(paramName.toString());
  params.set(paramName, value);
  return params.toString();
};
