import { DashboardIcon, GitHubLogoIcon } from "@radix-ui/react-icons";
import { NavLink } from "react-router-dom";
import { ThemeToggle } from "./theme-toggle";

export default function Header(): React.JSX.Element {
  return (
    <header className="absolute inset-x-0 top-0 z-50">
      <nav className="mx-auto flex max-w-7xl items-center justify-between p-6">
        <div className="flex flex-1">
          <NavLink to="/">
            <img className="h-8" src="/bore.svg" alt="Bore" />
          </NavLink>
        </div>
        <div className="flex items-center justify-end gap-4">
          <NavLink
            to="/dashboard"
            className={({ isActive }) =>
              `flex items-center gap-2 rounded-md px-3 py-2 font-medium text-sm transition-colors ${
                isActive
                  ? "bg-secondary text-foreground"
                  : "text-muted-foreground hover:bg-accent hover:text-foreground"
              }`
            }
          >
            <DashboardIcon className="h-4 w-4" />
            Dashboard
          </NavLink>
          <ThemeToggle />
          <a
            href="https://github.com/jkuri/bore"
            className="text-foreground transition-colors hover:text-muted-foreground"
          >
            <GitHubLogoIcon className="h-6 w-6" />
          </a>
        </div>
      </nav>
    </header>
  );
}
