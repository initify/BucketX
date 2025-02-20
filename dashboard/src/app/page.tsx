import Buckets from "./buckets/page";
import Sidebar from "./sidebar/page";
export default function Home() {
  return (
    <div className="flex min-h-screen w-full bg-gray-900">
      <Sidebar />
      <Buckets />
    </div>
  );
}
