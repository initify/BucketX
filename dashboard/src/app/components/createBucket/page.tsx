export default function CreateBucket({ setIsNewBucketOpen }: {
  setIsNewBucketOpen: (isNewBucketOpen: boolean) => void
}) {

  const handleCreateBucket = async (bucketName: string) => {
    try {
      const formData = new FormData();
      formData.append('bucket_name', bucketName);

      const res = await fetch("http://localhost:8080/api/v1/bucket", {
        method: "POST",
        body: formData,
      });

      if (!res.ok) {
        throw new Error(`Error: ${res.status} ${res.statusText}`);
      }

      console.log('Bucket created successfully');
    } catch (error) {
      console.error('Failed to create bucket:', error);
    } finally {
      setIsNewBucketOpen(false);
    }
  };

  return (
    <div className="fixed inset-0 z-50 bg-black/70 backdrop-blur-sm flex items-center justify-center p-4">
      <div className="bg-gray-900 rounded-xl shadow-2xl max-w-md w-full overflow-hidden border border-gray-800">
        <div className="p-6">
          <h2 className="text-xl font-semibold text-white mb-4">Create New Bucket</h2>
          <form onSubmit={(e) => {
            e.preventDefault();
            const formData = new FormData(e.currentTarget);
            handleCreateBucket(formData.get('bucketName') as string);
          }}>
            <div className="space-y-4">
              <div>
                <label htmlFor="bucketName" className="block text-sm font-medium text-gray-400 mb-1">
                  Bucket Name
                </label>
                <input
                  type="text"
                  id="bucketName"
                  name="bucketName"
                  required
                  className="w-full bg-gray-800 border border-gray-700 rounded-lg px-4 py-2 text-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="Enter bucket name"
                />
              </div>
              <div className="flex justify-end gap-3">
                <button
                  type="button"
                  onClick={() => setIsNewBucketOpen(false)}
                  className="px-4 py-2 text-gray-300 hover:text-white transition-colors"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  className="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors"
                >
                  Create Bucket
                </button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  )
}