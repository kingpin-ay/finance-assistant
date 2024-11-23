"use client";
import { createContext, ReactNode, useContext, useState } from "react";

interface User {
  email: string;
  firstname: string;
  lastname: string;
  user_id: number | null;
}

interface UserContextType {
  user: User;
  updateUser: (data: User) => void;
}

export const UserContext = createContext<UserContextType | undefined>(
  undefined
);

// Define the Provider's props
interface UserProviderProps {
  children: ReactNode;
}

// Provider component
export const UserProvider: React.FC<UserProviderProps> = ({ children }) => {
  const [user, setUser] = useState<User>({
    email: "",
    firstname: "",
    lastname: "",
    user_id: null,
  });

  // Function to update user data
  const updateUser = (data: User) => {
    setUser(data);
  };

  return (
    <UserContext.Provider value={{ user, updateUser }}>
      {children}
    </UserContext.Provider>
  );
};

// Custom hook for using the context
export const useUser = (): UserContextType => {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error("useUser must be used within a UserProvider");
  }
  return context;
};
