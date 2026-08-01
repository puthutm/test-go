import { SvgIconProps } from "@/types/svg-props";

export const SchoolIcon: React.FC<SvgIconProps> = ({
  color = "#495057",
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
        d="M8.00008 2L0.666748 6L3.33341 7.45333V11.4533L8.00008 14L12.6667 11.4533V7.45333L14.0001 6.72667V11.3333H15.3334V6L8.00008 2ZM12.5467 6L8.00008 8.48L3.45341 6L8.00008 3.52L12.5467 6ZM11.3334 10.66L8.00008 12.48L4.66675 10.66V8.18L8.00008 10L11.3334 8.18V10.66Z"
        fill={color}
      />
    </svg>
  );
};
