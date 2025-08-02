---
title: gobook build
---

# `gobook build`

Compiles your GoBook into static HTML files ready for deployment.

## Usage

```bash
gobook build
```

## Description

The `build` command transforms all your Markdown files into a complete static website. It processes all pages in the `pages/` directory, applies templates, generates navigation, and creates a production-ready site in the `_book/` directory.

### Build Process

1. **Pre-build setup**: Loads configuration and templates
2. **Page processing**: Converts Markdown to HTML
3. **Template application**: Applies HTML templates to each page
4. **Navigation generation**: Creates sidebar from `SUMMARY.md`
5. **Asset copying**: Includes stylesheets and other resources
6. **File output**: Generates static files in `_book/`

### Output Structure

After running `gobook build`, you'll have:

```
_book/
├── index.html          # Home page (from INDEX.md)
├── about.html          # Other pages (from *.md files)
├── guide.html
└── style.css
```