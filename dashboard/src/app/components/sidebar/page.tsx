'use client';
import { useState } from 'react';
import Image from 'next/image';
import Logo from '../../assets/logo.png';
import {
  FiFolder,
  FiKey,
  FiBook,
  FiShield,
  FiUsers,
  FiMenu,
} from 'react-icons/fi';

export type SidebarOption =
  | 'buckets'
  | 'access-keys'
  | 'documentation'
  | 'admin-buckets'
  | 'policies'
  | 'identity';

export default function Sidebar({ onViewChange }: {
  onViewChange: (view: "buckets" | "files" | "access-keys") => void
}) {
  const [isOpen, setIsOpen] = useState(false);
  const [selectedItem, setSelectedItem] = useState<SidebarOption>('buckets');

  const handleItemClick = (item: SidebarOption) => {
    setSelectedItem(item);
    
    // Map sidebar options to window views
    if (item === 'buckets') {
      onViewChange("buckets");
    } else if (item === 'access-keys') {
      onViewChange("access-keys");
    }
    // Other options can be handled as needed
  };

  return (
    <div
      className={`flex flex-col bg-gray-900 transition-all duration-300 ${
        isOpen ? 'w-[300px]' : 'w-[80px]'
      } overflow-hidden`}
    >
      <div className="border-b border-gray-800">
        {isOpen ? (
          <div className="px-6 h-16 flex justify-between items-center">
            <div className="flex items-center gap-2">
              <span className="text-lg font-semibold">
                <Image src={Logo} alt="logo" width={40} height={40} />
              </span>
              <span className="text-lg transition-opacity duration-200 ease-in-out whitespace-nowrap">
                OBJECT STORE
              </span>
            </div>
            <button
              onClick={() => setIsOpen(false)}
              className="text-gray-400 hover:text-white transition-colors"
            >
              <FiMenu size={20} />
            </button>
          </div>
        ) : (
          <div className="h-16 flex justify-center items-center">
            <button
              onClick={() => setIsOpen(true)}
              className="text-gray-400 hover:text-white transition-colors"
            >
              <FiMenu size={20} />
            </button>
          </div>
        )}
      </div>

      <div className="p-4">
        <div className="mb-6">
          {isOpen && (
            <h3 className="text-gray-400 text-sm mb-2 whitespace-nowrap">
              User
            </h3>
          )}
          <div className="space-y-1">
            <div
              onClick={() => handleItemClick('buckets')}
              className={`flex items-center ${isOpen ? 'gap-2 px-4' : 'justify-center'} py-2 ${selectedItem === 'buckets' ? 'bg-blue-900 text-white' : 'text-gray-300 hover:bg-gray-800'} rounded cursor-pointer`}
            >
              <FiFolder className="w-5 h-5 flex-shrink-0" />
              {isOpen && (
                <span className="transition-opacity duration-200 ease-in-out whitespace-nowrap">
                  Object Browser
                </span>
              )}
            </div>
            <div
              onClick={() => handleItemClick('access-keys')}
              className={`flex items-center ${isOpen ? 'gap-2 px-4' : 'justify-center'} py-2 ${selectedItem === 'access-keys' ? 'bg-blue-900 text-white' : 'text-gray-300 hover:bg-gray-800'} rounded cursor-pointer`}
            >
              <FiKey className="w-5 h-5 flex-shrink-0" />
              {isOpen && (
                <span className="transition-opacity duration-200 ease-in-out whitespace-nowrap">
                  Access Keys
                </span>
              )}
            </div>
            <div
              onClick={() => handleItemClick('documentation')}
              className={`flex items-center ${isOpen ? 'gap-2 px-4' : 'justify-center'} py-2 ${selectedItem === 'documentation' ? 'bg-blue-900 text-white' : 'text-gray-300 hover:bg-gray-800'} rounded cursor-pointer`}
            >
              <FiBook className="w-5 h-5 flex-shrink-0" />
              {isOpen && (
                <span className="transition-opacity duration-200 ease-in-out whitespace-nowrap">
                  Documentation
                </span>
              )}
            </div>
          </div>
        </div>

        <div className="mb-6">
          {isOpen && (
            <h3 className="text-gray-400 text-sm mb-2 whitespace-nowrap">
              Administrator
            </h3>
          )}
          <div className="space-y-1">
            <div
              onClick={() => handleItemClick('admin-buckets')}
              className={`flex items-center ${isOpen ? 'gap-2 px-4' : 'justify-center'} py-2 ${selectedItem === 'admin-buckets' ? 'bg-blue-900 text-white' : 'text-gray-300 hover:bg-gray-800'} rounded cursor-pointer`}
            >
              <FiFolder className="w-5 h-5 flex-shrink-0" />
              {isOpen && (
                <span className="transition-opacity duration-200 ease-in-out whitespace-nowrap">
                  Buckets
                </span>
              )}
            </div>
            <div
              onClick={() => handleItemClick('policies')}
              className={`flex items-center ${isOpen ? 'gap-2 px-4' : 'justify-center'} py-2 ${selectedItem === 'policies' ? 'bg-blue-900 text-white' : 'text-gray-300 hover:bg-gray-800'} rounded cursor-pointer`}
            >
              <FiShield className="w-5 h-5 flex-shrink-0" />
              {isOpen && (
                <span className="transition-opacity duration-200 ease-in-out whitespace-nowrap">
                  Policies
                </span>
              )}
            </div>
            <div
              onClick={() => handleItemClick('identity')}
              className={`flex items-center ${isOpen ? 'gap-2 px-4' : 'justify-center'} py-2 ${selectedItem === 'identity' ? 'bg-blue-900 text-white' : 'text-gray-300 hover:bg-gray-800'} rounded cursor-pointer`}
            >
              <FiUsers className="w-5 h-5 flex-shrink-0" />
              {isOpen && (
                <span className="transition-opacity duration-200 ease-in-out whitespace-nowrap">
                  Identity
                </span>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
