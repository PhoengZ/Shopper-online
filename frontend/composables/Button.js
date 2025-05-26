export const differentButtonSize=(size)=>{
    switch (size) {
        case "small":
      return "text-xs px-4 py-2 space-x-2";
    default:
      return "text-xl px-7 py-3 space-x-2";
    }
};

export const differentVariant = (theme) => {
  switch (theme) {
    case "first":
      return "bg-blue-600 text-white hover:bg-blue-700 rounded-lg shadow-md transition duration-300 ease-in-out transform hover:scale-105";
    case "second":
      return "bg-green-600 text-white hover:bg-green-700 rounded-lg shadow-md transition duration-300 ease-in-out transform hover:scale-105";
    case "third":
      return "bg-red-600 text-white hover:bg-red-700 rounded-lg shadow-md transition duration-300 ease-in-out transform hover:scale-105";
    case "circular": // New circular theme
      return "bg-purple-600 text-white hover:bg-purple-700 rounded-full w-12 h-12 flex items-center justify-center shadow-md transition duration-300 ease-in-out transform hover:scale-110";
    case "ghost": // Renamed 'fourth' for clarity
      return "bg-transparent text-gray-700 hover:text-blue-600 hover:underline transition duration-300 ease-in-out";
    default:
      return "bg-gray-600 text-white hover:bg-gray-700 rounded-lg shadow-md transition duration-300 ease-in-out transform hover:scale-105";
  }
};