"use client";

import React, { useState } from "react";
import Image from "next/image";
import Link from "next/link";
import { Dropdown, DropdownMenu, DropdownToggle, Form } from "reactstrap";

import FullScreenToggle from "./FullScrennToggle";
import ProfileDropdown from "@/components/layouts/header/ProfileDropdown";

import logoSm from "@/assets/images/logo-unsia-sm.png";
import logoDark from "@/assets/images/logo-unsia.png";
import logoLight from "@/assets/images/logo-unsia.png";
import logounsiaHome from "@/assets/images/logo-unsia-home.png";

const Header = ({ isHome }: { isHome?: boolean }) => {
  // Inside your component
  const sidebarVisibilitytype = "show";

  const [search, setSearch] = useState<boolean>(false);
  const toogleSearch = () => {
    setSearch(!search);
  };

  const toogleMenuBtn = () => {
    const windowSize = document.documentElement.clientWidth;
    const humberIcon = document.querySelector(".hamburger-icon") as HTMLElement;

    if (windowSize > 767) humberIcon.classList.toggle("open");

    //For collapse horizontal menu
    if (document.documentElement.getAttribute("data-layout") === "horizontal") {
      /* eslint-disable */
      document.body.classList.contains("menu")
        ? document.body.classList.remove("menu")
        : document.body.classList.add("menu");
    }

    //For collapse vertical and semibox menu
    if (
      sidebarVisibilitytype === "show" &&
      (document.documentElement.getAttribute("data-layout") === "vertical" ||
        document.documentElement.getAttribute("data-layout") === "semibox")
    ) {
      if (windowSize < 1025 && windowSize > 767) {
        /* eslint-disable */
        document.body.classList.remove("vertical-sidebar-enable");
        document.documentElement.getAttribute("data-sidebar-size") === "sm"
          ? document.documentElement.setAttribute("data-sidebar-size", "")
          : document.documentElement.setAttribute("data-sidebar-size", "sm");
      } else if (windowSize > 1025) {
        /* eslint-disable */
        document.body.classList.remove("vertical-sidebar-enable");
        document.documentElement.getAttribute("data-sidebar-size") === "lg"
          ? document.documentElement.setAttribute("data-sidebar-size", "sm")
          : document.documentElement.setAttribute("data-sidebar-size", "lg");
      } else if (windowSize <= 767) {
        document.body.classList.add("vertical-sidebar-enable");
        document.documentElement.setAttribute("data-sidebar-size", "lg");
      }
    }
  };

  return (
    <React.Fragment>
      <header
        id={isHome ? "" : "page-topbar"}
        style={{ backgroundColor: "white" }}
      >
        <div className="layout-width">
          <div className="navbar-header">
            <div className="d-flex">
              {isHome && (
                <Image src={logounsiaHome} alt="" height="32" width={"100"} />
              )}
              <div className="navbar-brand-box horizontal-logo">
                <Link href="/" className="logo logo-dark">
                  <span className="logo-sm">
                    <Image src={logoSm} alt="" height="22" />
                  </span>
                  <span className="logo-lg">
                    <Image src={logoDark} alt="" height="17" />
                  </span>
                </Link>

                <Link href="/" className="logo logo-light">
                  <span className="logo-sm">
                    <Image src={logoSm} alt="" height="22" />
                  </span>
                  <span className="logo-lg">
                    <Image src={logoLight} alt="" height="17" />
                  </span>
                </Link>
              </div>
              {!isHome && (
                <button
                  onClick={toogleMenuBtn}
                  type="button"
                  className="btn btn-sm px-3 fs-16 header-item vertical-menu-btn topnav-hamburger"
                  id="topnav-hamburger-icon"
                >
                  <span className="hamburger-icon">
                    <span></span>
                    <span></span>
                    <span></span>
                  </span>
                </button>
              )}

              {/* <SearchOption /> */}
            </div>

            <div className="d-flex align-items-center">
              <Dropdown
                isOpen={search}
                toggle={toogleSearch}
                className="d-md-none topbar-head-dropdown header-item"
              >
                <DropdownToggle
                  type="button"
                  tag="button"
                  className="btn btn-icon btn-topbar btn-ghost-secondary rounded-circle"
                >
                  <i className="bx bx-search fs-22"></i>
                </DropdownToggle>
                <DropdownMenu className="dropdown-menu-lg dropdown-menu-end p-0">
                  <Form className="p-3">
                    <div className="form-group m-0">
                      <div className="input-group">
                        <input
                          type="text"
                          className="form-control"
                          placeholder="Search ..."
                          aria-label="Recipient's username"
                        />
                        <button className="btn btn-primary" type="submit">
                          <i className="mdi mdi-magnify"></i>
                        </button>
                      </div>
                    </div>
                  </Form>
                </DropdownMenu>
              </Dropdown>
              <FullScreenToggle />
              {/* <NotificationDropdown /> */}
              <ProfileDropdown />
            </div>
          </div>
        </div>
      </header>
    </React.Fragment>
  );
};

export default Header;
