export const PaperWrapper = ({ children }: { children: React.ReactNode }) => {
  return (
    <div
      className="paper paper-yellow 
     shadow-xl"
    >
      <div className="tape-section"></div>
      {children}
      <div className="tape-section"></div>
    </div>
  );
};
