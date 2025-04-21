"use client";

import UserSelection from "@/components/UserSelection";
import IDInput from "@/components/IDInput";
import { useRide } from "@/provider/RideContext";

function Main() {
  const { role, userId } = useRide();
  if (!role) return <UserSelection />;
  if (!userId) return <IDInput />;
}

export default function Home() {
  return (
      <Main />
  );
}
