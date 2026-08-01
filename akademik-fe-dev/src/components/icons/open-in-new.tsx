import { SvgIconProps } from "@/types/svg-props";

export const OpenInNewIcon: React.FC<SvgIconProps> = ({
  color = "white",
  height = "16",
  width = "17",
  ...props
}) => {
  return (
    <svg
      width={width}
      height={height}
      viewBox="0 0 17 16"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      {...props}
    >
      <path
        d="M13.1667 12.6667H3.83333V3.33333H8.5V2H3.83333C3.09333 2 2.5 2.6 2.5 3.33333V12.6667C2.5 13.4 3.09333 14 3.83333 14H13.1667C13.9 14 14.5 13.4 14.5 12.6667V8H13.1667V12.6667ZM9.83333 2V3.33333H12.2267L5.67333 9.88667L6.61333 10.8267L13.1667 4.27333V6.66667H14.5V2H9.83333Z"
        fill={color}
      />
    </svg>
  );
};
