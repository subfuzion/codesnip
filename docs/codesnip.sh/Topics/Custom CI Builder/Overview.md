In this topic, we'll explore building a custom job for running a battery of tools and tests against source code in a pull request branch. There are many ways to handle this, including by using [GitHub Actions](https://github.com/features/actions), but for this scenario we'll take a look how you would accomplish this Cloud Build.

Let's take a look at the kinds of automated checks we might want passing on a branch before a pull request is ready :

- [ ] Policy checks, for example:
	- [ ] License (check that every source file header contains a license)
	- [ ] Copyright (check that every source file header contains a valid copyright)
- [ ] Quality checks, for example:
	- [ ] Lint (analyze code for potential errors, check that it conforms to stylistic and programmatic conventions)
	- [ ] Semantic analysis (analyze for vulnerabilities with a tool like [CodeQL](https://codeql.github.com/)
- [ ] Build test (check that the source that needs to be compiled[^1] successfully builds)
- [ ] Automated tests (check that all unit and end-to-end tests pass)

We can think of the orchestration of these steps as a workflow.[^2] Google Cloud provides a fully managed orchestration platform for task execution, called [Workflows](https://cloud.google.com/workflows), that runs tasks deployed to [Cloud Run](https://cloud.google.com/run) or [Cloud Run Functions](https://cloud.google.com/functions). However, Google Cloud already provides a managed workflow service specifically dedicated to continuous integration, called [Cloud Build](https://cloud.google.com/build).

A common use case for Cloud Build is to build a container image from source code. This container image is then used to deploy containers on Google Cloud (for example, services on Cloud Run). Cloud Build is able to build typical application services using [buildpacks](https://cloud.google.com/docs/buildpacks/), an open source technology for detecting the programming language of the source code, building it, and generating a container image that gets stored in [Artifact Registry](https://cloud.google.com/artifact-registry). For the typical use case, this is all that's needed.

There are three options (in order of increasing complexity) for customizing the standard workflow:

1. Customizing the [build configuration](https://cloud.google.com/build/docs/configuring-builds/create-basic-configuration) steps
2. Creating a custom [builder](https://cloud.google.com/build/docs/configuring-builds/use-community-and-custom-builders))
3. Creating a [custom buildpack](https://cloud.google.com/docs/buildpacks/build-run-image)

Knowing that options 2 and 3 are available is reassuring. With these facilities, just about any conceivable scenario can be tackled.

However, for our scenario, we don't actually need to [build and deploy containers](https://cloud.google.com/build/docs/deploying-builds/deploy-cloud-run). The out-of-the-box support for customizing configuration just to build and test source code meets our needs.






[^1]: Any source-to-target transformation process. This includes bundling; hermetically sealing; and machine code (like Go), bytecode (like Java), or language-to-language (like TypeScript-to-JavaScript) generation.
[^2]: The term *orchestration* is used because some or all of the steps of a workflow may be executed in parallel or in sequence (required when there are dependencies on previous steps). A successful workflow is one in which all of the individual steps have finished processing successfully.