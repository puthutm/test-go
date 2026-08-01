"use client";

import React, { useState, useEffect } from "react";
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
  const [blobUrl, setBlobUrl] = useState<string | null>(null);

  useEffect(() => {
    if (image && typeof window !== "undefined") {
      try {
        const objectUrl = URL.createObjectURL(image as Blob);
        setBlobUrl(objectUrl);
        return () => {
          URL.revokeObjectURL(objectUrl);
        };
      } catch (err) {
        console.error("Error creating Blob URL:", err);
      }
    }
  }, [image]);

  const finalSrc = src
    ? isStaticOrUrl
      ? src
      : blobUrl || defaultUser
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
