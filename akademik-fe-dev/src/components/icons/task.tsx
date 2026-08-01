import { SvgIconProps } from "@/types/svg-props";

export const TaskIcon: React.FC<SvgIconProps> = ({
  color = "#495057",
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
        d="M9.33332 1.33301H3.99999C3.26666 1.33301 2.67332 1.93301 2.67332 2.66634L2.66666 13.333C2.66666 14.0663 3.25999 14.6663 3.99332 14.6663H12C12.7333 14.6663 13.3333 14.0663 13.3333 13.333V5.33301L9.33332 1.33301ZM12 13.333H3.99999V2.66634H8.66666V5.99967H12V13.333ZM5.87999 8.69967L4.93332 9.63967L7.29332 11.9997L11.0667 8.22634L10.1267 7.28634L7.29999 10.113L5.87999 8.69967Z"
        fill={color}
      />
    </svg>
  );
};
