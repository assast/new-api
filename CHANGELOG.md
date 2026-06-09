# Changelog

## 2026-06-09

### Added

- Add a manual GitHub Container Registry image build workflow.
- Support converting OpenAI-compatible file content into Claude text and PDF document inputs.

### Fixed

- Preserve client-visible model names for mapped Claude relay responses while keeping upstream model names for internal usage accounting.
- Preserve client-visible model names when converting OpenAI-compatible responses to Claude format.
- Patch Claude streaming `message_delta` usage fields when upstream providers omit prompt/cache usage details.
