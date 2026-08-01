import { SvgIconProps } from "@/types/svg-props";

export const ModeEditIcon: React.FC<SvgIconProps> = ({
  color = "#0AB39C",
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
        d="M1.99927 14.0008H4.49927L11.8726 6.62751L9.3726 4.12751L1.99927 11.5008V14.0008ZM3.3326 12.0542L9.3726 6.01418L9.98593 6.62751L3.94593 12.6675H3.3326V12.0542Z"
        fill={color}
      />
      <path
        d="M12.2459 2.19418C11.9859 1.93418 11.5659 1.93418 11.3059 2.19418L10.0859 3.41418L12.5859 5.91418L13.8059 4.69418C14.0659 4.43418 14.0659 4.01418 13.8059 3.75418L12.2459 2.19418Z"
        fill={color}
      />
    </svg>
  );
};
