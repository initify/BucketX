import { useState, useRef } from 'react';
import { FiUpload, FiX, FiFile, FiCheckCircle } from 'react-icons/fi';

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
            
            {selectedFiles && selectedFiles.length > 0 ? (
              <div className="flex flex-col items-center">
                <div className="bg-blue-500/10 rounded-full p-4 mb-4">
                  <FiCheckCircle className="text-blue-400 w-10 h-10" />
                </div>
                <p className="text-gray-200 font-medium mb-2">{selectedFiles.length} {selectedFiles.length === 1 ? 'file' : 'files'} selected</p>
                <div className="flex flex-wrap justify-center gap-2 max-w-md mb-4">
                  {Array.from(selectedFiles).slice(0, 3).map((file, index) => (
                    <div key={index} className="flex items-center bg-gray-800 rounded-full px-3 py-1">
                      <FiFile className="text-blue-400 mr-2" />
                      <span className="text-gray-300 text-sm truncate max-w-[150px]">{file.name}</span>
                    </div>
                  ))}
                  {selectedFiles.length > 3 && (
                    <div className="flex items-center bg-gray-800 rounded-full px-3 py-1">
                      <span className="text-gray-300 text-sm">+{selectedFiles.length - 3} more</span>
                    </div>
                  )}
                </div>
                <label
                  htmlFor="file-upload"
                  className="text-blue-400 hover:text-blue-300 cursor-pointer text-sm"
                >
                  Change selection
                </label>
              </div>
            ) : (
              <label
                htmlFor="file-upload"
                className="cursor-pointer flex flex-col items-center"
              >
                <FiUpload className="text-gray-400 w-12 h-12 mb-4" />
                <p className="text-gray-300 mb-2">Drag and drop files here</p>
                <p className="text-gray-400 text-sm">or click to select files</p>
              </label>
            )}
          </div>
          
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
