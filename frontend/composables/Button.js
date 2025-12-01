export const differentButtonSize = (size) => {
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
      return "bg-primary-600 text-white hover:bg-primary-700 rounded-lg shadow-sm transition duration-300 ease-in-out transform hover:scale-[1.02] active:scale-95";
    case "second":
      return "bg-secondary-600 text-white hover:bg-secondary-700 rounded-lg shadow-sm transition duration-300 ease-in-out transform hover:scale-[1.02] active:scale-95";
    case "third":
      return "bg-red-600 text-white hover:bg-red-700 rounded-lg shadow-sm transition duration-300 ease-in-out transform hover:scale-[1.02] active:scale-95";
    case "circular":
      return "bg-neutral-100 text-neutral-600 hover:bg-neutral-200 rounded-full w-10 h-10 flex items-center justify-center shadow-sm transition duration-300 ease-in-out";
    case "fourth": // Keeping 'fourth' as it's used in LoginForm
    case "ghost":
      return "bg-transparent text-neutral-600 hover:text-primary-600 hover:bg-neutral-50 rounded-lg transition duration-300 ease-in-out";
    default:
      return "bg-neutral-800 text-white hover:bg-neutral-900 rounded-lg shadow-sm transition duration-300 ease-in-out transform hover:scale-[1.02] active:scale-95";
  }
};
