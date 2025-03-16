"use client";
import { useState } from 'react';
import Window from './components/window/page';
import Sidebar from './components/sidebar/page';

export default function Home() {
  const [currentView, setCurrentView] = useState<"buckets" | "files" | "access-keys">("buckets");

  const handleViewChange = (view: "buckets" | "files" | "access-keys") => {
    setCurrentView(view);
  };

  return (
    <div className="flex min-h-screen w-full bg-gray-900">
      <Sidebar onViewChange={handleViewChange} />
      <Window initialView={currentView} />
    </div>
  );
}
