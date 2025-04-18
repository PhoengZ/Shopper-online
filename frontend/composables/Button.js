export const differentButtonSize=(size)=>{
    switch (size) {
        case "small":
      return "text-xs px-4 py-2 space-x-2";
    default:
      return "text-xl px-7 py-3 space-x-2";
    }
};

export const differentVariant = (theme)=>{
    switch (theme){
        case "first":
            return "bg-blue-500 text-white hover:bg-blue-600 rounded-md border-2";
        case "second":
            return "bg-green-500 text-white hover:bg-green-600 rounded-md border-2";
        case "third":
            return "bg-red-500 text-white hover:bg-red-600 rounded-md border-2";    
        case "fourth":
            return "bg-none text-gray-500 hover:underline";
        default:
            return "bg-gray-500 text-white hover:bg-gray-600 rounded-md border-2";
    }
};