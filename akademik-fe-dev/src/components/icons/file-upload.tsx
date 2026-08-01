import { SvgIconProps } from "@/types/svg-props";

export const FileUploadIcon: React.FC<SvgIconProps> = ({
  color = "#495057",
  height = "17",
  width = "16",
  ...props
}) => {
  return (
    <svg
      width={width}
      height={height}
      viewBox="0 0 16 17"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      {...props}
    >
      <path
        d="M12.0003 10.5003V12.5003H4.00033V10.5003H2.66699V12.5003C2.66699 13.2337 3.26699 13.8337 4.00033 13.8337H12.0003C12.7337 13.8337 13.3337 13.2337 13.3337 12.5003V10.5003H12.0003ZM4.66699 6.50033L5.60699 7.44033L7.33366 5.72033V11.167H8.66699V5.72033L10.3937 7.44033L11.3337 6.50033L8.00033 3.16699L4.66699 6.50033Z"
        fill={color}
      />
    </svg>
  );
};
