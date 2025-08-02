---
title: gobook init-template
---

# `gobook init-template`

Initializes custom templates in your GoBook project by extracting the default templates.

## Usage

```bash
gobook init-template [path]
```

## Arguments

- `path` (optional): Directory where the templates will be created. If not specified, templates will be created in the current directory.

## Description

The `init-template` command extracts GoBook's default templates into a `template/` folder in your project, allowing you to customize the appearance and layout of your book. Once extracted, GoBook will automatically use your custom templates instead of the built-in defaults.

### Files Created

```
template/
├── index.html
└── style.css
```