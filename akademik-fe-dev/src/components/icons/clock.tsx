import { SvgIconProps } from "@/types/svg-props";

export const ClockIcon: React.FC<SvgIconProps> = ({
  color = "white",
  height = "17",
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
        xmlns="http://www.w3.org/2000/svg"
        d="M7.99331 1.83398C4.31331 1.83398 1.33331 4.82065 1.33331 8.50065C1.33331 12.1807 4.31331 15.1673 7.99331 15.1673C11.68 15.1673 14.6666 12.1807 14.6666 8.50065C14.6666 4.82065 11.68 1.83398 7.99331 1.83398ZM10.1933 11.6407L7.33331 8.77398V5.16732H8.66665V8.22732L11.14 10.7007L10.1933 11.6407Z"
        fill={color}
      />
    </svg>
  );
};
