"use client";

import { useEffect } from 'react';
import { initFaro } from '@/lib/faro';
import { useFaroSession } from '@/lib/hooks/use-faro-session';
import { faroPerformance } from '@/lib/faro-utils';

export function FaroProvider({ children }: { children: React.ReactNode }) {
  // Initialize Faro SDK
  useEffect(() => {
    initFaro();
    faroPerformance.measurePageLoad();
  }, []);

  // Track user session
  useFaroSession();

  return <>{children}</>;
}