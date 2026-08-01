import { SvgIconProps } from "@/types/svg-props";

export const VideoConferenceIcon: React.FC<SvgIconProps> = ({
  color = "#FFFFFF",
  height = "16",
  width = "16",
  ...props
}) => {
  return (
    <svg
      width={width}
      height={height}
      viewBox="0 0 16 16"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      {...props}
    >
      <path
        d="M10 5.83333V11.1667H3.33333V5.83333H10ZM10.6667 4.5H2.66667C2.3 4.5 2 4.8 2 5.16667V11.8333C2 12.2 2.3 12.5 2.66667 12.5H10.6667C11.0333 12.5 11.3333 12.2 11.3333 11.8333V9.5L14 12.1667V4.83333L11.3333 7.5V5.16667C11.3333 4.8 11.0333 4.5 10.6667 4.5Z"
        fill={color}
      />
    </svg>
  );
};
