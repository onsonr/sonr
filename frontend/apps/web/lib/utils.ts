import { ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

/**
 * The function "cn" takes in an array of class values and returns a string of concatenated class
 * names.
 * @param {ClassValue[]} inputs - The `inputs` parameter is a rest parameter that allows the function
 * to accept an arbitrary number of arguments. In this case, the arguments are expected to be class
 * names or class name arrays that will be concatenated and returned as a single string. The
 * `ClassValue` type indicates that the function expects inputs
 */
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}
