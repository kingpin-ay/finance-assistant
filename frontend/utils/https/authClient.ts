import { axiosGet, axiosPost, urlMaker } from "../helpers/request";

const POSTFIX = "auth";
// const URL =
//   process.env.NODE_ENV === "development"
//     ? process.env.DEV_URL
//     : process.env.DEV_URL;

interface AuthType {
  email: string;
  password: string;
}

class AuthClient {
  private static BASE_URL = `http://localhost:5001/${POSTFIX}`;
  private static instance: AuthClient | null = null;

  constructor() {
    if (AuthClient.instance) {
      return AuthClient.instance;
    }
    AuthClient.instance = this;
  }

  async login(values: AuthType) {
    const endPoint = "login";
    return await axiosPost(urlMaker(AuthClient.BASE_URL, endPoint), values);
  }
  async validate() {
    const endPoint = "validate";
    return await axiosGet(urlMaker(AuthClient.BASE_URL, endPoint));
  }
}

export const authClient = new AuthClient();
