import { SvgIconProps } from "@/types/svg-props";

export const InsertDriveFileIcon: React.FC<SvgIconProps> = ({
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
        d="M9.33341 1.33203H4.00008C3.26675 1.33203 2.67341 1.93203 2.67341 2.66536L2.66675 13.332C2.66675 14.0654 3.26008 14.6654 3.99341 14.6654H12.0001C12.7334 14.6654 13.3334 14.0654 13.3334 13.332V5.33203L9.33341 1.33203ZM4.00008 13.332V2.66536H8.66675V5.9987H12.0001V13.332H4.00008Z"
        fill={color}
      />
    </svg>
  );
};
