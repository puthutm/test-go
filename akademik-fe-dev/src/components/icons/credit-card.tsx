import { SvgIconProps } from "@/types/svg-props";

export const CreditCardIcon: React.FC<SvgIconProps> = ({
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
        d="M13.3333 2.6665H2.66659C1.92659 2.6665 1.33992 3.25984 1.33992 3.99984L1.33325 11.9998C1.33325 12.7398 1.92659 13.3332 2.66659 13.3332H13.3333C14.0733 13.3332 14.6666 12.7398 14.6666 11.9998V3.99984C14.6666 3.25984 14.0733 2.6665 13.3333 2.6665ZM13.3333 11.9998H2.66659V7.99984H13.3333V11.9998ZM13.3333 5.33317H2.66659V3.99984H13.3333V5.33317Z"
        fill={color}
      />
    </svg>
  );
};
