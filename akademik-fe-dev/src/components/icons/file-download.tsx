import { SvgIconProps } from "@/types/svg-props";

export const FileDownloadIcon: React.FC<SvgIconProps> = ({
  color = "#10487A",
  height = "20",
  width = "20",
  ...props
}) => {
  return (
    <svg
      width={width}
      height={height}
      viewBox="0 0 20 20"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      {...props}
    >
      <path
        d="M15 12.4987V14.9987H4.99998V12.4987H3.33331V14.9987C3.33331 15.9154 4.08331 16.6654 4.99998 16.6654H15C15.9166 16.6654 16.6666 15.9154 16.6666 14.9987V12.4987H15ZM14.1666 9.16536L12.9916 7.99036L10.8333 10.1404V3.33203H9.16665V10.1404L7.00831 7.99036L5.83331 9.16536L9.99998 13.332L14.1666 9.16536Z"
        fill={color}
      />
    </svg>
  );
};
