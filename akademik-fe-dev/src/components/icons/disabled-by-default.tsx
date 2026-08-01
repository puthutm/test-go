import { SvgIconProps } from "@/types/svg-props";

export const DisabledByDefaultIcon: React.FC<SvgIconProps> = ({
  color = "#F06548",
  height = "18",
  width = "18",
  ...props
}) => {
  return (
    <svg {...props} width={width} height={height} viewBox="0 0 18 18" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path d="M0.75 0.75V17.25H17.25V0.75H0.75ZM13.5833 12.2908L12.2908 13.5833L9 10.2925L5.70917 13.5833L4.41667 12.2908L7.7075 9L4.41667 5.70917L5.70917 4.41667L9 7.7075L12.2908 4.41667L13.5833 5.70917L10.2925 9L13.5833 12.2908Z" fill={color}/>
    </svg>
  );
};
