"use client";

import React, { useState } from "react";
import { signOut } from "next-auth/react";
import {
  Dropdown,
  DropdownItem,
  DropdownMenu,
  DropdownToggle,
} from "reactstrap";

import { useLogout } from "@/services/api/auth/logout";
import { ImageComponent } from "@/components/ui/image";
import { useGetProfile } from "@/services/api/sso/profile/use-get-profile";

const ProfileDropdown: React.FC = () => {
  //Dropdown Toggle
  const [isProfileDropdown, setIsProfileDropdown] = useState(false);

  const { data } = useGetProfile();

  const { handleLogout } = useLogout();

  const onLogout = async () => {
    try {
      await handleLogout();
      await signOut({
        callbackUrl: `${process.env.NEXT_PUBLIC_UI_SSO_URL}/home`,
      });
    } catch (e) {
      console.log(e);
      return e;
    }
  };

  const toggleProfileDropdown = () => {
    setIsProfileDropdown(!isProfileDropdown);
  };

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
              src={data?.data?.avatar as string}
              alt="Header Avatar"
              width={1000}
              height={1000}
            />
            <span className="text-start ms-xl-2">
              <span className="d-none d-xl-inline-block ms-1 fw-medium user-name-text">
                {data?.data.name}
              </span>
              <span className="d-none d-xl-block ms-1 fs-12 text-muted user-name-sub-text">
                @{data?.data.username}
              </span>
            </span>
          </span>
        </DropdownToggle>
        <DropdownMenu className="dropdown-menu-end">
          <DropdownItem className="p-0">
            <span className="dropdown-item" onClick={onLogout}>
              <i className="ri-home-8-line fs-16 align-middle me-1"></i>
              <span className="align-middle">Menu</span>
            </span>
          </DropdownItem>
          {/* <DropdownItem className="p-0">
            <span className="dropdown-item" onClick={onLogout}>
              <span className="align-middle" data-key="t-logout">
                <i className="ri-logout-box-line fs-16 align-middle me-1"></i>{" "}
                Logout
              </span>
            </span>
          </DropdownItem> */}
        </DropdownMenu>
      </Dropdown>
    </React.Fragment>
  );
};

export default ProfileDropdown;
