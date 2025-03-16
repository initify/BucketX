"use client";
import Buckets from "../buckets/page";
import CreateBucket from "../createBucket/page";
import Files from "../files/page";
import Upload from "../upload/page";
import Viewer from "../viewer/page";
import AccessKeys from "../accessKeys/page"; // Import the AccessKeys component
import { useState, useEffect } from "react";
import { FiSearch, FiArrowLeft, FiRefreshCw, FiUpload, FiSettings, FiPlus, FiKey } from 'react-icons/fi';

export default function Window({ initialView }: {
  initialView: "buckets" | "files" | "access-keys"
}) {
  const [filekey, setFilekey] = useState<string>("");
  const [selectedBucket, setSelectedBucket] = useState<string>("");
  const [view, setView] = useState<"buckets" | "files" | "access-keys">(initialView); // Add access-keys to the view options
  const [isUploadOpen, setIsUploadOpen] = useState(false);
  const [isNewBucketOpen, setIsNewBucketOpen] = useState(false);

  const handleBucketSelect = (bucketId: string) => {
    setSelectedBucket(bucketId);
    setView("files");
  };

  const handleBackToBuckets = () => {
    setSelectedBucket("");
    setView("buckets");
  };

  // Function to handle view changes from sidebar
  const handleViewChange = (newView: "buckets" | "files" | "access-keys") => {
    if (newView === "buckets") {
      setSelectedBucket("");
    }
    setView(newView);
  };

  // Add useEffect to sync view with parent state
  useEffect(() => {
    setView(initialView);
  }, [initialView]);

  return (
    <div className="flex flex-col bg-gray-900 flex-1 p-6 rounded-l-xl shadow-2xl">
      <div className="flex justify-between items-center mb-6">
        <div className="flex items-center gap-4">
          {(view === "files" || view === "access-keys") && (
            <button 
              onClick={handleBackToBuckets}
              className="mr-4 p-2 bg-gray-800 rounded-full hover:bg-gray-700 transition-colors"
            >
              <FiArrowLeft className="text-gray-300" />
            </button>
          )}
          <h1 className="text-2xl font-bold text-white">
            {view === "buckets" ? "Object Browser" : 
             view === "files" ? `Bucket: ${selectedBucket}` : 
             "Access Keys"}
          </h1>
        </div>
        
        <div className="flex items-center gap-3">
          <button className="p-2 bg-gray-800 rounded-full hover:bg-gray-700 transition-colors">
            <FiRefreshCw className="text-gray-300" />
          </button>
          {view !== "access-keys" && (
            <button 
              onClick={() => setIsUploadOpen(true)}
              className="p-2 bg-gray-800 rounded-full hover:bg-gray-700 transition-colors"
            >
              <FiUpload className="text-gray-300" />
            </button>
          )}
          {/* Update the access keys button */}
          <button 
            onClick={() => handleViewChange("access-keys")}
            className={`p-2 ${
              view === "access-keys" ? "bg-blue-600" : "bg-gray-800 hover:bg-gray-700"
            } rounded-full transition-colors`}
          >
            <FiKey className="text-gray-300" />
          </button>
          <button className="p-2 bg-gray-800 rounded-full hover:bg-gray-700 transition-colors">
            <FiSettings className="text-gray-300" />
          </button>
        </div>
      </div>

      {view !== "access-keys" && (
        <div className="relative mb-4">
          <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <FiSearch className="h-5 w-5 text-gray-400" />
          </div>
          <input
            type="text"
            placeholder={view === "buckets" ? "Search buckets..." : "Search files..."}
            className="w-full bg-gray-800 border border-gray-700 rounded-lg pl-10 pr-4 py-3 text-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />
        </div>
      )}

      {view === "buckets" && (
        <div className="flex justify-end mb-4">
          <button
            onClick={() => setIsNewBucketOpen(true)}
            className="flex items-center gap-2 px-3 py-1.5 bg-blue-600 hover:bg-blue-700 text-white rounded-md transition-colors text-sm font-medium"
          >
            <FiPlus className="w-4 h-4" />
            New
          </button>
        </div>
      )}

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

        {view === "access-keys" && (
          <AccessKeys />
        )}
      </div>
      
      {filekey && (
        <Viewer fileKey={filekey} setFilekey={setFilekey} />
      )}
      
      {isUploadOpen && <Upload setIsUploadOpen={setIsUploadOpen} />}
      {isNewBucketOpen && <CreateBucket setIsNewBucketOpen={setIsNewBucketOpen} />}
    </div>
  );
}
