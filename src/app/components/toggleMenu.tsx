"use client";
export default function ToggleMenu({ toggle, isOpen }: { toggle: () => void; isOpen: boolean }) {
  return (
    <button id="toggleMenu" onClick={toggle}>
      {isOpen ? "Close Menu" : "Open Menu"}
    </button>
  );
}
