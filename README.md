<h1 align="center" style="border-bottom: none">
    <img src="https://github.com/user-attachments/assets/91c06e38-4a5f-4111-bd33-80fcc2bd9627" alt="bucketX Logo" width="400"><br>bucketX
</h1>

<p align="center">
  A powerful, self-hosted cloud storage solution for managing image storage, optimization, and delivery
</p>

<div align="center">

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/TeamXSeven/bucketX)](https://goreportcard.com/report/github.com/TeamXSeven/bucketX)
[![Go Version](https://img.shields.io/github/go-mod/go-version/TeamXSeven/bucketX)](https://github.com/TeamXSeven/bucketX)
<!-- [![Docker Pulls](https://img.shields.io/docker/pulls/teamxseven/bucketx)](https://hub.docker.com/r/teamxseven/bucketx)
[![Documentation](https://img.shields.io/badge/docs-website-blue)](https://teamxseven.github.io/bucketX/) -->
[![Contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg)](https://github.com/TeamXSeven/bucketX/blob/main/CONTRIBUTING.md)

</div>

## ✨ Features

bucketX offers a complete solution for managing your image assets with enterprise-grade features:

- **🔄 Intelligent Deduplication**: Automatically detects and eliminates duplicate images using perceptual hashing
- **🗜️ Advanced Compression**: Optimizes image sizes without compromising visual quality
- **🔄 Format Conversion**: Converts images to modern formats like WebP, AVIF, and JPEG XL for optimal delivery
- **✂️ Dynamic Transformations**: Performs on-the-fly resizing and smart cropping based on content awareness
- **🔐 Secure Access Control**: Granular permission system with API keys and token authentication
- **⚡ Performance Optimization**: Built-in CDN integration and caching for blazing-fast delivery
- **📊 Usage Analytics**: Track storage, bandwidth, and transformation metrics
- **🔌 API-First Design**: Comprehensive REST API with excellent documentation

## Installation

```bash
git clone https://github.com/TeamXSeven/bucketX.git
```

## Usage

```bash
cd bucketX
go mod download
air # also can use 'go run main.go'
```

### Generate Swagger Docs

```bash
swag init -g main.go --parseDependency --parseInternal
```

### Fix go dependencies tree

```bash
go mod tidy
```

### Docker

```bash
docker compose up -d
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to
discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
