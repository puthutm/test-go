import { SvgIconProps } from "@/types/svg-props";

export const ChevronRightIcon: React.FC<SvgIconProps> = ({
  color = "#909090",
  height = "16",
  width = "16",
  ...props
}) => {
  return (
    <svg
      width={width}
      height={height}
      viewBox="0 0 13 14"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      {...props}
    >
      <path
        d="M6.47003 4L5.53003 4.94L8.58336 8L5.53003 11.06L6.47003 12L10.47 8L6.47003 4Z"
        fill={color}
      />
    </svg>
  );
};
