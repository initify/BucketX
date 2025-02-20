package utils

func FindFileType(fileExt string) string {
	if imageType := findImageType(fileExt); imageType != "" {
		return imageType
	}
	if videoType := findVideoType(fileExt); videoType != "" {
		return videoType
	}
	if audioType := findAudioType(fileExt); audioType != "" {
		return audioType
	}
	if documentType := findDocumentType(fileExt); documentType != "" {
		return documentType
	}
	if archiveType := findArchiveType(fileExt); archiveType != "" {
		return archiveType
	}
	if webType := findWebType(fileExt); webType != "" {
		return webType
	}
	return "application/octet-stream"
}

func findImageType(fileExt string) string {
	switch fileExt {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	case ".svg":
		return "image/svg+xml"
	case ".bmp":
		return "image/bmp"
	case ".tiff", ".tif":
		return "image/tiff"
	case ".ico":
		return "image/x-icon"
	default:
		return ""
	}
}

func findVideoType(fileExt string) string {
	switch fileExt {
	case ".mp4":
		return "video/mp4"
	case ".avi":
		return "video/x-msvideo"
	case ".mkv":
		return "video/x-matroska"
	case ".webm":
		return "video/webm"
	case ".mov":
		return "video/quicktime"
	case ".wmv":
		return "video/x-ms-wmv"
	case ".flv":
		return "video/x-flv"
	case ".m4v":
		return "video/x-m4v"
	default:
		return ""
	}
}

func findAudioType(fileExt string) string {
	switch fileExt {
	case ".mp3":
		return "audio/mpeg"
	case ".wav":
		return "audio/wav"
	case ".flac":
		return "audio/flac"
	case ".m4a":
		return "audio/mp4"
	case ".ogg":
		return "audio/ogg"
	case ".aac":
		return "audio/aac"
	case ".wma":
		return "audio/x-ms-wma"
	default:
		return ""
	}
}

func findDocumentType(fileExt string) string {
	switch fileExt {
	case ".txt":
		return "text/plain"
	case ".pdf":
		return "application/pdf"
	case ".doc", ".docx":
		return "application/msword"
	case ".xls", ".xlsx":
		return "application/vnd.ms-excel"
	case ".ppt", ".pptx":
		return "application/vnd.ms-powerpoint"
	case ".csv":
		return "text/csv"
	case ".rtf":
		return "application/rtf"
	default:
		return ""
	}
}

func findArchiveType(fileExt string) string {
	switch fileExt {
	case ".zip":
		return "application/zip"
	case ".rar":
		return "application/x-rar-compressed"
	case ".7z":
		return "application/x-7z-compressed"
	case ".tar":
		return "application/x-tar"
	case ".gz":
		return "application/gzip"
	default:
		return ""
	}
}

func findWebType(fileExt string) string {
	switch fileExt {
	case ".json":
		return "application/json"
	case ".xml":
		return "application/xml"
	case ".html", ".htm":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	default:
		return ""
	}
}
