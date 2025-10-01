import { GitHubLogoIcon } from '@radix-ui/react-icons';
import { NavLink } from 'react-router-dom';

export default function Header(): React.JSX.Element {
  return (
    <header className="absolute inset-x-0 top-0 z-50">
      <nav className="mx-auto flex max-w-4xl items-center justify-between p-4">
        <div className="flex flex-1">
          <NavLink to="/">
            <img className="h-8" src="/bore.svg" alt="Bore" />
          </NavLink>
        </div>
        <div className="flex items-center justify-end">
          <a
            href="https://github.com/jkuri/bore"
            className="text-zinc-900 transition-colors hover:text-zinc-700"
          >
            <GitHubLogoIcon className="h-6 w-6" />
          </a>
        </div>
      </nav>
    </header>
  );
}
