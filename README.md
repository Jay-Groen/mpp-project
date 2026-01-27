# MPP Project – D&D Character Manager

A Go-based application for managing Dungeons & Dragons 5e characters.  
The project supports both a **CLI interface** and an **HTML web interface**, and follows structured release management with CI, versioning, hotfixes, and rollback procedures.

---

## Features

### Character Management (CRU)
- Create new characters
- View character sheets
- Update level and proficiency bonus
- Ability scores based on Standard Array (15/14/13/12/10/8) + race modifiers
- Skill proficiency selection
- Automatic skill modifier calculation

### Equipment Management
- Add weapons, armor, and shields
- Enrich equipment data using the DnD 5e API
- Combat stat calculations (Armor Class, Initiative, Passive Perception)

### Spellcasting
- Add learned/prepared spells
- Spell slot management
- Spellcasting ability, spell save DC, and spell attack bonus

### Web Interface
- Individual character sheet view

### External API Integration
- Enrichment of spells and equipment via:
  https://www.dnd5eapi.co/
- Uses Go concurrency for parallel API requests
- Rate-limited to avoid excessive API usage

---

## Architecture

The project follows an Onion Architecture approach:

```
domain/          → Core business logic  
application/     → Use cases & services  
infrastructure/  → Database & external API  
presentation/    → CLI & Web interfaces  
```

This ensures:
- Clear separation of concerns  
- High testability  
- Maintainability  
- Structured layering  

---

## Running the Project

### Requirements
- Go 1.21+
- SQLite (uses `github.com/mattn/go-sqlite3`)

---

### Run CLI

```bash
go run . <command>
```

Available commands:

```
create
view
delete
equip
prepare-spell
learn-spell
test
html
```

Example:

```bash
go run . create
```

---

### Run Web Interface (Locally)

```bash
go run . html
```

Then open:

```
http://localhost:8081
```

---

## Running Tests

Run all unit tests:

```bash
go test ./...
```

---

## CI Pipeline

GitHub Actions automatically:

- Runs unit tests on every push
- Runs golangci-lint
- Performs static code analysis
- Validates build before releases

This ensures:
- Stable releases
- Consistent code quality
- Continuous improvement of coverage

---

## Versioning & Releases

The project follows **Semantic Versioning**:

```
MAJOR.MINOR.PATCH
```

Examples:

- `v0.2.0` → Feature release  
- `v0.2.1` → Hotfix (patch) release  
- `v0.3.0` → New feature release  

Each release includes:
- Updated documentation
- Linked issues
- Changelog entries
- GitHub release notes

---

## Hotfix & Rollback Strategy

### Hotfix
1. Bug is reported via GitHub Issue  
2. Fix implemented in a hotfix branch  
3. Patch version tagged (e.g., v0.2.1)  
4. GitHub Release published  

### Rollback
If a release is unstable:
- Revert to the last stable tag  
- Create rollback release if needed  
- Document rollback in release notes  

---

## Documentation

Documentation is updated for every release to support onboarding:

- `README.md` – Project overview & setup  
- `CHANGELOG.md` – Version history  
- `docs/development-setup.md`  
- `docs/testing.md`  
- `docs/releasing.md`  

---

## Deployment

The web application can be deployed as a **Web Service** (e.g., on Render).

When deploying:
- The application listens on the `PORT` environment variable  
- Static files are served via `/static/`  
- `/healthz` endpoint is available for monitoring  

---

## Development Workflow

1. Create user story or bug issue  
2. Assign milestone  
3. Implement feature/fix  
4. Run tests locally  
5. Merge via PR  
6. Tag release  
7. Update documentation  
8. Publish GitHub Release  
