import { useState, useRef } from 'react';
import { FiUpload, FiX, FiFile } from 'react-icons/fi';

export default function Upload({ setIsUploadOpen }: {
  setIsUploadOpen: (isUploadOpen: boolean) => void,
}) {
  const [selectedFiles, setSelectedFiles] = useState<FileList | null>(null);
  const [bucketId, setBucketId] = useState<string>('');
  const [fileKey, setFileKey] = useState<string>('');
  const [isDragging, setIsDragging] = useState(false);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSelectedFiles(e.target.files);
  };

  const handleDragOver = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setIsDragging(true);
  };

  const handleDragLeave = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setIsDragging(false);
  };

  const handleDrop = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setIsDragging(false);
    
    if (e.dataTransfer.files && e.dataTransfer.files.length > 0) {
      setSelectedFiles(e.dataTransfer.files);
      if (fileInputRef.current) {
        fileInputRef.current.files = e.dataTransfer.files;
      }
    }
  };

  const handleUpload = async () => {
    if (!selectedFiles || !bucketId || !fileKey) {
      alert('Please select files and provide bucket ID and file key.');
      return;
    }

    const formData = new FormData();
    Array.from(selectedFiles).forEach(file => {
      formData.append('file', file);
    });
    formData.append('bucket_id', bucketId);
    formData.append('file_key', fileKey);

    try {
      const res = await fetch("http://localhost:8080/api/v1/file", {
        method: "POST",
        body: formData,
      });

      if (!res.ok) {
        throw new Error(`Error: ${res.status} ${res.statusText}`);
      }

      console.log('Files uploaded successfully');
      setIsUploadOpen(false);
    } catch (error) {
      console.error('Failed to upload files:', error);
    }
  };

  return (
    <div className="fixed inset-0 z-50 bg-black/70 backdrop-blur-sm flex items-center justify-center p-4">
      <div className="bg-gray-900 rounded-xl shadow-2xl max-w-2xl w-full overflow-hidden border border-gray-800">
        <div className="flex justify-between items-center p-4 border-b border-gray-800">
          <h2 className="text-lg font-semibold text-white">Upload Files</h2>
          <button
            onClick={() => setIsUploadOpen(false)}
            className="p-2 bg-gray-800 hover:bg-gray-700 rounded-full transition-colors text-gray-300"
          >
            <FiX />
          </button>
        </div>
        <div className="p-6">
          <div 
            className={`border-2 border-dashed ${isDragging ? 'border-blue-500 bg-gray-800/50' : 'border-gray-700'} rounded-lg p-8 text-center transition-colors`}
            onDragOver={handleDragOver}
            onDragLeave={handleDragLeave}
            onDrop={handleDrop}
          >
            <input
              type="file"
              className="hidden"
              id="file-upload"
              ref={fileInputRef}
              multiple
              onChange={handleFileChange}
            />
            <label
              htmlFor="file-upload"
              className="cursor-pointer flex flex-col items-center"
            >
              <FiUpload className="text-gray-400 w-12 h-12 mb-4" />
              <p className="text-gray-300 mb-2">Drag and drop files here</p>
              <p className="text-gray-400 text-sm">or click to select files</p>
            </label>
          </div>
          
          {selectedFiles && selectedFiles.length > 0 && (
            <div className="mt-4 bg-gray-800 rounded-lg p-3">
              <h3 className="text-gray-300 text-sm font-medium mb-2">Selected Files:</h3>
              <div className="max-h-32 overflow-y-auto">
                {Array.from(selectedFiles).map((file, index) => (
                  <div key={index} className="flex items-center py-2 border-b border-gray-700 last:border-0">
                    <FiFile className="text-blue-400 mr-2" />
                    <span className="text-gray-300 text-sm truncate">{file.name}</span>
                    <span className="text-gray-500 text-xs ml-2">({(file.size / 1024).toFixed(1)} KB)</span>
                  </div>
                ))}
              </div>
            </div>
          )}
          
          <div className="mt-4 space-y-4">
            <input
              type="text"
              placeholder="Bucket ID"
              value={bucketId}
              onChange={(e) => setBucketId(e.target.value)}
              className="w-full bg-gray-800 border border-gray-700 rounded-lg px-4 py-2 text-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
            <input
              type="text"
              placeholder="File Key"
              value={fileKey}
              onChange={(e) => setFileKey(e.target.value)}
              className="w-full bg-gray-800 border border-gray-700 rounded-lg px-4 py-2 text-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
        </div>
        <div className="p-4 border-t border-gray-800 flex justify-end gap-3">
          <button
            onClick={() => setIsUploadOpen(false)}
            className="px-4 py-2 text-gray-300 hover:text-white transition-colors"
          >
            Cancel
          </button>
          <button
            onClick={handleUpload}
            className="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors"
          >
            Upload
          </button>
        </div>
      </div>
    </div>
  );
}
