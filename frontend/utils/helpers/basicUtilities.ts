export function baseUrlBuilder(path: string) {
  const URL =
    process.env.NEXT_PUBLIC_NODE_ENV === "development"
      ? process.env.NEXT_PUBLIC_DEV_URL
      : process.env.NEXT_PUBLIC_PROD_URL;
  return `${URL}/${path}`;
}
