---
title: gobook init
---

# `gobook init`

Creates a new GoBook project with the necessary basic structure.

## Usage

```bash
gobook init [path]
```

## Arguments

- `path` (optional): Directory where the project will be created. If not specified, the project will be created in the current directory.

## Description

The `init` command creates a minimal structure for your new book. During execution, GoBook will prompt you for the book name, which will be automatically added to the `_site.yml` configuration file.

### Files Created

```
book/
├── _site.yml
└── pages/
    └── INDEX.md
```

- **`_site.yml`**: YAML configuration file containing the site name and other settings
- **`pages/`**: Directory containing all your Markdown pages
- **`INDEX.md`**: Home page of your book with a welcome message