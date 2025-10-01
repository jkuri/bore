import { Outlet } from "react-router-dom";
import BgPattern from "./bg-pattern";
import Header from "./header";

export default function Layout(): React.JSX.Element {
  return (
    <div className="bg-white">
      <Header />
      <div className="relative isolate px-6 pt-4 lg:px-8">
        <BgPattern />
        <div className="mx-auto max-w-4xl py-32 sm:py-48">
          <Outlet />
        </div>
      </div>
    </div>
  );
}
