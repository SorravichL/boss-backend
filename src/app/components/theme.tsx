import { useState } from "react";

export default function Toggle() {

  const [isDark, setIsDark] = useState(false);
  
  const toggleTheme = () => {
    const currentTheme = document.documentElement.getAttribute("data-theme");
  
    if (currentTheme === "dark") {
      document.documentElement.setAttribute("data-theme", "light");
    } else {
      document.documentElement.setAttribute("data-theme", "dark");
    }
  }
  return (
    <div>
      <input
        type="checkbox"
        id="check"
        className="toggle"
        onChange={() => setIsDark(!isDark)}
        checked={isDark}
        onClick={toggleTheme}
      />
    </div>
  );
}
