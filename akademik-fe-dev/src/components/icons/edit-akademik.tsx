import { SvgIconProps } from "@/types/svg-props";

export const EditAkademikIcon: React.FC<SvgIconProps> = ({
  color = "#10487A",
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
        d="M12.6666 8.66634H8.66658V12.6663H7.33325V8.66634H3.33325V7.33301H7.33325V3.33301H8.66658V7.33301H12.6666V8.66634Z"
        fill={color}
      />
    </svg>
  );
};
