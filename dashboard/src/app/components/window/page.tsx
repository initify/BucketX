"use client";
import Buckets from "../buckets/page";
import Files from "../files/page";
import Viewer from "../viewer/page";
import { useState } from "react";
import { FiSearch, FiArrowLeft, FiRefreshCw, FiUpload, FiSettings } from 'react-icons/fi';

export default function Window() {
  const [filekey, setFilekey] = useState<string>("");
  const [selectedBucket, setSelectedBucket] = useState<string>("");
  const [view, setView] = useState<"buckets" | "files">("buckets");

  const handleBucketSelect = (bucketId: string) => {
    setSelectedBucket(bucketId);
    setView("files");
  };

  const handleBackToBuckets = () => {
    setSelectedBucket("");
    setView("buckets");
  };

  return (
    <div className="flex flex-col bg-gray-900 flex-1 p-6 rounded-l-xl shadow-2xl">
      <div className="flex justify-between items-center mb-6">
        <div className="flex items-center">
          {view === "files" && (
            <button 
              onClick={handleBackToBuckets}
              className="mr-4 p-2 bg-gray-800 rounded-full hover:bg-gray-700 transition-colors"
            >
              <FiArrowLeft className="text-gray-300" />
            </button>
          )}
          <h1 className="text-2xl font-bold text-white">
            {view === "buckets" ? "Object Browser" : `Bucket: ${selectedBucket}`}
          </h1>
        </div>
        
        <div className="flex items-center gap-3">
          <button className="p-2 bg-gray-800 rounded-full hover:bg-gray-700 transition-colors">
            <FiRefreshCw className="text-gray-300" />
          </button>
          <button className="p-2 bg-gray-800 rounded-full hover:bg-gray-700 transition-colors">
            <FiUpload className="text-gray-300" />
          </button>
          <button className="p-2 bg-gray-800 rounded-full hover:bg-gray-700 transition-colors">
            <FiSettings className="text-gray-300" />
          </button>
        </div>
      </div>

      <div className="relative mb-6">
        <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
          <FiSearch className="h-5 w-5 text-gray-400" />
        </div>
        <input
          type="text"
          placeholder={view === "buckets" ? "Search buckets..." : "Search files..."}
          className="w-full bg-gray-800 border border-gray-700 rounded-lg pl-10 pr-4 py-3 text-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        />
      </div>

      <div className="bg-gray-800 rounded-xl p-5 flex-1 overflow-hidden">
        {view === "buckets" && (
          <Buckets setSelectedBucket={handleBucketSelect} />
        )}
        
        {view === "files" && (
          <Files 
            setFilekey={setFilekey} 
            selectedBucket={selectedBucket} 
          />
        )}
      </div>
      
      {filekey && (
        <Viewer fileKey={filekey} setFilekey={setFilekey} />
      )}
    </div>
  );
}
