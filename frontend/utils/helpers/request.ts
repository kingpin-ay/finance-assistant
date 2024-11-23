import axios, { AxiosRequestConfig } from "axios";

export function urlMaker(baseUrl: string, endpoint: string) {
  return `${baseUrl}/${endpoint}`;
}

export async function axiosPost(
  url: string,
  data: any,
  extra?: AxiosRequestConfig<any>
) {
  const config: AxiosRequestConfig<any> = extra
    ? extra
    : {
        headers: {
          "Content-Type": "application/json",
        },
        withCredentials: true,
      };
  const response = await axios.post(url, data, config);
  return response.status == 200 ? response.data : null;
}

export async function axiosGet(url: string) {
  const config: AxiosRequestConfig<any> = {
    headers: {
      "Content-Type": "application/json",
    },
    withCredentials: true,
  };
  const response = await axios.get(url, config);
  return response.status == 200 ? response.data : null;
}
