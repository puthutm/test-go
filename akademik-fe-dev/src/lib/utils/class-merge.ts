import { type ClassValue, clsx } from "clsx";
import classnames from "classnames";

export const classMerge = (...inputs: ClassValue[]) => {
  return classnames(clsx(inputs));
};
