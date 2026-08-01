import { SvgIconProps } from "@/types/svg-props";

export const CheckBoxIcon: React.FC<SvgIconProps> = ({
  color = "#10487A",
  height = "18",
  width = "18",
  ...props
}) => {
  return (
    <svg
      {...props}
      width={width}
      height={height}
      viewBox="0 0 18 18"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M15.4167 0.75H2.58333C1.56583 0.75 0.75 1.575 0.75 2.58333V15.4167C0.75 16.425 1.56583 17.25 2.58333 17.25H15.4167C16.4342 17.25 17.25 16.425 17.25 15.4167V2.58333C17.25 1.575 16.4342 0.75 15.4167 0.75ZM7.16667 13.5833L2.58333 9L3.87583 7.7075L7.16667 10.9892L14.1242 4.03167L15.4167 5.33333L7.16667 13.5833Z"
        fill={color}
      />
    </svg>
  );
};
