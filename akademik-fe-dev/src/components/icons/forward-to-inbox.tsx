import { SvgIconProps } from "@/types/svg-props";

export const ForwardToInboxIcon: React.FC<SvgIconProps> = ({
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
        d="M13 1.66797H2.33333C1.6 1.66797 1 2.26797 1 3.0013V11.0013C1 11.7346 1.6 12.3346 2.33333 12.3346H8.33333V11.0013H2.33333V4.33464L7.66667 7.66797L13 4.33464V7.66797H14.3333V3.0013C14.3333 2.26797 13.7333 1.66797 13 1.66797ZM7.66667 6.33464L2.33333 3.0013H13L7.66667 6.33464ZM12.3333 9.0013L15 11.668L12.3333 14.3346V12.3346H9.66667V11.0013H12.3333V9.0013Z"
        fill={color}
      />
    </svg>
  );
};
