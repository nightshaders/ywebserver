# Introduction

Inside main/ and server/ you'll find code for a Go web-server.  In .www/ there is
a build system that will generate static assets that can be served by the Go web-server.


## Obtaining All Required Code

- Clone the repo `%> git clone <this repo>` to your $GOPATH src/ folder.
- Issue `go install web_pair/main` which should attempt to compile the web-server and all
  it's dependencies.  This being the first time compiling the code, you'll also need to
  issue `go get <missing-lib>` for each lib you don't currently have in your GOPATH.
- Next cd into .www and issue `npm install` and `bower install` (you may need install
  node `brew install node`, or bower: `npm install bower -f` for you can issue
  the first two commands).


## Install

At this point you should have all the required code.  You now need to build the
web-server and the static assets.

`cd` into the .www/ directory of the cloned repo and issue the command `%> gulp`.  If
you don't have this command you'll need to install it via `%> npm install gulp -g`.

From within $GOPATH, the web-server can be built with `%> go install web_pair/main` and
running the server typically can be done with `$GOPATH/bin/main web` from which will
start the web server listening on 9100 and serving files from inside
`$GOPATH/src/web_pair/.www/dest/_site` by default.


## License

See license file.

The use and distribution terms for this software are covered by the
[Eclipse Public License 1.0][EPL-1], which can be found in the file 'license' at the
root of this distribution. By using this software in any fashion, you are
agreeing to be bound by the terms of this license. You must not remove this
notice, or any other, from this software.


[EPL-1]: http://opensource.org/licenses/eclipse-1.0.txt

