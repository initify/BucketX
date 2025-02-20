export default function Buckets() {
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
            <th className="px-4 py-3 text-sm font-medium">Access</th>
          </tr>
        </thead>
        <tbody>
          {Array.from({ length: 5 }).map((_, i) => (
            <tr key={i} className="border-t border-gray-700 hover:bg-gray-800 cursor-pointer">
              <td className="px-4 py-3">bucketx</td>
              <td className="px-4 py-3">0</td>
              <td className="px-4 py-3">0.0 B</td>
              <td className="px-4 py-3">R/W</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
