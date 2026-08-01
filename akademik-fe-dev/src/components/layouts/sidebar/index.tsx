"use client";

import React, { useEffect, useCallback, useState } from "react";
import { usePathname } from "next/navigation";
import Link from "next/link";
import { Collapse, Container } from "reactstrap";
import Image from "next/image";
import { signOut } from "next-auth/react";

import { classMerge } from "@/lib/utils/class-merge";

import logoSm from "@/assets/images/logo-unsia-sm.png";
import logoDark from "@/assets/images/logo-unsia.png";
import logoLight from "@/assets/images/logo-unsia.png";
import { useGetSidebarMenu } from "@/services/api/sso/menu/use-get-sidebar-menu";
import { ImageComponent } from "@/components/ui/image";

const Sidebar = () => {
  const [isOpen, setIsOpen] = useState<string | undefined>("");
  const pathName = usePathname();

  const { data: dataSidebar, isLoading: isLoadingSidebar } =
    useGetSidebarMenu();

  if (dataSidebar?.status === 401) {
    signOut();
  }

  const resizeSidebarMenu = useCallback(() => {
    const windowSize = document.documentElement.clientWidth;
    const hamburgerIcon = document.querySelector(".hamburger-icon");

    if (windowSize >= 1025) {
      document.documentElement.setAttribute("data-sidebar-size", "lg");
      if (hamburgerIcon) {
        hamburgerIcon.classList.remove("open");
      }
    } else if (windowSize < 1025 && windowSize > 767) {
      document.documentElement.setAttribute("data-sidebar-size", "sm");
      if (hamburgerIcon) {
        hamburgerIcon.classList.add("open");
      }
    } else {
      document.documentElement.setAttribute("data-sidebar-size", "lg");
      if (hamburgerIcon) {
        hamburgerIcon.classList.add("open");
      }
    }
  }, []);

  const addEventListenerOnSmHoverMenu = () => {
    if (
      document.documentElement.getAttribute("data-sidebar-size") === "sm-hover"
    ) {
      document.documentElement.setAttribute(
        "data-sidebar-size",
        "sm-hover-active"
      );
    } else if (
      document.documentElement.getAttribute("data-sidebar-size") ===
      "sm-hover-active"
    ) {
      document.documentElement.setAttribute("data-sidebar-size", "sm-hover");
    } else {
      document.documentElement.setAttribute("data-sidebar-size", "sm-hover");
    }
  };

  useEffect(() => {
    window.addEventListener("resize", resizeSidebarMenu, true);
    return () => {
      window.removeEventListener("resize", resizeSidebarMenu);
    };
  }, [resizeSidebarMenu]);

  useEffect(() => {
    const findActivePath = (items: any[]): string | undefined => {
      for (const item of items) {
        if (item.path && pathName.startsWith(item.path)) {
          if (item.children) {
            const activeChild = item.children.find((child: any) =>
              pathName.startsWith(child.path)
            );
            if (activeChild) return item.path;
          }

          return item.path;
        }

        if (item.children) {
          const activeChildPath = findActivePath(item.children);
          if (activeChildPath) {
            return item.path;
          }
        }
      }
      return undefined;
    };

    const activePath = findActivePath(dataSidebar?.data || []);

    setIsOpen(activePath);
  }, [pathName, dataSidebar]);

  useEffect(() => {
    const verticalOverlay = document.getElementsByClassName("vertical-overlay");
    if (verticalOverlay && verticalOverlay[0]) {
      verticalOverlay[0].addEventListener("click", function () {
        document.body.classList.remove("vertical-sidebar-enable");
      });
    }
  });

  useEffect(() => {
    const htmlElement = document.documentElement;

    const updateNavbarMenu = () => {
      const sidebarSize = htmlElement.getAttribute("data-sidebar-size");

      const navbarMenu = document.querySelector(".navbar-menu");
      if (sidebarSize === "lg") {
        navbarMenu?.classList.add("overflow-y-auto");
      } else {
        navbarMenu?.classList.remove("overflow-y-auto");
      }
    };

    updateNavbarMenu();

    const observer = new MutationObserver(updateNavbarMenu);
    observer.observe(htmlElement, {
      attributes: true,
      attributeFilter: ["data-sidebar-size"],
    });

    return () => observer.disconnect();
  }, []);

  return (
    <React.Fragment>
      <div className="app-menu navbar-menu">
        <div className="navbar-brand-box">
          <Link href="/dashboard" className="logo logo-dark">
            <span className="logo-sm">
              <Image src={logoSm} alt="" height="22" />
            </span>
            <span className="logo-lg">
              <Image src={logoDark} alt="" height="17" />
            </span>
          </Link>

          <Link href="/dashboard" className="logo logo-light">
            <span className="logo-sm">
              <Image src={logoSm} alt="" height="22" />
            </span>
            <span className="logo-lg">
              <Image src={logoLight} alt="" height="30" width={100} />
            </span>
          </Link>
          <button
            onClick={addEventListenerOnSmHoverMenu}
            type="button"
            className="btn btn-sm p-0 fs-20 header-item float-end btn-vertical-sm-hover"
            id="vertical-hover"
          >
            <i className="ri-record-circle-line"></i>
          </button>
        </div>
        <Container fluid>
          <div id="two-column-menu"></div>
          <ul className="navbar-nav " id="navbar-nav">
            <li className="menu-title">
              <span data-key="t-menu">Menu</span>
            </li>
            {isLoadingSidebar ? (
              <div className="d-flex flex-column gap-3 placeholder-glow px-3">
                {Array.from({ length: 10 }).map((_, index) => (
                  <span
                    className="placeholder w-100 rounded"
                    style={{ height: "25px" }}
                    key={index}
                  ></span>
                ))}
              </div>
            ) : (
              dataSidebar?.data?.map((item) => (
                <React.Fragment key={item.id}>
                  {item.children ? (
                    <li className="nav-item">
                      <a
                        onClick={(e) => {
                          e.preventDefault();
                          setIsOpen(
                            isOpen === item.path ? undefined : item.path
                          );
                        }}
                        className={classMerge(
                          isOpen?.includes(item.path) && "active",
                          "nav-link menu-link"
                        )}
                        role="button"
                        data-bs-toggle="collapse"
                        aria-expanded={isOpen?.includes(item.path)}
                      >
                        <i>
                          <ImageComponent
                            src={item.icon}
                            alt={item.label}
                            width={18}
                            height={18}
                          />
                        </i>
                        <span data-key="t-apps">{item.label}</span>
                      </a>
                      <Collapse
                        className="menu-dropdown"
                        isOpen={isOpen?.includes(item.path)}
                        id="sidebarApps"
                      >
                        <ul className="nav nav-sm flex-column">
                          {item.children.map((child) => (
                            <li className="nav-item" key={child.id}>
                              <Link
                                href={child.path ? child.path : "/#"}
                                className={classMerge(
                                  pathName.includes(child.path) && "active",
                                  "nav-link"
                                )}
                              >
                                {child.label}
                              </Link>
                            </li>
                          ))}
                        </ul>
                      </Collapse>
                    </li>
                  ) : (
                    <li className="nav-item">
                      <Link
                        className={classMerge(
                          pathName.includes(item.path) && "active",
                          "nav-link menu-link"
                        )}
                        href={item.path}
                      >
                        <i>
                          <ImageComponent
                            src={item.icon}
                            alt={item.label}
                            width={18}
                            height={18}
                          />
                        </i>
                        <span>{item.label}</span>
                      </Link>
                    </li>
                  )}
                </React.Fragment>
              ))
            )}
          </ul>
        </Container>
      </div>
      <div className="vertical-overlay"></div>
    </React.Fragment>
  );
};

export default Sidebar;
