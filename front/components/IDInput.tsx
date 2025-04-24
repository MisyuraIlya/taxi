"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { useRide } from "@/provider/RideContext";
import { useRouter } from "next/navigation";
import { useRideStore } from "@/store/ride.store";

export default function IDInput() {
  const { role, setUserId} = useRide();
  const { longitude, setLongitude, latitude, setLatitude } = useRideStore()
  const [id, setId] = useState("");
  const router = useRouter()

  const handleCLick = () => {
    router.push("/match");
    setUserId(id)
  }

  const test = () => {
    if (role === "driver") {
      setUserId('driver123')
    }
    if( role === "client") {
      setUserId('client123')
    }
    setLatitude(37.7723)
    setLongitude(-122.4182)
    router.push("/match");
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
      {role === "client" && (      
        <p>
          test : client123 <br/>
          latitude: 37.7723 <br/>
          longitude: -122.4182 <br/>
        </p>
        )
      }

      {role === "driver" && (      
        <p>
          test: driver123 <br/>
          latitude: 37.773 <br/>
          longitude: -122.4182 <br/>
        </p>
        )
      }

      <Button onClick={() => test()}>
        test
      </Button>
      <Button onClick={() => handleCLick()} disabled={!id}>
        Continue
      </Button>
    </div>
  );
}