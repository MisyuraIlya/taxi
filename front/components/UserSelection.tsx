"use client";

import { Button } from "@/components/ui/button";

export default function UserSelection({ onSelect }: { onSelect: (role: "driver" | "client") => void }) {
  return (
    <div className="flex items-center justify-center h-screen space-x-4">
      <Button className="px-6 py-3" onClick={() => onSelect("driver")}>
        I'm a Driver
      </Button>
      <Button className="px-6 py-3" onClick={() => onSelect("client")}>
        I'm a Client
      </Button>
    </div>
  );
}
