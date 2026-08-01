import { SvgIconProps } from "@/types/svg-props";

export const ReplayIcon: React.FC<SvgIconProps> = ({
  color = "#10487A",
  height = "16",
  width = "16",
  ...props
}) => {
  return (
    <svg {...props} width={width} height={height} viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path d="M8.00008 3.99999V1.33333L4.66675 4.66666L8.00008 8V5.33333C10.2067 5.33333 12.0001 7.12666 12.0001 9.33333C12.0001 11.54 10.2067 13.3333 8.00008 13.3333C5.79341 13.3333 4.00008 11.54 4.00008 9.33333H2.66675C2.66675 12.28 5.05341 14.6667 8.00008 14.6667C10.9467 14.6667 13.3334 12.28 13.3334 9.33333C13.3334 6.38666 10.9467 3.99999 8.00008 3.99999Z" fill={color}/>
    </svg>
  );
};
