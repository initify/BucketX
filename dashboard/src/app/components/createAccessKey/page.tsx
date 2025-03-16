import { useState } from 'react';
import { FiX, FiCopy, FiCheck } from 'react-icons/fi';

export default function CreateAccessKey({ 
  setIsCreateKeyOpen,
  onKeyCreated 
}: {
  setIsCreateKeyOpen: (isCreateKeyOpen: boolean) => void,
  onKeyCreated: () => void
}) {
  const [keyName, setKeyName] = useState('');
  const [newKey, setNewKey] = useState<string | null>(null);
  const [copied, setCopied] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleCreateKey = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError(null);

    try {
      const formData = new FormData();
      formData.append('key_name', keyName);

      const res = await fetch("http://localhost:8080/api/v1/keys", {
        method: "POST",
        body: formData,
      });

      if (!res.ok) {
        throw new Error(`Error: ${res.status} ${res.statusText}`);
      }

      const data = await res.json();
      setNewKey(data.key);
      onKeyCreated();
    } catch (error) {
      console.error('Failed to create access key:', error);
      setError('Failed to create access key. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  const handleCopyKey = () => {
    if (!newKey) return;
    
    navigator.clipboard.writeText(newKey)
      .then(() => {
        setCopied(true);
        setTimeout(() => setCopied(false), 2000);
      })
      .catch(err => {
        console.error('Failed to copy key:', err);
      });
  };

  const handleClose = () => {
    if (newKey) {
      const confirmClose = confirm('Make sure you have copied your access key. It won\'t be shown again. Are you sure you want to close?');
      if (!confirmClose) return;
    }
    setIsCreateKeyOpen(false);
  };

  return (
    <div className="fixed inset-0 z-50 bg-black/70 backdrop-blur-sm flex items-center justify-center p-4">
      <div className="bg-gray-900 rounded-xl shadow-2xl max-w-md w-full overflow-hidden border border-gray-800">
        <div className="flex justify-between items-center p-4 border-b border-gray-800">
          <h2 className="text-xl font-semibold text-white">Create Access Key</h2>
          <button
            onClick={handleClose}
            className="p-2 bg-gray-800 hover:bg-gray-700 rounded-full transition-colors text-gray-300"
          >
            <FiX />
          </button>
        </div>
        
        <div className="p-6">
          {newKey ? (
            <div className="space-y-4">
              <div className="bg-green-900/20 border border-green-800 rounded-lg p-4 text-center">
                <p className="text-green-300 mb-2">Your access key has been created!</p>
                <p className="text-gray-400 text-sm">Make sure to copy this key now. You won't be able to see it again.</p>
              </div>
              
              <div className="bg-gray-800 p-3 rounded-lg">
                <div className="flex items-center justify-between">
                  <code className="text-blue-400 font-mono text-sm break-all">{newKey}</code>
                  <button
                    onClick={handleCopyKey}
                    className="ml-2 p-2 bg-gray-700 hover:bg-gray-600 rounded-full transition-colors"
                    title={copied ? "Copied!" : "Copy to clipboard"}
                  >
                    {copied ? <FiCheck className="text-green-400" /> : <FiCopy className="text-gray-300" />}
                  </button>
                </div>
              </div>
              
              <button
                onClick={handleClose}
                className="w-full px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors"
              >
                Done
              </button>
            </div>
          ) : (
            <form onSubmit={handleCreateKey}>
              <div className="space-y-4">
                <div>
                  <label htmlFor="keyName" className="block text-sm font-medium text-gray-400 mb-1">
                    Key Name (Optional)
                  </label>
                  <input
                    type="text"
                    id="keyName"
                    value={keyName}
                    onChange={(e) => setKeyName(e.target.value)}
                    className="w-full bg-gray-800 border border-gray-700 rounded-lg px-4 py-2 text-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    placeholder="Enter a name for this key"
                  />
                </div>
                
                {error && (
                  <div className="bg-red-900/20 border border-red-800 text-red-200 p-3 rounded-lg text-sm">
                    {error}
                  </div>
                )}
                
                <div className="flex justify-end gap-3">
                  <button
                    type="button"
                    onClick={handleClose}
                    className="px-4 py-2 text-gray-300 hover:text-white transition-colors"
                  >
                    Cancel
                  </button>
                  <button
                    type="submit"
                    disabled={isLoading}
                    className="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors flex items-center gap-2"
                  >
                    {isLoading ? (
                      <>
                        <div className="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
                        Creating...
                      </>
                    ) : (
                      'Create Key'
                    )}
                  </button>
                </div>
              </div>
            </form>
          )}
        </div>
      </div>
    </div>
  );
}