# bucketX

<img src="https://github.com/user-attachments/assets/a862a6da-992b-4a6b-afc0-69a2874ae9c0" alt="bucketX Image">

Self hosted cloud storage

A Powerful, open-source solution for managing image storage, optimization, and
delivery

## Core Features

- Deduplication: Automatically detects and eliminates duplicate images.

- Compression: Optimize image sizes without compromising quality.

- Format Conversion: Convert images to modern formats like WebP, AVIF, and more.

- Resizing and Cropping: On-the-fly transformations with smart cropping.

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
