import { SvgIconProps } from "@/types/svg-props";

export const WarningIcon: React.FC<SvgIconProps> = ({
  color = "#A66900",
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
        d="M7.99987 1.83398C4.31987 1.83398 1.3332 4.82065 1.3332 8.50065C1.3332 12.1807 4.31987 15.1673 7.99987 15.1673C11.6799 15.1673 14.6665 12.1807 14.6665 8.50065C14.6665 4.82065 11.6799 1.83398 7.99987 1.83398ZM8.66654 11.834H7.3332V10.5007H8.66654V11.834ZM8.66654 9.16732H7.3332V5.16732H8.66654V9.16732Z"
        fill={color}
      />
    </svg>
  );
};
