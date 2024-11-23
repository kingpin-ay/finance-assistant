import { AxiosRequestConfig } from "axios";
import { baseUrlBuilder } from "../helpers/basicUtilities";
import { axiosGet, axiosPost, urlMaker } from "../helpers/request";

const POSTFIX = "auth";

interface AuthType {
  email: string;
  password: string;
}

class AuthClient {
  private static BASE_URL = baseUrlBuilder(POSTFIX);
  private static instance: AuthClient | null = null;

  constructor() {
    if (AuthClient.instance) {
      return AuthClient.instance;
    }
    AuthClient.instance = this;
  }

  async login(values: AuthType) {
    const endPoint = "login";
    return await axiosPost(
      urlMaker(AuthClient.BASE_URL, endPoint),
      values,
    );
  }
  async validate() {
    const endPoint = "validate";
    return await axiosGet(urlMaker(AuthClient.BASE_URL, endPoint));
  }
}

export const authClient = new AuthClient();
