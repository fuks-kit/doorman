# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Fixed

- Fixed side channel attack vulnerability

## [2.1.0] - 2023-11-24

### Added

- Implemented access log for fuks app

### Changed

- Changed server port from 50051 to 44888

## [2.0.0]

### Changed

- Set open door time to 4 seconds

### Fixed

- Fix door cmd bug that caused the door to open and close immediately

## [1.1.4] - 2023-11-31

### Fixed

- Fix tls error in `doorman.service`

### Added

- Some logs

## [1.1.3] - 2023-10-31

### Added

- Some error logs

## [1.1.2] - 2023-10-24

### Security

- Fix dependency vulnerabilities