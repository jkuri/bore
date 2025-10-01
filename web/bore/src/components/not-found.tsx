import { ArrowRightIcon, RocketIcon } from "@radix-ui/react-icons";
import { NavLink, useSearchParams } from "react-router-dom";

export default function NotFound(): React.JSX.Element {
  const [searchParams] = useSearchParams();
  const tunnelID = searchParams.get("tunnelID") as string;

  return (
    <div className="space-y-8 text-center">
      <div className="space-y-4">
        <h1 className="font-bold text-5xl text-zinc-900 tracking-tight sm:text-6xl">
          Tunnel Not Found
        </h1>

        <div className="flex items-center justify-center">
          <p className="text-lg text-zinc-600 leading-8">
            Tunnel with ID
            <span className="mx-2 inline-flex items-center gap-2 rounded-md border border-zinc-200 bg-zinc-50 px-3 py-1.5 font-medium text-sm text-zinc-900">
              <RocketIcon className="h-4 w-4" />
              {tunnelID}
            </span>
            was not found.
          </p>
        </div>

        <p className="text-zinc-600">
          The tunnel may have expired or the ID is incorrect.
        </p>
      </div>

      <div className="flex flex-col items-center justify-center gap-3 sm:flex-row">
        <NavLink to="/">
          <span className="inline-flex h-10 items-center justify-center gap-2 rounded-md bg-zinc-900 px-8 font-medium text-sm text-zinc-50 transition-colors hover:bg-zinc-800 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-zinc-950">
            Go back home
          </span>
        </NavLink>
        <a
          href="https://github.com/jkuri/bore/blob/master/README.md"
          className="inline-flex h-10 items-center justify-center gap-2 rounded-md border border-zinc-200 bg-white px-8 font-medium text-sm text-zinc-900 transition-colors hover:bg-zinc-50 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-zinc-950"
        >
          View documentation
          <ArrowRightIcon className="h-4 w-4" />
        </a>
      </div>
    </div>
  );
}
