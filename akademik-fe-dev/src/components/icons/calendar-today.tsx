import { SvgIconProps } from "@/types/svg-props";

export const CalendarTodayIcon: React.FC<SvgIconProps> = ({
  color = "#909090",
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
        d="M13.3334 2.49984H12.6667V1.1665H11.3334V2.49984H4.66671V1.1665H3.33337V2.49984H2.66671C1.93337 2.49984 1.33337 3.09984 1.33337 3.83317V14.4998C1.33337 15.2332 1.93337 15.8332 2.66671 15.8332H13.3334C14.0667 15.8332 14.6667 15.2332 14.6667 14.4998V3.83317C14.6667 3.09984 14.0667 2.49984 13.3334 2.49984ZM13.3334 14.4998H2.66671V7.1665H13.3334V14.4998ZM13.3334 5.83317H2.66671V3.83317H13.3334V5.83317Z"
        fill={color}
      />
    </svg>
  );
};
