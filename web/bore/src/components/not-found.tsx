import { ArrowRightIcon, RocketIcon } from "@radix-ui/react-icons";
import { NavLink, useSearchParams } from "react-router-dom";

export default function NotFound(): React.JSX.Element {
  const [searchParams] = useSearchParams();
  const tunnelID = searchParams.get("tunnelID") as string;

  return (
    <div className="space-y-8 text-center">
      <div className="space-y-4">
        <h1 className="font-bold text-5xl text-foreground tracking-tight sm:text-6xl">
          Tunnel Not Found
        </h1>

        <div className="flex items-center justify-center">
          <p className="text-lg text-muted-foreground leading-8">
            Tunnel with ID
            <span className="mx-2 inline-flex items-center gap-2 rounded-md border bg-secondary px-3 py-1.5 font-medium text-foreground text-sm">
              <RocketIcon className="h-4 w-4" />
              {tunnelID}
            </span>
            was not found.
          </p>
        </div>

        <p className="text-muted-foreground">
          The tunnel may have expired or the ID is incorrect.
        </p>
      </div>

      <div className="flex flex-col items-center justify-center gap-3 sm:flex-row">
        <NavLink to="/">
          <span className="inline-flex h-10 items-center justify-center gap-2 rounded-md bg-primary px-8 font-medium text-primary-foreground text-sm transition-colors hover:bg-primary/90 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring">
            Go back home
          </span>
        </NavLink>
        <a
          href="https://github.com/jkuri/bore/blob/master/README.md"
          className="inline-flex h-10 items-center justify-center gap-2 rounded-md border bg-background px-8 font-medium text-foreground text-sm transition-colors hover:bg-accent focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring"
        >
          View documentation
          <ArrowRightIcon className="h-4 w-4" />
        </a>
      </div>
    </div>
  );
}
