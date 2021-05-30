# bore

Reverse HTTP/TCP proxy tunnel via secure SSH connections.

## Installation

First, clone repository

```sh
git clone https://github.com/jkuri/bore.git
```

Then install client:

```sh
make install_dependencies
make
cp ./build/bore /usr/local/bin/bore
```

This will compile and install `bore` client locally.

## Establish tunnel on hosted bore.network

Let's say you are running HTTP server locally on port 6500, then command would be:

```sh
bore -s bore.network -p 2200 -ls localhost -lp 6500
```

2200 is port where bore-server is running and localhost:6500 is local HTTP server.

Example output:

```sh
bore -s bore.network -p 2200 -ls localhost -lp 6500

Generated HTTP URL: http://918574de.bore.network
Generated HTTPS URL: https://918574de.bore.network
Direct TCP: tcp://bore.network:60637
```

Then open generated URL in the browser to check if it works, then share the URL if needed.

You can also request custom id instead of randomly generated one:

```sh
bore -lp 6500 -id myapp

Generated HTTP URL: http://myapp.bore.network
Generated HTTPS URL: https://myapp.bore.network
Direct TCP: tcp://bore.network:55474
```

If custom requested ID is already taken, then random id is used.

You can also specify custom remote bind listening port, which is useful for using direct TCP connection:

```sh
bore -lp 6500 -bp 55000

Generated HTTP URL: http://fe2d57f3.bore.network
Generated HTTPS URL: https://fe2d57f3.bore.network
Direct TCP: tcp://bore.network:55000
```

Note that for hosted bore you need to specify port in range 55000-65000.

If port is already taken, random port is used.

## Running Server

### Run Compilation

```sh
make install_dependencies
make
```

### Running bore-server example

```sh
BORE_DOMAIN=bore.network BORE_HTTPADDR=0.0.0.0:80 BORE_SSHADDR=0.0.0.0:2200 ./build/bore-server
```

This will generate initial config at `~/bore/bore-server.yaml` with values you provided over environment variables.

## License

```license
MIT License

Copyright (c) 2020-2021 Jan Kuri <jkuri88@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
