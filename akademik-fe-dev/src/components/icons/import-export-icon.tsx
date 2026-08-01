import { SvgIconProps } from "@/types/svg-props";

export const ImportExportIcon: React.FC<SvgIconProps> = ({
  color = "#10487A",
  height = "20",
  width = "20",
  ...props
}) => {
  return (
    <svg
      width={width}
      height={height}
      viewBox="0 0 20 20"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      {...props}
    >
      <path
        d="M7.50008 2.5L4.16675 5.825H6.66675V11.6667H8.33341V5.825H10.8334L7.50008 2.5ZM13.3334 14.175V8.33333H11.6667V14.175H9.16675L12.5001 17.5L15.8334 14.175H13.3334Z"
        fill={color}
      />
    </svg>
  );
};
