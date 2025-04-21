'use client'
import { Button } from '@/components/ui/button';
import React, { createContext, useContext } from 'react';
import { toast } from "sonner"
import { Toaster } from "@/components/ui/sonner"

interface NotificationContext {
    message: (message: string, description: string) => void;
}

const NotificationContext = createContext<NotificationContext | null>(null);

const useNoitifcation = () => {
    const context = useContext(NotificationContext)
    if(!context) {
        throw new Error("useNotification must be used within NotificationProvider");
    }
    return context;
}

interface NotificationProviderProps {
    children: React.ReactNode;
}

const NotificationProvider: React.FC<NotificationProviderProps> = ({ children }) => {

    const message = (message: string, description: string) => {
        toast(message, {
            description: description,
            action: {
                label: "Undo",
                onClick: () => console.log("Undo"),
            },
        });
    }

    return (
        <NotificationContext.Provider value={{message}}>
            {children}
            <Toaster />
        </NotificationContext.Provider>
    );
};

export { NotificationProvider, useNoitifcation };