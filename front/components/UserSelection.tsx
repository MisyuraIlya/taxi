"use client";

import { Button } from "@/components/ui/button";
import { useRide } from "@/provider/RideContext";

export default function UserSelection() {
  const { setRole } = useRide();
  return (
    <div className="flex items-center justify-center h-screen space-x-4">
      <Button className="px-6 py-3" onClick={() => setRole("driver")}>
        I'm a Driver
      </Button>
      <Button className="px-6 py-3" onClick={() => setRole("client")}>
        I'm a Client
      </Button>
    </div>
  );
}
