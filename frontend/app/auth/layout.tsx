const layout = ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) => {
  return (
    <div className="h-screen w-screen p-8 flex items-center justify-center gap-4">
      <LeftBar />
      {children}
    </div>
  );
};

function LeftBar() {
  return (
    <div className="w-1/2 h-full bg-gray-200 rounded-lg">
      <h1></h1>
    </div>
  );
}

export default layout;
