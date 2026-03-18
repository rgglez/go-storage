# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/)
and this project adheres to [Semantic Versioning](https://semver.org/).

## v0.4.0 - 2021-10-23

- feat(services/memory): Move services memory back (#912)
- ci(*): Upgrade minimum version to Go 1.16 (#916)

## [v0.3.0] - 2021-09-13

### Added

- feat: Implement IoCallback support (#23)
- feat: Add storage_features and default_storage_pairs support

### Changed

- ci: Enable auto merge for Dependabot
- ci: Upgrade fetch-metadata
- docs: Update README (#27)

## [v0.2.0] - 2021-08-23

### Added

- ci: Add intergration tests (#15)
- feat: Add support for Copier, Mover, Appender, Direr (#14)

### Fixed

- fix: stat root path will return ErrObjectNotExist (#17)
- fix: Object size calculated incorrectly while short write (#21)

### Refactor

- refactor: Add name field in object (#19)

## v0.1.0 - 2021-07-26

### Added

- Implement memory services.

[v0.3.0]: https://github.com/rgglez/go-service-memory/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/rgglez/go-service-memory/compare/v0.1.0...v0.2.0 


## License

Copyright 2024 go-storage authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
