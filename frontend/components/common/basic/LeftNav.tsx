"use client";
import {
  SidebarHeader,
  SidebarContent,
  SidebarGroup,
  SidebarFooter,
  Sidebar,
  SidebarGroupLabel,
  SidebarGroupContent,
  SidebarMenu,
  SidebarMenuItem,
  SidebarMenuButton,
} from "@/components/ui/sidebar";
import React from "react";
import {
  Calendar,
  ChevronUp,
  Home,
  Inbox,
  Search,
  Settings,
  User2,
} from "lucide-react";
import Link from "next/link";
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
} from "@/components/ui/dropdown-menu";
import { User, useUser } from "@/utils/contexts/AuthContext";
import { fullNameConstructor } from "@/utils/helpers/string";
import { redirect } from "next/navigation";

const items = [
  {
    title: "Dashboard",
    url: "/main/dashboard",
    icon: Home,
  },
  {
    title: "Assests",
    url: "/main/assests",
    icon: Inbox,
  },
  {
    title: "Calendar",
    url: "/main/calendar",
    icon: Calendar,
  },
  {
    title: "Stratergy",
    url: "/main/stratergy",
    icon: Search,
  },
  {
    title: "Settings",
    url: "/main/settings",
    icon: Settings,
  },
];

const LeftNav = () => {
  return (
    <Sidebar>
      <SideBarHeaderMain />
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Application</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild>
                    <Link href={item.url}>
                      <item.icon />
                      <span>{item.title}</span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SideBarFooterMain />
    </Sidebar>
  );
};

const SideBarHeaderMain = () => {
  return (
    <SidebarHeader>
      <h1 className="text-3xl font-bold text-balance text-center">Financier</h1>
    </SidebarHeader>
  );
};

const SideBarFooterMain = () => {
  const { user, updateUser } = useUser();

  const signOut = () => {
    const nullUser: User = {
      email: "",
      firstname: "",
      lastname: "",
      user_id: null,
    };
    updateUser(nullUser);
    redirect("/auth/login");
  };

  return (
    <SidebarFooter>
      <SidebarMenu>
        <SidebarMenuItem>
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <SidebarMenuButton>
                <User2 /> {fullNameConstructor(user?.firstname, user?.lastname)}
                <ChevronUp className="ml-auto" />
              </SidebarMenuButton>
            </DropdownMenuTrigger>
            <DropdownMenuContent
              side="top"
              className="w-[--radix-popper-anchor-width]"
            >
              <DropdownMenuItem>
                <span>Account</span>
              </DropdownMenuItem>
              <DropdownMenuItem onClick={signOut}>
                <span>Sign out</span>
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </SidebarMenuItem>
      </SidebarMenu>
    </SidebarFooter>
  );
};

export default LeftNav;
