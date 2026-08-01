import { SvgIconProps } from "@/types/svg-props";

export const AddIcon: React.FC<SvgIconProps> = ({
  color = "#10487A",
  height = "16",
  width = "16",
  ...props
}) => {
  return (
    <svg {...props} width={width} height={height} viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path d="M12.6668 8.66671H8.66683V12.6667H7.3335V8.66671H3.3335V7.33337H7.3335V3.33337H8.66683V7.33337H12.6668V8.66671Z" fill={color}/>
    </svg>
  );
};
