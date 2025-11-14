import {
  ArrowRightIcon,
  DownloadIcon,
  LockClosedIcon,
  RocketIcon,
} from "@radix-ui/react-icons";
import {
  FlaskConical,
  GitFork,
  Globe,
  Home,
  Monitor,
  Presentation,
  Server,
  Terminal,
} from "lucide-react";

export default function Landing(): React.JSX.Element {
  return (
    <div className="space-y-20">
      <div className="space-y-8">
        <div className="flex justify-center">
          <div className="inline-flex items-center gap-2 rounded-md border bg-secondary px-3 py-1 text-muted-foreground text-sm">
            <span className="relative flex h-2 w-2">
              <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-muted-foreground opacity-75"></span>
              <span className="relative inline-flex h-2 w-2 rounded-full bg-muted-foreground"></span>
            </span>
            <span>bore v0.5.0 now available</span>
          </div>
        </div>

        <div className="space-y-4 text-center">
          <h1 className="font-bold text-4xl text-foreground tracking-tight sm:text-6xl">
            Expose local servers
            <br />
            behind NAT and firewalls
          </h1>

          <p className="mx-auto max-w-2xl text-lg text-muted-foreground leading-8">
            A simple command-line tool for creating secure SSH tunnels to share
            your local development server with anyone, anywhere.
          </p>
        </div>

        <div className="flex flex-col items-center justify-center gap-3 sm:flex-row">
          <a
            href="https://github.com/jkuri/bore/releases"
            className="inline-flex h-10 items-center justify-center gap-2 rounded-md bg-primary px-8 font-medium text-primary-foreground text-sm transition-colors hover:bg-primary/90 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring"
          >
            <DownloadIcon className="h-4 w-4" />
            Get Started
          </a>
          <a
            href="https://github.com/jkuri/bore"
            className="inline-flex h-10 items-center justify-center gap-2 rounded-md border bg-background px-8 font-medium text-foreground text-sm transition-colors hover:bg-accent focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring"
          >
            View on GitHub
            <ArrowRightIcon className="h-4 w-4" />
          </a>
        </div>
      </div>

      <div className="mx-auto max-w-3xl">
        <div className="overflow-hidden rounded-lg border bg-muted">
          <div className="border-b bg-card px-4 py-2">
            <div className="flex items-center gap-2">
              <div className="flex gap-1.5">
                <div className="h-3 w-3 rounded-full bg-muted-foreground/30"></div>
                <div className="h-3 w-3 rounded-full bg-muted-foreground/30"></div>
                <div className="h-3 w-3 rounded-full bg-muted-foreground/30"></div>
              </div>
              <span className="ml-2 text-muted-foreground text-xs">
                bore.digital
              </span>
            </div>
          </div>
          <div className="p-6 font-mono text-sm leading-relaxed">
            <div className="space-y-3">
              <div className="flex gap-2">
                <span className="text-muted-foreground">$</span>
                <span className="text-foreground">bore -lp 3000 -id myapp</span>
              </div>
              <div className="space-y-1 text-muted-foreground">
                <div>Welcome to bore server 0.5.0 at bore.digital</div>
                <div className="flex items-center gap-5">
                  <div className="text-muted-foreground/70">→ HTTP</div>
                  <div>http://myapp.bore.digital</div>
                </div>
                <div className="flex items-center gap-3">
                  <div className="text-muted-foreground/70">→ HTTPS</div>
                  <div>https://myapp.bore.digital</div>
                </div>
                <div className="flex items-center gap-7">
                  <div className="text-muted-foreground/70">→ TCP</div>
                  <div>tcp://bore.digital:64746</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="space-y-8">
        <div className="space-y-2 text-center">
          <h2 className="font-bold text-3xl text-foreground tracking-tight">
            Features
          </h2>
          <p className="text-muted-foreground">
            Everything you need for secure tunneling
          </p>
        </div>

        <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
          <div className="space-y-2 rounded-lg border bg-card p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-primary">
              <LockClosedIcon className="h-5 w-5 text-primary-foreground" />
            </div>
            <h3 className="font-semibold text-foreground">Secure by default</h3>
            <p className="text-muted-foreground text-sm">
              Built on SSH protocol with end-to-end encryption for all tunnels.
            </p>
          </div>

          <div className="space-y-2 rounded-lg border bg-card p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-primary">
              <RocketIcon className="h-5 w-5 text-primary-foreground" />
            </div>
            <h3 className="font-semibold text-foreground">Fast & reliable</h3>
            <p className="text-muted-foreground text-sm">
              Minimal overhead with optimized performance for low latency.
            </p>
          </div>

          <div className="space-y-2 rounded-lg border bg-card p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-primary">
              <Terminal className="h-5 w-5 text-primary-foreground" />
            </div>
            <h3 className="font-semibold text-foreground">Simple CLI</h3>
            <p className="text-muted-foreground text-sm">
              One command to expose your server. No configuration required.
            </p>
          </div>

          <div className="space-y-2 rounded-lg border bg-card p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-primary">
              <Globe className="h-5 w-5 text-primary-foreground" />
            </div>
            <h3 className="font-semibold text-foreground">HTTP & TCP</h3>
            <p className="text-muted-foreground text-sm">
              Support for any protocol. Perfect for web apps and databases.
            </p>
          </div>

          <div className="space-y-2 rounded-lg border bg-card p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-primary">
              <Server className="h-5 w-5 text-primary-foreground" />
            </div>
            <h3 className="font-semibold text-foreground">Self-hosted</h3>
            <p className="text-muted-foreground text-sm">
              Run your own server with full control over your infrastructure.
            </p>
          </div>

          <div className="space-y-2 rounded-lg border bg-card p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-primary">
              <GitFork className="h-5 w-5 text-primary-foreground" />
            </div>
            <h3 className="font-semibold text-foreground">Open source</h3>
            <p className="text-muted-foreground text-sm">
              Free and open source. Inspect, contribute, and customize freely.
            </p>
          </div>
        </div>
      </div>

      <div className="space-y-8">
        <div className="space-y-2 text-center">
          <h2 className="font-bold text-3xl text-foreground tracking-tight">
            Use cases
          </h2>
          <p className="text-muted-foreground">
            Perfect for developers, designers, and teams
          </p>
        </div>

        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
          <div className="space-y-2 rounded-lg border bg-card p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-secondary">
              <Monitor className="h-5 w-5 text-foreground" />
            </div>
            <h3 className="font-semibold text-foreground">Development</h3>
            <p className="text-muted-foreground text-sm">
              Share localhost with clients and teammates
            </p>
          </div>

          <div className="space-y-2 rounded-lg border bg-card p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-secondary">
              <Presentation className="h-5 w-5 text-foreground" />
            </div>
            <h3 className="font-semibold text-foreground">Demos</h3>
            <p className="text-muted-foreground text-sm">
              Show work in progress instantly
            </p>
          </div>

          <div className="space-y-2 rounded-lg border bg-card p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-secondary">
              <FlaskConical className="h-5 w-5 text-foreground" />
            </div>
            <h3 className="font-semibold text-foreground">Testing</h3>
            <p className="text-muted-foreground text-sm">
              Test webhooks and integrations
            </p>
          </div>

          <div className="space-y-2 rounded-lg border bg-card p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-secondary">
              <Home className="h-5 w-5 text-foreground" />
            </div>
            <h3 className="font-semibold text-foreground">Home labs</h3>
            <p className="text-muted-foreground text-sm">
              Access home servers remotely
            </p>
          </div>
        </div>
      </div>

      <div className="space-y-4 text-center">
        <h2 className="font-bold text-2xl text-foreground tracking-tight">
          Ready to get started?
        </h2>
        <p className="text-muted-foreground">
          Download bore and start tunneling in seconds
        </p>
        <div className="flex flex-col items-center justify-center gap-3 sm:flex-row">
          <a
            href="https://github.com/jkuri/bore/releases"
            className="inline-flex h-10 items-center justify-center gap-2 rounded-md bg-primary px-8 font-medium text-primary-foreground text-sm transition-colors hover:bg-primary/90"
          >
            <DownloadIcon className="h-4 w-4" />
            Download
          </a>
          <a
            href="https://github.com/jkuri/bore/blob/master/README.md"
            className="inline-flex h-10 items-center justify-center gap-2 rounded-md border bg-background px-8 font-medium text-foreground text-sm transition-colors hover:bg-accent"
          >
            Documentation
            <ArrowRightIcon className="h-4 w-4" />
          </a>
        </div>
      </div>
    </div>
  );
}
