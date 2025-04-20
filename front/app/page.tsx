"use client";

import { useState } from "react";
import UserSelection from "@/components/UserSelection";
import IDInput from "@/components/IDInput";
import MapPage from "@/components/MapPage";

export default function Home() {
  const [role, setRole] = useState<"driver" | "client" | null>(null);
  const [userId, setUserId] = useState<string | null>(null);

  if (!role) return <UserSelection onSelect={setRole} />;
  if (!userId) return <IDInput role={role} onSubmit={setUserId} />;
  return <MapPage role={role} userId={userId} />;
}
