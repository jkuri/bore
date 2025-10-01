import { RocketIcon } from "@radix-ui/react-icons";
import { NavLink, useSearchParams } from "react-router-dom";

export default function NotFound(): React.JSX.Element {
  const [searchParams] = useSearchParams();
  const tunnelID = searchParams.get("tunnelID") as string;

  return (
    <div className="text-center">
      <h1 className="font-bold text-5xl text-zinc-900 tracking-tight">
        Tunnel Not Found!
      </h1>
      <div className="inline-flex items-center justify-center">
        <p className="mt-6 text-gray-600 text-lg leading-8">
          Tunnel with ID
          <span className="mx-2 inline-flex items-center rounded-full border bg-white px-3 py-2 font-bold text-gray-700 text-xs leading-sm">
            <RocketIcon className="mr-2" />
            {tunnelID}
          </span>
          not found.
        </p>
      </div>
      <div className="mt-10 flex items-center justify-center gap-x-6">
        <NavLink to="/">
          <span className="rounded-md bg-blue-600 px-3.5 py-2.5 font-semibold text-sm text-white shadow-sm hover:bg-blue-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-blue-500 focus-visible:outline-offset-2">
            Go back
          </span>
        </NavLink>
        <a
          href="https://github.com/jkuri/bore/blob/master/README.md"
          className="font-semibold text-gray-900 text-sm leading-6"
        >
          Show docs →
        </a>
      </div>
    </div>
  );
}
