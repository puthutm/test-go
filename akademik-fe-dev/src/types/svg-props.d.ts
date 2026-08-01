import { SVGProps } from "react";

export interface SvgIconProps extends SVGProps<SVGSVGElement> {
  width?: string;
  height?: string;
  color?: string;
  className?: string;
}
