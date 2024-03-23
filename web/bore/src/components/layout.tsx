import { Outlet } from 'react-router-dom';
import Header from './header';

export default function Layout(): React.JSX.Element {
  return (
    <div className="bg-white">
      <Header />
      <div className="relative isolate px-6 pt-4 lg:px-8">
        <div className="mx-auto max-w-2xl py-32 sm:py-48">
          <Outlet />
        </div>
      </div>
    </div>
  );
}
