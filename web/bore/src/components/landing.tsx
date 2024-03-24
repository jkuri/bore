import { ArrowRightIcon, DownloadIcon } from '@radix-ui/react-icons';

export default function Landing(): React.JSX.Element {
  return (
    <>
      <div className="hidden sm:mb-8 sm:flex sm:justify-center">
        <div className="relative rounded-full bg-white px-3 py-1 text-sm leading-6 text-gray-600 ring-1 ring-gray-900/10 hover:ring-gray-900/20">
          bore v0.4.2 now released.{' '}
          <a
            href="https://github.com/jkuri/bore/releases"
            className="font-semibold text-pink-500"
          >
            <span className="absolute inset-0" />
            Download here <span>&rarr;</span>
          </a>
        </div>
      </div>
      <div className="text-center">
        <h1 className="tracking-tigh bg-gradient-to-r from-green-500 to-blue-500 bg-clip-text text-4xl font-bold text-transparent sm:text-6xl">
          Expose Yourself to the World!
        </h1>
        <p className="mt-8 text-lg leading-8 text-zinc-900">
          Rolling behind NAT and quickly want to show your work to a collegue,
          download bore client and make secure tunnel over SSH protocol.
        </p>
        <div className="mt-16 flex items-center justify-center gap-x-6">
          <a
            href="https://github.com/jkuri/bore/releases"
            className="inline-flex items-center rounded-md bg-green-600 px-4 py-2.5 text-sm font-semibold text-white shadow-sm transition-colors hover:bg-green-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-green-500"
          >
            Download
            <DownloadIcon className="ml-2" />
          </a>
          <a
            href="https://github.com/jkuri/bore/blob/master/README.md"
            className="inline-flex items-center rounded-md border border-gray-200 bg-white px-4 py-2.5 text-sm font-semibold text-zinc-700 shadow-sm transition-colors hover:bg-gray-50 focus:outline-none focus:ring-4 focus:ring-gray-100"
          >
            Readme
            <ArrowRightIcon className="ml-2" />
          </a>
        </div>
      </div>
    </>
  );
}
