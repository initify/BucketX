"use client";
import Image from 'next/image';
import { useState, useEffect } from 'react';
import { FiX, FiDownload, FiZoomIn, FiZoomOut, FiRotateCw } from 'react-icons/fi';

export default function Viewer({ fileKey, setFilekey }:
  {
    fileKey: string,
    setFilekey: (filekey: string) => void
  }
) {
  const [zoom, setZoom] = useState(1);
  const [rotation, setRotation] = useState(0);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    setIsLoading(true);
  }, [fileKey]);

  if (!fileKey) return null;

  const handleZoomIn = () => setZoom(prev => Math.min(prev + 0.25, 3));
  const handleZoomOut = () => setZoom(prev => Math.max(prev - 0.25, 0.5));
  const handleRotate = () => setRotation(prev => (prev + 90) % 360);

  return (
    <div className="fixed inset-0 z-50 bg-black/70 backdrop-blur-sm flex items-center justify-center p-4">
      <div className="bg-gray-900 rounded-xl shadow-2xl max-w-5xl w-full overflow-hidden border border-gray-800">
        <div className="flex justify-between items-center p-4 border-b border-gray-800">
          <h2 className="text-lg font-semibold text-white">File Preview</h2>
          <div className="flex items-center gap-3">
            <button
              onClick={handleZoomOut}
              className="p-2 bg-gray-800 hover:bg-gray-700 rounded-full transition-colors text-gray-300"
              title="Zoom out"
            >
              <FiZoomOut />
            </button>
            <span className="text-gray-400 text-sm">{Math.round(zoom * 100)}%</span>
            <button
              onClick={handleZoomIn}
              className="p-2 bg-gray-800 hover:bg-gray-700 rounded-full transition-colors text-gray-300"
              title="Zoom in"
            >
              <FiZoomIn />
            </button>
            <button
              onClick={handleRotate}
              className="p-2 bg-gray-800 hover:bg-gray-700 rounded-full transition-colors text-gray-300"
              title="Rotate"
            >
              <FiRotateCw />
            </button>
            <button
              onClick={() => window.open(`http://localhost:8080/api/v1/file/${fileKey}`, '_blank')}
              className="p-2 bg-gray-800 hover:bg-gray-700 rounded-full transition-colors text-gray-300"
              title="Download"
            >
              <FiDownload />
            </button>
            <button
              onClick={() => setFilekey('')}
              className="p-2 bg-gray-800 hover:bg-gray-700 rounded-full transition-colors text-gray-300"
              title="Close"
            >
              <FiX />
            </button>
          </div>
        </div>
        <div className="p-6 bg-gray-800 flex items-center justify-center" style={{ height: '70vh' }}>
          {isLoading && (
            <div className="absolute inset-0 flex items-center justify-center bg-gray-900 bg-opacity-50 z-10">
              <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
            </div>
          )}
          <div
            className="relative w-full h-full flex items-center justify-center"
            style={{
              transform: `scale(${zoom}) rotate(${rotation}deg)`,
              transition: 'transform 0.3s ease-in-out'
            }}
          >
            <div className="relative w-4/5 h-4/5">
              <Image
                src={`http://localhost:8080/api/v1/file/${fileKey}`}
                alt="File Preview"
                fill
                className="object-contain"
                onLoadingComplete={() => setIsLoading(false)}
                priority
                unoptimized={true}
              />
            </div>
          </div>
        </div>
        <div className="p-4 border-t border-gray-800 flex justify-between items-center">
          <div className="text-gray-400 text-sm truncate">
            File Key: <span className="text-blue-400 font-mono">{fileKey}</span>
          </div>
          <button
            onClick={() => setFilekey('')}
            className="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors"
          >
            Close
          </button>
        </div>
      </div>
    </div>
  );
}
