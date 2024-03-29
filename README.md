# Mazzaroth-XDR

[![CircleCI](https://circleci.com/gh/kochavalabs/mazzaroth-xdr.svg?style=svg)](https://circleci.com/gh/kochavalabs/mazzaroth-xdr)

**[Self Link](https://github.com/kochavalabs/mazzaroth-xdr)**

Mazzaroth-xdr is a library that defines the core XDR data structures used by
Mazzaroth. This includes the definitions of the blockchain, its underlying
data structures and the objects stored in the StateDB.

All objects are defined in the idl directory as `.x` files and the code for
each language is generated.

We currently support code generation for 3 languages: javascript, rust and go.
The generated code has the following dependencies:

- go: [go-xdr](https://github.com/stellar/go-xdr)
- rust: [xdr-rs-serialize](https://github.com/kochavalabs/xdr-rs-serialize)
- javascript: [xdr-js-serialize](https://github.com/kochavalabs/xdr-js-serialize)

The generated code is created with the help of [xdr-codegen](https://github.com/kochavalabs/xdr-codegen)
and should not be modified manually.

## Generating Code

If any changes are made to the idl files the code can generated by running the
npm build script after updating the xdr-codegen.

First clone or update xdr-codegen

```console
git clone git@github.com:kochavalabs/xdr-codegen.git
cd xdr-codegen
git pull
```

Then run the npm build script to generate the code for all languages:

```console
npm install
npm run build
```
