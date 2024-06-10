# Changelog

## [0.3.1](https://github.com/defenseunicorns/lula/compare/v0.3.0...v0.3.1) (2024-06-10)


### Bug Fixes

* **commitlint:** pinned dependency issue by extracting into package(-… ([#454](https://github.com/defenseunicorns/lula/issues/454)) ([17ac8ca](https://github.com/defenseunicorns/lula/commit/17ac8ca76c081f2d385a8dca4bd35ced8cbbf70d))
* **generate:** resolve parent flag options properly ([#442](https://github.com/defenseunicorns/lula/issues/442)) ([5850115](https://github.com/defenseunicorns/lula/commit/585011577ffcc9426a694296d631d29bc88b1f99))
* **validate:** fix related observations when empty ([#448](https://github.com/defenseunicorns/lula/issues/448)) ([f6f6993](https://github.com/defenseunicorns/lula/commit/f6f699312b1dfaa1190292143a039c105f5d5293))


### Miscellaneous

* **actions:** fix code scanning alerts ([#446](https://github.com/defenseunicorns/lula/issues/446)) ([aa568c7](https://github.com/defenseunicorns/lula/commit/aa568c76f9667195a0f106a4762f3ff315dc1666))
* **deps:** update dependency linkinator to v6.0.5 ([#458](https://github.com/defenseunicorns/lula/issues/458)) ([dfa1cbe](https://github.com/defenseunicorns/lula/commit/dfa1cbef30c40ccdba1b46915fd63aebc080f6e0))
* **deps:** update dependency prettier to v3.3.1 ([#459](https://github.com/defenseunicorns/lula/issues/459)) ([c999b78](https://github.com/defenseunicorns/lula/commit/c999b78c945e73ea1c5e93d269341e7c2c953e17))
* **deps:** update github/codeql-action action to v3.25.7 ([#452](https://github.com/defenseunicorns/lula/issues/452)) ([2583eea](https://github.com/defenseunicorns/lula/commit/2583eea6200f305c74ae214576188c424d39f749))
* **deps:** update goreleaser/goreleaser-action action to v6 ([#464](https://github.com/defenseunicorns/lula/issues/464)) ([e74b9d5](https://github.com/defenseunicorns/lula/commit/e74b9d53add588a4ddb83138b04e4e9ca263dae6))
* **deps:** update module github.com/kyverno/kyverno-json to v0.0.3 ([#453](https://github.com/defenseunicorns/lula/issues/453)) ([1dc96e8](https://github.com/defenseunicorns/lula/commit/1dc96e832aa29a1d126a2b4f4d620aa006c8fec1))
* **deps:** update module github.com/open-policy-agent/opa to v0.65.0 ([#451](https://github.com/defenseunicorns/lula/issues/451)) ([7867a3c](https://github.com/defenseunicorns/lula/commit/7867a3c36826eaa264106434f993f37597ffd8e7))

## [0.3.0](https://github.com/defenseunicorns/lula/compare/v0.2.1...v0.3.0) (2024-05-24)


### ⚠ BREAKING CHANGES

* #388 update types to use pointers ([#410](https://github.com/defenseunicorns/lula/issues/410))
* #367 compiling external/remote validations ([#384](https://github.com/defenseunicorns/lula/issues/384))

### refactor

* [#388](https://github.com/defenseunicorns/lula/issues/388) update types to use pointers ([#410](https://github.com/defenseunicorns/lula/issues/410)) ([9c51d56](https://github.com/defenseunicorns/lula/commit/9c51d5693565022f353c2739c97fac2686d78ce4))


### Features

* [#367](https://github.com/defenseunicorns/lula/issues/367) compiling external/remote validations ([#384](https://github.com/defenseunicorns/lula/issues/384)) ([8bb42b0](https://github.com/defenseunicorns/lula/commit/8bb42b02f6da2670830f11a1d2e1e5367c2b7d09))
* **oscal:** merge on write consolidation ([#407](https://github.com/defenseunicorns/lula/pull/407)) ([ef2f9f5](https://github.com/defenseunicorns/lula/commit/ef2f9f536ac8809785ca03325ae56575bbe0361c))
* **compose:** add ability to pull and compose import component defs ([#406](https://github.com/defenseunicorns/lula/pull/406)) ([ddf919a](https://github.com/defenseunicorns/lula/commit/ddf919a43995703f782f8667eb7305363b95b3cd))
* **generate:** add generate command and initial component generation ([#401](https://github.com/defenseunicorns/lula/issues/401)) ([918299a](https://github.com/defenseunicorns/lula/commit/918299ad397363e0d3580cd15b92ddf09044ce05))
* **dev:** added Observation logging to dev validate ([#396](https://github.com/defenseunicorns/lula/pull/396)) ([c32027e](https://github.com/defenseunicorns/lula/commit/c32027eafbf65e1cf69b3a72acac3a51c4f35656))
* **dev:** dev validate with optional resources file input ([#394](https://github.com/defenseunicorns/lula/pull/394)) ([f034a97](https://github.com/defenseunicorns/lula/commit/f034a976d20d10fe5ec660433e4554a02f76158c))
* **validate:** validation store/cache ([#373](https://github.com/defenseunicorns/lula/issues/373)) ([751982f](https://github.com/defenseunicorns/lula/commit/751982f5d4eee60a6412eed5e554c86a683ecb7a))


### Bug Fixes

* **deps:** consolidate use of goyaml pkg ([#422](https://github.com/defenseunicorns/lula/issues/422)) ([d1abbcc](https://github.com/defenseunicorns/lula/commit/d1abbcc052fd1f1ff2e57265e54a1b005ec66641))
* **deps:** controller runtime ([#379](https://github.com/defenseunicorns/lula/issues/379)) ([7d3aec3](https://github.com/defenseunicorns/lula/commit/7d3aec3e7e94652d524d1e40d62c61736ca1e12e))
* **dev:** updated result condition to match satisfaction logic ([#400](https://github.com/defenseunicorns/lula/issues/400)) ([5feda9d](https://github.com/defenseunicorns/lula/commit/5feda9dde93b270e9d2cebc5ee40ec21ab4b1c4c))
* **validate:** validation errors mapped to observations ([#405](https://github.com/defenseunicorns/lula/pull/405)) ([39e5ebd](https://github.com/defenseunicorns/lula/commit/39e5ebd45d4a9cfc0918899ed647192d8bcf0952))
* **validate:** fix order of assessment-results results ([#437](https://github.com/defenseunicorns/lula/issues/437)) ([a8db208](https://github.com/defenseunicorns/lula/commit/a8db20862f9f1bf7c269cd75839823f58b7c9541))


### Miscellaneous

* **actions:** [#420](https://github.com/defenseunicorns/lula/issues/420) update release process with release-please ([#421](https://github.com/defenseunicorns/lula/issues/421)) ([a372df0](https://github.com/defenseunicorns/lula/commit/a372df0e0316067a3adb62d02c95433d37930ec5))
* **deps:** bump golang.org/x/net from 0.22.0 to 0.23.0 ([#378](https://github.com/defenseunicorns/lula/issues/378)) ([8088bd0](https://github.com/defenseunicorns/lula/commit/8088bd0d38c89768a7ee4eae7e12edea3ce4f35e))
* **deps:** Update actions/checkout action to v4.1.3 ([#382](https://github.com/defenseunicorns/lula/issues/382)) ([08eed39](https://github.com/defenseunicorns/lula/commit/08eed39078dc38e79b5b5483b2cdd0770a00594d))
* **deps:** Update actions/download-artifact action to v4.1.6 ([#376](https://github.com/defenseunicorns/lula/issues/376)) ([2982b36](https://github.com/defenseunicorns/lula/commit/2982b3659d00671bdd6eac0a29ad25aca0e7da30))
* **deps:** Update actions/download-artifact action to v4.1.7 ([#387](https://github.com/defenseunicorns/lula/issues/387)) ([92064e6](https://github.com/defenseunicorns/lula/commit/92064e6184581c116d0fd4fd07521c852211ebcf))
* **deps:** Update actions/upload-artifact action to v4.3.2 ([#377](https://github.com/defenseunicorns/lula/issues/377)) ([f575f82](https://github.com/defenseunicorns/lula/commit/f575f82a3b00d2c5260b0887391638faa427b8ee))
* **deps:** Update actions/upload-artifact action to v4.3.3 ([#383](https://github.com/defenseunicorns/lula/issues/383)) ([26f1f32](https://github.com/defenseunicorns/lula/commit/26f1f32d54e1fe64f528ca3441cbd39055dc8ee2))
* **deps:** update anchore/sbom-action action to v0.16.0 ([#426](https://github.com/defenseunicorns/lula/issues/426)) ([a2063a5](https://github.com/defenseunicorns/lula/commit/a2063a5d47e9d816901cda1f908c6e6d68b53442))
* **deps:** update github/codeql-action action to v3.25.6 ([#425](https://github.com/defenseunicorns/lula/issues/425)) ([9ef1703](https://github.com/defenseunicorns/lula/commit/9ef1703f0180f05bec97b1fa5f894fae5d4627f6))
* **deps:** update golang to version 1.22.3 ([#423](https://github.com/defenseunicorns/lula/issues/423)) ([aa8cab7](https://github.com/defenseunicorns/lula/commit/aa8cab732053ed588ec7c6d83895e3d6f0ecf7f4))
* **deps:** update kubernetes packages to v0.30.1 ([#417](https://github.com/defenseunicorns/lula/issues/417)) ([e47a04d](https://github.com/defenseunicorns/lula/commit/e47a04d4df9a9f3f14a157716120cf5f12714d5c))
* **deps:** Update module github.com/defenseunicorns/go-oscal to v0.3.2 ([#380](https://github.com/defenseunicorns/lula/issues/380)) ([03aa969](https://github.com/defenseunicorns/lula/commit/03aa969ff102111325c0b45b91e7c3543c15cf16))
* **deps:** update module github.com/defenseunicorns/go-oscal to v0.4.0 ([#429](https://github.com/defenseunicorns/lula/issues/429)) ([4ff7775](https://github.com/defenseunicorns/lula/commit/4ff7775890113407c60240ca21382dfca0eb102c))
* **deps:** update module github.com/defenseunicorns/go-oscal to v0.4.1 ([#435](https://github.com/defenseunicorns/lula/issues/435)) ([4570658](https://github.com/defenseunicorns/lula/commit/4570658fcc3d20d7c4d118e89626b5d81267af91))
* **deps:** update module github.com/hashicorp/go-version to v1.7.0 ([#438](https://github.com/defenseunicorns/lula/issues/438)) ([4f6de9b](https://github.com/defenseunicorns/lula/commit/4f6de9b0cff92530a6d0e34697bfa694d4e86f33))
* **deps:** update module sigs.k8s.io/e2e-framework to v0.4.0 ([#419](https://github.com/defenseunicorns/lula/issues/419)) ([890a7d8](https://github.com/defenseunicorns/lula/commit/890a7d8d7a5da123b6dff23b75e1b390ff7ca716))
* **renovate:** update config to handle conventional commit titles ([#428](https://github.com/defenseunicorns/lula/issues/428)) ([5f4139a](https://github.com/defenseunicorns/lula/commit/5f4139a3b6df6fd5ba4c1ee7f4e04dd50f23be1f))
