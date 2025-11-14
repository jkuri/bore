import { Outlet, useLocation } from "react-router-dom";
import Header from "./header";

export default function Layout(): React.JSX.Element {
  const location = useLocation();
  const isDashboard = location.pathname === "/dashboard";

  return (
    <div className="min-h-screen">
      <Header />
      {isDashboard ? (
        <div className="pt-16">
          <Outlet />
        </div>
      ) : (
        <div className="relative isolate px-6 pt-4 lg:px-8">
          <div className="mx-auto max-w-7xl py-32 sm:py-48">
            <Outlet />
          </div>
        </div>
      )}
    </div>
  );
}
