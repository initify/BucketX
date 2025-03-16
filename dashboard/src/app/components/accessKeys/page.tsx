"use client";
import { useState, useEffect } from 'react';
import { FiKey, FiPlus, FiTrash2, FiCopy } from 'react-icons/fi';
import CreateAccessKey from '../createAccessKey/page';

export default function AccessKeys() {
  const [accessKeys, setAccessKeys] = useState<any[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [isCreateKeyOpen, setIsCreateKeyOpen] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [copiedKeyId, setCopiedKeyId] = useState<string | null>(null);

  useEffect(() => {
    fetchAccessKeys();
  }, []);

  const fetchAccessKeys = async () => {
    setIsLoading(true);
    try {
      // Temporary dummy data
      await new Promise(resolve => setTimeout(resolve, 1000)); // Simulate API delay
      setAccessKeys([
        {
          id: "1",
          name: "Test Key",
          key: "AKIAIOSFODNN7EXAMPLE",
          created_at: "2024-02-15T12:00:00Z"
        },
        {
          id: "2",
          name: "Backup Key",
          key: "AKIAR4VEXAMPLEKEY",
          created_at: "2024-01-30T09:30:00Z"
        }
      ]);
    } catch (error) {
      console.error('Failed to fetch access keys:', error);
      setError('Failed to load access keys. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  const handleCopyKey = (keyId: string, keyValue: string) => {
    navigator.clipboard.writeText(keyValue)
      .then(() => {
        setCopiedKeyId(keyId);
        setTimeout(() => setCopiedKeyId(null), 2000);
      })
      .catch(err => {
        console.error('Failed to copy key:', err);
      });
  };

  const handleDeleteKey = async (keyId: string) => {
    if (!confirm('Are you sure you want to delete this access key?')) return;
    
    try {
      // Simulate deletion by filtering local data
      setAccessKeys(prev => prev.filter(key => key.id !== keyId));
    } catch (error) {
      console.error('Failed to delete access key:', error);
      setError('Failed to delete access key. Please try again.');
    }
  };

  return (
    <div className="flex flex-col h-full">
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-xl font-semibold text-white">Access Keys</h2>
        <button
          onClick={() => setIsCreateKeyOpen(true)}
          className="flex items-center gap-2 px-3 py-1.5 bg-blue-600 hover:bg-blue-700 text-white rounded-md transition-colors text-sm font-medium"
        >
          <FiPlus className="w-4 h-4" />
          New Key
        </button>
      </div>

      {isLoading ? (
        <div className="flex items-center justify-center h-64">
          <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
        </div>
      ) : error ? (
        <div className="bg-red-900/20 border border-red-800 text-red-200 p-4 rounded-lg">
          {error}
        </div>
      ) : accessKeys.length === 0 ? (
        <div className="flex flex-col items-center justify-center h-64 text-gray-400">
          <FiKey className="w-12 h-12 mb-4 opacity-50" />
          <p className="text-lg mb-2">No access keys found</p>
          <p className="text-sm">Create a new access key to get started</p>
        </div>
      ) : (
        <div className="grid gap-4">
          {accessKeys.map((key) => (
            <div key={key.id} className="bg-gray-900 border border-gray-800 rounded-lg p-4 flex justify-between items-center">
              <div>
                <div className="flex items-center gap-2 mb-1">
                  <FiKey className="text-blue-400" />
                  <span className="text-gray-200 font-medium">{key.name || 'Unnamed Key'}</span>
                </div>
                <div className="flex items-center">
                  <code className="bg-gray-800 px-2 py-1 rounded text-gray-400 text-sm font-mono">
                    {key.key.substring(0, 8)}...{key.key.substring(key.key.length - 8)}
                  </code>
                  <button 
                    onClick={() => handleCopyKey(key.id, key.key)}
                    className="ml-2 text-gray-500 hover:text-gray-300"
                    title="Copy key"
                  >
                    <FiCopy className="w-4 h-4" />
                  </button>
                </div>
                <p className="text-gray-500 text-xs mt-1">Created: {new Date(key.created_at).toLocaleString()}</p>
              </div>
              <button
                onClick={() => handleDeleteKey(key.id)}
                className="p-2 text-gray-500 hover:text-red-400 transition-colors"
                title="Delete key"
              >
                <FiTrash2 />
              </button>
            </div>
          ))}
        </div>
      )}

      {isCreateKeyOpen && <CreateAccessKey setIsCreateKeyOpen={setIsCreateKeyOpen} onKeyCreated={fetchAccessKeys} />}
    </div>
  );
}