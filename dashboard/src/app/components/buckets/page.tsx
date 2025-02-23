"use client";
import { useEffect, useState } from "react";

interface Bucket {
  bucket_name: string;
  file_count: number;
  size: number;
}

export default function Buckets() {
  const [buckets, setBuckets] = useState<Bucket[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchBuckets = async () => {
      try {
        setIsLoading(true);
        setError(null);
        const res = await fetch("http://localhost:8080/api/v1/buckets");
        if (!res.ok) {
          throw new Error(`HTTP error! status: ${res.status}`);
        }
        const data = await res.json();
        setBuckets(data.buckets);
      } catch (err) {
        setError(err instanceof Error ? err.message : 'Failed to fetch buckets');
        setBuckets([]);
      } finally {
        setIsLoading(false);
      }
    };
    fetchBuckets();
  }, []);

  if (isLoading) {
    return <div className="flex-1 flex items-center justify-center">Loading...</div>;
  }

  if (error) {
    return <div className="flex-1 flex items-center justify-center text-red-500">{error}</div>;
  }

  return (
    <div className="flex flex-col gap-2 bg-gray-500 flex-1 p-4">
      <div className="px-4 py-6 bg-black text-2xl">
        <p>Object Browser</p>
        <div className="mt-4">
          <input
            type="text"
            placeholder="Filter Buckets"
            className="w-full bg-gray-800 border border-gray-700 rounded px-4 py-2 text-base"
          />
        </div>
      </div>

      <table className="w-full text-left">
        <thead className="bg-gray-800">
          <tr>
            <th className="px-4 py-3 text-sm font-medium">Name</th>
            <th className="px-4 py-3 text-sm font-medium">Objects</th>
            <th className="px-4 py-3 text-sm font-medium">Size</th>
          </tr>
        </thead>
        <tbody>
          {buckets.map((bucket, i) => (
            <tr key={i} className="border-t border-gray-700 hover:bg-gray-800 cursor-pointer">
              <td className="px-4 py-3">{bucket.bucket_name}</td>
              <td className="px-4 py-3">{bucket.file_count}</td>
              <td className="px-4 py-3">{bucket.size} B</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
