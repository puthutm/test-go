"use client";

import React, { useState } from "react";
import { signOut } from "next-auth/react";
import {
  Dropdown,
  DropdownItem,
  DropdownMenu,
  DropdownToggle,
} from "reactstrap";

import { HomeIcon } from "../icons/home";
import { useLogout } from "@/services/api/auth/logout";

const ProfileDropdown: React.FC = () => {
  const [isProfileDropdown, setIsProfileDropdown] = useState(false);

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
          {/* <Profile /> */}
        </DropdownToggle>
        <DropdownMenu className="dropdown-menu-end">
          <DropdownItem className="p-0">
            <span className="dropdown-item" onClick={onLogout}>
              <i className="me-1">
                <HomeIcon />
              </i>
              <span className="align-middle">Menu</span>
            </span>
          </DropdownItem>
        </DropdownMenu>
      </Dropdown>
    </React.Fragment>
  );
};

export default ProfileDropdown;
