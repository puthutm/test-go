import { SvgIconProps } from "@/types/svg-props";

export const UploadIcon: React.FC<SvgIconProps> = ({
  color = "#909090",
  height = "18",
  width = "18",
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
        d="M15.6225 12.9997V15.4997H5.62248V12.9997H3.95581V15.4997C3.95581 16.4163 4.70581 17.1663 5.62248 17.1663H15.6225C16.5391 17.1663 17.2891 16.4163 17.2891 15.4997V12.9997H15.6225ZM6.45581 7.99967L7.63081 9.17467L9.78914 7.02467V13.833H11.4558V7.02467L13.6141 9.17467L14.7891 7.99967L10.6225 3.83301L6.45581 7.99967Z"
        fill={color}
      />
    </svg>
  );
};
