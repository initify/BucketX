"use client";
import { useEffect, useState } from "react";
import { FiFolder, FiHardDrive, FiDatabase } from "react-icons/fi";

interface Bucket {
  bucket_name: string;
  file_count: number;
  size: number;
}

export default function Buckets({ setSelectedBucket }: {
  setSelectedBucket: (bucketId: string) => void
}) {
  const [buckets, setBuckets] = useState<Bucket[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchBuckets = async () => {
      setLoading(true);
      try {
        const res = await fetch("http://localhost:8080/api/v1/buckets");
        if (!res.ok) {
          throw new Error(`HTTP error! status: ${res.status}`);
        }
        const data = await res.json();
        setBuckets(data.buckets);
      } catch (err) {
        setBuckets([]);
      } finally {
        setLoading(false);
      }
    };
    fetchBuckets();
  }, []);

  const formatSize = (bytes: number) => {
    if (bytes === 0) return '0 B';
    const k = 1024;
    const sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
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
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {buckets.map((bucket, i) => (
          <div
            key={i}
            className="bg-gray-900 rounded-xl p-4 cursor-pointer hover:bg-gray-700 transition-all duration-200 border border-gray-800"
            onClick={() => setSelectedBucket(bucket.bucket_name)}
          >
            <div className="flex items-center mb-3">
              <div className="p-3 bg-blue-500 bg-opacity-20 rounded-lg mr-3">
                <FiDatabase className="text-blue-400 text-xl" />
              </div>
              <h3 className="text-lg font-medium text-white truncate">{bucket.bucket_name}</h3>
            </div>
            <div className="flex justify-between text-sm text-gray-400">
              <div className="flex items-center">
                <FiFolder className="mr-1" />
                <span>{bucket.file_count} objects</span>
              </div>
              <div className="flex items-center">
                <FiHardDrive className="mr-1" />
                <span>{formatSize(bucket.size)}</span>
              </div>
            </div>
          </div>
        ))}
      </div>

      {buckets.length === 0 && (
        <div className="text-center py-10">
          <div className="inline-flex items-center justify-center w-16 h-16 rounded-full bg-gray-700 mb-4">
            <FiDatabase className="text-gray-400 text-2xl" />
          </div>
          <h3 className="text-lg font-medium text-white mb-2">No buckets found</h3>
          <p className="text-gray-400">Create a bucket to get started</p>
        </div>
      )}
    </div>
  );
}
