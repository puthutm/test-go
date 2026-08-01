import { SvgIconProps } from "@/types/svg-props";

export const ContentCopyIcon: React.FC<SvgIconProps> = ({
  color = "#10487A",
  height = "16",
  width = "14",
  ...props
}) => {
  return (
    <svg {...props} width={width} height={height} viewBox="0 0 14 16" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path d="M10.0003 0.666672H2.00033C1.26699 0.666672 0.666992 1.26667 0.666992 2.00001V11.3333H2.00033V2.00001H10.0003V0.666672ZM12.0003 3.33334H4.66699C3.93366 3.33334 3.33366 3.93334 3.33366 4.66667V14C3.33366 14.7333 3.93366 15.3333 4.66699 15.3333H12.0003C12.7337 15.3333 13.3337 14.7333 13.3337 14V4.66667C13.3337 3.93334 12.7337 3.33334 12.0003 3.33334ZM12.0003 14H4.66699V4.66667H12.0003V14Z" fill={color}/>
    </svg>
  );
};
