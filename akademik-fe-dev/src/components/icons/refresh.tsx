import { SvgIconProps } from "@/types/svg-props";

export const RefreshIcon: React.FC<SvgIconProps> = ({
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
        d="M7.99984 4.00001V1.33334L4.6665 4.66668L7.99984 8.00001V5.33334C10.2065 5.33334 11.9998 7.12668 11.9998 9.33334C11.9998 11.54 10.2065 13.3333 7.99984 13.3333C5.79317 13.3333 3.99984 11.54 3.99984 9.33334H2.6665C2.6665 12.28 5.05317 14.6667 7.99984 14.6667C10.9465 14.6667 13.3332 12.28 13.3332 9.33334C13.3332 6.38668 10.9465 4.00001 7.99984 4.00001Z"
        fill={color}
      />
    </svg>
  );
};
