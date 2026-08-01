import { SvgIconProps } from "@/types/svg-props";

export const FolderIcon: React.FC<SvgIconProps> = ({
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
        d="M6.11325 3.99984L7.44659 5.33317H13.3333V11.9998H2.66659V3.99984H6.11325ZM6.66659 2.6665H2.66659C1.93325 2.6665 1.33992 3.2665 1.33992 3.99984L1.33325 11.9998C1.33325 12.7332 1.93325 13.3332 2.66659 13.3332H13.3333C14.0666 13.3332 14.6666 12.7332 14.6666 11.9998V5.33317C14.6666 4.59984 14.0666 3.99984 13.3333 3.99984H7.99992L6.66659 2.6665Z"
        fill={color}
      />
    </svg>
  );
};
