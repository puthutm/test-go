"use client";

import Image, { StaticImageData } from "next/image";

import { useGetFileStorage } from "@/services/api/sso/file-storage";
import user from "@/assets/images/logo-unsia-sm.png";

interface ImageComponentProps {
  src: string | StaticImageData;
  alt: string;
  width?: number;
  height?: number;
  className?: string;
}

export const ImageComponent = ({
  src,
  alt,
  height,
  width,
  className,
}: ImageComponentProps) => {
  const { data: image } = useGetFileStorage(src as string);

  const url = image ? URL.createObjectURL(image as Blob) : null;

  return (
    <Image
      src={(url as string) || user}
      alt={alt}
      width={width}
      height={height}
      className={className}
    />
  );
};
