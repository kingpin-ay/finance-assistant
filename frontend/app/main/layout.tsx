"use client";
import { useUser } from "@/utils/contexts/AuthContext";
import { redirect } from "next/navigation";
import React, { useLayoutEffect } from "react";

const layout = ({ children }: { children: React.ReactNode }) => {
  const { user } = useUser();
  useLayoutEffect(() => {
    if (!user || !user.email) {
      // Redirect to the login page if user data is missing
      redirect("/auth/login");
    }
  }, [user]);
  return <div>{children}</div>;
};

export default layout;
