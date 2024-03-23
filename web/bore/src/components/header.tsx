import { GitHubLogoIcon } from '@radix-ui/react-icons';

export default function Header(): React.JSX.Element {
  return (
    <header className="absolute inset-x-0 top-0 z-50">
      <nav className="flex items-center justify-between p-3 lg:px-8">
        <div className="flex flex-1">
          <a href="/" className="-m-1.5 p-1.5">
            <span className="sr-only">Bore</span>
            <img className="h-10 w-auto" src="/bore-logo.svg" alt="" />
          </a>
        </div>
        <div className="flex items-center justify-end">
          <a
            href="https://github.com/jkuri/bore"
            className="text-sm font-semibold leading-6 text-gray-900"
          >
            <GitHubLogoIcon className="h-8 w-8" />
          </a>
        </div>
      </nav>
    </header>
  );
}
