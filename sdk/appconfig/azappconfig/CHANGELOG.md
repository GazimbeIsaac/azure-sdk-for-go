# Release History

## 0.3.0 (Unreleased)

### Features Added

### Breaking Changes
* Changed argument semantics of `AddSetting`, `DeleteSetting`, `GetSetting`, `SetSetting`, and `SetReadOnly`.

### Bugs Fixed

### Other Changes

## 0.2.0 (2022-04-20)

### Breaking Changes
* Upgraded to latest `azcore` which requires Go 1.18 or later.
* Renamed method `ListRevisions` to `NewListRevisionsPager` and removed `ListRevisionsPager` type.

### Other Changes
* Regenerated internal code with latest code generator.

### Bugs Fixed
* Fixed authentication in Germany West Central using connection string (#17424).

## 0.1.0 (2022-03-09)

### Features Added
* Initial release
