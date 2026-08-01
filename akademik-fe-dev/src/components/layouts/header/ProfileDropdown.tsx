"use client";

import React, { useState } from "react";
import { signOut, useSession } from "next-auth/react";
import {
  Dropdown,
  DropdownItem,
  DropdownMenu,
  DropdownToggle,
} from "reactstrap";

import { useLogout } from "@/services/api/auth/logout";
import { ImageComponent } from "@/components/ui/image";
import { useGetProfile } from "@/services/api/sso/profile/use-get-profile";
import defaultAvatar from "@/assets/images/users/avatar-1.jpg";

const ProfileDropdown: React.FC = () => {
  const [isProfileDropdown, setIsProfileDropdown] = useState(false);

  const { data: session } = useSession();
  const { data } = useGetProfile();

  const { handleLogout } = useLogout();

  const onLogout = async () => {
    try {
      await handleLogout().catch(() => {});
    } catch (e) {
      console.log(e);
    } finally {
      await signOut({
        callbackUrl: "/",
      });
    }
  };

  const toggleProfileDropdown = () => {
    setIsProfileDropdown(!isProfileDropdown);
  };

  const name = data?.data?.name || session?.user?.name || "User";
  const username = data?.data?.username || session?.user?.email?.split("@")[0] || "user";
  const avatar = data?.data?.avatar || defaultAvatar;

  return (
    <React.Fragment>
      <Dropdown
        style={{ backgroundColor: "white" }}
        isOpen={isProfileDropdown}
        toggle={toggleProfileDropdown}
        className="ms-sm-3 header-item topbar-user"
      >
        <DropdownToggle tag="button" type="button" className="btn">
          <span className="d-flex align-items-center">
            <ImageComponent
              className="rounded-circle header-profile-user"
              src={avatar}
              alt="Header Avatar"
              width={35}
              height={35}
            />
            <span className="text-start ms-xl-2">
              <span className="d-none d-xl-inline-block ms-1 fw-medium user-name-text">
                {name}
              </span>
              <span className="d-none d-xl-block ms-1 fs-12 text-muted user-name-sub-text">
                @{username}
              </span>
            </span>
          </span>
        </DropdownToggle>
        <DropdownMenu className="dropdown-menu-end">
          <DropdownItem className="p-0">
            <span
              className="dropdown-item d-flex align-items-center text-danger fw-semibold"
              onClick={onLogout}
              style={{ cursor: "pointer" }}
            >
              <i className="ri-logout-box-r-line fs-18 align-middle me-2"></i>
              <span className="align-middle">Logout / Keluar</span>
            </span>
          </DropdownItem>
        </DropdownMenu>
      </Dropdown>
    </React.Fragment>
  );
};

export default ProfileDropdown;
