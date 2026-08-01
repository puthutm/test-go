import { SvgIconProps } from "@/types/svg-props";

export const ScoreIcon: React.FC<SvgIconProps> = ({
  color = "#10487A",
  height = "13",
  width = "12",
  ...props
}) => {
  return (
    <svg {...props} width={width} height={height} viewBox="0 0 12 13" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path d="M10.6667 0.5H1.33333C0.6 0.5 0 1.1 0 1.83333V11.1667C0 11.9 0.6 12.5 1.33333 12.5H10.6667C11.4 12.5 12 11.9 12 11.1667V1.83333C12 1.1 11.4 0.5 10.6667 0.5ZM10.6667 11.1667H1.33333L4 8.5L6.66667 11.1667L10.6667 7.16667V11.1667ZM10.6667 5.5L6.66667 9.5L4 6.83333L1.33333 9.5V1.83333H10.6667V5.5ZM7 4.5V2.5H6V6.5H7V4.5ZM9.46667 6.5L8.13333 4.5L9.46667 2.5H8.33333L7 4.5L8.33333 6.5H9.46667ZM5.33333 5.5H3.66667V5H5.33333V2.5H2.66667V3.5H4.33333V4H2.66667V6.5H5.33333V5.5Z" fill={color}/>
    </svg>
  );
};
