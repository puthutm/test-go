"use client";
import React, { useState } from "react";
import Image from "next/image";
import {
  Collapse,
  Navbar,
  NavbarToggler,
  NavbarBrand,
  Nav,
  NavItem,
  NavLink,
  DropdownToggle,
  Dropdown,
  DropdownMenu,
  DropdownItem,
} from "reactstrap";
import { signOut } from "next-auth/react";
import { usePathname } from "next/navigation";

import styles from "@/styles/navbar.module.css";

import logoUnsia from "@/assets/images/logo-unsia-with-text.svg";
import { useLogout } from "@/services/api/auth/logout";
import { useGetProfile } from "@/services/api/sso/profile/use-get-profile";
import { ImageComponent } from "./image";

const navData = [
  {
    title: "Beranda",
    href: "/home",
  },
  {
    title: "Mahasiswa",
    href: "/student",
  },
  {
    title: "Keuangan",
    href: "#",
  },
  {
    title: "Perkuliahan",
    href: "#a",
  },
];

export function NavbarStudent() {
  const [isOpen, setIsOpen] = useState(false);
  const [isProfileDropdown, setIsProfileDropdown] = useState(false);
  const [isProfileDropdownMobile, setIsProfileDropdownMobile] = useState(false);
  const pathname = usePathname();

  const { handleLogout } = useLogout();

  const { data: profile } = useGetProfile();

  const onLogout = async () => {
    try {
      await handleLogout();
      await signOut({
        callbackUrl: `${process.env.NEXT_PUBLIC_UI_SSO_URL}/home`,
      });
    } catch (e) {
      console.log(e);
      return;
    }
  };

  const toggleProfileDropdown = () => {
    setIsProfileDropdown(!isProfileDropdown);
  };

  const toggle = () => setIsOpen(!isOpen);
  return (
    <div
      className="border-bottom border-3 bg-light fixed-top "
      style={{ zIndex: 10, height: "70px" }}
    >
      <Navbar
        expand="lg"
        className="mx-auto container container-sm container-md py-1"
      >
        {/* logo */}
        <NavbarBrand href="/">
          {/* logo */}
          <div className="navbar-brand d-flex gap-2 align-items-center">
            <Image
              src={logoUnsia}
              alt="unsia"
              width={100}
              height={32}
              className="navbar-logo pe-2"
            />
          </div>
        </NavbarBrand>
        {/* center */}
        <div className="d-flex order-3 order-lg-3 justify-content-end d-lg-none ms-2">
          <NavbarToggler onClick={toggle} />
        </div>

        {/* collapse */}
        <Collapse
          isOpen={isOpen}
          navbar
          className=" flex-grow-1 order-4 order-lg-1 mt-3"
        >
          <Nav className="mx-auto gap-0 gap-lg-4 bg-light px-3 rounded" navbar>
            {navData.map((nav, index) => (
              <NavItem key={index}>
                <NavLink
                  href={nav.href}
                  className={styles.nav_item}
                  style={{
                    color: pathname.includes(nav.href) ? "#10487A" : "",
                    textDecoration: pathname.includes(nav.href)
                      ? "underline"
                      : "",
                    fontWeight: 500,
                  }}
                >
                  {nav.title}
                </NavLink>
              </NavItem>
            ))}
            <Dropdown
              isOpen={isProfileDropdownMobile}
              toggle={() =>
                setIsProfileDropdownMobile(!isProfileDropdownMobile)
              }
              className="order-2 d-block d-lg-none"
            >
              <DropdownToggle tag="button" type="button" className="btn px-0">
                <div className={styles.nav_profile}>
                  <ImageComponent
                    src={profile?.data?.avatar as string}
                    width={200}
                    height={200}
                    alt="profile"
                  />
                  <div className={styles.profile_wrapper}>
                    <p className={`${styles.profile_name} text-truncate`}>
                      {profile?.data?.name}
                    </p>
                    <p className={`${styles.profile_nim} text-truncate`}>
                      NIM heres
                    </p>
                  </div>
                </div>
              </DropdownToggle>
              <DropdownMenu className="dropdown-menu-end">
                <DropdownItem className="p-0" onClick={onLogout}>
                  <span className="dropdown-item">
                    <i className="ri-home-8-line fs-16 align-middle me-1"></i>
                    <span className="align-middle">Menu</span>
                  </span>
                </DropdownItem>
              </DropdownMenu>
            </Dropdown>
          </Nav>
        </Collapse>

        {/* rigth */}
        <Dropdown
          isOpen={isProfileDropdown}
          toggle={toggleProfileDropdown}
          className="ms-sm-3 order-2 d-none d-lg-block"
        >
          <DropdownToggle tag="button" type="button" className="btn">
            <div className={styles.nav_profile}>
              <ImageComponent
                src={profile?.data?.avatar as string}
                width={200}
                height={200}
                alt="profile"
              />
              <div className={styles.profile_wrapper}>
                <p className={`${styles.profile_name} text-truncate`}>
                  {profile?.data?.name}
                </p>
                <p className={`${styles.profile_nim} text-truncate`}>
                  NIM here
                </p>
              </div>
            </div>
          </DropdownToggle>
          <DropdownMenu className="dropdown-menu-end">
            <DropdownItem className="p-0" onClick={onLogout}>
              <span className="dropdown-item">
                <i className="ri-home-8-line fs-16 align-middle me-1"></i>
                <span className="align-middle">Menu</span>
              </span>
            </DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </Navbar>
    </div>
  );
}
