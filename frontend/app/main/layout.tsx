import { cookies } from "next/headers";
import LeftNav from "@/components/common/basic/LeftNav";
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import React from "react";

const layout = async ({ children }: { children: React.ReactNode }) => {
  const cookieStore = await cookies();
  const defaultOpen = cookieStore.get("sidebar:state")?.value === "true";
  return (
    <SidebarProvider defaultOpen={defaultOpen}>
      <LeftNav />
      <div className="flex h-screen w-screen">
        {/* <SidebarTrigger /> */}
        <div className="p-8">{children}</div>
      </div>
    </SidebarProvider>
  );
};

export default layout;
