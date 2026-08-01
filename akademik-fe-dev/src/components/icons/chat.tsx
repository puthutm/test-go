import { SvgIconProps } from "@/types/svg-props";

export const ChatIcon: React.FC<SvgIconProps> = ({
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
        d="M2.66659 2.66732H13.3333V10.6673H3.44659L2.66659 11.4473V2.66732ZM2.66659 1.33398C1.93325 1.33398 1.33992 1.93398 1.33992 2.66732L1.33325 14.6673L3.99992 12.0007H13.3333C14.0666 12.0007 14.6666 11.4007 14.6666 10.6673V2.66732C14.6666 1.93398 14.0666 1.33398 13.3333 1.33398H2.66659ZM3.99992 8.00065H9.33325V9.33398H3.99992V8.00065ZM3.99992 6.00065H11.9999V7.33398H3.99992V6.00065ZM3.99992 4.00065H11.9999V5.33398H3.99992V4.00065Z"
        fill={color}
      />
    </svg>
  );
};
