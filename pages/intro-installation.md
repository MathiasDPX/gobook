---
title: Installation
---

# Installation

## Prerequisites

- Go 1.22.2 or later installed on your system
- Git (for cloning the repository)

## Install GoBook

You can install GoBook in two ways:

### Option 1: Install from source

```bash
git clone https://github.com/MathiasDPX/gobook
cd gobook
go install
```

### Option 2: Direct install (recommended)

```bash
go install github.com/MathiasDPX/gobook@latest
```

## Verify Installation

Check that GoBook is installed correctly:

```bash
gobook -h
```

You should see the help output with available commands.

## Quick Start

Create your first book:

```bash
gobook init my-book
cd my-book
gobook serve
```

Open your browser and go to `http://localhost:8080` to see your book!