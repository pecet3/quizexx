export const PaperWrapper = ({ children }: { children: React.ReactNode }) => {
  return (
    <div
      className="paper paper-yellow 
     shadow-lg shadow-gray-100 w-full md:max-w-lg max-w-sm"
    >
      <div className="tape-section"></div>
      {children}
      <div className="tape-section"></div>
    </div>
  );
};
export const LittlePaperWrapper = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  return (
    <div
      className="paper paper-yellow 
     shadow-lg shadow-gray-100 "
    >
      <div className="top-tape"></div>
      <div className="pt-10 p-2"> {children}</div>
    </div>
  );
};
export const PatternWrapper = ({ children }: { children: React.ReactNode }) => {
  return (
    <div
      className="paper m-0
     w-full md:max-w-lg max-w-sm"
    >
      <div className="tape-section"></div>

      {children}
      <div className="tape-section"></div>
    </div>
  );
};
