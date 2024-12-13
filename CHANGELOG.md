# Changelog

## [1.1.0](https://github.com/vergissberlin/thinkport-api/compare/v1.0.0...v1.1.0) (2024-12-13)


### Features

* **member.go:** add hash field to MemberStruct to uniquely identify members ([8ddbe01](https://github.com/vergissberlin/thinkport-api/commit/8ddbe013729f990f814d175df9b149b9bcf8f318))
* **member.go:** update MembersCount to include CEO in manager count ([8ddbe01](https://github.com/vergissberlin/thinkport-api/commit/8ddbe013729f990f814d175df9b149b9bcf8f318))
* **openapi.json:** add production environment URL for Encore ([060da72](https://github.com/vergissberlin/thinkport-api/commit/060da727901b9638f064492376344a504c75872a))
* **utility.go:** add hashMember function to generate unique hash for members ([8ddbe01](https://github.com/vergissberlin/thinkport-api/commit/8ddbe013729f990f814d175df9b149b9bcf8f318))


### Bug Fixes

* **build.yml:** update encore login command to use --auth-key instead of --token for authentication ([656411c](https://github.com/vergissberlin/thinkport-api/commit/656411c15f96b1f2d9c86f09a4dfca9e20084214))
* **encore.gen.go:** change Member function parameter from name to hash for unique identification ([8ddbe01](https://github.com/vergissberlin/thinkport-api/commit/8ddbe013729f990f814d175df9b149b9bcf8f318))
