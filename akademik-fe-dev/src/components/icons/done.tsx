import { SvgIconProps } from "@/types/svg-props";

export const DoneIcon: React.FC<SvgIconProps> = ({
  color = "#0AB39C",
  height = "12",
  width = "16",
  ...props
}) => {
  return (
    <svg {...props} width={width} height={height} viewBox="0 0 16 12" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path d="M5.33268 9.24999L1.83268 5.74999L0.666016 6.91666L5.33268 11.5833L15.3327 1.58332L14.166 0.416656L5.33268 9.24999Z" fill={color}/>
    </svg>
  );
};
