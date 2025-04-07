"use client";

import Link from "next/link";
import { motion, AnimatePresence } from "framer-motion";
import { useMenu } from "./MenuContext";

export default function Menu() {
  const { isOpen } = useMenu();
  return (
    <AnimatePresence>
      {isOpen && (
        <motion.div
          className="menu"
          initial={{ opacity: 0, y: -50 }}
          animate={{
            opacity: 1,
            y: 0,
            transition: { duration: 0.5, ease: "easeInOut", delay: 0.4 },
          }}
          exit={{
            opacity: 0,
            y: -50,
            transition: { duration: 0.5, ease: "easeInOut" },
          }}
        >
          <div className="menu-content">
            <div className="menu-items">
              <nav>
                <ol>
                  <li><div className="item-inner"><Link href="/"><h1>Home</h1></Link></div></li>
                  <li><div className="item-inner"><Link href="/work"><h1>Work</h1></Link></div></li>
                  <li><div className="item-inner"><Link href="/about"><h1>About</h1></Link></div></li>
                  <li><div className="item-inner"><Link href="/contact"><h1>Contact</h1></Link></div></li>
                </ol>
              </nav>
            </div>
            <div className="menu-socials">
              <ul>
                <div className="menu-social-inside">
                  <li>
                    <div className="item-inner">
                      <div className="link-wrapper">
                        <div className="link">
                          <a href="https://www.instagram.com/sorravich._/" target="_blank" rel="noreferrer">Instagram</a>
                        </div>
                        <div className="link-underline"></div>
                      </div>
                    </div>
                  </li>
                  <li>
                    <div className="item-inner">
                      <div className="link-wrapper">
                        <div className="link">
                          <a href="https://github.com/SorravichL" target="_blank" rel="noreferrer">Github</a>
                        </div>
                        <div className="link-underline"></div>
                      </div>
                    </div>
                  </li>
                </div>
              </ul>
            </div>
          </div>
        </motion.div>
      )}
    </AnimatePresence>
  );
}
