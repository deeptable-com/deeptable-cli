# Changelog

## 0.1.0-beta.3 (2026-04-30)

Full Changelog: [v0.1.0-beta.2...v0.1.0-beta.3](https://github.com/deeptable-com/deeptable-cli/compare/v0.1.0-beta.2...v0.1.0-beta.3)

### Features

* allow `-` as value representing stdin to binary-only file parameters in CLIs ([e30bb3a](https://github.com/deeptable-com/deeptable-cli/commit/e30bb3a93f299cc74a5fd022fd6d7bc4ad16bebf))
* better error message if scheme forgotten in CLI `*_BASE_URL`/`--base-url` ([3ee457a](https://github.com/deeptable-com/deeptable-cli/commit/3ee457aa17ed94997cc134b99c77e28a85900a82))
* binary-only parameters become CLI flags that take filenames only ([482d722](https://github.com/deeptable-com/deeptable-cli/commit/482d722a794806e5c60ab124dee0459170397bc9))
* **cli:** add `--raw-output`/`-r` option to print raw (non-JSON) strings ([177c28f](https://github.com/deeptable-com/deeptable-cli/commit/177c28f629b61bb326788545144976e2e54e46a3))
* **cli:** alias parameters in data with `x-stainless-cli-data-alias` ([9a763d3](https://github.com/deeptable-com/deeptable-cli/commit/9a763d3b2a19e9e1fb60cab9da8a9562a0760177))
* **cli:** send filename and content type when reading input from files ([4116cf0](https://github.com/deeptable-com/deeptable-cli/commit/4116cf01842502beea08e26368f9d99a5049bb06))
* set CLI flag constant values automatically where `x-stainless-const` is set ([d5a50f3](https://github.com/deeptable-com/deeptable-cli/commit/d5a50f395b1a844f07824fb1f2e4bdfcd22e354c))
* support passing path and query params over stdin ([8e27218](https://github.com/deeptable-com/deeptable-cli/commit/8e27218906f07ebd77ee8ba790eec504bbcfbe7f))


### Bug Fixes

* cli no longer hangs when stdin is attached to a pipe with empty input ([0063fc1](https://github.com/deeptable-com/deeptable-cli/commit/0063fc108a77d22495846cc6106592c22761792c))
* **cli:** correctly load zsh autocompletion ([9e91ddf](https://github.com/deeptable-com/deeptable-cli/commit/9e91ddf6fb36e634f56a23a9a01102e548073eb4))
* fall back to main branch if linking fails in CI ([ec4eb09](https://github.com/deeptable-com/deeptable-cli/commit/ec4eb09e2f76d4a75fdfec1f2b407d9ad8afb67a))
* fix for failing to drop invalid module replace in link script ([b3ea45a](https://github.com/deeptable-com/deeptable-cli/commit/b3ea45a6f1c7bc4111bf68507f061b642d6bb789))
* fix for off-by-one error in pagination logic ([d7142f9](https://github.com/deeptable-com/deeptable-cli/commit/d7142f9c3552886983d8590e9b47df4b774ffb49))
* fix quoting typo ([006840d](https://github.com/deeptable-com/deeptable-cli/commit/006840d8fc41deeb53de45b2ab1c18fd877480aa))
* flags for nullable body scalar fields are strictly typed ([a98eb25](https://github.com/deeptable-com/deeptable-cli/commit/a98eb255536c876836413ba333773b7b54929fb9))
* handle empty data set using `--format explore` ([76e8aa7](https://github.com/deeptable-com/deeptable-cli/commit/76e8aa73ea6dbaceb088b835888f26688dc6f360))
* use `RawJSON` when iterating items with `--format explore` in the CLI ([0806fe8](https://github.com/deeptable-com/deeptable-cli/commit/0806fe8ed9e8355d4f06f028de673e541d091105))


### Chores

* add documentation for ./scripts/link ([a8008f2](https://github.com/deeptable-com/deeptable-cli/commit/a8008f2a7b3a064ad21d86ab8800d6696c8d455e))
* **ci:** skip lint on metadata-only changes ([0be1028](https://github.com/deeptable-com/deeptable-cli/commit/0be1028ca0f24e44cd5fe74b046ad514b47d08a4))
* **ci:** support manually triggering release workflow ([92efacd](https://github.com/deeptable-com/deeptable-cli/commit/92efacd686131b1245b8f25467daaba4c650d602))
* **cli:** additional test cases for `ShowJSONIterator` ([f882462](https://github.com/deeptable-com/deeptable-cli/commit/f8824621a6ffc592f70e7cefe0c0ded4ce892888))
* **cli:** fall back to JSON when using default "explore" with non-TTY ([22db507](https://github.com/deeptable-com/deeptable-cli/commit/22db507c72c8578edaea3d1295d08b22229430bd))
* **cli:** let `--format raw` be used in conjunction with `--transform` ([eca1743](https://github.com/deeptable-com/deeptable-cli/commit/eca1743bead6e107e3c839d96245005d1401a0c7))
* **cli:** switch long lists of positional args over to param structs ([5d7b97a](https://github.com/deeptable-com/deeptable-cli/commit/5d7b97aad32a194a00a4f43c95197bbd018413b2))
* **cli:** use `ShowJSONOpts` as argument to `formatJSON` instead of many positionals ([8eb757d](https://github.com/deeptable-com/deeptable-cli/commit/8eb757d4dc41a21f87cdeda972597ada022a3580))
* **internal:** codegen related update ([ab0b024](https://github.com/deeptable-com/deeptable-cli/commit/ab0b024de44d535fb1402bdfd82f048d943ff65d))
* **internal:** codegen related update ([636523c](https://github.com/deeptable-com/deeptable-cli/commit/636523c180805486a19a0687901d9b4c27212e19))
* **internal:** codegen related update ([6241318](https://github.com/deeptable-com/deeptable-cli/commit/6241318a939febb87359001b98fe890de2797099))
* **internal:** codegen related update ([aad810e](https://github.com/deeptable-com/deeptable-cli/commit/aad810e7b632dd112c095cf2d428c23ee4a96a99))
* **internal:** more robust bootstrap script ([3e2edb5](https://github.com/deeptable-com/deeptable-cli/commit/3e2edb59bd63cdb77f3128feaf7c1a68e1b0c77c))
* **internal:** update gitignore ([428cb27](https://github.com/deeptable-com/deeptable-cli/commit/428cb27be73809d9bb17a02134a302a08a0f9520))
* mark all CLI-related tests in Go with `t.Parallel()` ([f895692](https://github.com/deeptable-com/deeptable-cli/commit/f8956929608c4ee645de258a11fc982cdf1f35e6))
* modify CLI tests to inject stdout so mutating `os.Stdout` isn't necessary ([1fae171](https://github.com/deeptable-com/deeptable-cli/commit/1fae171f023fc2529341f3b2ee4704dceefeeff1))
* omit full usage information when missing required CLI parameters ([33fb84c](https://github.com/deeptable-com/deeptable-cli/commit/33fb84c343af94a2a673652f9da1655f9ea20d2e))
* switch some CLI Go tests from `os.Chdir` to `t.Chdir` ([1c11dc7](https://github.com/deeptable-com/deeptable-cli/commit/1c11dc73b06973991ce1c129a2010e54bf4f599e))

## 0.1.0-beta.2 (2026-03-19)

Full Changelog: [v0.1.0-beta.1...v0.1.0-beta.2](https://github.com/deeptable-com/deeptable-cli/compare/v0.1.0-beta.1...v0.1.0-beta.2)

### Bug Fixes

* avoid reading from stdin unless request body is form encoded or json ([e168622](https://github.com/deeptable-com/deeptable-cli/commit/e168622e79411b925ab38e911560a5d459e7b922))
* improve linking behavior when developing on a branch not in the Go SDK ([9980d7a](https://github.com/deeptable-com/deeptable-cli/commit/9980d7a4948747695738d43a6e0e64dc7ea6d914))


### Chores

* ignore macOS .DS_Store ([d17f356](https://github.com/deeptable-com/deeptable-cli/commit/d17f356671441e706be01cc8b039dface0bccf09))

## 0.1.0-beta.1 (2026-03-18)

Full Changelog: [v0.0.1...v0.1.0-beta.1](https://github.com/deeptable-com/deeptable-cli/compare/v0.0.1...v0.1.0-beta.1)

### Features

* **api:** manual updates ([795ee8d](https://github.com/deeptable-com/deeptable-cli/commit/795ee8d40763c4f3252e4db38a0ee87ae869dae0))


### Bug Fixes

* better support passing client args in any position ([e81495a](https://github.com/deeptable-com/deeptable-cli/commit/e81495a55aa32fb41abe246fc2b0c102d8d8724e))
* improved workflow for developing on branches ([dc933ae](https://github.com/deeptable-com/deeptable-cli/commit/dc933aed0bbd3574566b10a0824c1205d4bdb924))
* no longer require an API key when building on production repos ([21d2dd9](https://github.com/deeptable-com/deeptable-cli/commit/21d2dd9f25472c0414079b92ca2e8f17691f1c5c))


### Chores

* configure new SDK language ([060a31a](https://github.com/deeptable-com/deeptable-cli/commit/060a31a94880eaa4b90cab441f2c645b513b93b4))
* **internal:** tweak CI branches ([86d668d](https://github.com/deeptable-com/deeptable-cli/commit/86d668d07914caef5bbd4e2a9665c65152346276))
* update SDK settings ([6723ba5](https://github.com/deeptable-com/deeptable-cli/commit/6723ba5cc2f2de5ab6c7d40ab84e52047e0d22e4))
