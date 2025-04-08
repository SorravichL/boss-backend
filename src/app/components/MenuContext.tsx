// components/MenuContext.tsx
"use client";

import { createContext, useContext, useState } from "react";

interface MenuContextType {
    isOpen: boolean;
    setIsOpen: React.Dispatch<React.SetStateAction<boolean>>;
    toggleMenu:()=>void;
  }

  const MenuContext = createContext<MenuContextType | undefined>(undefined);

  export const useMenu = () => {
    const context = useContext(MenuContext);
    if (!context) {
      throw new Error("useMenu must be used within a MenuProvider");
    }
    return context;
  };

  export const MenuProvider = ({ children }: { children: React.ReactNode }) => {
    const [isOpen, setIsOpen] = useState(false);

    const toggleMenu = () => setIsOpen(prev => !prev);
  
    return (
      <MenuContext.Provider value={{ isOpen, setIsOpen, toggleMenu }}>
        {children}
      </MenuContext.Provider>
    );
  };