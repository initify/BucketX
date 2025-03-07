"use client";
import { useEffect, useState } from "react";
import { FiFile, FiImage, FiFileText, FiVideo, FiMusic, FiArchive, FiCode } from "react-icons/fi";

interface File {
  FileKey: string;
  FileMetadata: {
    BucketId: string;
    Filename: string;
    FileType: string;
    Hash: string;
  }
}

export default function Files({ setFilekey, selectedBucket }: {
  setFilekey: (filekey: string) => void,
  selectedBucket: string
}) {
  const [files, setFiles] = useState<File[]>([]);
  const [allFiles, setAllFiles] = useState<File[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchFiles = async () => {
      setLoading(true);
      try {
        const res = await fetch("http://localhost:8080/api/v1/files");
        if (!res.ok) {
          throw new Error(`HTTP error! status: ${res.status}`);
        }
        const data = await res.json();
        setAllFiles(data.files);
        setFiles(data.files);
      } catch (err) {
        setAllFiles([]);
        setFiles([]);
      } finally {
        setLoading(false);
      }
    };
    fetchFiles();
  }, []);

  useEffect(() => {
    if (selectedBucket) {
      const filteredFiles = allFiles.filter(
        file => file.FileMetadata.BucketId === selectedBucket
      );
      setFiles(filteredFiles);
    } else {
      setFiles(allFiles);
    }
  }, [selectedBucket, allFiles]);

  if (!selectedBucket) return null;

  const getFileIcon = (mimeType: string) => {
    const type = mimeType.split('/')[0];
    switch (type) {
      case "image":
        return <FiImage className="text-green-400" />;
      case "video":
        return <FiVideo className="text-purple-400" />;
      case "audio":
        return <FiMusic className="text-yellow-400" />;
      case "application":
        if (mimeType.includes("pdf") || mimeType.includes("msword") || mimeType.includes("vnd.ms-excel") || mimeType.includes("vnd.ms-powerpoint")) {
          return <FiFileText className="text-blue-400" />;
        }
        if (mimeType.includes("zip") || mimeType.includes("x-rar-compressed") || mimeType.includes("x-7z-compressed") || mimeType.includes("x-tar") || mimeType.includes("gzip")) {
          return <FiArchive className="text-orange-400" />;
        }
        if (mimeType.includes("json") || mimeType.includes("xml") || mimeType.includes("javascript")) {
          return <FiCode className="text-pink-400" />;
        }
        break;
      case "text":
        return <FiFileText className="text-blue-400" />;
      default:
        return <FiFile className="text-gray-400" />;
    }
  };

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
      </div>
    );
  }

  return (
    <div>
      {files.length > 0 ? (
        <div className="overflow-x-auto">
          <table className="w-full text-left">
            <thead>
              <tr className="border-b border-gray-700">
                <th className="px-4 py-3 text-sm font-medium text-gray-400">Type</th>
                <th className="px-4 py-3 text-sm font-medium text-gray-400">Filename</th>
                <th className="px-4 py-3 text-sm font-medium text-gray-400">File Key</th>
                <th className="px-4 py-3 text-sm font-medium text-gray-400">Hash</th>
              </tr>
            </thead>
            <tbody>
              {files.map((file, i) => (
                <tr
                  key={i}
                  className="border-b border-gray-700 hover:bg-gray-700 cursor-pointer transition-colors"
                  onClick={() => setFilekey(file.FileKey)}
                >
                  <td className="px-4 py-4">
                    <div className="p-2 bg-gray-800 rounded-lg inline-block">
                      {getFileIcon(file.FileMetadata.FileType)}
                    </div>
                  </td>
                  <td className="px-4 py-4 text-white font-medium">{file.FileMetadata.Filename}</td>
                  <td className="px-4 py-4 text-gray-400">{file.FileKey}</td>
                  <td className="px-4 py-4 text-gray-400 font-mono text-xs">{file.FileMetadata.Hash.substring(0, 10)}...</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      ) : (
        <div className="text-center py-10">
          <div className="inline-flex items-center justify-center w-16 h-16 rounded-full bg-gray-700 mb-4">
            <FiFile className="text-gray-400 text-2xl" />
          </div>
          <h3 className="text-lg font-medium text-white mb-2">No files found</h3>
          <p className="text-gray-400">Upload files to this bucket to get started</p>
        </div>
      )}
    </div>
  );
}
