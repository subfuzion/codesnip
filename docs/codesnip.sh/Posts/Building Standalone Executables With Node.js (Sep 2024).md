---
publish: "true"
permalink: /posts/building-standalone-nodejs-executables
date: 2024-09-24
tags:
  - nodejs
  - executable
  - bundle
---
# Building Standalone Executables With Node.js

Node.js has experimental support for building a [single executable application](https://github.com/nodejs/single-executable), or SEA, which is what the team calls a standalone executable that can be distributed to [supported platforms](https://nodejs.org/api/single-executable-applications.html#platform-support).

The way it works is that a single *blob* is injected into the `node` binary for a supported platform. This single blob can be a JavaScript file, including multi-file JavaScript source that has been bundled with a tool like Webpack. If the blob is present, then the `node` binary will execute the script in the blob.

The recipe for building a SEA is spelled out in the [docs](https://nodejs.org/api/single-executable-applications.html#single-executable-applications). Unfortunately, compared to Deno and Bun.js, the developer experience is pretty rough at the moment.

I wrote a script to automate building a SEA for Linux and Darwin targets (I also took a stab at the Windows target, but it's not tested) to smooth out the experience.

I demonstrate a simple project that uses Webpack for bundling. You can use the project as a starting point for your own project.

The project is straightforward enough, but it assumes the following:

* You're running on the platform you want to create the executable for (generating cross-platform SEAs is possible, however)
* Your main module is `bin/app.js` (if not, modify `entry` in `webpack.config.js`)
* Intermediate output (the bundle and the blob it's injected into) will be generated under `.build`
* The executable will be generated under `dist/{target}`

The GitHub repo is here:
https://github.com/subfuzion/nodejs-executable-app

What's next? I plan to create and share a standalone tool (a SEA, of course!) with a CLI for prompts and support for at least `esbuild` and `swc` bundling. I also have some thoughts about contributing to Node.js to pitch in and help enhance the SEA devex a bit, if I can carve out some time.