"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { useRide } from "@/provider/RideContext";
import { useRouter } from "next/navigation";

export default function IDInput() {
  const { role, setUserId } = useRide();
  const [id, setId] = useState("");
  const router = useRouter()

  const handleCLick = () => {
    router.push("/match");
    setUserId(id)
  }

  return (
    <div className="flex flex-col items-center justify-center h-screen space-y-4">
      <h2 className="text-2xl font-semibold">Enter your {role} ID</h2>
      <input
        value={id}
        onChange={(e) => setId(e.target.value)}
        className="border rounded-lg p-2 w-64"
        placeholder="Your ID"
      />
      <Button onClick={() => handleCLick()} disabled={!id}>
        Continue
      </Button>
    </div>
  );
}