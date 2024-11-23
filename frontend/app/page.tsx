"use client";
import { useUser } from "@/utils/contexts/AuthContext";
import { redirect } from "next/navigation";
import { useLayoutEffect } from "react";

export default function Home() {
  const { user } = useUser();
  useLayoutEffect(() => {
    console.log("user", user);
    if (!user || !user.email) {
      // Redirect to the login page if user data is missing
      redirect("/auth/login");
    } else {
      redirect("/main/dashboard");
    }
  }, [user]);
  return <div></div>;
}
