import React from "react";

export const MainWrapper = ({ children }: { children: React.ReactNode }) => {
  return (
    <main className="mx-auto px-4 sm:px-6 md:px-8 lg:px-10 xl:px-12 h-screen">
      {children}
    </main>
  );
};
