---
publish: "true"
permalink: /posts/built-in-typescript-support-with-nodejs
date: 2024-09-24
tags:
  - nodejs
  - typescript
  - biome
---
# Built-in TypeScript Support with Node.js

Node.js 22.6.0 adds initial (experimental), lightweight support for TypeScript. This post shows how you can use type stripping to skip transpiling, but still benefit from type checking.

Deno and Bun.js already provide a native TypeScript experience so that you can skip an explicit transpiling step (generating JavaScript files from TypeScript source code).

What's nice about this is that it lets you maintain a simple JavaScript-type of workflow (no explicit intermediate transpilation step) while also letting you code in TypeScript to get the benefits of static typing in development.

Keep in mind that the benefits of TypeScript static typing are only reaped during development. These can include helping you to catch errors earlier during programming, improving code readability, enhancing your IDE programming experience, and making it easier to refactor code.

JavaScript, however, is dynamically typed and the JavaScript runtime doesn't run TypeScript. When TypeScript is transformed (whether transparently for you or through an explicit transpilation process) into JavaScript, static type information must be erased.

Although currently experimental, with the introduction of the `--experimental-strip-types` flag, Node.js now joins the ranks of Deno and Bun.js.

```
node --experimental-strip-types ...
```

If you  want to suppress the experimental warning, you can use the following command:

```
node --experimental-strip-types --disable-warning=ExperimentalWarning ...
```

You can also set these options in the `NODE_OPTIONS` environment variable:

```
NODE_OPTIONS="--experimental-strip-types --disable-warning=ExperimentalWarning" node ...
```

## Dependency-free Node.js TypeScript support

Here's a dependency-free example:

`package.json`
```json
{  
  "type": "module",  
  "engines": {  
    "node": ">=22.6.0"  
  },  
  "scripts": {
    "start": "node ./bin/main.ts",  
    "test": "node --test ./lib/**/*.test.ts"
  }  
}
```

Where are the options? This works because the Node.js options have been stored in an [npm config file](https://docs.npmjs.com/cli/v10/using-npm/config#node-options) so there's a bit less clutter in `package.json`.

`.npmrc`
```
node-options=--experimental-strip-types --disable-warning=ExperimentalWarning
```

If I wanted to install the main module as an executable called `myapp` in the path,  it would look like this:

`package.json`
```json
{  
  "type": "module",  
  "engines": {  
    "node": ">=22.6.0"  
  },  
  "bin": {  
    "myapp": "bin/main.ts"  
  },
  "scripts": {
    "test": "node --test ./lib/**/*.test.ts"
  }  
}
```

But for this to work, the shebang line in the main module needs to look like this:

```typescript
#!/usr/bin/env node --experimental-strip-types --disable-warning=ExperimentalWarning  
console.log(`Hello, ${process.argv[2] || 'World'}!`);
```

## But wait, what about those static type checking benefits?

With the `--experimental-strip-types` you can skip the transpilation step and directly run your `.ts` files. But Node.js doesn't do any actual type checking — if you want those static typing benefits previously mentioned, that's the job of a type checking tool, like the TypeScript compiler (`tsc`).

An IDE like Visual Studio Code or WebStorm can help provide you with type checking in the editor, but you can standardize on the type checking support you want by installing type checking and linting tools as *development dependencies*.

`package.json`
```json
{  
  "type": "module",  
  "engines": {  
    "node": ">=22.6.0"  
  },  
  "bin": {  
    "myapp": "bin/main.ts"  
  },
  "scripts": {  
    "test": "node --test ./lib/**/*.test.ts",  
    "check": "biome check && tsc",  
    "fix": "biome check --write && tsc"  
  },    
  "devDependencies": {  
    "@biomejs/biome": "1.9.2",  
    "@types/node": "22.5.5",  
    "typescript": "5.6.2"  
  }
}
```

What are these dependencies for?

`typescript` provides `tsc` for type checking the source code. It will report errors but won't generate any transformed JavaScript output because of the `noEmit` setting stored in a [TypeScript configuration](https://www.typescriptlang.org/tsconfig/) file. The following is the bare minimum you'll need.

`tsconfig.json`
```json
{  
  "compilerOptions": {  
    "allowImportingTsExtensions": true,  
    "esModuleInterop": true,  
    "module": "nodenext",  
    "noEmit": true
  }  
}
```

We are strictly using `tsc` for providing us with type checking. We don't need to generate JavaScript because  `node --experimental-strip-types` flag will handle erasing types before executing the code.

`@biomejs/biome` provides another level of checking support with recommended lint rules for TypeScript. Conveniently, Biome also provides useful formatting support. [Biome](https://biomejs.dev/) should work fine with its default rules, but it's also configurable to suit specific preferences.

These two dependencies are leveraged by the `check` and `fix` package scripts. Generally `fix` is all you'll need to format and then analyze your code.

Finally, `@types/node` simply provides convenient type information for core Node.js types that will provide an enhanced experience in IDEs like Visual Studio Code and WebStorm.

## Putting all of this into practice

For a more fleshed out example, you can look at this repo: [subfuzion/mc](https://github.com/subfuzion/mc). This is a project for a CLI tool I'm prototyping for running tests on cloud infrastructure. The prototype is a work in progress, but there's enough there to highlight the Node.js TypeScript support I've described in this post.

After cloning the repo and changing to the `mc` directory, run the following (make sure you're running Node.js 22.6.0 or later):

```
npm install
npm link
```

This will install `mc` in your path.

You're running a TypeScript-only project on Node.js! After you're finished browsing the code and trying it out, you can remove the executable from your path:

```
npm unlink
```
