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
          <div className="inline-flex items-center gap-2 rounded-md border border-zinc-200 bg-zinc-50 px-3 py-1 text-sm text-zinc-600">
            <span className="relative flex h-2 w-2">
              <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-zinc-400 opacity-75"></span>
              <span className="relative inline-flex h-2 w-2 rounded-full bg-zinc-500"></span>
            </span>
            <span>bore v0.4.2 now available</span>
          </div>
        </div>

        <div className="space-y-4 text-center">
          <h1 className="font-bold text-4xl text-zinc-900 tracking-tight sm:text-6xl">
            Expose local servers
            <br />
            behind NAT and firewalls
          </h1>

          <p className="mx-auto max-w-2xl text-lg text-zinc-600 leading-8">
            A simple command-line tool for creating secure SSH tunnels to share
            your local development server with anyone, anywhere.
          </p>
        </div>

        <div className="flex flex-col items-center justify-center gap-3 sm:flex-row">
          <a
            href="https://github.com/jkuri/bore/releases"
            className="inline-flex h-10 items-center justify-center gap-2 rounded-md bg-zinc-900 px-8 font-medium text-sm text-zinc-50 transition-colors hover:bg-zinc-800 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-zinc-950"
          >
            <DownloadIcon className="h-4 w-4" />
            Get Started
          </a>
          <a
            href="https://github.com/jkuri/bore"
            className="inline-flex h-10 items-center justify-center gap-2 rounded-md border border-zinc-200 bg-white px-8 font-medium text-sm text-zinc-900 transition-colors hover:bg-zinc-50 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-zinc-950"
          >
            View on GitHub
            <ArrowRightIcon className="h-4 w-4" />
          </a>
        </div>
      </div>

      <div className="mx-auto max-w-3xl">
        <div className="overflow-hidden rounded-lg border border-zinc-200 bg-zinc-50">
          <div className="border-zinc-200 border-b bg-white px-4 py-2">
            <div className="flex items-center gap-2">
              <div className="flex gap-1.5">
                <div className="h-3 w-3 rounded-full bg-zinc-300"></div>
                <div className="h-3 w-3 rounded-full bg-zinc-300"></div>
                <div className="h-3 w-3 rounded-full bg-zinc-300"></div>
              </div>
              <span className="ml-2 text-xs text-zinc-500">bore.digital</span>
            </div>
          </div>
          <div className="p-6 font-mono text-sm leading-relaxed">
            <div className="space-y-3">
              <div className="flex gap-2">
                <span className="text-zinc-500">$</span>
                <span className="text-zinc-900">bore -lp 3000 -id myapp</span>
              </div>
              <div className="space-y-1 text-zinc-600">
                <div>Welcome to bore server 0.4.2 at bore.digital</div>
                <div className="flex items-center gap-5">
                  <div className="text-zinc-500">→ HTTP</div>
                  <div>http://myapp.bore.digital</div>
                </div>
                <div className="flex items-center gap-3">
                  <div className="text-zinc-500">→ HTTPS</div>
                  <div>https://myapp.bore.digital</div>
                </div>
                <div className="flex items-center gap-7">
                  <div className="text-zinc-500">→ TCP</div>
                  <div>tcp://bore.digital:64746</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="space-y-8">
        <div className="space-y-2 text-center">
          <h2 className="font-bold text-3xl text-zinc-900 tracking-tight">
            Features
          </h2>
          <p className="text-zinc-600">
            Everything you need for secure tunneling
          </p>
        </div>

        <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
          <div className="space-y-2 rounded-lg border border-zinc-200 bg-white p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-zinc-900">
              <LockClosedIcon className="h-5 w-5 text-zinc-50" />
            </div>
            <h3 className="font-semibold text-zinc-900">Secure by default</h3>
            <p className="text-sm text-zinc-600">
              Built on SSH protocol with end-to-end encryption for all tunnels.
            </p>
          </div>

          <div className="space-y-2 rounded-lg border border-zinc-200 bg-white p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-zinc-900">
              <RocketIcon className="h-5 w-5 text-zinc-50" />
            </div>
            <h3 className="font-semibold text-zinc-900">Fast & reliable</h3>
            <p className="text-sm text-zinc-600">
              Minimal overhead with optimized performance for low latency.
            </p>
          </div>

          <div className="space-y-2 rounded-lg border border-zinc-200 bg-white p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-zinc-900">
              <Terminal className="h-5 w-5 text-zinc-50" />
            </div>
            <h3 className="font-semibold text-zinc-900">Simple CLI</h3>
            <p className="text-sm text-zinc-600">
              One command to expose your server. No configuration required.
            </p>
          </div>

          <div className="space-y-2 rounded-lg border border-zinc-200 bg-white p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-zinc-900">
              <Globe className="h-5 w-5 text-zinc-50" />
            </div>
            <h3 className="font-semibold text-zinc-900">HTTP & TCP</h3>
            <p className="text-sm text-zinc-600">
              Support for any protocol. Perfect for web apps and databases.
            </p>
          </div>

          <div className="space-y-2 rounded-lg border border-zinc-200 bg-white p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-zinc-900">
              <Server className="h-5 w-5 text-zinc-50" />
            </div>
            <h3 className="font-semibold text-zinc-900">Self-hosted</h3>
            <p className="text-sm text-zinc-600">
              Run your own server with full control over your infrastructure.
            </p>
          </div>

          <div className="space-y-2 rounded-lg border border-zinc-200 bg-white p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-zinc-900">
              <GitFork className="h-5 w-5 text-zinc-50" />
            </div>
            <h3 className="font-semibold text-zinc-900">Open source</h3>
            <p className="text-sm text-zinc-600">
              Free and open source. Inspect, contribute, and customize freely.
            </p>
          </div>
        </div>
      </div>

      <div className="space-y-8">
        <div className="space-y-2 text-center">
          <h2 className="font-bold text-3xl text-zinc-900 tracking-tight">
            Use cases
          </h2>
          <p className="text-zinc-600">
            Perfect for developers, designers, and teams
          </p>
        </div>

        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
          <div className="space-y-2 rounded-lg border border-zinc-200 bg-white p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-zinc-100">
              <Monitor className="h-5 w-5 text-zinc-900" />
            </div>
            <h3 className="font-semibold text-zinc-900">Development</h3>
            <p className="text-sm text-zinc-600">
              Share localhost with clients and teammates
            </p>
          </div>

          <div className="space-y-2 rounded-lg border border-zinc-200 bg-white p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-zinc-100">
              <Presentation className="h-5 w-5 text-zinc-900" />
            </div>
            <h3 className="font-semibold text-zinc-900">Demos</h3>
            <p className="text-sm text-zinc-600">
              Show work in progress instantly
            </p>
          </div>

          <div className="space-y-2 rounded-lg border border-zinc-200 bg-white p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-zinc-100">
              <FlaskConical className="h-5 w-5 text-zinc-900" />
            </div>
            <h3 className="font-semibold text-zinc-900">Testing</h3>
            <p className="text-sm text-zinc-600">
              Test webhooks and integrations
            </p>
          </div>

          <div className="space-y-2 rounded-lg border border-zinc-200 bg-white p-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-md bg-zinc-100">
              <Home className="h-5 w-5 text-zinc-900" />
            </div>
            <h3 className="font-semibold text-zinc-900">Home labs</h3>
            <p className="text-sm text-zinc-600">
              Access home servers remotely
            </p>
          </div>
        </div>
      </div>

      <div className="space-y-4 text-center">
        <h2 className="font-bold text-2xl text-zinc-900 tracking-tight">
          Ready to get started?
        </h2>
        <p className="text-zinc-600">
          Download bore and start tunneling in seconds
        </p>
        <div className="flex flex-col items-center justify-center gap-3 sm:flex-row">
          <a
            href="https://github.com/jkuri/bore/releases"
            className="inline-flex h-10 items-center justify-center gap-2 rounded-md bg-zinc-900 px-8 font-medium text-sm text-zinc-50 transition-colors hover:bg-zinc-800"
          >
            <DownloadIcon className="h-4 w-4" />
            Download
          </a>
          <a
            href="https://github.com/jkuri/bore/blob/master/README.md"
            className="inline-flex h-10 items-center justify-center gap-2 rounded-md border border-zinc-200 bg-white px-8 font-medium text-sm text-zinc-900 transition-colors hover:bg-zinc-50"
          >
            Documentation
            <ArrowRightIcon className="h-4 w-4" />
          </a>
        </div>
      </div>
    </div>
  );
}
