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
  const isStaticOrUrl =
    !src ||
    typeof src !== "string" ||
    src.startsWith("/") ||
    src.startsWith("http://") ||
    src.startsWith("https://") ||
    src.startsWith("data:");

  const { data: image } = useGetFileStorage(isStaticOrUrl ? "" : (src as string));

  const url = isStaticOrUrl
    ? (src as string)
    : image
    ? URL.createObjectURL(image as Blob)
    : null;

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
