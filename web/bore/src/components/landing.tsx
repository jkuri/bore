export default function Landing(): React.JSX.Element {
  return (
    <>
      <div className="hidden sm:mb-8 sm:flex sm:justify-center">
        <div className="relative rounded-full bg-white px-3 py-1 text-sm leading-6 text-gray-600 ring-1 ring-gray-900/10 hover:ring-gray-900/20">
          bore v0.4.2 now released.{' '}
          <a
            href="https://github.com/jkuri/bore/releases"
            className="font-semibold text-blue-600"
          >
            <span className="absolute inset-0" />
            Download here <span>&rarr;</span>
          </a>
        </div>
      </div>
      <div className="text-center">
        <h1 className="text-4xl font-bold tracking-tight text-blue-600 sm:text-6xl">
          Expose Yourself to the World!
        </h1>
        <p className="mt-6 text-lg leading-8 text-gray-600">
          Rolling behind NAT and quickly want to show your work to a collegue,
          download bore client and make secure tunnel over SSH protocol.
        </p>
        <div className="mt-10 flex items-center justify-center gap-x-6">
          <a
            href="https://github.com/jkuri/bore/releases"
            className="rounded-md bg-blue-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-blue-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-500"
          >
            Download
          </a>
          <a
            href="https://github.com/jkuri/bore/blob/master/README.md"
            className="text-sm font-semibold leading-6 text-gray-900"
          >
            Show docs →
          </a>
        </div>
      </div>
    </>
  );
}
