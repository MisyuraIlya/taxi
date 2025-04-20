"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";

export default function IDInput({ role, onSubmit }: { role: string; onSubmit: (id: string) => void }) {
  const [id, setId] = useState("");
  return (
    <div className="flex flex-col items-center justify-center h-screen space-y-4">
      <h2 className="text-2xl font-semibold">Enter your {role} ID</h2>
      <input
        value={id}
        onChange={(e) => setId(e.target.value)}
        className="border rounded-lg p-2 w-64"
        placeholder="Your ID"
      />
      <Button onClick={() => onSubmit(id)} disabled={!id}>
        Continue
      </Button>
    </div>
  );
}