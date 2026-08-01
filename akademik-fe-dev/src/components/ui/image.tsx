"use client";

import Image, { StaticImageData } from "next/image";

import { useGetFileStorage } from "@/services/api/sso/file-storage";
import defaultUser from "@/assets/images/logo-unsia-sm.png";

interface ImageComponentProps {
  src?: string | StaticImageData | null;
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
  const isStaticOrUrl =
    !src ||
    typeof src !== "string" ||
    src.startsWith("http://") ||
    src.startsWith("https://") ||
    src.startsWith("data:");

  const { data: image } = useGetFileStorage(isStaticOrUrl ? "" : (src as string));

  const finalSrc = src
    ? isStaticOrUrl
      ? src
      : image
      ? URL.createObjectURL(image as Blob)
      : defaultUser
    : defaultUser;

  return (
    <Image
      src={finalSrc}
      alt={alt || "Image"}
      width={width || 35}
      height={height || 35}
      className={className}
    />
  );
};
