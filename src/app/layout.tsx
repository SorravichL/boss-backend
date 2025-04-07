"use client"
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import { MenuProvider, useMenu } from "./components/MenuContext";
import Header from "./components/header";
import Menu from "./components/menu";
import { motion, AnimatePresence } from "framer-motion";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});


function ContentWrapper({ children }: { children: React.ReactNode }) {
  const { isOpen } = useMenu();

  return (
    <AnimatePresence mode="wait">
      <motion.div
        key={isOpen ? "menu" : "page"}
        initial={{ opacity: 0, y: -50 }}
        animate={{ y: isOpen ? -50 : 0, opacity: isOpen ? 0 : 1 ,transition: { duration: 0.5, ease: "easeInOut", delay: 0.4 },}}
        exit={{ opacity: 0, y: -50,transition: { duration: 0.5, ease: "easeInOut" }, }}
        className="page-wrapper"
      >
        {children}
      </motion.div>
    </AnimatePresence>
  );
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <MenuProvider>
      <html lang="en">
        <body
          className={`${geistSans.variable} ${geistMono.variable} antialiased`}
        >
          <Header />
          <Menu />

          <div className="canvas"></div>

          <main className="container">
            <div className="opacity-100 transform-none">
              <div className="data-scroll-container">
                <div className="data-scroll-section">
                  <div className="page-content">
                    <ContentWrapper>{children}</ContentWrapper>
                  </div>
                </div>
              </div>
            </div>
          </main>

          <div className="noise"></div>
          <div className="cursor-dot"></div>
          <div className="cursor"></div>
        </body>
      </html>
    </MenuProvider>
  );
}
