<h1 align="center" style="border-bottom: none">
    <img src="https://github.com/user-attachments/assets/3ace43e2-d175-4d7f-8e4d-0df36b586f83" alt="bucketX Logo" width="200"><br>Bucket X
</h1>

<p align="center">
  A powerful, self-hosted cloud storage solution for managing file storage, optimization, and delivery
</p>

<div align="center">

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/TeamXSeven/bucketX)](https://goreportcard.com/report/github.com/TeamXSeven/bucketX)
[![Go Version](https://img.shields.io/github/go-mod/go-version/TeamXSeven/bucketX)](https://github.com/TeamXSeven/bucketX)
[![Documentation](https://img.shields.io/badge/docs-website-blue)](https://teamxseven.github.io/bucketX/)
[![Contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg)](https://github.com/TeamXSeven/bucketX/blob/main/CONTRIBUTING.md)
<!-- [![Docker Pulls](https://img.shields.io/docker/pulls/teamxseven/bucketx)](https://hub.docker.com/r/teamxseven/bucketx) -->

</div>

## ✨ Features

bucketX offers a complete solution for managing your image assets with enterprise-grade features:

- Intelligent Deduplication**: Automatically detects and eliminates duplicate images using perceptual hashing
- Advanced Compression**: Optimizes image sizes without compromising visual quality
- Format Conversion**: Converts images to modern formats like WebP, AVIF, and JPEG XL for optimal delivery
- Dynamic Transformations**: Performs on-the-fly resizing and smart cropping based on content awareness
- Secure Access Control**: Granular permission system with API keys and token authentication
- Performance Optimization**: Built-in CDN integration and caching for blazing-fast delivery
- Usage Analytics**: Track storage, bandwidth, and transformation metrics
- API-First Design**: Comprehensive REST API with excellent documentation

<h1 align="center" style="border-bottom: none">
    <img src="https://github.com/user-attachments/assets/2ccfea75-8953-4823-ac4a-a3b2744be5df" alt="diagram">
</h1>

## Dashboard Interface

The dashboard serves as the central interface for managing storage operations. From here, you can:

- **Upload Images & Files** – Easily upload and manage your objects within different buckets.  
- **Create & Manage Buckets** – Organize your stored data efficiently by creating and managing buckets.  
- **Generate Access Keys** – Create secure access keys for authentication and permission management.  
- **View Existing Objects** – Browse through all stored objects and their metadata in a structured view.  
- **Manage User Access & Permissions** – Control who can access specific objects and perform operations.  

This intuitive interface allows seamless interaction with your storage system, providing a streamlined experience for users.

<h1 align="center" style="border-bottom: none">
    <img src="https://github.com/user-attachments/assets/ca628016-f1fc-4901-9344-61d6f19cdc97" alt="dashboard">
</h1>

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

### Installation

```bash
curl -fsSL https://raw.githubusercontent.com/initify/BucketX/main/install.sh | bash
```

https://github.com/user-attachments/assets/51a19bcb-2049-4a82-b110-16810c4853ba


## Contributing

Pull requests are welcome. For major changes, please open an issue first to
discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
