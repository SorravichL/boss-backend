"use client";

import { useMenu } from "./MenuContext";
import Toggle from "./theme";
import ToggleMenu from "./toggleMenu";

const Header = () => {
  const { isOpen, toggleMenu } = useMenu();

  return (
    <header>
      <div className="header-content">
        <div className="left-part">
          <div className="logo"></div>
          <h5 className="full-identity">
            <span>Sorravich Lakngoenchai</span>
          </h5>
        </div>
        <div className="right-part">
          <ul className="flex">
            <li className="theme-switch"><Toggle /></li>
            <li className="menu-toggle">
              <ToggleMenu toggle={toggleMenu} isOpen={isOpen} />
            </li>
          </ul>
        </div>
      </div>
    </header>
  );
};

export default Header;
