# Mazzaroth-xdr

Mazzaroth-xdr is a library that contains the XDR objects used by Mazzaroth.

All objects are defined in the idl directory as `.x` files and the code for each language is generated.

We currently support code generation for 3 languages: javascript, rust and go.
The generated code has the following dependencies:

- go: [go-xdr](https://github.com/stellar/go-xdr)
- rust: [xdr-rs-serialize](https://github.com/kochavalabs/xdr-rs-serialize)
- javascript: [xdr-js-serialize](https://github.com/kochavalabs/xdr-js-serialize)

The generated code is created with the help of [xdr-codegen](https://github.com/kochavalabs/xdr-codegen) and should not be modified.

## Generating Code

If any changes are made to the idl files the code can generated by running the npm build script after updating the xdr-codegen.

First clone or update xdr-codegen

```console
git clone git@github.com:kochavalabs/xdr-codegen.git
cd xdr-codegen
git pull
```

Then run the npm build script to generate the code for all languagues:

```console
npm install
npm run build
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
