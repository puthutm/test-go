import Image from "next/image";

import unsiaTowerBg from "@/assets/images/unsia-tower.png";

export const BannerProfile = () => {
  return (
    <div className="profile-foreground position-relative mt-n4">
      <div className="profile-wid-bg">
        <Image src={unsiaTowerBg} alt="" className="profile-wid-img" />
      </div>
    </div>
  );
};
